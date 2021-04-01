package mikrotik

import (
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceInterfaceGre() *schema.Resource {
	return &schema.Resource{
		Create: resourceInterfaceGreCreate,
		Read:   resourceInterfaceGreRead,
		Update: resourceInterfaceGreUpdate,
		Delete: resourceInterfaceGreDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"allow_fast_path": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"clamp_tcp_mss": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"copy_from": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"disabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"dont_fragment": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "no",
			},
			"dscp": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "inherit",
			},
			"ipsec_secret": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"keepalive": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "10s,10",
			},
			"local_address": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"mtu": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "auto",
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"remote_address": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

type InterfaceGre struct {
	Id              string `mikrotik:".id"`
	Allow_fast_path bool   `mikrotik:".allow-fast-path"`
	Clamp_tcp_mss   bool   `mikrotik:".clamp-tcp-mss"`
	Comment         string `mikrotik:".comment"`
	Copy_from       string `mikrotik:".copy-from"`
	Disabled        bool   `mikrotik:".disabled"`
	Dont_fragment   string `mikrotik:".dont-fragment"`
	Dscp            string `mikrotik:".dscp"`
	Ipsec_secret    string `mikrotik:".ipsec-secret"`
	Keepalive       string `mikrotik:".keepalive"`
	Local_address   string `mikrotik:".local-address"`
	Mtu             string `mikrotik:".mtu"`
	Name            string `mikrotik:".name"`
	Remote_address  string `mikrotik:".remote-address"`
}

func resourceInterfaceGreCreate(d *schema.ResourceData, m interface{}) error {
	allow_fast_path := d.Get("allow_fast_path").(bool)
	clamp_tcp_mss := d.Get("clamp_tcp_mss").(bool)
	comment := d.Get("comment").(string)
	copy_from := d.Get("copy_from").(string)
	disabled := d.Get("disabled").(bool)
	dont_fragment := d.Get("dont_fragment").(string)
	dscp := d.Get("dscp").(string)
	ipsec_secret := d.Get("ipsec_secret").(string)
	keepalive := d.Get("keepalive").(string)
	local_address := d.Get("local_address").(string)
	mtu := d.Get("mtu").(string)
	name := d.Get("name").(string)
	remote_address := d.Get("remote_address").(string)

	c := m.(mikrotikConfig)

	greif, err := c.AddInterfaceGre(allow_fast_path, clamp_tcp_mss, comment, copy_from, disabled, dont_fragment, dscp, ipsec_secret, keepalive, local_address, mtu, name, remote_address)

	if err != nil {
		return err
	}

	writeStateInterfaceGre(greif, d)
	return nil
}

func resourceInterfaceGreRead(d *schema.ResourceData, m interface{}) error {
	c := m.(mikrotikConfig)

	greif, err := c.FindInterfaceGre(d.Id())

	if err != nil {
		d.SetId("")
		return nil
	}

	if greif == nil {
		d.SetId("")
		return nil
	}

	writeStateInterfaceGre(greif, d)
	return nil
}

func resourceInterfaceGreUpdate(d *schema.ResourceData, m interface{}) error {
	allow_fast_path := d.Get("allow_fast_path").(bool)
	clamp_tcp_mss := d.Get("clamp_tcp_mss").(bool)
	comment := d.Get("comment").(string)
	copy_from := d.Get("copy_from").(string)
	disabled := d.Get("disabled").(bool)
	dont_fragment := d.Get("dont_fragment").(string)
	dscp := d.Get("dscp").(string)
	ipsec_secret := d.Get("ipsec_secret").(string)
	keepalive := d.Get("keepalive").(string)
	local_address := d.Get("local_address").(string)
	mtu := d.Get("mtu").(string)
	name := d.Get("name").(string)
	remote_address := d.Get("remote_address").(string)

	c := m.(mikrotikConfig)

	greif, err := c.UpdateInterfaceGre(d.Id(), allow_fast_path, clamp_tcp_mss, comment, copy_from, disabled, dont_fragment, dscp, ipsec_secret, keepalive, local_address, mtu, name, remote_address)

	if err != nil {
		return err
	}

	writeStateInterfaceGre(greif, d)
	return nil
}

func resourceInterfaceGreDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(mikrotikConfig)

	err := c.DeleteInterfaceGre(d.Id())

	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}

