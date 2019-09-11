package add

import (
	"github.td.teradata.com/ck250037/stackilackey/cmd"
)

type attr struct {
}

/*
Attr will add a global attribute for all nodes.
Parameters

	{attr=string}

	Name of the attribute

	{value=string}

	Value of the attribute

	[shadow=boolean]

	If set to true, then set the 'shadow' value (only readable by root
	and apache).
*/
func (attr *attr) Attr(attrName, value string, shadow bool) ([]byte, error) {
	var shadowstr string
	if shadow == true {
		shadowstr = "true"
	} else {
		shadowstr = "false"
	}

	argKeys := []string{"attr", "value", "shadow"}
	argValues := []interface{}{attrName, value, shadowstr}
	baseCommand := "add attr"
	c, err := cmd.ArgsExpander(baseCommand, argKeys, argValues)
	if err != nil {
		return nil, err
	}

	return cmd.RunCommand(c)
}
