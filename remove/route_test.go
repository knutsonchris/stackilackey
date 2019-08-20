package remove

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.td.teradata.com/ck250037/stacki-lackey-2/add"
	"github.td.teradata.com/ck250037/stacki-lackey-2/listcommand"
)

func TestRoute_Route(t *testing.T) {
	// before we can test the removal of a route we need to add one
	add := add.Add{}
	_, err := add.Route.Route("192.168.0.1", "eth0", "private", "255.255.255.0")
	if err != nil {
		t.Fatalf("remove route set up failed. add route failed with error %s", err)
	}

	// make sure it was added correctly
	list := listcommand.List{}
	routes, err := list.Route.Route()
	if err != nil {
		t.Fatalf("remove route set up failed. list route failed with error %s", err)
	}
	var found bool
	expected := listcommand.Route{
		Network:       "192.168.0.1",
		Gateway:       "eth0",
		InterfaceName: "private",
		Netmask:       "255.255.255.0",
	}
	for _, route := range routes {
		if cmp.Equal(route, expected) {
			found = true
		}
	}
	if !found {
		t.Fatalf("remove route failed. unable to find expected \n%+v\n in list of routes\n%+v", expected, routes)
	}

	// now we can remove it
	r := route{}
	_, err = r.Route("192.168.0.1")
	if err != nil {
		t.Fatalf("remove route failed with error %s", err)
	}

	// make sure it was removed correctly
	routes, err = list.Route.Route()
	if err != nil {
		t.Fatalf("remove route check failed. list route failed with error %s", err)
	}
	for _, route := range routes {
		if cmp.Equal(route, expected) {
			t.Fatal("remove route check failed. did not expectto find test route after remove")
		}
	}
}
