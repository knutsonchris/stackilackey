package remove

import (
	"fmt"

	"github.td.teradata.com/ck250037/stackilackey/cmd"
)

type bootaction struct {
}

/*
Bootaction will remove a boot action specification from the system
Arguments

	{action}

	The label name for the boot action. You can see the boot action label
	names by executing: 'stack list bootaction'.


Parameters

	[os=string]

	Specify the 'os' (e.g., 'redhat', 'sles', etc.)

	[type=string]

	The 'type' parameter should be either 'os' or 'install'.
*/
func (b *bootaction) Bootaction(actionName, OS, actionType string) ([]byte, error) {
	argKeys := []string{"os", "type"}
	argValues := []interface{}{OS, actionType}
	baseCommand := fmt.Sprintf("remove bootaction %s", actionName)

	c, err := cmd.ArgsExpander(baseCommand, argKeys, argValues)
	if err != nil {
		return nil, err
	}

	return cmd.RunCommand(c)
}