func (mikrotikClient mikrotikConfig) AddInterfaceGre(allow_fast_path bool, clamp_tcp_mss bool, comment string, copy_from string, disabled bool, dont_fragment string, dscp string, ipsec_secret string, keepalive string, local_address string, mtu string, name string, remote_address string) (*InterfaceGre, error) {
	c, err := mikrotikClient.getMikrotikClient()

	if err != nil {
		return nil, err
	}

	cmd := []string{
		"/interface/gre/add",
	}
	cmd = append(cmd, "=allow-fast-path="+boolToMikrotikBool(allow_fast_path))
	cmd = append(cmd, "=clamp-tcp-mss="+boolToMikrotikBool(clamp_tcp_mss))
	if comment != "" {
		cmd = append(cmd, "=comment="+comment)
	}
	if copy_from != "" {
		cmd = append(cmd, "=copy-from="+copy_from)
	}
	cmd = append(cmd, "=disabled="+boolToMikrotikBool(disabled))
	cmd = append(cmd, "=dont-fragment="+dont_fragment)
	if dscp != "" {
		cmd = append(cmd, "=dscp="+dscp)
	}
	if ipsec_secret != "" {
		cmd = append(cmd, "=ipsec-secret="+ipsec_secret)
	}
	if keepalive != "" {
		cmd = append(cmd, "=keepalive="+keepalive)
	}
	if local_address != "" {
		cmd = append(cmd, "=local-address="+local_address)
	}
	if mtu != "" {
		cmd = append(cmd, "=mtu="+mtu)
	}
	if name != "" {
		cmd = append(cmd, "=name="+name)
	}
	if remote_address != "" {
		cmd = append(cmd, "=remote-address="+remote_address)
	}

	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	r, err := c.RunArgs(cmd)

	log.Printf("[DEBUG] gre interface creation response: `%v`", r)

	if err != nil {
		return nil, err
	}

	id := r.Done.Map["ret"]

	return mikrotikClient.FindInterfaceGre(id)
}

func (mikrotikClient mikrotikConfig) UpdateInterfaceGre(id string, allow_fast_path bool, clamp_tcp_mss bool, comment string, copy_from string, disabled bool, dont_fragment string, dscp string, ipsec_secret string, keepalive string, local_address string, mtu string, name string, remote_address string) (*InterfaceGre, error) {
	c, err := mikrotikClient.getMikrotikClient()

	if err != nil {
		return nil, err
	}

	cmd := []string{
		"/interface/gre/set",
		fmt.Sprintf("=.id=%s", id),
	}
	cmd = append(cmd, "=allow-fast-path="+boolToMikrotikBool(allow_fast_path))
	cmd = append(cmd, "=clamp-tcp-mss="+boolToMikrotikBool(clamp_tcp_mss))
	if comment != "" {
		cmd = append(cmd, "=comment="+comment)
	}
	if copy_from != "" {
		cmd = append(cmd, "=copy-from="+copy_from)
	}
	cmd = append(cmd, "=disabled="+boolToMikrotikBool(disabled))
	cmd = append(cmd, "=dont-fragment="+dont_fragment)
	if dscp != "" {
		cmd = append(cmd, "=dscp="+dscp)
	}
	if ipsec_secret != "" {
		cmd = append(cmd, "=ipsec-secret="+ipsec_secret)
	}
	if keepalive != "" {
		cmd = append(cmd, "=keepalive="+keepalive)
	}
	if local_address != "" {
		cmd = append(cmd, "=local-address="+local_address)
	}
	if mtu != "" {
		cmd = append(cmd, "=mtu="+mtu)
	}
	if name != "" {
		cmd = append(cmd, "=name="+name)
	}
	if remote_address != "" {
		cmd = append(cmd, "=remote-address="+remote_address)
	}

	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	r, err := c.RunArgs(cmd)

	log.Printf("[DEBUG] gre interface update response: `%v`", r)

	if err != nil {
		return nil, err
	}

	return mikrotikClient.FindInterfaceGre(id)
}

func (mikrotikClient mikrotikConfig) DeleteInterfaceGre(id string) error {
	c, err := mikrotikClient.getMikrotikClient()

	if err != nil {
		return err
	}

	cmd := []string{
		"/interface/gre/remove",
		fmt.Sprintf("=.id=%s", id),
	}

	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	r, err := c.RunArgs(cmd)

	log.Printf("[DEBUG] gre interface delete response: `%v`", r)

	return err
}

func (mikrotikClient mikrotikConfig) FindInterfaceGre(id string) (*InterfaceGre, error) {
	c, err := mikrotikClient.getMikrotikClient()

	if err != nil {
		return nil, err
	}

	cmd := strings.Split(fmt.Sprintf("/interface/gre/print ?.id=%s", id), " ")
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	r, err := c.RunArgs(cmd)

	log.Printf("[DEBUG] gre interface response: %v", r)

	if err != nil {
		return nil, err
	}

	greif := InterfaceGre{}
	err = Unmarshal(*r, &greif)

	if err != nil {
		return nil, err
	}

	if greif.Id == "" {
		return nil, NewNotFound(fmt.Sprintf("gre interface `%s`not found", id))
	}

	return &greif, nil
}

func writeStateInterfaceGre(greif *InterfaceGre, d *schema.ResourceData) error {
	d.SetId(greif.Id)
	d.Set("Allow_fast_path", greif.Allow_fast_path)
	d.Set("Clamp_tcp_mss", greif.Clamp_tcp_mss)
	d.Set("Comment", greif.Comment)
	d.Set("Copy_from", greif.Copy_from)
	d.Set("Disabled", greif.Disabled)
	d.Set("Dont_fragment", greif.Dont_fragment)
	d.Set("Dscp", greif.Dscp)
	d.Set("Ipsec_secret", greif.Ipsec_secret)
	d.Set("Keepalive", greif.Keepalive)
	d.Set("Local_address", greif.Local_address)
	d.Set("Mtu", greif.Mtu)
	d.Set("Name", greif.Name)
	d.Set("Remote_address", greif.Remote_address)
	return nil
}
