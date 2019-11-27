package remove

import (
	"fmt"

	"github.com/knutsonchris/stackilackey/cmd"
)

type cart struct {
}

/*
Cart will  remove a cart from both the database and the filesystem
Arguments

	{cart ...}

	List of carts.
*/
func (cart *cart) Cart(cartName string) ([]byte, error) {
	c := fmt.Sprintf("remove cart %s", cartName)
	return cmd.RunCommand(c)
}
