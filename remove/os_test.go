package remove

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.td.teradata.com/ck250037/stacki-lackey-2/add"
	"github.td.teradata.com/ck250037/stacki-lackey-2/listcommand"
)

func TestOS_OS(t *testing.T) {
	// we have no add OS command to add a dummy OS, so we'll just remove ubuntu
	// make sure it is there first
	list := listcommand.List{}
	oses, err := list.OS.OS("")
	if err != nil {
		t.Fatalf("remove os set up failed. list os failed with error %s", err)
	}
	var found bool
	for _, os := range oses {
		if os.OSName == "ubuntu" {
			found = true
		}
	}
	if !found {
		t.Fatal("remove os set up failed. unable to find ubuntu in list of oses. find another os to test remove")
	}

	// now we can remove it
	o := os{}
	_, err = o.OS("ubuntu")
	if err != nil {
		t.Fatalf("remove os failed with error %s", err)
	}

	// make sure it was removed correctly
	oses, err = list.OS.OS("")
	if err != nil {
		t.Fatalf("remove os check failed. list os failed with error %s", err)
	}
	for _, os := range oses {
		if os.OSName == "ubuntu" {
			t.Fatal("remove os failed. did not expect to find os ubuntu after remove")
		}
	}
}

func TestOS_Attr(t *testing.T) {
	// before we test the removal of an OS attr we need to add one
	add := add.Add{}
	_, err := add.OS.Attr("sles", "testosattr", "testvalue", false)
	if err != nil {
		t.Fatalf("remove os attr set up failed. add os attr failed with error %s", err)
	}

	// make sure it was added correctly
	list := listcommand.List{}
	attrs, err := list.OS.Attr("sles")
	if err != nil {
		t.Fatalf("remove os attr set up failed. list os attr failed with error %s", err)
	}
	expected := listcommand.Attr{
		Scope:    "os",
		Type:     "var",
		AttrName: "testosattr",
		Value:    "testvalue",
	}
	var found bool
	for _, attr := range attrs {
		if cmp.Equal(attr, expected) {
			found = true
		}
	}
	if !found {
		t.Fatalf("remove os attr set up failed. unable to find test attr\n%+v\n in list of attrs\n%+v", expected, attrs)
	}

	// now we can remove it
	o := os{}
	_, err = o.Attr("sles", "testosattr")
	if err != nil {
		t.Fatalf("remove os attr failed with error %s", err)
	}

	// make sure it was removed correctly
	attrs, err = list.OS.Attr("sles")
	if err != nil {
		t.Fatalf("remove os attr check failed. list os attr failed with error %s", err)
	}
	for _, attr := range attrs {
		if cmp.Equal(attr, expected) {
			t.Fatal("remove os attr check failed. did not expect to find test attr after remove")
		}
	}
}

func TestOS_Firewall(t *testing.T) {
	// before we can remove an OS firewall we need to add one
	add := add.Add{}
	_, err := add.OS.Firewall("sles", "ACCEPT", "FORWARD", "udp", "www", "testComment", "-m state", "private", "private", "testosfirewall", "filter")
	if err != nil {
		t.Fatalf("remove os firewall set up failed. add os firewall failed with error %s", err)
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
		Name:          "testosfirewall",
		Table:         "filter",
		Source:        "O",
		Type:          "var",
	}
	list := listcommand.List{}
	firewalls, err := list.OS.Firewall("sles")
	if err != nil {
		t.Fatalf("remove os firewall set up failed. list os firewall failed with error %s", err)
	}
	var found bool
	for _, firewall := range firewalls {
		if cmp.Equal(firewall, expected) {
			found = true
		}
	}
	if !found {
		t.Fatalf("remove os firewall set  up failed. unable to find expected rule\n%+v\n in list of rules\n%+v", expected, firewalls)
	}

	// now we can remove it
	o := os{}
	_, err = o.Firewall("sles", "testosfirewall")
	if err != nil {
		t.Fatalf("remove os firewall failed with error %s", err)
	}

	// make sure it was removed correctly
	firewalls, err = list.OS.Firewall("sles")
	if err != nil {
		t.Fatalf("remove os firewall check failed. list os firewall failed with error %s", err)
	}
	for _, firewall := range firewalls {
		if cmp.Equal(firewall, expected) {
			t.Fatal("remove os firewall check failed. did not expect to find test rule after remove")
		}
	}
}

func TestOS_Route(t *testing.T) {
	// before we can test the removal of an OS route we need to add one
	add := add.Add{}
	_, err := add.OS.Route("sles", "192.168.0.1", "private", "eth0", "255.255.255.0")
	if err != nil {
		t.Fatalf("remove os route set up failed. add os route failed with error %s", err)
	}
	// make sure that it was added correctly
	expected := listcommand.Route{
		Network:       "192.168.0.1",
		Subnet:        "private",
		InterfaceName: "eth0",
		Netmask:       "255.255.255.0",
	}
	list := listcommand.List{}
	routes, err := list.OS.Route("sles")
	if err != nil {
		t.Fatalf("remove os route set up failed. list os route failed with error %s", err)
	}
	var found bool
	for _, route := range routes {
		if cmp.Equal(route, expected) {
			found = true
		}
	}
	if !found {
		t.Fatalf("remove os route set up failed. unable to find expected route\n%+v\n in list of os routes\n%+v", expected, routes)
	}

	// now we can remove it
	o := os{}
	_, err = o.Route("sles", "192.168.0.1")
	if err != nil {
		t.Fatalf("remove os route failed with error %s", err)
	}

	// make sure it was removed correctly
	routes, err = list.OS.Route("sles")
	if err != nil {
		t.Fatalf("remove os route check failed. list os route failed with error %s", err)
	}
	for _, route := range routes {
		if cmp.Equal(route, expected) {
			t.Fatal("remove os route check failed. did not expect to find test route after remove")
		}
	}
}

func TestOS_StorageController(t *testing.T) {
	// before we can test the removal of an os storage controller we need to add one
	add := add.Add{}
	_, err := add.OS.StorageController("redhat", "9", 30, 31, 5, 34)
	if err != nil {
		t.Fatalf("remove os storage controller set up failed. add storage controller failed with error %s", err)
	}

	// make sure it was added correctly
	list := listcommand.List{}
	controllers, err := list.OS.StorageController("redhat")
	if err != nil {
		t.Fatalf("remove os storage controller set up failed. list os storage controller failed with error %s", err)
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
		t.Fatal("remove os storage controller set up failed. unable to find test storage controller after add")
	}

	// now we can remove it
	o := os{}
	_, err = o.StorageController("redhat", "34", "", "")
	if err != nil {
		t.Fatalf("remove os storage controller failed with error %s", err)
	}

	// make sure it was removed correctly
	controllers, err = list.OS.StorageController("redhat")
	if err != nil {
		t.Fatalf("remove os storage controller check failed. list os storage controller failed with error %s", err)
	}
	for _, controller := range controllers {
		if cmp.Equal(controller, expected) {
			t.Fatal("remove os storage controller failed. did not expect to find test controller after remove")
		}
	}
}
