package mikrotik

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var origGreName = "gre1"
var updatedGreName = "gre2"
var origGreComment = "comment1"
var updatedGreComment = "comment2"
var origGreRemoteAddr = "1.1.1.1"
var updatedGreRemoteAddr = "2.2.2.2"

func TestAccMikrotikResourceInterfaceGre_create(t *testing.T) {
	resourceName := "mikrotik_interface_gre.autotest"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMikrotikInterfaceGreDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccInterfaceGre(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccInterfaceGreExists(resourceName),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "name", origGreName),
					resource.TestCheckResourceAttr(resourceName, "comment", origGreComment),
					resource.TestCheckResourceAttr(resourceName, "remote_address", origGreRemoteAddr),
				),
			},
		},
	})
}

func TestAccMikrotikResourceInterfaceGre_update(t *testing.T) {
	resourceName := "mikrotik_interface_gre.autotest"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMikrotikInterfaceGreDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccInterfaceGre(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccInterfaceGreExists(resourceName),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "name", origGreName),
					resource.TestCheckResourceAttr(resourceName, "comment", origGreComment),
					resource.TestCheckResourceAttr(resourceName, "remote_address", origGreRemoteAddr),
				),
			},
			{
				Config: testAccInterfaceGreUpdateName(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccInterfaceGreExists(resourceName),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "name", updatedGreName),
					resource.TestCheckResourceAttr(resourceName, "comment", origGreComment),
					resource.TestCheckResourceAttr(resourceName, "remote_address", origGreRemoteAddr),
				),
			},
			{
				Config: testAccInterfaceGreUpdatedComment(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccInterfaceGreExists(resourceName),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "name", updatedGreName),
					resource.TestCheckResourceAttr(resourceName, "comment", updatedGreComment),
					resource.TestCheckResourceAttr(resourceName, "remote_address", origGreRemoteAddr),
				),
			},
			{
				Config: testAccInterfaceGreUpdatedRemoteAddress(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccInterfaceGreExists(resourceName),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "name", updatedGreName),
					resource.TestCheckResourceAttr(resourceName, "comment", updatedGreComment),
					resource.TestCheckResourceAttr(resourceName, "remote_address", updatedGreRemoteAddr),
				),
			},
		},
	})
}

func testAccInterfaceGreExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("mikrotik_interface_gre does not exist in the statefile")
		}

		c := NewClient(GetConfigFromEnv())

		greIf, err := c.FindInterfaceGre(rs.Primary.ID)

		if err != nil {
			return fmt.Errorf("Unable to get the gre interface with error: %v", err)
		}

		if greIf == nil {
			return fmt.Errorf("Unable to get the gre interface")
		}

		if greIf.Id == rs.Primary.ID {
			return nil
		}
		return nil
	}
}

func testAccInterfaceGre() string {
	return fmt.Sprintf(`
resource "mikrotik_interface_gre" "autotest" {
	name = "%s"
	comment = "%s"
	remote_address = "%s"
}
`, origGreName, origGreComment, origGreRemoteAddr)
}

func testAccInterfaceGreUpdateName() string {
	return fmt.Sprintf(`
resource "mikrotik_interface_gre" "autotest" {
	name = "%s"
	comment = "%s"
	remote_address = "%s"
}
`, updatedGreName, origGreComment, origGreRemoteAddr)
}

func testAccInterfaceGreUpdatedComment() string {
	return fmt.Sprintf(`
resource "mikrotik_interface_gre" "autotest" {
	name = "%s"
	comment = "%s"
	remote_address = "%s"
}
`, updatedGreName, updatedGreComment, origGreRemoteAddr)
}

func testAccInterfaceGreUpdatedRemoteAddress() string {
	return fmt.Sprintf(`
resource "mikrotik_interface_gre" "autotest" {
	name = "%s"
	comment = "%s"
	remote_address = "%s"
}
`, updatedGreName, updatedGreComment, updatedGreRemoteAddr)
}

