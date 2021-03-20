package mikrotik

import (
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceIpFirewallAddressList() *schema.Resource {
	return &schema.Resource{
		Create: resourceIpFirewallAddressListCreate,
		Read:   resourceIpFirewallAddressListRead,
		Update: resourceIpFirewallAddressListUpdate,
		Delete: resourceIpFirewallAddressListDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"address": {
				Type:     schema.TypeString,
				Required: true,
			},
			"list": {
				Type:     schema.TypeString,
				Required: true,
			},
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceIpFirewallAddressListCreate(d *schema.ResourceData, m interface{}) error {
	address := d.Get("address").(string)
	list := d.Get("list").(string)
	comment := d.Get("comment").(string)

	c := m.(mikrotikConfig)

	addresslist, err := c.AddIpFirewallAddressList(address, list, comment)

	if err != nil {
		return err
	}

	writeStateIpFirewallAddressList(addresslist, d)
	return nil
}

func resourceIpFirewallAddressListRead(d *schema.ResourceData, m interface{}) error {
	c := m.(mikrotikConfig)

	addresslist, err := c.FindIpFirewallAddressList(d.Id())

	if err != nil {
		d.SetId("")
		return nil
	}

	if addresslist == nil {
		d.SetId("")
		return nil
	}

	writeStateIpFirewallAddressList(addresslist, d)
	return nil
}

func resourceIpFirewallAddressListUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(mikrotikConfig)

	address := d.Get("address").(string)
	list := d.Get("list").(string)
	comment := d.Get("comment").(string)

	addresslist, err := c.UpdateIpFirewallAddressList(d.Id(), address, list, comment)

	if err != nil {
		return err
	}

	writeStateIpFirewallAddressList(addresslist, d)
	return nil
}

func resourceIpFirewallAddressListDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(mikrotikConfig)

	err := c.DeleteIpFirewallAddressList(d.Id())

	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}

func writeStateIpFirewallAddressList(addresslist *IpFirewallAddressList, d *schema.ResourceData) error {
	d.SetId(addresslist.Id)
	d.Set("address", addresslist.Address)
	d.Set("list", addresslist.List)
	d.Set("comment", addresslist.Comment)
	return nil
}

type IpFirewallAddressList struct {
	Id      string `mikrotik:".id"`
	Address string `mikrotik:".address"`
	List    string `mikrotik:".list"`
	Comment string `mikrotik:".comment"`
}

func (mikrotikClient mikrotikConfig) AddIpFirewallAddressList(address string, list string, comment string) (*IpFirewallAddressList, error) {
	c, err := mikrotikClient.getMikrotikClient()

	if err != nil {
		return nil, err
	}

	cmd := []string{
		"/ip/firewall/address-list/add",
		fmt.Sprintf("=address=%s", address),
		fmt.Sprintf("=list=%s", list),
		fmt.Sprintf("=comment=%s", comment),
	}

	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	r, err := c.RunArgs(cmd)

	log.Printf("[DEBUG] ip firewall address-list creation response: `%v`", r)

	if err != nil {
		return nil, err
	}

	id := r.Done.Map["ret"]

	return mikrotikClient.FindIpFirewallAddressList(id)
}

func (mikrotikClient mikrotikConfig) FindIpFirewallAddressList(id string) (*IpFirewallAddressList, error) {
	c, err := mikrotikClient.getMikrotikClient()

	if err != nil {
		return nil, err
	}

	cmd := strings.Split(fmt.Sprintf("/ip/firewall/address-list/print ?.id=%s", id), " ")
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	r, err := c.RunArgs(cmd)

	log.Printf("[DEBUG] ip firewall address-list response: %v", r)

	if err != nil {
		return nil, err
	}

	addresslist := IpFirewallAddressList{}
	err = Unmarshal(*r, &addresslist)

	if err != nil {
		return nil, err
	}

	if addresslist.Id == "" {
		return nil, NewNotFound(fmt.Sprintf("ip firewall address-list `%s`not found", id))
	}

	return &addresslist, nil
}

func (mikrotikClient mikrotikConfig) UpdateIpFirewallAddressList(id string, address string, list string, comment string) (*IpFirewallAddressList, error) {
	c, err := mikrotikClient.getMikrotikClient()

	if err != nil {
		return nil, err
	}

	cmd := []string{
		"/ip/firewall/address-list/set",
		fmt.Sprintf("=.id=%s", id),
		fmt.Sprintf("=address=%s", address),
		fmt.Sprintf("=list=%s", list),
		fmt.Sprintf("=comment=%s", comment),
	}

	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	r, err := c.RunArgs(cmd)

	log.Printf("[DEBUG] ip firewall address-list update response: `%v`", r)

	if err != nil {
		return nil, err
	}

	return mikrotikClient.FindIpFirewallAddressList(id)
}

func (mikrotikClient mikrotikConfig) DeleteIpFirewallAddressList(id string) error {
	c, err := mikrotikClient.getMikrotikClient()

	if err != nil {
		return err
	}

	cmd := []string{
		"/ip/firewall/address-list/remove",
		fmt.Sprintf("=.id=%s", id),
	}

	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	r, err := c.RunArgs(cmd)

	log.Printf("[DEBUG] ip firewall address-list delete response: `%v`", r)

	return err
}
