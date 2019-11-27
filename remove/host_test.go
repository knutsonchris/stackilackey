package remove

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/knutsonchris/stackilackey/add"
	"github.com/knutsonchris/stackilackey/listcommand"
)

// TODO: host storage

func TestHost_Host(t *testing.T) {
	// before we test remove host let's add a new one
	add := add.Add{}
	_, err := add.Host.Host("testremovehost", "backend", "default", "", "1", "2")
	if err != nil {
		t.Fatalf("remove host set up failed. add host failed with error %s", err)
	}

	// make sure it was added correctly
	expected := listcommand.Host{
		Name:          "testremovehost",
		Rack:          "1",
		Rank:          "2",
		Appliance:     "backend",
		OS:            "sles", // TODO: make this test more sophisticated, include multiple OSes
		Box:           "default",
		OSAction:      "default",
		InstallAction: "default",
	}
	list := listcommand.List{}
	hosts, err := list.Host.Host("testremovehost")
	if err != nil {
		t.Fatalf("remove host set up failed. list host failed with error %s", err)
	}
	var found bool
	for _, host := range hosts {
		if cmp.Equal(host, expected) {
			found = true
		}
	}
	if !found {
		t.Fatalf("remove host set up failed. unable to find expected test host \n%+v in list of hosts \n%+v", expected, hosts)
	}

	// now we can remove it
	h := host{}
	_, err = h.Host("testremovehost")
	if err != nil {
		t.Fatalf("remove host failed with error %s", err)
	}

	// make sure it was removed correctly
	hosts, err = list.Host.Host("")
	if err != nil {
		t.Fatalf("remove host check failed. list host failed with error %s", err)
	}
	for _, host := range hosts {
		if cmp.Equal(host, expected) {
			t.Fatalf("remove host check failed. did not expect to find test host after remove")
		}
	}
}

func TestHost_Attr(t *testing.T) {
	// before we remove a host we need to add one and add an attr to it
	add := add.Add{}
	_, err := add.Host.Host("testremovehostattrhost", "backend", "default", "", "1", "2")
	if err != nil {
		t.Fatalf("remove host attr set up failed. add host failed with error %s", err)
	}
	_, err = add.Host.Attr("testremovehostattrhost", "testremovehostattr", "testremovehostattrvalue", false)
	if err != nil {
		t.Fatalf("remove host attr set up failed. add host attr failed with error %s", err)
	}

	// check to make sure it is there
	list := listcommand.List{}
	attrs, err := list.Host.Attr("testremovehostattrhost")
	if err != nil {
		t.Fatalf("remove host attr set up failed. list host attr failed with error %s", err)
	}
	var found bool
	for _, attr := range attrs {
		if attr.AttrName == "testremovehostattr" {
			found = true
		}
	}
	if !found {
		t.Fatal("remove host attr set up failed. unable to find test attr after add")
	}

	// now we can remove it
	h := host{}
	_, err = h.Attr("testremovehostattrhost", "testremovehostattr")
	if err != nil {
		t.Fatalf("remove host attr failed with error %s", err)
	}

	// make sure it was removed correctly
	attrs, err = list.Host.Attr("testremovehostattrhost")
	if err != nil {
		t.Fatalf("remove host attr check failed. list host attr failed with error %s", err)
	}
	for _, attr := range attrs {
		if attr.AttrName == "testremovehostattr" {
			t.Fatal("remove host attr check failed. did not expect to find test attr after remove")
		}
	}

	// to clean up after ourselves, remove the test host
	_, err = h.Host("testremovehostattrhost")
	if err != nil {
		t.Fatal("remove host attr check failed. unable to remove test host after attr removal")
	}
}

func TestHost_Firewall(t *testing.T) {
	// before we can test remove host firewall, we need to add a host and a host firewall
	add := add.Add{}
	_, err := add.Host.Host("testremovehostfirewallhost", "backend", "default", "", "1", "2")
	if err != nil {
		t.Fatalf("remove host firewall set up failed. add host failed with error %s", err)
	}
	_, err = add.Host.Firewall("testremovehostfirewallhost", "ACCEPT", "FORWARD", "udp", "www", "testComment", "-m state", "private", "private", "testRemoveHostFirewall", "filter")
	if err != nil {
		t.Fatalf("remove host firewall set up failed. add host firewall failed with error %s", err)
	}

	// make sure it was added correctly
	expected := listcommand.Firewall{
		Action:        "ACCEPT",
		Chain:         "FORWARD",
		Protocol:      "udp",
		Service:       "www",
		Comment:       "testComment",
		Flags:         "-m state",
		Network:       "private",
		OutputNetwork: "private",
		Name:          "testRemoveHostFirewall",
		Table:         "filter",
		Source:        "H",
		Type:          "var",
	}
	list := listcommand.List{}
	firewalls, err := list.Host.Firewall("testremovehostfirewallhost")
	if err != nil {
		t.Fatalf("remove host firewall set up failed. list host firewall failed with error %s", err)
	}
	var found bool
	for _, firewall := range firewalls {
		if cmp.Equal(firewall, expected) {
			found = true
		}
	}
	if !found {
		t.Fatal("remove host firewall set up failed. unable to find host firewall after add")
	}

	// now we can remove it
	h := host{}
	_, err = h.Firewall("testremovehostfirewallhost", "testRemoveHostFirewall")
	if err != nil {
		t.Fatalf("remove host firewall failed with error %s", err)
	}

	// make sure it was removed correctly
	firewalls, err = list.Host.Firewall("testremovehostfirewallhost")
	if err != nil {
		t.Fatalf("remove host firewall check failed. list host firewall failed with error %s", err)
	}
	for _, firewall := range firewalls {
		if cmp.Equal(firewall, expected) {
			t.Fatal("remove host firewall check failed. did not expect to find test firewall after remove")
		}
	}

	// to clean up after ourselves, remove the test host
	_, err = h.Host("testremovehostfirewallhost")
	if err != nil {
		t.Fatal("remove host firewall check failed. unable to remove test host after firewall removal")
	}
}

