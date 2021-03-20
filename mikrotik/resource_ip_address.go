package mikrotik

import (
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceIpAddress() *schema.Resource {
	return &schema.Resource{
		Create: resourceIpAddressCreate,
		Read:   resourceIpAddressRead,
		Update: resourceIpAddressUpdate,
		Delete: resourceIpAddressDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"address": {
				Type:     schema.TypeString,
				Required: true,
			},
			"network": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"interface": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceIpAddressCreate(d *schema.ResourceData, m interface{}) error {
	address := d.Get("address").(string)
	ifname := d.Get("interface").(string)

	c := m.(mikrotikConfig)

	ipaddr, err := c.AddIpAddress(address, ifname)

	if err != nil {
		return err
	}

	writeStateIpAddress(ipaddr, d)
	return nil
}

func resourceIpAddressRead(d *schema.ResourceData, m interface{}) error {
	c := m.(mikrotikConfig)

	ipaddr, err := c.FindIpAddress(d.Id())

	if err != nil {
		d.SetId("")
		return nil
	}

	if ipaddr == nil {
		d.SetId("")
		return nil
	}

	writeStateIpAddress(ipaddr, d)
	return nil
}

func resourceIpAddressUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(mikrotikConfig)

	address := d.Get("address").(string)
	ifname := d.Get("interface").(string)

	ipaddr, err := c.UpdateIpAddress(d.Id(), address, ifname)

	if err != nil {
		return err
	}

	writeStateIpAddress(ipaddr, d)
	return nil
}

func resourceIpAddressDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(mikrotikConfig)

	err := c.DeleteIpAddress(d.Id())

	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}

func writeStateIpAddress(ipaddr *IpAddress, d *schema.ResourceData) error {
	d.SetId(ipaddr.Id)
	d.Set("address", ipaddr.Address)
	d.Set("network", ipaddr.Network)
	d.Set("interface", ipaddr.Interface)
	return nil
}

type IpAddress struct {
	Id        string `mikrotik:".id"`
	Address   string `mikrotik:".address"`
	Network   string `mikrotik:".network"`
	Interface string `mikrotik:".interface"`
}

func (mikrotikClient mikrotikConfig) AddIpAddress(address string, ifname string) (*IpAddress, error) {
	c, err := mikrotikClient.getMikrotikClient()

	if err != nil {
		return nil, err
	}

	cmd := []string{
		"/ip/address/add",
		fmt.Sprintf("=address=%s", address),
		fmt.Sprintf("=interface=%s", ifname),
	}

	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	r, err := c.RunArgs(cmd)

	log.Printf("[DEBUG] ip address creation response: `%v`", r)

	if err != nil {
		return nil, err
	}

	id := r.Done.Map["ret"]

	return mikrotikClient.FindIpAddress(id)
}

func (mikrotikClient mikrotikConfig) FindIpAddress(id string) (*IpAddress, error) {
	c, err := mikrotikClient.getMikrotikClient()

	if err != nil {
		return nil, err
	}

	cmd := strings.Split(fmt.Sprintf("/ip/address/print ?.id=%s", id), " ")
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	r, err := c.RunArgs(cmd)

	log.Printf("[DEBUG] ip address response: %v", r)

	if err != nil {
		return nil, err
	}

	ipaddr := IpAddress{}
	err = Unmarshal(*r, &ipaddr)

	if err != nil {
		return nil, err
	}

	if ipaddr.Id == "" {
		return nil, NewNotFound(fmt.Sprintf("ip address `%s`not found", id))
	}

	return &ipaddr, nil
}

func (mikrotikClient mikrotikConfig) UpdateIpAddress(id string, address string, ifname string) (*IpAddress, error) {
	c, err := mikrotikClient.getMikrotikClient()

	if err != nil {
		return nil, err
	}

	cmd := []string{
		"/ip/address/set",
		fmt.Sprintf("=.id=%s", id),
		fmt.Sprintf("=address=%s", address),
		fmt.Sprintf("=interface=%s", ifname),
	}

	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	r, err := c.RunArgs(cmd)

	log.Printf("[DEBUG] ip address update response: `%v`", r)

	if err != nil {
		return nil, err
	}

	return mikrotikClient.FindIpAddress(id)
}

func (mikrotikClient mikrotikConfig) DeleteIpAddress(id string) error {
	c, err := mikrotikClient.getMikrotikClient()

	if err != nil {
		return err
	}

	cmd := []string{
		"/ip/address/remove",
		fmt.Sprintf("=.id=%s", id),
	}

	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	r, err := c.RunArgs(cmd)

	log.Printf("[DEBUG] ip address delete response: `%v`", r)

	return err
}
