package add

import (
	"fmt"

	"github.td.teradata.com/ck250037/stackilackey/cmd"
)

type group struct {
}

/*
Group will add a group to the databse. Groups are generic sets of hosts, they have no semantics other
than enforcing a common group membership. Hosts may belong to zero or more groups, and by default belong to none.
Arguments

	{group}

	The name of the group to be created.
*/
func (group *group) Group(groupName string) ([]byte, error) {
	c := fmt.Sprintf("add group %s", groupName)
	return cmd.RunCommand(c)
}
