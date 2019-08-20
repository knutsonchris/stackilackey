package listcommand

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.td.teradata.com/ck250037/stacki-lackey-2/add"
	"github.td.teradata.com/ck250037/stacki-lackey-2/remove"
)

func TestAppliance_Appliance(t *testing.T) {
	// run a list command for a specific appliance and make sure we got the right one
	a := Appliance{}
	appliances, err := a.Appliance("backend")
	if err != nil {
		t.Fatalf("list applaince backend failed with error %s", err)
	}
	if len(appliances) == 0 {
		t.Fatal("list applaince backend check failed. unable to find any applainces named backend")
	}
	if appliances[0].ApplianceName != "backend" {
		t.Errorf("list appliance backend failed. listed backend but got %s", appliances[0].ApplianceName)
	}
}

func TestAppliance_Attr(t *testing.T) {
	// run a list command for a specific appliance attr and make sure we get the right one
	a := Appliance{}
	attrs, err := a.Attr("backend", "kickstartable")
	if err != nil {
		t.Fatalf("list appliance attr failed with error %s", err)
	}
	if len(attrs) == 0 {
		t.Fatal("list appliance attr check failed. unable to find any applaince attrs for named kickstartable on the applaince named backend")
	}
	if attrs[0].AttrName != "kickstartable" {
		t.Errorf("list appliance attr backend attr=kickstartable failed. listed kickstartable but got %s", attrs[0].AttrName)
	}
}

func TestAppliance_Firewall(t *testing.T) {
	// before we run the appliance firewall test we need to add one
	a := add.Add{}
	_, err := a.Appliance.Firewall("backend", "ACCEPT", "FORWARD", "udp", "www", "testComment", "-m state", "private", "private", "testApplianceFirewall", "filter")
	if err != nil {
		t.Fatalf("list appliance firewall set up failed. add appliance firewall failed with error %s", err)
	}

	// now test if we can find it
	app := Appliance{}
	firewalls, err := app.Firewall("backend")
	if err != nil {
		t.Fatalf("list appliance firewall failed with error %s", err)
	}
	expected := Firewall{
		Action:        "ACCEPT",
		Chain:         "FORWARD",
		Protocol:      "udp",
		Service:       "www",
		Comment:       "testComment",
		Flags:         "-m state",
		Network:       "private",
		OutputNetwork: "private",
		Name:          "testApplianceFirewall",
		Table:         "filter",
		Source:        "A",
		Type:          "var",
	}

	for _, firewall := range firewalls {
		if cmp.Equal(firewall, expected) {
			return
		}
	}
	t.Fatalf("list appliance firewall failed. unable to find expected firewall rule:\n %+v \nin list of firewall rules: \n %+v", expected, firewalls)

	// clean up our test and remove it
	remove := remove.Remove{}
	_, err = remove.Appliance.Firewall("backend", "testApplianceFirewall")
	if err != nil {
		t.Fatalf("list appliance firewall cleanup failed. remove appliance firewall failed with error %s", err)
	}

	// make sure the rule was removed successfully
	firewalls, err = app.Firewall("backend")
	if err != nil {
		t.Fatalf("list applaince firewall cleanup failed. list appliance firewall failed with error %s", err)
	}
	for _, firewall := range firewalls {
		if cmp.Equal(firewall, expected) {
			t.Fatal("list appliance firewall cleanup failed. did not expect to find test firewall rule after remove")
		}
	}
}

func TestApplaince_Route(t *testing.T) {
	// before we run the appliance route test we need to add one
	a := add.Add{}
	_, err := a.Appliance.Route("backend", "192.168.0.1", "private", "eth0", "255.255.255.0")
	if err != nil {
		t.Fatalf("list appliance route set up failed. add appliance route failed with error %s", err)
	}
	app := Appliance{}
	routes, err := app.Route("backend")
	if err != nil {
		t.Fatalf("list appliance route failed with error %s", err)
	}
	expected := Route{
		Network:       "192.168.0.1",
		Subnet:        "private",
		InterfaceName: "eth0",
		Netmask:       "255.255.255.0",
	}

	var found bool
	for _, route := range routes {
		if cmp.Equal(route, expected) {
			found = true
		}
	}
	if !found {
		t.Fatalf("\nlist appliance route failed. unable to find expected backend route:\n%+v \nin list of routes: \n%+v", expected, routes)
	}

	// clean up for the next test
	remove := remove.Remove{}
	_, err = remove.Appliance.Route("backend", "192.168.0.1")
	if err != nil {
		t.Fatalf("list appliance route clean up failed. remove appliance route failed with error %s", err)
	}
}

func TestAppliance_StorageController(t *testing.T) {
	// before we run the appliance storage controller test we need to add one
	a := add.Add{}
	_, err := a.Appliance.StorageController("backend", 9, 30, 31, 5, "34")
	if err != nil {
		t.Fatalf("list storage controller set up failed. add storage controller failed with error %s", err)
	}
	app := Appliance{}
	controllers, err := app.StorageController("backend")
	if err != nil {
		t.Fatalf("list storage controller failed with error %s", err)
	}
	expected := Controller{
		ArrayID:   "9",
		Adapter:   "30",
		Enclosure: "31",
		RaidLevel: "5",
		Slot:      "34",
		Options:   "",
	}
	var found bool
	for _, controller := range controllers {
		if cmp.Equal(controller, expected) {
			found = true
		}
	}
	if !found {
		t.Fatalf("\nlist storage controller failed. unable to find expected controller: \n%+v \nin list of controllers: \n %+v", expected, controllers)
	}

	// clean up for the next test
	remove := remove.Remove{}
	_, err = remove.Appliance.StorageController("backend", "34")
	if err != nil {
		t.Fatalf("list appliance storage controller clean up failed. remove applaince storage controller failed with error %s", err)
	}
}

func TestAppliance_XML(t *testing.T) {
	// run a list command for an appliance XML and make sure that we got XML for the correct appliance
	a := Appliance{}
	xmls, err := a.XML("backend")
	if err != nil {
		t.Fatalf("list appliance xml failed wit error %s", err)
	}
	if xmls[0].Col0 != "backend" {
		t.Errorf("list appliance xml backend failed. listed backend but got %s", xmls[0].Col0)
	}
}
