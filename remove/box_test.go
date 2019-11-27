package remove

import (
	"testing"

	"github.com/knutsonchris/stackilackey/add"
	"github.com/knutsonchris/stackilackey/listcommand"
)

func TestBox_box(t *testing.T) {
	// before we can test remove box we need to add one
	add := add.Add{}
	_, err := add.Box.Box("testremovebox", "")
	if err != nil {
		t.Fatalf("remove box set up failed. add box failed with error  %s", err)
	}

	// make sure it was added correctly
	list := listcommand.List{}
	boxes, err := list.Box.Box("testremovebox")
	if err != nil {
		t.Fatalf("remove box set up failed. list box failed with error %s", err)
	}
	var found bool
	for _, box := range boxes {
		if box.Name == "testremovebox" {
			found = true
		}
	}
	if !found {
		t.Fatal("remove box set up failed. unable to find test box after add")
	}

	// now we can remove it
	b := box{}
	_, err = b.Box("testremovebox")
	if err != nil {
		t.Fatalf("remove box failed with error %s", err)
	}

	// make sure it was removed correctly
	boxes, err = list.Box.Box("")
	if err != nil {
		t.Fatalf("remove box check failed. list box failed with error %s", err)
	}
	for _, box := range boxes {
		if box.Name == "testremovebox" {
			t.Fatal("remove box check failed. did not expect to find text box after remove")
		}
	}

}
