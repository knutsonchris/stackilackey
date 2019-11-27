package listcommand

import (
	"testing"

	"github.com/knutsonchris/stackilackey/add"
)

func TestCart_Cart(t *testing.T) {
	// before we run the list cart test we need to add one
	a := add.Add{}
	_, err := a.Cart.Cart("testCart", "", "", "", "", "", "", "", false)
	if err != nil {
		t.Fatalf("list cart set up failed. add cart failed with error %s", err)
	}
	c := Cart{}
	carts, err := c.Cart("testCart")
	if err != nil {
		t.Fatalf("list cart failed with error %s", err)
	}
	for _, cart := range carts {
		if cart.Name == "testCart" {
			// TODO: tear down this add cart test with remove cart
			return
		}
	}
	t.Errorf("list cart failed. unable to find expected cart")
}
