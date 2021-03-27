package mikrotik

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var origAction = "accept"
var updatedAction = "drop"
var origChain = "chain1"
var updatedChain = "chain2"
var origSrcAddr = "192.168.0.0/24"
var updatedSrcAddr = "192.68.1.0/24"
var origDstAddr = "172.16.0.0/24"
var updatedDstAddr = "172.16.1.0/24"

func TestAccMikrotikResourceIpFirewallFilter_create(t *testing.T) {
	resourceName := "mikrotik_ip_firewall_filter.autotest"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMikrotikIpFirewallFilterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccIpFirewallFilter(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccIpFirewallFilterExists(resourceName),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "action", origAction),
					resource.TestCheckResourceAttr(resourceName, "chain", origChain),
					resource.TestCheckResourceAttr(resourceName, "src_address", origSrcAddr),
					resource.TestCheckResourceAttr(resourceName, "dst_address", origDstAddr),
				),
			},
		},
	})
}

func TestAccMikrotikResourceIpFirewallFilter_update(t *testing.T) {
	resourceName := "mikrotik_ip_firewall_filter.autotest"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMikrotikIpFirewallFilterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccIpFirewallFilter(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccIpFirewallFilterExists(resourceName),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "action", origAction),
					resource.TestCheckResourceAttr(resourceName, "chain", origChain),
					resource.TestCheckResourceAttr(resourceName, "src_address", origSrcAddr),
					resource.TestCheckResourceAttr(resourceName, "dst_address", origDstAddr),
				),
			},
			{
				Config: testAccIpFirewallFilterUpdateAction(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccIpFirewallFilterExists(resourceName),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "action", updatedAction),
					resource.TestCheckResourceAttr(resourceName, "chain", origChain),
					resource.TestCheckResourceAttr(resourceName, "src_address", origSrcAddr),
					resource.TestCheckResourceAttr(resourceName, "dst_address", origDstAddr),
				),
			},
			{
				Config: testAccIpFirewallFilterUpdateChain(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccIpFirewallFilterExists(resourceName),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "action", updatedAction),
					resource.TestCheckResourceAttr(resourceName, "chain", updatedChain),
					resource.TestCheckResourceAttr(resourceName, "src_address", origSrcAddr),
					resource.TestCheckResourceAttr(resourceName, "dst_address", origDstAddr),
				),
			},
			{
				Config: testAccIpFirewallFilterUpdateSrcAddress(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccIpFirewallFilterExists(resourceName),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "action", updatedAction),
					resource.TestCheckResourceAttr(resourceName, "chain", updatedChain),
					resource.TestCheckResourceAttr(resourceName, "src_address", updatedSrcAddr),
					resource.TestCheckResourceAttr(resourceName, "dst_address", origDstAddr),
				),
			},
			{
				Config: testAccIpFirewallFilterUpdateDstAddress(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccIpFirewallFilterExists(resourceName),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "action", updatedAction),
					resource.TestCheckResourceAttr(resourceName, "chain", updatedChain),
					resource.TestCheckResourceAttr(resourceName, "src_address", updatedSrcAddr),
					resource.TestCheckResourceAttr(resourceName, "dst_address", updatedDstAddr),
				),
			},
		},
	})
}

func testAccIpFirewallFilterExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("mikrotik_ip_firewall_filter does not exist in the statefile")
		}

		c := NewClient(GetConfigFromEnv())

		fwlist, err := c.FindIpFirewallFilter(rs.Primary.ID)

		if err != nil {
			return fmt.Errorf("Unable to get the ip firewall filter with error: %v", err)
		}

		if fwlist == nil {
			return fmt.Errorf("Unable to get the ip firewall filter")
		}

		if fwlist.Id == rs.Primary.ID {
			return nil
		}
		return nil
	}
}

func testAccIpFirewallFilter() string {
	return fmt.Sprintf(`
resource "mikrotik_ip_firewall_filter" "autotest" {
	action = "%s"
	chain = "%s"
	src_address = "%s"
	dst_address = "%s"
}
`, origAction, origChain, origSrcAddr, origDstAddr)
}

func testAccIpFirewallFilterUpdateAction() string {
	return fmt.Sprintf(`
resource "mikrotik_ip_firewall_filter" "autotest" {
	action = "%s"
	chain = "%s"
	src_address = "%s"
	dst_address = "%s"
}
`, updatedAction, origChain, origSrcAddr, origDstAddr)
}

func testAccIpFirewallFilterUpdateChain() string {
	return fmt.Sprintf(`
resource "mikrotik_ip_firewall_filter" "autotest" {
	action = "%s"
	chain = "%s"
	src_address = "%s"
	dst_address = "%s"
}
`, updatedAction, updatedChain, origSrcAddr, origDstAddr)
}

func testAccIpFirewallFilterUpdateSrcAddress() string {
	return fmt.Sprintf(`
resource "mikrotik_ip_firewall_filter" "autotest" {
	action = "%s"
	chain = "%s"
	src_address = "%s"
	dst_address = "%s"
}
`, updatedAction, updatedChain, updatedSrcAddr, origDstAddr)
}

func testAccIpFirewallFilterUpdateDstAddress() string {
	return fmt.Sprintf(`
resource "mikrotik_ip_firewall_filter" "autotest" {
	action = "%s"
	chain = "%s"
	src_address = "%s"
	dst_address = "%s"
}
`, updatedAction, updatedChain, updatedSrcAddr, updatedDstAddr)
}

func testAccCheckMikrotikIpFirewallFilterDestroy(s *terraform.State) error {
	c := NewClient(GetConfigFromEnv())
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "mikrotik_ip_firewall_filter" {
			continue
		}

		fwlist, err := c.FindIpFirewallFilter(rs.Primary.ID)

		_, ok := err.(*NotFound)
		if !ok && err != nil {
			return err
		}

		if fwlist != nil {
			return fmt.Errorf("ip firewall filter (%s) still exists", fwlist.Id)
		}
	}
	return nil
}
