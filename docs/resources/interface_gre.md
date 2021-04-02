# mikrotik_ip_address

Creates a gre interface on the mikrotik device

## Example Usage

```hcl
resource "mikrotik_interface_gre" "gre1" {
  remote_address = "3.3.3.3"
  name = "gre1" 
}```

## Argument Reference

* allow_fast_path - (Optional, defaults to true)
* clamp_tcp_mss - (Optional, defaults to true)
* comment - (Optional) Comment/description for the interface
* copy_from - (Optional)
* disabled (Optional, defaults to false)
* dont_fragment - (Optional, defaults to no)
* dscp - (Optional, defaults to inherit)
* ipsec_secret - (Optional)
* keepalive - (Optional, defaults to 10s,10)
* local_address - (Optional)
* mtu - (Optional, defaults to auto)
* name - (Optional)
* remote_address - (Required)

## Attributes Reference

https://help.mikrotik.com/docs/display/ROS/GRE

## Import Reference

```bash
terraform import mikrotik_interface_gre.gre1 *12
```

Last argument (*d) is a mikrotik internal id which can be obtained via CLI:

```bash
[admin@MikroTik] /interface gre> :put [find where remote-address=2.2.2.2"]
*12
```