func testAccCheckMikrotikInterfaceGreDestroy(s *terraform.State) error {
	c := NewClient(GetConfigFromEnv())
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "mikrotik_interface_gre" {
			continue
		}

		greIf, err := c.FindInterfaceGre(rs.Primary.ID)

		_, ok := err.(*NotFound)
		if !ok && err != nil {
			return err
		}

		if greIf != nil {
			return fmt.Errorf("gre interface (%s) still exists", greIf.Id)
		}
	}
	return nil
}

func TestAccMikrotikResourceInterfaceGre_add_delete(t *testing.T) {
	c := NewClient(GetConfigFromEnv())

	var allow_fast_path (bool) = true
	var clamp_tcp_mss (bool) = true
	var comment (string) = "comment1"
	var copy_from (string) = ""
	var disabled (bool) = false
	var dont_fragment (string) = "no"
	var dscp (string) = "inherit"
	var ipsec_secret (string) = ""
	var keepalive (string) = "10s,10"
	var local_address (string) = ""
	var mtu (string) = "auto"
	var name (string) = "gre_test"
	var remote_address (string) = "2.2.2.2"

	expectedGreIf := &InterfaceGre{
		Allow_fast_path: allow_fast_path,
		Clamp_tcp_mss:   clamp_tcp_mss,
		Comment:         comment,
		Copy_from:       copy_from,
		Disabled:        disabled,
		Dont_fragment:   dont_fragment,
		Dscp:            dscp,
		Ipsec_secret:    ipsec_secret,
		Keepalive:       keepalive,
		Local_address:   local_address,
		Mtu:             mtu,
		Name:            name,
		Remote_address:  remote_address,
	}

	greif, err := c.AddInterfaceGre(allow_fast_path, clamp_tcp_mss, comment, copy_from, disabled, dont_fragment, dscp, ipsec_secret, keepalive, local_address, mtu, name, remote_address)
	if err != nil {
		t.Errorf("Error creating a gre interface with: %v", err)
	}

	if len(greif.Id) < 1 {
		t.Errorf("The created gre interface does not have an Id: %v", greif)
	}

	if greif.Allow_fast_path && greif.Allow_fast_path != expectedGreIf.Allow_fast_path {
		t.Errorf("The allow-fast-path field do not match. actual: %v expected: %v", greif.Allow_fast_path, expectedGreIf.Allow_fast_path)
	}
	if greif.Clamp_tcp_mss && greif.Clamp_tcp_mss != expectedGreIf.Clamp_tcp_mss {
		t.Errorf("The clamp-tcp-mss field do not match. actual: %v expected: %v", greif.Clamp_tcp_mss, expectedGreIf.Clamp_tcp_mss)
	}
	if strings.Compare(greif.Comment, expectedGreIf.Comment) != 0 {
		t.Errorf("The comment field do not match. actual: %v expected: %v", greif.Comment, expectedGreIf.Comment)
	}
	if strings.Compare(greif.Copy_from, expectedGreIf.Copy_from) != 0 {
		t.Errorf("The copy-from field do not match. actual: %v expected: %v", greif.Copy_from, expectedGreIf.Copy_from)
	}
	if greif.Disabled != expectedGreIf.Disabled {
		t.Errorf("The disabled field do not match. actual: %v expected: %v", greif.Disabled, expectedGreIf.Disabled)
	}
	if greif.Dont_fragment != "" && strings.Compare(greif.Dont_fragment, expectedGreIf.Dont_fragment) != 0 {
		t.Errorf("The dont-fragment field do not match. actual: %v expected: %v", greif.Dont_fragment, expectedGreIf.Dont_fragment)
	}
	if strings.Compare(greif.Dscp, expectedGreIf.Dscp) != 0 {
		t.Errorf("The dscp field do not match. actual: %v expected: %v", greif.Dscp, expectedGreIf.Dscp)
	}
	if strings.Compare(greif.Ipsec_secret, expectedGreIf.Ipsec_secret) != 0 {
		t.Errorf("The ipsec-secret field do not match. actual: %v expected: %v", greif.Ipsec_secret, expectedGreIf.Ipsec_secret)
	}
	if strings.Compare(greif.Keepalive, expectedGreIf.Keepalive) != 0 {
		t.Errorf("The keepalive field do not match. actual: %v expected: %v", greif.Keepalive, expectedGreIf.Keepalive)
	}
	if strings.Compare(greif.Local_address, expectedGreIf.Local_address) != 0 {
		t.Errorf("The local-address field do not match. actual: %v expected: %v", greif.Local_address, expectedGreIf.Local_address)
	}
	if strings.Compare(greif.Mtu, expectedGreIf.Mtu) != 0 {
		t.Errorf("The mtu field do not match. actual: %v expected: %v", greif.Mtu, expectedGreIf.Mtu)
	}
	if strings.Compare(greif.Name, expectedGreIf.Name) != 0 {
		t.Errorf("The name field do not match. actual: %v expected: %v", greif.Name, expectedGreIf.Name)
	}
	if greif.Remote_address != "" && strings.Compare(greif.Remote_address, expectedGreIf.Remote_address) != 0 {
		t.Errorf("The remote-address field do not match. actual: %v expected: %v", greif.Remote_address, expectedGreIf.Remote_address)
	}

	foundGreIf, err := c.FindInterfaceGre(greif.Id)

	if err != nil {
		t.Errorf("Error getting gre interface with: %v", err)
	}

	if !reflect.DeepEqual(greif, foundGreIf) {
		t.Errorf("created gre interface and found gre interface do not match. actual: %v expected: %v", foundGreIf, greif)
	}

	err = c.DeleteInterfaceGre(greif.Id)

	if err != nil {
		t.Errorf("Error deleting gre interface with: %v", err)
	}
}

