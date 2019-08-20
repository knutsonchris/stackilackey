package listcommand

import (
	"encoding/json"

	"github.td.teradata.com/ck250037/stackilackey/cmd"
)

// Route represents a route stored in the stacki database
type Route struct {
	Network       string `json:"network"`
	Netmask       string `json:"netmask"`
	Gateway       string `json:"gateway"`
	Subnet        string `json:"subnet"`
	InterfaceName string `json:"interface"`
}

/*
Route will list the global routes
*/
func (route *Route) Route() ([]Route, error) {
	c := "list route"
	b, err := cmd.RunCommand(c)
	if err != nil {
		return nil, err
	}
	routes := []Route{}
	err = json.Unmarshal(b, &routes)
	if err != nil {
		// it may have been just an empty output from the Frontend
		nullOutput := NullOutput{}
		err = json.Unmarshal(b, &nullOutput)
		if err != nil {
			// if we still can't recognize the output, return an error
			return nil, err
		}
		return routes, err
	}
	return routes, err
}
