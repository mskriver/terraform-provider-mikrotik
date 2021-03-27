package mikrotik

import (
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceIpFirewallFilter() *schema.Resource {
	return &schema.Resource{
		Create: resourceIpFirewallFilterCreate,
		Read:   resourceIpFirewallFilterRead,
		Update: resourceIpFirewallFilterUpdate,
		Delete: resourceIpFirewallFilterDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"action": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"address_list": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"address_list_timeout": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"chain": {
				Type:     schema.TypeString,
				Required: true,
			},
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"connection_bytes": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"connection_limit": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"connection_mark": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"connection_nat_state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"connection_rate": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"connection_state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"connection_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"content": {
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
			"dscp": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"dst_address": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"dst_address_list": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"dst_address_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"dst_limit": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"dst_port": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"fragment": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"hotspot": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"icmp_options": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"in_bridge_port": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"in_bridge_port_list": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"in_interface": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"in_interface_list": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ingress_priority": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ipsec_policy": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ipv4_options": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"jump_target": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"layer7_protocol": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"log": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"log_prefix": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"nth": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"out_bridge_port": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"out_bridge_port_list": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"p2p": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"packet_mark": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"packet_size": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"per_connection_classifier": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"place_before": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"port": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"priority": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"protocol": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"psd": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"random": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"reject_with": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"routing_mark": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"routing_table": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"src_address": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"src_address_list": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"src_address_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"src_mac_address": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"src_port": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"tcp_flags": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"tcp_mss": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"tls_host": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ttl": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

type IpFirewallFilter struct {
	Id                        string `mikrotik:".id"`
	Action                    string `mikrotik:".action"`
	Address_list              string `mikrotik:".address_list"`
	Address_list_timeout      string `mikrotik:".address_list_timeout"`
	Chain                     string `mikrotik:".chain"`
	Comment                   string `mikrotik:".comment"`
	Connection_bytes          string `mikrotik:".connection_bytes"`
	Connection_limit          string `mikrotik:".connection_limit"`
	Connection_mark           string `mikrotik:".connection_mark"`
	Connection_nat_state      string `mikrotik:".connection_nat_state"`
	Connection_rate           string `mikrotik:".connection_rate"`
	Connection_state          string `mikrotik:".connection_state"`
	Connection_type           string `mikrotik:".connection_type"`
	Content                   string `mikrotik:".content"`
	Copy_from                 string `mikrotik:".copy_from"`
	Disabled                  bool   `mikrotik:".disabled"`
	Dscp                      string `mikrotik:".dscp"`
	Dst_address               string `mikrotik:".dst_address"`
	Dst_address_list          string `mikrotik:".dst_address_list"`
	Dst_address_type          string `mikrotik:".dst_address_type"`
	Dst_limit                 string `mikrotik:".dst_limit"`
	Dst_port                  string `mikrotik:".dst_port"`
	Fragment                  bool   `mikrotik:".fragment"`
	Hotspot                   string `mikrotik:".hotspot"`
	Icmp_options              string `mikrotik:".icmp_options"`
	In_bridge_port            string `mikrotik:".in_bridge_port"`
	In_bridge_port_list       string `mikrotik:".in_bridge_port_list"`
	In_interface              string `mikrotik:".in_interface"`
	In_interface_list         string `mikrotik:".in_interface_list"`
	Ingress_priority          string `mikrotik:".ingress_priority"`
	Ipsec_policy              string `mikrotik:".ipsec_policy"`
	Ipv4_options              string `mikrotik:".ipv4_options"`
	Jump_target               string `mikrotik:".jump_target"`
	Layer7_protocol           string `mikrotik:".layer7_protocol"`
	Limit                     string `mikrotik:".limit"`
	Log                       bool   `mikrotik:".log"`
	Log_prefix                string `mikrotik:".log_prefix"`
	Nth                       string `mikrotik:".nht"`
	Out_bridge_port           string `mikrotik:".out_bridge_port"`
	Out_bridge_port_list      string `mikrotik:".out_bridge_port_list"`
	P2p                       string `mikrotik:".p2p"`
	Packet_mark               string `mikrotik:".packet_mark"`
	Packet_size               string `mikrotik:".packet_size"`
	Per_connection_classifier string `mikrotik:".per_connection_classifier"`
	Place_before              string `mikrotik:".place_before"`
	Port                      string `mikrotik:".port"`
	Priority                  string `mikrotik:".priority"`
	Protocol                  string `mikrotik:".protocol"`
	Psd                       string `mikrotik:".psd"`
	Random                    string `mikrotik:".random"`
	Reject_with               string `mikrotik:".reject_with"`
	Routing_mark              string `mikrotik:".routing_mark"`
	Routing_table             string `mikrotik:".routing_table"`
	Src_address               string `mikrotik:".src_address"`
	Src_address_list          string `mikrotik:".src_address_list"`
	Src_address_type          string `mikrotik:".src_address_type"`
	Src_mac_address           string `mikrotik:".src_mac_address"`
	Src_port                  string `mikrotik:".src_port"`
	Tcp_flags                 string `mikrotik:".tcp_flags"`
	Tcp_mss                   string `mikrotik:".tcp_mss"`
	Time                      string `mikrotik:".time"`
	Tls_host                  string `mikrotik:".tls_host"`
	Ttl                       string `mikrotik:".tls"`
}

func resourceIpFirewallFilterCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(mikrotikConfig)

	filterlist, err := c.AddIpFirewallFilter(d)

	if err != nil {
		return err
	}

	writeStateIpFirewallFilter(filterlist, d)
	return nil
}

func resourceIpFirewallFilterRead(d *schema.ResourceData, m interface{}) error {
	c := m.(mikrotikConfig)

	filterlist, err := c.FindIpFirewallFilter(d.Id())

	if err != nil {
		d.SetId("")
		return nil
	}

	if filterlist == nil {
		d.SetId("")
		return nil
	}

	writeStateIpFirewallFilter(filterlist, d)
	return nil
}

func resourceIpFirewallFilterUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(mikrotikConfig)

	filterlist, err := c.UpdateIpFirewallFilter(d.Id(), d)

	if err != nil {
		return err
	}

	writeStateIpFirewallFilter(filterlist, d)
	return nil
}

func resourceIpFirewallFilterDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(mikrotikConfig)

	err := c.DeleteIpFirewallFilter(d.Id())

	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}

