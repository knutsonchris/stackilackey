package remove

import (
	"fmt"

	"github.td.teradata.com/ck250037/stackilackey/cmd"
)

type network struct {
}

/*
Network will remove a network definition from the system. If there are still nodes defined in the database that are assigned to the network name you
are trying to remove, the command will not remove the network definition and print a message saying it cannot remove the network.
*/
func (n *network) Network(networkName string) ([]byte, error) {
	c := fmt.Sprintf("remove network %s", networkName)
	return cmd.RunCommand(c)
}
