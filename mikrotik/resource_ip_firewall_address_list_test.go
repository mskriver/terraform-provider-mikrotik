package mikrotik

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var originalAddress string = "192.168.1.0/24"
var originalList string = "list1"
var originalComment string = "comment1"
var updatedAddress string = "192.168.2.0/24"
var updatedList string = "list2"
var updatedComment string = "comment2"

func TestAccMikrotikResourceIpFirewallAddressList_create(t *testing.T) {
	resourceName := "mikrotik_ip_firewall_address_list.autotest"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMikrotikIpFirewallAddressListDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccIpFirewallAddressList(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccIpFirewallAddressListExists(resourceName),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "address", originalAddress),
					resource.TestCheckResourceAttr(resourceName, "list", originalList),
					resource.TestCheckResourceAttr(resourceName, "comment", originalComment),
				),
			},
		},
	})
}

func TestAccMikrotikResourceIpFirewallAddressList_update(t *testing.T) {
	resourceName := "mikrotik_ip_firewall_address_list.autotest"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMikrotikIpFirewallAddressListDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccIpFirewallAddressList(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccIpFirewallAddressListExists(resourceName),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "address", originalAddress),
					resource.TestCheckResourceAttr(resourceName, "list", originalList),
					resource.TestCheckResourceAttr(resourceName, "comment", originalComment),
				),
			},
			{
				Config: testAccIpFirewallAddressListUpdateAddress(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccIpFirewallAddressListExists(resourceName),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "address", updatedAddress),
					resource.TestCheckResourceAttr(resourceName, "list", originalList),
					resource.TestCheckResourceAttr(resourceName, "comment", originalComment),
				),
			},
			{
				Config: testAccIpFirewallAddressListUpdateList(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccIpFirewallAddressListExists(resourceName),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "address", updatedAddress),
					resource.TestCheckResourceAttr(resourceName, "list", updatedList),
					resource.TestCheckResourceAttr(resourceName, "comment", originalComment),
				),
			},
			{
				Config: testAccIpFirewallAddressListUpdateComment(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccIpFirewallAddressListExists(resourceName),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "address", updatedAddress),
					resource.TestCheckResourceAttr(resourceName, "list", updatedList),
					resource.TestCheckResourceAttr(resourceName, "comment", updatedComment),
				),
			},
		},
	})
}

func testAccIpFirewallAddressListExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("mikrotik_ip_firewall_address_list does not exist in the statefile")
		}

		c := NewClient(GetConfigFromEnv())

		fwlist, err := c.FindIpFirewallAddressList(rs.Primary.ID)

		if err != nil {
			return fmt.Errorf("Unable to get the ip address with error: %v", err)
		}

		if fwlist == nil {
			return fmt.Errorf("Unable to get the ip address")
		}

		if fwlist.Id == rs.Primary.ID {
			return nil
		}
		return nil
	}
}

func testAccIpFirewallAddressList() string {
	return fmt.Sprintf(`
resource "mikrotik_ip_firewall_address_list" "autotest" {
	address = "%s"
	list = "%s"
	comment = "%s"
}
`, originalAddress, originalList, originalComment)
}

func testAccIpFirewallAddressListUpdateAddress() string {
	return fmt.Sprintf(`
resource "mikrotik_ip_firewall_address_list" "autotest" {
	address = "%s"
	list = "%s"
	comment = "%s"
}
`, updatedAddress, originalList, originalComment)
}

func testAccIpFirewallAddressListUpdateList() string {
	return fmt.Sprintf(`
resource "mikrotik_ip_firewall_address_list" "autotest" {
	address = "%s"
	list = "%s"
	comment = "%s"
}
`, updatedAddress, updatedList, originalComment)
}

func testAccIpFirewallAddressListUpdateComment() string {
	return fmt.Sprintf(`
resource "mikrotik_ip_firewall_address_list" "autotest" {
	address = "%s"
	list = "%s"
	comment = "%s"
}
`, updatedAddress, updatedList, updatedComment)
}

func testAccCheckMikrotikIpFirewallAddressListDestroy(s *terraform.State) error {
	c := NewClient(GetConfigFromEnv())
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "mikrotik_ip_firewall_address_list" {
			continue
		}

		fwlist, err := c.FindIpFirewallAddressList(rs.Primary.ID)

		_, ok := err.(*NotFound)
		if !ok && err != nil {
			return err
		}

		if fwlist != nil {
			return fmt.Errorf("ip firewall address-list (%s) still exists", fwlist.Id)
		}
	}
	return nil
}

