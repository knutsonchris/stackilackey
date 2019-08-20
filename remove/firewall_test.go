package remove

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.td.teradata.com/ck250037/stackilackey/add"
	"github.td.teradata.com/ck250037/stackilackey/listcommand"
)

func TestFirewall_Firewall(t *testing.T) {
	// before we can test removing a firewall rule we need to add one
	add := add.Add{}
	_, err := add.Firewall.Firewall("ACCEPT", "FORWARD", "udp", "www", "testComment", "-m state", "private", "private", "testRemoveFirewall", "filter")
	if err != nil {
		t.Fatalf("remove firewalll set up failed. add firewall failed with error %s", err)
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
		Name:          "testRemoveFirewall",
		Table:         "filter",
		Source:        "G",
		Type:          "var",
	}

	var found bool
	list := listcommand.List{}
	firewalls, err := list.Firewall.Firewall()
	for _, firewall := range firewalls {
		if cmp.Equal(firewall, expected) {
			found = true
		}
	}
	if !found {
		t.Fatal("remove firewall set up failed. unable to find test rule after add")
	}

	// now we can remove it
	f := firewall{}
	_, err = f.Firewall("testRemoveFirewall")
	if err != nil {
		t.Fatalf("remove firewall failed with error %s", err)
	}

	// check if it was removed correctly
	firewalls, err = list.Firewall.Firewall()
	for _, firewall := range firewalls {
		if cmp.Equal(firewall, expected) {
			t.Fatal("remove firewall check failed. did not expect to find test rule after remove")
		}
	}
}
