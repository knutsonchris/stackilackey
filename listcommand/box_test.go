package listcommand

import "testing"

func TestBox_Box(t *testing.T) {
	// run a list box command for a specific box and make sure we get the correct one
	b := Box{}
	boxes, err := b.Box("default")
	if err != nil {
		t.Fatalf("list box failed with error %s", err)
	}
	if boxes[0].Name != "default" {
		t.Fatalf("list box failed. listed default and got %s", boxes[0].Name)
	}
}
