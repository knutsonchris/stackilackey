package add

import (
	"fmt"

	"github.td.teradata.com/ck250037/stackilackey/cmd"
)

type box struct {
}

/*
Box will add a box specification to the database.
Arguments

	{box}

	Name of the new box.


Parameters

	[os=string]

	OS associated with the box. Default is the native os (e.g., 'redhat', 'sles').
*/
func (box *box) Box(boxName, os string) ([]byte, error) {

	argkeys := []string{"os"}
	argvalues := []interface{}{os}
	baseCommand := fmt.Sprintf("add box %s", boxName)

	c, err := cmd.ArgsExpander(baseCommand, argkeys, argvalues)
	if err != nil {
		return nil, err
	}

	return cmd.RunCommand(c)
}
