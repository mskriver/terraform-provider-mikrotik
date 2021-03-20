package mikrotik

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var originalIpAddress string = "192.168.1.1/24"
var originalInterface string = "ether1"
var updatedIpAddress string = "192.168.2.1/24"
var updatedInterface string = "ether2"

func TestAccMikrotikIpAddress_create(t *testing.T) {
	resourceName := "mikrotik_ip_address.autotest"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMikrotikIpAddressDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccIpAddress(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccIpAddressExists(resourceName),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "address", originalIpAddress),
					resource.TestCheckResourceAttr(resourceName, "interface", originalInterface),
				),
			},
		},
	})
}

func TestAccMikrotikIpAddress_update(t *testing.T) {
	resourceName := "mikrotik_ip_address.autotest"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMikrotikIpAddressDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccIpAddress(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccIpAddressExists(resourceName),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "address", originalIpAddress),
					resource.TestCheckResourceAttr(resourceName, "interface", originalInterface),
				),
			},
			{
				Config: testAccIpAddressUpdateAddress(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccIpAddressExists(resourceName),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "address", updatedIpAddress),
					resource.TestCheckResourceAttr(resourceName, "interface", originalInterface),
				),
			},
			{
				Config: testAccIpAddressUpdateInterface(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccIpAddressExists(resourceName),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "address", updatedIpAddress),
					resource.TestCheckResourceAttr(resourceName, "interface", updatedInterface),
				),
			},
		},
	})
}

func testAccIpAddressExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("mikrotik_ip_address does not exist in the statefile")
		}

		c := NewClient(GetConfigFromEnv())

		ipaddr, err := c.FindIpAddress(rs.Primary.ID)

		if err != nil {
			return fmt.Errorf("Unable to get the ip address with error: %v", err)
		}

		if ipaddr == nil {
			return fmt.Errorf("Unable to get the ip address")
		}

		if ipaddr.Id == rs.Primary.ID {
			return nil
		}
		return nil
	}
}

func testAccIpAddress() string {
	return fmt.Sprintf(`
resource "mikrotik_ip_address" "autotest" {
	address = "%s"
	interface = "%s"
}
`, originalIpAddress, originalInterface)
}

func testAccIpAddressUpdateAddress() string {
	return fmt.Sprintf(`
	resource "mikrotik_ip_address" "autotest" {
		address = "%s"
		interface = "%s"
	}
`, updatedIpAddress, originalInterface)
}

func testAccIpAddressUpdateInterface() string {
	return fmt.Sprintf(`
	resource "mikrotik_ip_address" "autotest" {
		address = "%s"
		interface = "%s"
	}
`, updatedIpAddress, updatedInterface)
}

func testAccCheckMikrotikIpAddressDestroy(s *terraform.State) error {
	c := NewClient(GetConfigFromEnv())
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "mikrotik_ip_address" {
			continue
		}

		ipaddr, err := c.FindIpAddress(rs.Primary.ID)

		_, ok := err.(*NotFound)
		if !ok && err != nil {
			return err
		}

		if ipaddr != nil {
			return fmt.Errorf("ip address (%s) still exists", ipaddr.Id)
		}
	}
	return nil
}

