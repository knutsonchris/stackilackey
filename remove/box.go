package remove

import (
	"fmt"

	"github.td.teradata.com/ck250037/stacki-lackey-2/cmd"
)

type box struct {
}

/*
Box will remove a box specification from the system
Arguments

	{box}

	A list of boxes to remove.  Boxes must not have any hosts assigned.
*/
func (box *box) Box(boxName string) ([]byte, error) {
	c := fmt.Sprintf("remove box %s", boxName)
	return cmd.RunCommand(c)
}
