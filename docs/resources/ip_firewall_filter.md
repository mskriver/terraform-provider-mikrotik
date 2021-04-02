# mikrotik_ip_firewall_filter

Creates a IPv4 firewall filter entry on the mikrotik device

## Example Usage

```hcl
resource "mikrotik_ip_firewall_filter" "fw_filter1" {
  chain = "chain1"
  src_address = "1.1.1.0/24"
  dst_address = "2.2.2.2"
  action = "accept"
  comment = "Sample filter"
}```

## Argument Reference

* action - (Optional) - Action for the firewall filter entry
* address_list - (Optional)
* address_list_timeout - (Optional)
* chain - (Required) - Chain which the filter entry will belong to
* comment - (Optional) Comment/description for the filter entry
* connection_bytes - (Optional)
* connection_limit - (Optional)
* connection_mark - (Optional)
* connection_nat_state - (Optional)
* connection_rate - (Optional)
* connection_state - (Optional)
* connection_type - (Optional)
* content - (Optional)
* copy_from - (Optional)
* disabled (Optional, defaults to false)
* dscp - (Optional)
* dst_address - (Optional)
* dst_address_list - (Optional)
* dst_address_type - (Optional)
* dst_limit - (Optional)
* dst_port - (Optional)
* fragment - (Optional)
* hotspot - (Optional)
* icmp_options - (Optional)
* in_bridge_port - (Optional)
* in_bridge_port_list - (Optional)
* in_interface - (Optional)
* in_interface_list - (Optional)
* ingress_priority - (Optional)
* ipsec_policy - (Optional)
* ipv4_options - (Optional)
* jump_target - (Optional)
* layer7_protocol - (Optional)
* limit - (Optional)
* log - (Optional)
* log_prefix - (Optional)
* nth - (Optional)
* out_bridge_port - (Optional)
* out_bridge_port_list - (Optional)
* p2p - (Optional)
* packet_mark - (Optional)
* packet_size" - (Optional)
* per_connection_classifier - (Optional)
* place_before - (Optional)
* port - (Optional)
priority - (Optional)
protocol - (Optional)
psd - (Optional)
random - (Optional)
reject_with - (Optional)
routing_mark - (Optional)
routing_table - (Optional)
src_address - (Optional)
src_address_list - (Optional)
src_address_type - (Optional)
src_mac_address - (Optional)
src_port - (Optional)
tcp_flags - (Optional)
tcp_mss - (Optional)
time - (Optional)
tls_host - (Optional)
ttl - (Optional)

## Attributes Reference

https://help.mikrotik.com/docs/display/ROS/Filter

## Import Reference

```bash
terraform import mikrotik_ip_firewall_filter.fw_filter1 *3f
```

Last argument (*d) is a mikrotik internal id which can be obtained via CLI:

```bash
[admin@MikroTik] /ip firewall filter> :put [find where dst-address="2.2.2.2"]
*3f
```
