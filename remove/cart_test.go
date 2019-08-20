package remove

import (
	"testing"

	"github.td.teradata.com/ck250037/stackilackey/add"
	"github.td.teradata.com/ck250037/stackilackey/listcommand"
)

func TestCart_Cart(t *testing.T) {
	// before we can test the removal of a cart we need to add one
	add := add.Add{}
	_, err := add.Cart.Cart("testremovecart", "", "", "", "", "", "", "", true)
	if err != nil {
		t.Fatalf("remove cart set up failed. add cart failed with error %s", err)
	}

	// make sure it was added correctly
	list := listcommand.List{}
	carts, err := list.Cart.Cart("testremovecart")
	if err != nil {
		t.Fatalf("remove cart set up failed. list cart failed with error %s", err)
	}
	var found bool
	for _, cart := range carts {
		if cart.Name == "testremovecart" {
			found = true
		}
	}
	if !found {
		t.Fatal("remove cart set up failed. unable to find expected cart after add")
	}

	// now we can remove it
	c := cart{}
	_, err = c.Cart("testremovecart")
	if err != nil {
		t.Fatalf("remove cart failed with error %s", err)
	}

	// make sure it was removed correctly
	carts, err = list.Cart.Cart("")
	if err != nil {
		t.Fatalf("remove cart check failed. list cart failed with error %s", err)
	}
	for _, cart := range carts {
		if cart.Name == "testremovecart" {
			t.Fatalf("test remove cart check failed. did not expect to find test cart after remove")
		}
	}
}
