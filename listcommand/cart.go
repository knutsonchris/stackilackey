package listcommand

import (
	"encoding/json"
	"fmt"

	"github.com/knutsonchris/stackilackey/cmd"
)

// Cart represents a cart in the Stacki database
type Cart struct {
	Name  string `json:"name"`
	Boxes string `json:"boxes"`
}

/*
Cart will return a list of carts. This should be the cart base name (e.g., stacki, os, etc.). If no carts are listed, then status for all the carts are listed.
Parameters

	{expanded=string}

	Include the source url of the cart.
*/
func (cart *Cart) Cart(cartName string) ([]Cart, error) {
	c := fmt.Sprintf("list cart %s", cartName)
	b, err := cmd.RunCommand(c)
	if err != nil {
		return nil, err
	}
	carts := []Cart{}
	err = json.Unmarshal(b, &carts)
	if err != nil {
		// it may have been just an empty output from the Frontend
		nullOutput := NullOutput{}
		err = json.Unmarshal(b, &nullOutput)
		if err != nil {
			// if we still can't recognize the output, return an error
			return nil, err
		}
		return carts, err
	}
	return carts, err
}
