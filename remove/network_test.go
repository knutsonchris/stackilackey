package remove

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.td.teradata.com/ck250037/stacki-lackey-2/add"
	"github.td.teradata.com/ck250037/stacki-lackey-2/listcommand"
)

func TestNetwork_Network(t *testing.T) {
	// before we can test the removal of a network we must fist add one
	add := add.Add{}
	_, err := add.Network.Network("testnetwork", "192.168.1.0", "255.255.255.0", "eth0", "2000", "", false, false)
	if err != nil {
		t.Fatalf("remove network set up failed. add network failed with error %s", err)
	}

	// make sure it was added properly
	list := listcommand.List{}
	networks, err := list.Network.Network("testnetwork")
	expected := listcommand.Network{
		NetworkName: "testnetwork",
		Address:     "192.168.1.0",
		Mask:        "255.255.255.0",
		Gateway:     "eth0",
		MTU:         2000,
		Zone:        "testnetwork",
		DNS:         false,
		PXE:         false,
	}
	var found bool
	for _, network := range networks {
		if cmp.Equal(network, expected) {
			found = true
		}
	}
	if !found {
		t.Fatalf("remove network set up failed. unable to find test network\n%+v\nin list of networks\n%+v", expected, networks)
	}

	// now we can remove it
	n := network{}
	_, err = n.Network("testnetwork")
	if err != nil {
		t.Fatalf("remove network failed with error %s", err)
	}

	// make sure it was removed correctly
	networks, err = list.Network.Network("")
	if err != nil {
		t.Fatalf("remove network check failed. list network failed with error %s", err)
	}
	for _, network := range networks {
		if cmp.Equal(network, expected) {
			t.Fatal("remove network check failed. did not expect to find test network after remove")
		}
	}
}
