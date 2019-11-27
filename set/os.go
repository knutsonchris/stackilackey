package set

import (
	"fmt"

	"github.com/knutsonchris/stackilackey/cmd"
)

type os struct {
}

/*
Attr will set an attribute to an os and set the associated values.
Arguments

	[os ...]

	Name of os


Parameters

	{attr=string}

	Name of the attribute

	{value=string}

	Value of the attribute

	[shadow=boolean]

	If set to true, then set the 'shadow' value (only readable by root
	and apache).
*/
func (os *os) Attr(osName, attr, value string, shadow bool) ([]byte, error) {
	var shadowstr string
	if shadow == true {
		shadowstr = "true"
	} else {
		shadowstr = "false"
	}
	args := []interface{}{osName, attr, value, shadowstr}
	c := fmt.Sprintf("set os attr %s attr='%s' value='%s' shadow='%s'", args...)
	return cmd.RunCommand(c)
}
