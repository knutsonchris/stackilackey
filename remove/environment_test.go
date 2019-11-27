package remove

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/knutsonchris/stackilackey/add"
	"github.com/knutsonchris/stackilackey/listcommand"
)

func TestEnvironment_Environment(t *testing.T) {
	// before we can test the removal of an environment we need to add one
	add := add.Add{}
	_, err := add.Environment.Environment("testremoveenvironment")
	if err != nil {
		t.Fatalf("remove environment set up failed. add environment failed with error %s", err)
	}

	// make sure it was added correctly
	list := listcommand.List{}
	environments, err := list.Environment.Environment("testremoveenvironment")
	if err != nil {
		t.Fatalf("remove environment set up failed. list environment failed with error %s", err)
	}
	var found bool
	for _, env := range environments {
		if env.EnvironmentName == "testremoveenvironment" {
			found = true
		}
	}
	if !found {
		t.Fatal("remove environment set up failed. unable to find test environment in list of environments")
	}

	// now we can remove it
	e := environment{}
	_, err = e.Environment("testremoveenvironment")
	if err != nil {
		t.Fatalf("remove environment failed with error %s", err)
	}

	// make sure it was removed correctly
	environments, err = list.Environment.Environment("")
	if err != nil {
		t.Fatalf("remove environment check failed. list environment failed with error %s", err)
	}
	for _, env := range environments {
		if env.EnvironmentName == "testremoveenvironment" {
			t.Fatal("remove environment check failed. did not expect to find test environment after remove")
		}
	}
}

func TestEnvironment_Attr(t *testing.T) {
	// before we can test the removal of an environment attr, we need to to add an env and an attr
	add := add.Add{}
	_, err := add.Environment.Environment("testremoveenvironmentattr")
	if err != nil {
		t.Fatalf("remove environment attr set up failed. add environment failed with error %s", err)
	}
	_, err = add.Environment.Attr("testremoveenvironmentattr", "testattr", "testvalue", false)
	if err != nil {
		t.Fatalf("remove environemnt attr set up failed. add environment failed with error %s", err)
	}

	// make sure it was added correctly
	list := listcommand.List{}
	attrs, err := list.Environment.Attr("testremoveenvironmentattr", "testattr")
	if err != nil {
		t.Fatalf("remove environment attr set up failed. list environment failed with error %s", err)
	}
	expected := listcommand.Attr{
		AttrName: "testattr",
		Value:    "testvalue",
		Scope:    "environment",
		Type:     "var",
	}
	var found bool
	for _, attr := range attrs {
		if cmp.Equal(attr, expected) {
			found = true
		}
	}
	if !found {
		t.Fatal("remove environemnt attr set up failed. unable to find test environment after add")
	}

	// now we can remove it
	e := environment{}
	_, err = e.Attr("testremoveenvironmentattr", "testattr")
	if err != nil {
		t.Fatalf("remove environment attr failed with error %s", err)
	}

	// make sure it was removed correctly
	attrs, err = list.Environment.Attr("testremoveenvironmentattr", "testattr")
	if err != nil {
		t.Fatalf("remove environment attr check failed. list environment failed with error %s", err)
	}
	for _, attr := range attrs {
		if cmp.Equal(attr, expected) {
			t.Fatal("remove environment attr failed. did not expect to find test attr after remove")
		}
	}

	// to clean up after ourselves, remove the test environment also
	_, err = e.Environment("testremoveenvironmentattr")
	if err != nil {
		t.Fatal("remove environment attr failed. unable to remove the test environment after attr removal")
	}
}