func TestHost_Route(t *testing.T) {
	// before we can test remove route, we need to add a host and a host route
	add := add.Add{}
	_, err := add.Host.Host("testremovehostroutehost", "backend", "default", "", "1", "2")
	if err != nil {
		t.Fatalf("remove host route set up failed. add host failed with error %s", err)
	}
	_, err = add.Host.Route("testremovehostroutehost", "192.168.0.1", "eth0", "", "", "")
	if err != nil {
		t.Fatalf("remove host route set up failed. add host route failed with error %s", err)
	}

	// make sure that it was added correctly
	list := listcommand.List{}
	routes, err := list.Host.Route("testremovehostroutehost")
	if err != nil {
		t.Fatalf("remove host route set up failed. list host route failed with error %s", err)
	}
	var found bool
	for _, route := range routes {
		if route.Network == "192.168.0.1" {
			found = true
		}
	}
	if !found {
		t.Fatal("remove host route set up failed. unable to find expected host route after add")
	}

	// now we can remove it
	h := host{}
	_, err = h.Route("testremovehostroutehost", "192.168.0.1", "false")
	if err != nil {
		t.Fatalf("remove host route failed with error %s", err)
	}

	// make sure it was removed correctly
	routes, err = list.Host.Route("testremovehostroutehost")
	if err != nil {
		t.Fatalf("remove host route check failed. list host route failed with error %s", err)
	}
	for _, route := range routes {
		if route.Network == "192.168.0.1" {
			t.Fatalf("remove host route check. did not expect to find test route after remove")
		}
	}

	// to clean up after ourselves, remove the test host
	_, err = h.Host("testremovehostroutehost")
	if err != nil {
		t.Fatalf("remove host route check. unable to remove test host after route removal")
	}
}

func TestHost_StorageController(t *testing.T) {
	// before we can test the removal of a host storage controller we must add a host and a storage controller
	add := add.Add{}
	_, err := add.Host.Host("testhostremovestoragecontroller", "backend", "default", "", "1", "2")
	if err != nil {
		t.Fatalf("remove host storage controller set up failed. add host failed with error %s", err)
	}
	_, err = add.Host.StorageController("testhostremovestoragecontroller", 99, 98, 97, 5, 96)
	if err != nil {
		t.Fatalf("remove host storage controller set ip failed. add host storage controller failed with error %s", err)
	}

	// make sure it was added correctly
	list := listcommand.List{}
	controllers, err := list.Host.StorageController("testhostremovestoragecontroller")
	if err != nil {
		t.Fatalf("remove host storage controller set up failed. list host storage controller failed with error %s", err)
	}
	expected := listcommand.Controller{
		ArrayID:   "99",
		Adapter:   "98",
		Enclosure: "97",
		RaidLevel: "5",
		Slot:      "96",
	}
	var found bool
	for _, controller := range controllers {
		if cmp.Equal(controller, expected) {
			found = true
		}
	}
	if !found {
		t.Fatal("remove host storage controller set up failed. unable to find test controller after add")
	}

	// now we can remove it
	h := host{}
	_, err = h.StorageController("testhostremovestoragecontroller", "96")
	if err != nil {
		t.Fatalf("remove host storage controller failed with error %s", err)
	}

	// make sure it was removed correctly
	controllers, err = list.Host.StorageController("testhostremovestoragecontroller")
	if err != nil {
		t.Fatalf("remove host storage controller check failed. list host storage controller failed with error %s", err)
	}
	for _, controller := range controllers {
		if cmp.Equal(controller, expected) {
			t.Fatal("remove host storage controller check failed. did not expect to find test controller after remove")
		}
	}

	// to clean up after ourselves, remove the test host
	_, err = h.Host("testhostremovestoragecontroller")
	if err != nil {
		t.Fatal("remove host storage controller check failed. unable to remove test host after storage controller removal")
	}
}
