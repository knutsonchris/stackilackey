package set

import (
	"fmt"

	"github.com/knutsonchris/stackilackey/cmd"
)

type environment struct {
}

/*
Attr will set an attribute to an environment and set the associated values.
Arguments

	{environment ...}

	Name of environment


Parameters

	{attr=string}

	Name of the attribute

	{value=string}

	Value of the attribute

	[shadow=boolean]

	If set to true, then set the 'shadow' value (only readable by root
	and apache).
*/
func (environment *environment) Attr(environmentName, attr, value string, shadow bool) ([]byte, error) {
	var shadowstr string
	if shadow == true {
		shadowstr = "true"
	} else {
		shadowstr = "false"
	}
	a := []interface{}{environmentName, attr, value, shadowstr}
	c := fmt.Sprintf("set environment attr %s attr='%s' value='%s' shadow='%s'", a...)
	return cmd.RunCommand(c)
}
