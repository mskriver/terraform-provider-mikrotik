# mikrotik_ip_address

Creates a IPv4 interface on the mikrotik device

## Example Usage

```hcl
resource "mikrotik_ip_address" "network1" {
  address = "192.168.1.1/24"
  interface = "ether1"
}
```

## Argument Reference

* address - (Required) The IP address of the interface to be created
* interface - (Required) Interface name of the interface on which the IP address will be configured

## Attributes Reference

## Import Reference

```bash
terraform import mikrotik_ip_address.network1 *19
```

Last argument (*19) is a mikrotik internal id which can be obtained via CLI:

```bash
[admin@MikroTik] /ip address> :put [find where address="192.168.1.1/24"]
*19
```
