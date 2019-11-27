package set

import (
	"fmt"

	"github.com/knutsonchris/stackilackey/cmd"
)

type appliance struct {
}

/*
Attr will set an attiubute to an appliance and set the asssociated values.
Arguments

	[appliance ...]

	Name of appliance


Parameters

	{attr=string}

	Name of the attribute

	{value=string}

	Value of the attribute

	[shadow=boolean]

	If set to true, then set the 'shadow' value (only readable by root
	and apache).
*/
func (applaince *appliance) Attr(applianceName, attr, value string, shadow bool) ([]byte, error) {
	var shadowstr string
	if shadow == true {
		shadowstr = "true"
	} else {
		shadowstr = "false"
	}
	args := []interface{}{applianceName, attr, value, shadowstr}
	c := fmt.Sprintf("set appliance attr %s attr='%s' value='%s' shadow='%s'", args...)
	return cmd.RunCommand(c)
}