func (mikrotikClient mikrotikConfig) AddIpFirewallFilter(d *schema.ResourceData) (*IpFirewallFilter, error) {
	c, err := mikrotikClient.getMikrotikClient()

	if err != nil {
		return nil, err
	}

	cmd_attributes := FormatIpFirewallFilterCommand(d)

	cmd := []string{
		"/ip/firewall/filter/add",
	}
	cmd = append(cmd, cmd_attributes...)

	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	r, err := c.RunArgs(cmd)

	log.Printf("[DEBUG] ip firewall filter creation response: `%v`", r)

	if err != nil {
		return nil, err
	}

	id := r.Done.Map["ret"]

	return mikrotikClient.FindIpFirewallFilter(id)
}

func (mikrotikClient mikrotikConfig) UpdateIpFirewallFilter(id string, d *schema.ResourceData) (*IpFirewallFilter, error) {
	c, err := mikrotikClient.getMikrotikClient()

	if err != nil {
		return nil, err
	}

	cmd_attributes := FormatIpFirewallFilterCommand(d)

	cmd := []string{
		"/ip/firewall/filter/set",
		fmt.Sprintf("=.id=" + id),
	}
	cmd = append(cmd, cmd_attributes...)

	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	r, err := c.RunArgs(cmd)

	log.Printf("[DEBUG] ip firewall filter update response: `%v`", r)

	if err != nil {
		return nil, err
	}

	return mikrotikClient.FindIpFirewallFilter(id)
}

func (mikrotikClient mikrotikConfig) DeleteIpFirewallFilter(id string) error {
	c, err := mikrotikClient.getMikrotikClient()

	if err != nil {
		return err
	}

	cmd := []string{
		"/ip/firewall/filter/remove",
		fmt.Sprintf("=.id=" + id),
	}

	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	r, err := c.RunArgs(cmd)

	log.Printf("[DEBUG] ip firewall filter delete response: `%v`", r)

	return err
}

