package add

import (
	"fmt"

	"github.td.teradata.com/ck250037/stacki-lackey-2/cmd"
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
	args := []interface{}{attrName, value, shadowstr}
	c := fmt.Sprintf("add attr attr='%s' value='%s' shadow='%s'", args...)
	return cmd.RunCommand(c)
}
