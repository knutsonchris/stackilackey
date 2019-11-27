package listcommand

import (
	"encoding/json"
	"fmt"

	"github.com/knutsonchris/stackilackey/cmd"
)

// Box represents a collection of pallets and carts in the Stacki Database
type Box struct {
	Name    string `json:"name"`
	OS      string `json:"os"`
	Pallets string `json:"pallets"`
	Carts   string `json:"carts"`
}

/*
Box will list the pallets and carts that are associated with boxes.
Arguments

	[box ...]

	Optional list of box names.
*/
func (box *Box) Box(boxName string) ([]Box, error) {
	c := fmt.Sprintf("list box %s", boxName)
	b, err := cmd.RunCommand(c)
	if err != nil {
		return nil, err
	}
	boxes := []Box{}
	err = json.Unmarshal(b, &boxes)
	if err != nil {
		// it may have been just an empty output from the Frontend
		nullOutput := NullOutput{}
		err = json.Unmarshal(b, &nullOutput)
		if err != nil {
			// if we still can't recognize the output, return an error
			return nil, err
		}
		return boxes, err
	}
	return boxes, err
}