func FormatIpFirewallFilterCommand(d *schema.ResourceData) []string {
	var cmd_string []string

	if d.Get("action") != "" {
		cmd_string = append(cmd_string, "=action="+d.Get("action").(string))
	}
	if d.Get("address_list") != "" {
		cmd_string = append(cmd_string, "=address-list="+d.Get("address_list").(string))
	}
	if d.Get("address_list_timeout") != "" {
		cmd_string = append(cmd_string, "=address-list-timeout="+d.Get("address_list_timeout").(string))
	}
	if d.Get("chain") != "" {
		cmd_string = append(cmd_string, "=chain="+d.Get("chain").(string))
	}
	if d.Get("comment") != "" {
		cmd_string = append(cmd_string, "=comment=\""+d.Get("comment").(string)+"\"")
	}
	if d.Get("connection_bytes") != "" {
		cmd_string = append(cmd_string, "=connection-bytes="+d.Get("connection_bytes").(string))
	}
	if d.Get("connection_limit") != "" {
		cmd_string = append(cmd_string, "=connection-limit="+d.Get("connection_limit").(string))
	}
	if d.Get("connection_mark") != "" {
		cmd_string = append(cmd_string, "=connection-mark="+d.Get("connection_mark").(string))
	}
	if d.Get("connection_nat_state") != "" {
		cmd_string = append(cmd_string, "=connection-nat-state="+d.Get("connection_nat_state").(string))
	}
	if d.Get("connection_rate") != "" {
		cmd_string = append(cmd_string, "=connection-rate="+d.Get("connection_rate").(string))
	}
	if d.Get("connection_state") != "" {
		cmd_string = append(cmd_string, "=connection-state="+d.Get("connection_state").(string))
	}
	if d.Get("connection_type") != "" {
		cmd_string = append(cmd_string, "=connection-type="+d.Get("connection_type").(string))
	}
	if d.Get("content") != "" {
		cmd_string = append(cmd_string, "=content="+d.Get("content").(string))
	}
	if d.Get("copy_from") != "" {
		cmd_string = append(cmd_string, "=copy-from="+d.Get("copy_from").(string))
	}
	if d.Get("disabled") != "" {
		cmd_string = append(cmd_string, "=disabled="+boolToMikrotikBool(d.Get("disabled").(bool)))
	}
	if d.Get("dscp") != "" {
		cmd_string = append(cmd_string, "=dscp="+d.Get("dscp").(string))
	}
	if d.Get("dst_address") != "" {
		cmd_string = append(cmd_string, "=dst-address="+d.Get("dst_address").(string))
	}
	if d.Get("dst_address_list") != "" {
		cmd_string = append(cmd_string, "=dst-address-list="+d.Get("dst_address_list").(string))
	}
	if d.Get("dst_address_type") != "" {
		cmd_string = append(cmd_string, "=dst-address-type="+d.Get("dst_address_type").(string))
	}
	if d.Get("dst_limit") != "" {
		cmd_string = append(cmd_string, "=dst-limit="+d.Get("dst_limit").(string))
	}
	if d.Get("dst_port") != "" {
		cmd_string = append(cmd_string, "=dst-port="+d.Get("dst_port").(string))
	}
	if d.Get("fragment") != "" {
		cmd_string = append(cmd_string, "=fragment="+boolToMikrotikBool(d.Get("fragment").(bool)))
	}
	if d.Get("hotspot") != "" {
		cmd_string = append(cmd_string, "=hotspot="+d.Get("hotspot").(string))
	}
	if d.Get("icmp_options") != "" {
		cmd_string = append(cmd_string, "=icmp-options="+d.Get("icmp_options").(string))
	}
	if d.Get("in_bridge_port") != "" {
		cmd_string = append(cmd_string, "=in-bridge-port="+d.Get("in_bridge_port").(string))
	}
	if d.Get("in_bridge_port_list") != "" {
		cmd_string = append(cmd_string, "=in-bridge-port-list="+d.Get("in_bridge_port_list").(string))
	}
	if d.Get("in_interface") != "" {
		cmd_string = append(cmd_string, "=in-interface="+d.Get("in_interface").(string))
	}
	if d.Get("in_interface_list") != "" {
		cmd_string = append(cmd_string, "=in-interface-list="+d.Get("in_interface_list").(string))
	}
	if d.Get("ingress_priority") != "" {
		cmd_string = append(cmd_string, "=ingress-priority="+d.Get("ingress_priority").(string))
	}
	if d.Get("ipsec_policy") != "" {
		cmd_string = append(cmd_string, "=ipsec-policy="+d.Get("ipsec_policy").(string))
	}
	if d.Get("ipv4_options") != "" {
		cmd_string = append(cmd_string, "=ipv4-options="+d.Get("ipv4_options").(string))
	}
	if d.Get("jump_target") != "" {
		cmd_string = append(cmd_string, "=jump-target="+d.Get("jump_target").(string))
	}
	if d.Get("layer7_protocol") != "" {
		cmd_string = append(cmd_string, "=layer7-protocol="+d.Get("layer7_protocol").(string))
	}
	if d.Get("limit") != "" {
		cmd_string = append(cmd_string, "=limit="+d.Get("limit").(string))
	}
	if d.Get("log") != "" {
		cmd_string = append(cmd_string, "=log="+boolToMikrotikBool(d.Get("log").(bool)))
	}
	if d.Get("log_prefix") != "" {
		cmd_string = append(cmd_string, "=log-prefix="+d.Get("log_prefix").(string))
	}
	if d.Get("nth") != "" {
		cmd_string = append(cmd_string, "=nth="+d.Get("nth").(string))
	}
	if d.Get("out_bridge_port") != "" {
		cmd_string = append(cmd_string, "=out-bridge-port="+d.Get("out_bridge_port").(string))
	}
	if d.Get("out_bridge_port_list") != "" {
		cmd_string = append(cmd_string, "=out-bridge-port-list="+d.Get("out_bridge_port_list").(string))
	}
	if d.Get("p2p") != "" {
		cmd_string = append(cmd_string, "=p2p="+d.Get("p2p").(string))
	}
	if d.Get("packet_mark") != "" {
		cmd_string = append(cmd_string, "=packet-mark="+d.Get("packet_mark").(string))
	}
	if d.Get("packet_size") != "" {
		cmd_string = append(cmd_string, "=packet-size="+d.Get("packet_size").(string))
	}
	if d.Get("per_connection_classifier") != "" {
		cmd_string = append(cmd_string, "=per-connection-classifier="+d.Get("per_connection_classifier").(string))
	}
	if d.Get("place_before") != "" {
		cmd_string = append(cmd_string, "=plase-before="+d.Get("place_before").(string))
	}
	if d.Get("port") != "" {
		cmd_string = append(cmd_string, "=port="+d.Get("port").(string))
	}
	if d.Get("priority") != "" {
		cmd_string = append(cmd_string, "=priority="+d.Get("priority").(string))
	}
	if d.Get("protocol") != "" {
		cmd_string = append(cmd_string, "=protocol="+d.Get("protocol").(string))
	}
	if d.Get("psd") != "" {
		cmd_string = append(cmd_string, "=psd="+d.Get("psd").(string))
	}
	if d.Get("random") != "" {
		cmd_string = append(cmd_string, "=random="+d.Get("random").(string))
	}
	if d.Get("reject_with") != "" {
		cmd_string = append(cmd_string, "=reject-with="+d.Get("reject_with").(string))
	}
	if d.Get("routing_mark") != "" {
		cmd_string = append(cmd_string, "=routing-mark="+d.Get("routing_mark").(string))
	}
	if d.Get("routing_table") != "" {
		cmd_string = append(cmd_string, "=routing-table="+d.Get("routing_table").(string))
	}
	if d.Get("src_address") != "" {
		cmd_string = append(cmd_string, "=src-address="+d.Get("src_address").(string))
	}
	if d.Get("src_address_list") != "" {
		cmd_string = append(cmd_string, "=src-address-list="+d.Get("src_address_list").(string))
	}
	if d.Get("src_address_type") != "" {
		cmd_string = append(cmd_string, "=src-address-type="+d.Get("src_address_type").(string))
	}
	if d.Get("src_mac_address") != "" {
		cmd_string = append(cmd_string, "=src-mac-address="+d.Get("src_mac_address").(string))
	}
	if d.Get("src_port") != "" {
		cmd_string = append(cmd_string, "=src-port="+d.Get("src_port").(string))
	}
	if d.Get("tcp_flags") != "" {
		cmd_string = append(cmd_string, "=tcp-flags="+d.Get("tcp_flags").(string))
	}
	if d.Get("tcp_mss") != "" {
		cmd_string = append(cmd_string, "=tcp-mss="+d.Get("tcp_mss").(string))
	}
	if d.Get("time") != "" {
		cmd_string = append(cmd_string, "=time="+d.Get("time").(string))
	}
	if d.Get("tls_host") != "" {
		cmd_string = append(cmd_string, "=tls-host="+d.Get("tls_host").(string))
	}
	if d.Get("ttl") != "" {
		cmd_string = append(cmd_string, "=ttl="+d.Get("ttl").(string))
	}

	return cmd_string
}