func TestAccMikrotikResourceInterfaceGre_add_update_delete(t *testing.T) {
	c := NewClient(GetConfigFromEnv())

	var initial_allow_fast_path (bool) = true
	var updated_allow_fast_path (bool) = false
	var initial_clamp_tcp_mss (bool) = true
	var updated_clamp_tcp_mss (bool) = false
	var initial_comment (string) = "comment1"
	var updated_comment (string) = "comment123"
	var copy_from (string) = ""
	var initial_disabled (bool) = false
	var updated_disabled (bool) = true
	var dont_fragment (string) = "no"
	var dscp (string) = "inherit"
	var ipsec_secret (string) = ""
	var keepalive (string) = "10s,10"
	var local_address (string) = ""
	var mtu (string) = "auto"
	var initial_name (string) = "gre_test"
	var updated_name (string) = "gre_test123"
	var initial_remote_address (string) = "2.2.2.2"
	var updated_remote_address (string) = "23.23.23.23"

	expectedGreIf := &InterfaceGre{
		Allow_fast_path: updated_allow_fast_path,
		Clamp_tcp_mss:   updated_clamp_tcp_mss,
		Comment:         updated_comment,
		Copy_from:       copy_from,
		Disabled:        updated_disabled,
		Dont_fragment:   dont_fragment,
		Dscp:            dscp,
		Ipsec_secret:    ipsec_secret,
		Keepalive:       keepalive,
		Local_address:   local_address,
		Mtu:             mtu,
		Name:            updated_name,
		Remote_address:  updated_remote_address,
	}

	initial_greif, err := c.AddInterfaceGre(initial_allow_fast_path, initial_clamp_tcp_mss, initial_comment, copy_from, initial_disabled, dont_fragment, dscp, ipsec_secret, keepalive, local_address, mtu, initial_name, initial_remote_address)
	if err != nil {
		t.Errorf("Error creating a gre interface with: %v", err)
	}

	if len(initial_greif.Id) < 1 {
		t.Errorf("The created gre interface does not have an Id: %v", initial_greif)
	}

	updated_greif, err := c.UpdateInterfaceGre(initial_greif.Id, updated_allow_fast_path, updated_clamp_tcp_mss, updated_comment, copy_from, updated_disabled, dont_fragment, dscp, ipsec_secret, keepalive, local_address, mtu, updated_name, updated_remote_address)

	if err != nil {
		t.Errorf("Error updating the gre interface with: %v", err)
	}

	if updated_greif.Allow_fast_path && updated_greif.Allow_fast_path != expectedGreIf.Allow_fast_path {
		t.Errorf("The allow-fast-path field do not match. actual: %v expected: %v", updated_greif.Allow_fast_path, expectedGreIf.Allow_fast_path)
	}
	if updated_greif.Clamp_tcp_mss && updated_greif.Clamp_tcp_mss != expectedGreIf.Clamp_tcp_mss {
		t.Errorf("The clamp-tcp-mss field do not match. actual: %v expected: %v", updated_greif.Clamp_tcp_mss, expectedGreIf.Clamp_tcp_mss)
	}
	if strings.Compare(updated_greif.Comment, expectedGreIf.Comment) != 0 {
		t.Errorf("The comment field do not match. actual: %v expected: %v", updated_greif.Comment, expectedGreIf.Comment)
	}
	if strings.Compare(updated_greif.Copy_from, expectedGreIf.Copy_from) != 0 {
		t.Errorf("The copy-from field do not match. actual: %v expected: %v", updated_greif.Copy_from, expectedGreIf.Copy_from)
	}
	if updated_greif.Disabled != expectedGreIf.Disabled {
		t.Errorf("The disabled field do not match. actual: %v expected: %v", updated_greif.Disabled, expectedGreIf.Disabled)
	}
	if updated_greif.Dont_fragment != "" && strings.Compare(updated_greif.Dont_fragment, expectedGreIf.Dont_fragment) != 0 {
		t.Errorf("The dont-fragment field do not match. actual: %v expected: %v", updated_greif.Dont_fragment, expectedGreIf.Dont_fragment)
	}
	if strings.Compare(updated_greif.Dscp, expectedGreIf.Dscp) != 0 {
		t.Errorf("The dscp field do not match. actual: %v expected: %v", updated_greif.Dscp, expectedGreIf.Dscp)
	}
	if strings.Compare(updated_greif.Ipsec_secret, expectedGreIf.Ipsec_secret) != 0 {
		t.Errorf("The ipsec-secret field do not match. actual: %v expected: %v", updated_greif.Ipsec_secret, expectedGreIf.Ipsec_secret)
	}
	if strings.Compare(updated_greif.Keepalive, expectedGreIf.Keepalive) != 0 {
		t.Errorf("The keepalive field do not match. actual: %v expected: %v", updated_greif.Keepalive, expectedGreIf.Keepalive)
	}
	if strings.Compare(updated_greif.Local_address, expectedGreIf.Local_address) != 0 {
		t.Errorf("The local-address field do not match. actual: %v expected: %v", updated_greif.Local_address, expectedGreIf.Local_address)
	}
	if strings.Compare(updated_greif.Mtu, expectedGreIf.Mtu) != 0 {
		t.Errorf("The mtu field do not match. actual: %v expected: %v", updated_greif.Mtu, expectedGreIf.Mtu)
	}
	if strings.Compare(updated_greif.Name, expectedGreIf.Name) != 0 {
		t.Errorf("The name field do not match. actual: %v expected: %v", updated_greif.Name, expectedGreIf.Name)
	}
	if updated_greif.Remote_address != "" && strings.Compare(updated_greif.Remote_address, expectedGreIf.Remote_address) != 0 {
		t.Errorf("The remote-address field do not match. actual: %v expected: %v", updated_greif.Remote_address, expectedGreIf.Remote_address)
	}

	foundGreIf, err := c.FindInterfaceGre(updated_greif.Id)

	if err != nil {
		t.Errorf("Error getting gre interface with: %v", err)
	}

	if !reflect.DeepEqual(updated_greif, foundGreIf) {
		t.Errorf("created gre interface and found gre interface do not match. actual: %v expected: %v", foundGreIf, updated_greif)
	}

	err = c.DeleteInterfaceGre(updated_greif.Id)

	if err != nil {
		t.Errorf("Error deleting gre interface with: %v", err)
	}
}

func TestAccMikrotikResourceInterfaceGre_find_nonexisting(t *testing.T) {
	c := NewClient(GetConfigFromEnv())

	greifId := "Invalid Id"
	_, err := c.FindInterfaceGre(greifId)

	expectedErrStr := fmt.Sprintf("gre interface `%s`not found", greifId)
	if err == nil || err.Error() != expectedErrStr {
		t.Errorf("client should have received error indicating the following gre interface `%s`was not found. Instead error was nil", greifId)
	}

}