func TestEnvironment_Firewall(t *testing.T) {
	// before we can test the removal of an environment firewall we first need to add an environment and a firewall
	add := add.Add{}
	_, err := add.Environment.Environment("testremoveenvironmentfirewall")
	if err != nil {
		t.Fatalf("remove environment firewall set up failed. add environment failed with error %s", err)
	}
	_, err = add.Environment.Firewall("testremoveenvironmentfirewall", "ACCEPT", "FORWARD", "udp", "www", "testComment", "-m state", "private", "private", "testenvironmentfirewall", "filter")
	if err != nil {
		t.Fatalf("remove environment firewall set up failed. add environment firewall failed with error %s", err)
	}

	// make sure it was added correctly
	list := listcommand.List{}
	firewalls, err := list.Environment.Firewall("testremoveenvironmentfirewall")
	if err != nil {
		t.Fatalf("remove environment firewall setup failed. list environment firewall failed with error %s", err)
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
		Name:          "testenvironmentfirewall",
		Table:         "filter",
		Source:        "E",
		Type:          "var",
	}

	var found bool
	for _, firewall := range firewalls {
		if cmp.Equal(firewall, expected) {
			found = true
		}
	}

	if !found {
		t.Fatal("remove environment firewall set up failed. unable to find expected firewall rule after add")
	}

	// now we can remove it
	e := environment{}
	_, err = e.Firewall("testremoveenvironmentfirewall", "testenvironmentfirewall")
	if err != nil {
		t.Fatalf("remove environment firewall failed with error %s", err)
	}

	// make sure it was removed properly
	firewalls, err = list.Environment.Firewall("testremoveenvironmentfirewall")
	if err != nil {
		t.Fatalf("remove environment firewall check failed. list appliance firewall failed with error %s", err)
	}
	for _, firewall := range firewalls {
		if cmp.Equal(firewall, expected) {
			t.Fatal("remove environment firewall check failed. did not expected to find supposedly removed firewall rule")
		}
	}

	// to clean up after ourselves, remove the test environment too
	_, err = e.Environment("testremoveenvironmentfirewall")
	if err != nil {
		t.Fatal("remove environment firewall check failed. unable to remove the test environment after firewall removal")
	}
}

func TestEnvironment_Route(t *testing.T) {
	// before we can test the removal of an environment route we need to first add one
	add := add.Add{}
	_, err := add.Environment.Environment("testremoveenvironmentroute")
	if err != nil {
		t.Fatalf("remove environment route set up failed. add environment failed with error %s", err)
	}
	_, err = add.Environment.Route("testremoveenvironmentroute", "192.168.0.1", "private", "eth0", "255.255.255.0")
	if err != nil {
		t.Fatalf("remove environment route set up failed. add environment route failed with error %s", err)
	}

	// make sure it was added correctly
	list := listcommand.List{}
	routes, err := list.Environment.Route("testremoveenvironmentroute")
	if err != nil {
		t.Fatalf("remove environemnt route set up failed. list environment route failed with error %s", err)
	}
	expected := listcommand.Route{
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
		t.Fatal("remove environment route failed. unable to find expected route after add")
	}

	// now we can remove it
	e := environment{}
	_, err = e.Route("testremoveenvironmentroute", "192.168.0.1")
	if err != nil {
		t.Fatalf("remove environment route failed with error %s", err)
	}

	// make sure it was removed correctly
	routes, err = list.Environment.Route("testremoveenvironmentroute")
	if err != nil {
		t.Fatalf("remove environemnt route check failed. list environment route failed with error %s", err)
	}
	for _, route := range routes {
		if cmp.Equal(route, expected) {
			t.Fatal("remove environment route failed. did not expect to find test route after remove")
		}
	}

	// to clean up after ourselves, we need to remove the test environment also
	_, err = e.Environment("testremoveenvironmentroute")
	if err != nil {
		t.Fatal("remove environment route failed. unable to remove test environment after route removal")
	}
}

func TestEnvironment_StorageController(t *testing.T) {
	// before we can test the removal of an environment storage contoller we need to add one
	add := add.Add{}
	_, err := add.Environment.Environment("testremoveenvcontroller")
	if err != nil {
		t.Fatalf("remove environment storage controller set up failed. add environment failed with error %s", err)
	}
	_, err = add.Environment.StorageController("testremoveenvcontroller", "9", 30, 31, 5, 34)
	if err != nil {
		t.Fatalf("remove environment storage controller set up failed. add environment storage controller failed with error %s", err)
	}

	// make sure it was added correctly
	list := listcommand.List{}
	controllers, err := list.Environment.StorageController("testremoveenvcontroller")
	if err != nil {
		t.Fatalf("remove environment storage controller set up failed. list storage controller failed with error %s", err)
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
		t.Fatal("remove environment storage controller set up failed. unable to find test storage controller after add")
	}

	// now we can remove it
	e := environment{}
	_, err = e.StorageController("testremoveenvcontroller", "34")
	if err != nil {
		t.Fatalf("remove environment storage controller failed with error %s", err)
	}

	// make sure it was removed correctly
	controllers, err = list.Environment.StorageController("testremoveenvcontroller")
	if err != nil {
		t.Fatalf("remove environment storage controller check failed. list storage controller failed with error %s", err)
	}
	for _, controller := range controllers {
		if cmp.Equal(controller, expected) {
			t.Fatal("remove environment storage controller failed. did not expect to find test controller after remove")
		}
	}

	// to clean up after ourselves, remove the test environment also
	_, err = e.Environment("testremoveenvcontroller")
	if err != nil {
		t.Fatal("remove environment storage controller failed. unable to remove test environment after storage controller removal")
	}
}