func TestAccMikrotikResourceIpFirewallAddressList_add_delete(t *testing.T) {
	c := NewClient(GetConfigFromEnv())

	address := "192.168.1.0/24"
	list := "list1"
	comment := "comment1"

	expectedIpFirewallAddressList := &IpFirewallAddressList{
		Address: address,
		List:    list,
		Comment: comment,
	}
	fwlist, err := c.AddIpFirewallAddressList(
		address,
		list,
		comment,
	)

	if err != nil {
		t.Errorf("Error creating a ip firewall address-list with: %v", err)
	}

	if len(fwlist.Id) < 1 {
		t.Errorf("The created ip firewall address-list does not have an Id: %v", fwlist)
	}

	if strings.Compare(fwlist.Address, expectedIpFirewallAddressList.Address) != 0 {
		t.Errorf("The address field do not match. actual: %v expected: %v", fwlist.Address, expectedIpFirewallAddressList.Address)
	}

	if strings.Compare(fwlist.List, expectedIpFirewallAddressList.List) != 0 {
		t.Errorf("The list field do not match. actual: %v expected: %v", fwlist.List, expectedIpFirewallAddressList.List)
	}

	if strings.Compare(fwlist.Comment, expectedIpFirewallAddressList.Comment) != 0 {
		t.Errorf("The comment field do not match. actual: %v expected: %v", fwlist.Comment, expectedIpFirewallAddressList.Comment)
	}

	foundfwlist, err := c.FindIpFirewallAddressList(fwlist.Id)

	if err != nil {
		t.Errorf("Error getting ip firewall address-list with: %v", err)
	}

	if !reflect.DeepEqual(fwlist, foundfwlist) {
		t.Errorf("created ip firewall address-list and found adress-list do not match. actual: %v expected: %v", foundfwlist, fwlist)
	}

	err = c.DeleteIpFirewallAddressList(fwlist.Id)

	if err != nil {
		t.Errorf("Error deleting ip firewall address-list with: %v", err)
	}
}

func TestAccMikrotikResourceIpFirewallAddressList_add_update_delete(t *testing.T) {
	c := NewClient(GetConfigFromEnv())

	initial_address := "192.168.1.0/24"
	updated_address := "192.168.2.0/24"
	initial_list := "list1"
	updated_list := "list2"
	initial_comment := "comment1"
	updated_comment := "comment2"

	expectedIpFirewallAddressList := &IpFirewallAddressList{
		Address: updated_address,
		List:    updated_list,
		Comment: updated_comment,
	}

	init_fwlist, err := c.AddIpFirewallAddressList(
		initial_address,
		initial_list,
		initial_comment,
	)

	if err != nil {
		t.Errorf("Error creating a ip firewall address-list with: %v", err)
	}

	if len(init_fwlist.Id) < 1 {
		t.Errorf("The created ip firewall address-list does not have an Id: %v", init_fwlist)
	}

	updated_fwlist, err := c.UpdateIpFirewallAddressList(init_fwlist.Id, updated_address, updated_list, updated_comment)

	if err != nil {
		t.Errorf("Error updating the ip firewall address-list with: %v", err)
	}

	if strings.Compare(updated_fwlist.Address, expectedIpFirewallAddressList.Address) != 0 {
		t.Errorf("The address field do not match. actual: %v expected: %v", updated_fwlist.Address, expectedIpFirewallAddressList.Address)
	}

	if strings.Compare(updated_fwlist.List, expectedIpFirewallAddressList.List) != 0 {
		t.Errorf("The list field do not match. actual: %v expected: %v", updated_fwlist.List, expectedIpFirewallAddressList.List)
	}

	if strings.Compare(updated_fwlist.Comment, expectedIpFirewallAddressList.Comment) != 0 {
		t.Errorf("The comment field do not match. actual: %v expected: %v", updated_fwlist.Comment, expectedIpFirewallAddressList.Comment)
	}

	foundFwList, err := c.FindIpFirewallAddressList(updated_fwlist.Id)

	if err != nil {
		t.Errorf("Error getting ip addr with: %v", err)
	}

	if !reflect.DeepEqual(updated_fwlist, foundFwList) {
		t.Errorf("Created ip firewall address-list and found address-list do not match. actual: %v expected: %v", foundFwList, updated_fwlist)
	}

	err = c.DeleteIpFirewallAddressList(updated_fwlist.Id)

	if err != nil {
		t.Errorf("Error deleting ip firewall address-list with: %v", err)
	}
}

func TestAccMikrotikResourceIpFirewallAddressList_find_nonexisting(t *testing.T) {
	c := NewClient(GetConfigFromEnv())

	fwlistId := "Invalid id"
	_, err := c.FindIpFirewallAddressList(fwlistId)

	expectedErrStr := fmt.Sprintf("ip firewall address-list `%s`not found", fwlistId)
	if err == nil || err.Error() != expectedErrStr {
		t.Errorf("client should have received error indicating the following ip firewall address-list `%s`was not found. Instead error was nil", fwlistId)
	}
}
