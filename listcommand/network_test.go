package listcommand

import "testing"

func TestNetwork_Network(t *testing.T) {
	// run a list network command and make sure there is at least one listed
	n := Network{}
	networks, err := n.Network("")
	if err != nil {
		t.Fatalf("list network failed with error %s", err)
	}
	if len(networks) == 0 {
		t.Error("list network failed. unable to find any networks on the frontend")
	}
}
