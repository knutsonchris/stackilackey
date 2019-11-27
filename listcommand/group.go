package listcommand

import (
	"encoding/json"

	"github.com/knutsonchris/stackilackey/cmd"
)

// Group represents a group and the hosts that belong to it in the Stacki database
type Group struct {
	GroupName string `json:"group"`
	Hosts     string `json:"hosts"`
}

/*
Group will list the current groups and the number of member hosts in each.
*/
func (group *Group) Group() ([]Group, error) {
	c := "list group"
	b, err := cmd.RunCommand(c)
	if err != nil {
		return nil, err
	}
	groups := []Group{}
	err = json.Unmarshal(b, &groups)
	if err != nil {
		// it may have been just an empty output from the Frontend
		nullOutput := NullOutput{}
		err = json.Unmarshal(b, &nullOutput)
		if err != nil {
			// if we still can't recognize the output, return an error
			return nil, err
		}
		return groups, err
	}
	return groups, err
}
