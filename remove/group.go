package remove

import (
	"fmt"

	"github.com/knutsonchris/stackilackey/cmd"
)

type group struct {
}

/*
Group will remove a group
Groups are generic sets of hosts, they have no semantics other than enforcing a common group membership. Hosts may belong to
zero or more groups, and by default belong to none.

Only groups without member hosts can be removed.
*/
func (g *group) Group(groupName string) ([]byte, error) {
	c := fmt.Sprintf("remove group %s", groupName)
	return cmd.RunCommand(c)
}
