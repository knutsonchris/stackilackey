package remove

import (
	"fmt"

	"github.td.teradata.com/ck250037/stacki-lackey-2/cmd"
)

type route struct {
}

/*
Route will remove a global static route
Parameters

	{address=string}

	The address of the static route to remove.
*/
func (route *route) Route(address string) ([]byte, error) {
	c := fmt.Sprintf("remove route address='%s'", address)
	return cmd.RunCommand(c)
}
