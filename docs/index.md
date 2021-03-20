# Mikrotik Provider

The mikrotik provider is used to interact with the resources supported by RouterOS.
The provider needs to be configured with the proper credentials before it can be used.

Please note that this provider uses the RouterOS API where credentials are sent in clear text.
More information about the RouterOS API can be found at <https://wiki.mikrotik.com/wiki/Manual:API>

This provider is inspired by <https://github.com/ddelnano/terraform-provider-mikrotik>

## Requirements

* RouterOS v6.45.2+ (It may work with other versions but it is untested against other versions!)

## Using the provider

This provider is on the terraform registry so you only need to reference it in your terraform code (example below).

## Example Usage

```hcl
# Configure the mikrotik Provider
provider "mikrotik" {
  host = "hostname-of-server:8728"     # Or set MIKROTIK_HOST environment variable
  username = "<username>"              # Or set MIKROTIK_USER environment variable
  password = "<password>"              # Or set MIKROTIK_PASSWORD environment variable
}
```
