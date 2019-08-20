package set

import (
	"fmt"

	"github.td.teradata.com/ck250037/stacki-lackey-2/cmd"
)

type api struct {
}

/*
UserAdmin will set or unset admin priveleges of a user.
Arguments

	{Username}

	Username of user for which to set / unset
	the admin privileges.


Parameters

	[admin=string]

	Set or unset admin privileges.
*/
func (api *api) UserAdmin(user, admin string) ([]byte, error) {
	c := fmt.Sprintf("set api user admin %s admin='%s'", user, admin)
	return cmd.RunCommand(c)
}