func (mikrotikClient mikrotikConfig) FindIpFirewallFilter(id string) (*IpFirewallFilter, error) {
	c, err := mikrotikClient.getMikrotikClient()

	if err != nil {
		return nil, err
	}

	cmd := strings.Split(fmt.Sprintf("/ip/firewall/filter/print ?.id="+id), " ")
	log.Printf("[INFO] Running the mikrotik command: `%s`", cmd)
	r, err := c.RunArgs(cmd)

	log.Printf("[DEBUG] ip firewall filter response: %v", r)

	if err != nil {
		return nil, err
	}

	filterlist := IpFirewallFilter{}
	err = Unmarshal(*r, &filterlist)

	if err != nil {
		return nil, err
	}

	if filterlist.Id == "" {
		return nil, NewNotFound(fmt.Sprintf("ip firewall filter `%s`not found", id))
	}

	return &filterlist, nil
}

func writeStateIpFirewallFilter(filterlist *IpFirewallFilter, d *schema.ResourceData) error {
	d.SetId(filterlist.Id)
	d.Set("Action", filterlist.Action)
	d.Set("Address_list", filterlist.Address_list)
	d.Set("Address_list_timeout", filterlist.Address_list_timeout)
	d.Set("Chain", filterlist.Chain)
	d.Set("Comment", filterlist.Comment)
	d.Set("Connection_bytes", filterlist.Connection_bytes)
	d.Set("Connection_limit", filterlist.Connection_limit)
	d.Set("Connection_mark", filterlist.Connection_mark)
	d.Set("Connection_nat_state", filterlist.Connection_nat_state)
	d.Set("Connection_rate", filterlist.Connection_rate)
	d.Set("Connection_state", filterlist.Connection_state)
	d.Set("Connection_type", filterlist.Connection_type)
	d.Set("Content", filterlist.Content)
	d.Set("Copy_from", filterlist.Copy_from)
	d.Set("Disabled", filterlist.Disabled)
	d.Set("Dscp", filterlist.Dscp)
	d.Set("Dst_address", filterlist.Dst_address)
	d.Set("Dst_address_list", filterlist.Dst_address_list)
	d.Set("Dst_address_type", filterlist.Dst_address_type)
	d.Set("Dst_limit", filterlist.Dst_limit)
	d.Set("Dst_port", filterlist.Dst_port)
	d.Set("Fragment", filterlist.Fragment)
	d.Set("Hotspot", filterlist.Hotspot)
	d.Set("Icmp_options", filterlist.Icmp_options)
	d.Set("In_bridge_port", filterlist.In_bridge_port)
	d.Set("In_bridge_port_list", filterlist.In_bridge_port_list)
	d.Set("In_interface", filterlist.In_interface)
	d.Set("In_interface_list", filterlist.In_interface_list)
	d.Set("Ingress_priority", filterlist.Ingress_priority)
	d.Set("Ipsec_policy", filterlist.Ipsec_policy)
	d.Set("Ipv4_options", filterlist.Ipv4_options)
	d.Set("Jump_target", filterlist.Jump_target)
	d.Set("Layer7_protocol", filterlist.Layer7_protocol)
	d.Set("Limit", filterlist.Limit)
	d.Set("Log", filterlist.Log)
	d.Set("Log_prefix", filterlist.Log_prefix)
	d.Set("Nth", filterlist.Nth)
	d.Set("Out_bridge_port", filterlist.Out_bridge_port)
	d.Set("Out_bridge_port_list", filterlist.Out_bridge_port_list)
	d.Set("P2p", filterlist.P2p)
	d.Set("Packet_mark", filterlist.Packet_mark)
	d.Set("Packet_size", filterlist.Packet_size)
	d.Set("Per_connection_classifier", filterlist.Per_connection_classifier)
	d.Set("Place_before", filterlist.Place_before)
	d.Set("Port", filterlist.Port)
	d.Set("Priority", filterlist.Priority)
	d.Set("Protocol", filterlist.Protocol)
	d.Set("Psd", filterlist.Psd)
	d.Set("Random", filterlist.Random)
	d.Set("Reject_with", filterlist.Reject_with)
	d.Set("Routing_mark", filterlist.Routing_mark)
	d.Set("Routing_table", filterlist.Routing_table)
	d.Set("Src_address", filterlist.Src_address)
	d.Set("Src_address_list", filterlist.Src_address_list)
	d.Set("Src_address_type", filterlist.Src_address_type)
	d.Set("Src_mac_address", filterlist.Src_mac_address)
	d.Set("Src_port", filterlist.Src_port)
	d.Set("Tcp_flags", filterlist.Tcp_flags)
	d.Set("Tcp_mss", filterlist.Tcp_mss)
	d.Set("Time", filterlist.Time)
	d.Set("Tls_host", filterlist.Tls_host)
	d.Set("Ttl", filterlist.Ttl)

	return nil
}
