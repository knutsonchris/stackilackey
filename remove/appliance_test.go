package remove

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.td.teradata.com/ck250037/stackilackey/add"
	"github.td.teradata.com/ck250037/stackilackey/listcommand"
)

func TestAppliance_Appliance(t *testing.T) {
	// before we can test remove appliance we need to add one
	add := add.Add{}
	_, err := add.Appliance.Appliance("testRemoveAppliance", "backend", true)
	if err != nil {
		t.Fatalf("remove appliance set up failed. add appliance failed with error %s", err)
	}

	// make sure it was added correctly
	list := listcommand.List{}
	appliances, err := list.Appliance.Appliance("testRemoveAppliance")
	if err != nil {
		t.Fatalf("remove appliance set up failed. list appliance failed with error %s", err)
	}
	var applianceFound bool
	for _, appliance := range appliances {
		if appliance.ApplianceName == "testRemoveAppliance" {
			applianceFound = true
		}
	}
	if !applianceFound {
		t.Fatal("remoev applaince set up failed. unable to find expected testRemoveAppliance")
	}

	// now we can remove it
	a := appliance{}
	_, err = a.Appliance("testRemoveAppliance")
	if err != nil {
		t.Fatalf("remove appliance failed with error %s", err)
	}

	// now make sure it is gone. attempting to list a non-existant applaince will result in an error
	appliances, err = list.Appliance.Appliance("")
	if err != nil {
		t.Fatalf("remove appliance check failed. list appliance failed with error %s", err)
	}
	for _, applaince := range appliances {
		if applaince.ApplianceName == "testRemoveAppliance" {
			t.Fatal("remove appliance failed. did not expect to find testRemoveAppliance")
		}
	}
}

func TestAppliance_Attr(t *testing.T) {
	// kickstartable is a default attr on the default applaince backend. we can test removing it without ever adding it
	a := appliance{}
	_, err := a.Attr("backend", "kickstartable")
	if err != nil {
		t.Fatalf("remove applaince attr failed with error %s", err)
	}

	// make sure it is removes
	list := listcommand.List{}
	attrs, err := list.Appliance.Attr("backend", "")
	if err != nil {
		t.Fatalf("remove applaince attr check failed. list applaince attr failed with err %s", err)
	}
	for _, attr := range attrs {
		if attr.AttrName == "kickstartable" {
			t.Fatal("remove applaince attr check failed. did not expect to find kickstartable")
		}
	}

	// reset the state of our test environemnt by re-adding the attr
	add := add.Add{}
	_, err = add.Appliance.Attr("backend", "kickstartable", "true", false)
	if err != nil {
		t.Errorf("remove appliance attr backend reset failed. add applaince attr failed with error %s", err)
	}
}

func TestApplaince_Firewall(t *testing.T) {
	// before we can test the removal of an applaince firewall, we must add one
	add := add.Add{}
	_, err := add.Appliance.Firewall("backend", "ACCEPT", "FORWARD", "udp", "www", "testComment", "-m state", "private", "private", "testApplianceFirewall", "filter")
	/*
		 i suppose it does not matter if the rule already exist. we are testing that the remove works correctly, so if it is already here then get rid of it
		if err != nil {
			t.Fatalf("remove appliance firewall set up failed. add appliance firewall failed with error %s", err)
		}
	*/

	// make sure it was added correctly
	list := listcommand.List{}
	firewalls, err := list.Appliance.Firewall("backend")
	if err != nil {
		t.Fatalf("remove appliance firewall setup failed. list applaince firewall failed with error %s", err)
	}

	expected := listcommand.Firewall{
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

	var found bool
	for _, firewall := range firewalls {
		if cmp.Equal(firewall, expected) {
			found = true
		}
	}

	if !found {
		t.Fatal("remove applaince firewall set up failed. unable to find expected firewall rule after add")
	}

	// now we can remove it
	a := appliance{}
	_, err = a.Firewall("backend", "testApplianceFirewall")
	if err != nil {
		t.Fatalf("remove appliance firewall failed with error %s", err)
	}

	// make sure it was removed properly
	firewalls, err = list.Appliance.Firewall("backend")
	if err != nil {
		t.Fatalf("remove applaince firewall check failed. list appliance firewall failed with error %s", err)
	}
	for _, firewall := range firewalls {
		if cmp.Equal(firewall, expected) {
			t.Fatal("remove appliance attr check failed. did not expected to find supposedly removed firewall rule")
		}
	}
}

func TestAppliance_Route(t *testing.T) {
	// before we can test the removal of an appliance route, we must add one
	add := add.Add{}
	_, err := add.Appliance.Route("backend", "192.168.0.1", "private", "eth0", "255.255.255.0")
	if err != nil {
		t.Fatalf("remove appliance route set up failed. add appliance route failed with error %s", err)
	}

	// make sure it was added correctly
	expected := listcommand.Route{
		Network:       "192.168.0.1",
		Subnet:        "private",
		InterfaceName: "eth0",
		Netmask:       "255.255.255.0",
	}
	list := listcommand.List{}
	routes, err := list.Appliance.Route("backend")
	if err != nil {
		t.Fatalf("remove appliance route set up failed. list appliance route failed with error %s", err)
	}
	var found bool
	for _, route := range routes {
		if cmp.Equal(route, expected) {
			found = true
		}
	}
	if !found {
		t.Fatal("remove appliance route set up failed. test applaince route not found after add")
	}

	// now we can remove it
	app := appliance{}
	_, err = app.Route("backend", "192.168.0.1")
	if err != nil {
		t.Fatalf("remove appliance route failed with error %s", err)
	}

	// make sure the remove was successful
	routes, err = list.Appliance.Route("backend")
	if err != nil {
		t.Fatalf("remove appliance check failed. list appliance route failed with error %s", err)
	}
	for _, route := range routes {
		if cmp.Equal(route, expected) {
			t.Fatal("remove appliance chack failed. did not expect to find test appliance route after remove")
		}
	}
}

func TestAppliance_StorageController(t *testing.T) {
	// before we can test the removal of an appliance storage controller we need to add one
	add := add.Add{}
	_, err := add.Appliance.StorageController("backend", 9, 30, 31, 5, "34")
	if err != nil {
		t.Fatalf("remove appliance storage controller set up failed. add storage controller failed with error %s", err)
	}

	// make sure it was added correctly
	list := listcommand.List{}
	controllers, err := list.Appliance.StorageController("backend")
	if err != nil {
		t.Fatalf("remove applaince storage controller set up failed. list storage controller failed with error %s", err)
	}
	expected := listcommand.Controller{
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
		t.Fatal("remove appliance storage controller set up failed. unable to find test storage controller after add")
	}

	// now we can remove it
	app := appliance{}
	_, err = app.StorageController("backend", "34")
	if err != nil {
		t.Fatalf("remove appliance storage controller failed with error %s", err)
	}

	// make sure it was removed properly
	controllers, err = list.Appliance.StorageController("backend")
	if err != nil {
		t.Fatalf("remove applaince storage controller check failed. list storage controller failed with error %s", err)
	}
	for _, controller := range controllers {
		if cmp.Equal(controller, expected) {
			t.Fatal("remove appliance storage controller check failed. did not expect to find test storage controller after remove")
		}
	}
}
