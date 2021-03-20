# mikrotik_ip_address

Creates a IPv4 ip firewall address-list entry on the mikrotik device

## Example Usage

```hcl
resource "mikrotik_ip_firewall_address-list" "list_entry_1" {
  address = "1.1.1.1"
  list = "list1"
  comment = "test"
}
```

## Argument Reference

* address - (Required) The IP address or subnet of the entry to be created
* list - (Required) Name of the list address-list which the address/subnet should be added to
* comment - (Optional) Comment/description for the address-list entry

## Attributes Reference

## Import Reference

```bash
terraform import mikrotik_ip_address.list_entry_1 *d
```

Last argument (*d) is a mikrotik internal id which can be obtained via CLI:

```bash
[admin@MikroTik] /ip firewall address-list> :put [find where address="1.1.1.1"]
*d
```
