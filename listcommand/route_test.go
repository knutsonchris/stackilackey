package listcommand

import "testing"

func TestRoute_Route(t *testing.T) {
	// run a list route command and make sure that we get at least one back
	r := Route{}
	routes, err := r.Route()
	if err != nil {
		t.Fatalf("list route failed with error %s", err)
	}
	if len(routes) == 0 {
		t.Error("list route failed. unable to find any routes on the frontend")
	}
}
