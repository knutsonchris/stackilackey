package set

import (
	"fmt"

	"github.com/knutsonchris/stackilackey/cmd"
)

type access struct {
}

/*
Access will set an access control pattern.
Parameters

	{command=string}

	Command Pattern.

	{group=string}

	Group name / ID for access.
*/
func (access *access) Access(command, group string) ([]byte, error) {
	c := fmt.Sprintf("set access command='%s' group='%s", command, group)
	return cmd.RunCommand(c)
}
