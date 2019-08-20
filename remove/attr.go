package remove

import (
	"fmt"

	"github.td.teradata.com/ck250037/stacki-lackey-2/cmd"
)

type attr struct {
}

/*
Attr will remove a glbal attribute
Parameters

	{attr=string}

	The attribute name that should be removed.
*/
func (attr *attr) Attr(attrName string) ([]byte, error) {
	c := fmt.Sprintf("remove attr attr='%s'", attrName)
	return cmd.RunCommand(c)
}