func TestAddIpAddressAndDeleteIpAddress(t *testing.T) {
	c := NewClient(GetConfigFromEnv())

	address := "1.1.1.1/24"
	ifname := "ether1"
	network := "1.1.1.0"
	expectedIpAddress := &IpAddress{
		Address:   address,
		Interface: ifname,
		Network:   network,
	}
	ipaddr, err := c.AddIpAddress(
		address,
		ifname,
	)

	if err != nil {
		t.Errorf("Error creating a ip address with: %v", err)
	}

	if len(ipaddr.Id) < 1 {
		t.Errorf("The created ip address does not have an Id: %v", ipaddr)
	}

	if strings.Compare(ipaddr.Address, expectedIpAddress.Address) != 0 {
		t.Errorf("The address field do not match. actual: %v expected: %v", ipaddr.Address, expectedIpAddress.Address)
	}

	if strings.Compare(ipaddr.Interface, expectedIpAddress.Interface) != 0 {
		t.Errorf("The interface field do not match. actual: %v expected: %v", ipaddr.Interface, expectedIpAddress.Interface)
	}

	if strings.Compare(ipaddr.Network, expectedIpAddress.Network) != 0 {
		t.Errorf("The network field do not match. actual: %v expected: %v", ipaddr.Network, expectedIpAddress.Network)
	}

	foundIpAddress, err := c.FindIpAddress(ipaddr.Id)

	if err != nil {
		t.Errorf("Error getting ip addr with: %v", err)
	}

	if !reflect.DeepEqual(ipaddr, foundIpAddress) {
		t.Errorf("Created ip address and found ip address do not match. actual: %v expected: %v", foundIpAddress, ipaddr)
	}

	err = c.DeleteIpAddress(ipaddr.Id)

	if err != nil {
		t.Errorf("Error deleting ip address with: %v", err)
	}
}

func TestAdd_Update_DeleteIpAddress(t *testing.T) {
	c := NewClient(GetConfigFromEnv())

	initialAddress := "1.1.1.1/24"
	updatedAddress := "1.1.1.2/24"
	ifname := "ether1"
	updatedifname := "ether2"
	network := "1.1.1.0"
	expectedIpAddress := &IpAddress{
		Address:   updatedAddress,
		Interface: updatedifname,
		Network:   network,
	}

	init_ipaddr, err := c.AddIpAddress(
		initialAddress,
		ifname,
	)

	if err != nil {
		t.Errorf("Error creating a ip address with: %v", err)
	}

	if len(init_ipaddr.Id) < 1 {
		t.Errorf("The created ip address does not have an Id: %v", init_ipaddr)
	}

	updated_ipaddr, err := c.UpdateIpAddress(init_ipaddr.Id, updatedAddress, updatedifname)

	if err != nil {
		t.Errorf("Error updating the ip address with: %v", err)
	}

	if strings.Compare(updated_ipaddr.Address, expectedIpAddress.Address) != 0 {
		t.Errorf("The address field do not match. actual: %v expected: %v", updated_ipaddr.Address, expectedIpAddress.Address)
	}

	if strings.Compare(updated_ipaddr.Interface, expectedIpAddress.Interface) != 0 {
		t.Errorf("The interface field do not match. actual: %v expected: %v", updated_ipaddr.Interface, expectedIpAddress.Interface)
	}

	if strings.Compare(updated_ipaddr.Network, expectedIpAddress.Network) != 0 {
		t.Errorf("The network field do not match. actual: %v expected: %v", updated_ipaddr.Network, expectedIpAddress.Network)
	}

	foundIpAddress, err := c.FindIpAddress(updated_ipaddr.Id)

	if err != nil {
		t.Errorf("Error getting ip addr with: %v", err)
	}

	if !reflect.DeepEqual(updated_ipaddr, foundIpAddress) {
		t.Errorf("Created ip address and found ip address do not match. actual: %v expected: %v", foundIpAddress, updated_ipaddr)
	}

	err = c.DeleteIpAddress(updated_ipaddr.Id)

	if err != nil {
		t.Errorf("Error deleting ip address with: %v", err)
	}
}

func TestFindIpAddress_forNonExistingIpAddress(t *testing.T) {
	c := NewClient(GetConfigFromEnv())

	ipaddrId := "Invalid id"
	_, err := c.FindIpAddress(ipaddrId)

	expectedErrStr := fmt.Sprintf("ip address `%s`not found", ipaddrId)
	if err == nil || err.Error() != expectedErrStr {
		t.Errorf("client should have received error indicating the following ip address `%s`was not found. Instead error was nil", ipaddrId)
	}
}
