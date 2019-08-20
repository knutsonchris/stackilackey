package listcommand

import (
	"encoding/json"
	"fmt"

	"github.td.teradata.com/ck250037/stackilackey/cmd"
)

// API represents the list api subcommand
type API struct {
}

// APICommand represents a Stacki API command. Used for listing blacklist and sudo commands
type APICommand struct {
	Owner   string `json:"owner"`
	Command string `json:"command"`
}

// APIGroup represents a list of users allowed to access the API framework
type APIGroup struct {
	Group string `json:"group"`
	Users string `json:"users"`
}

// APIGroupPerms represents API group permissions
type APIGroupPerms struct {
	Group   string `json:"group"`
	Command string `json:"commanf"`
}

// APIUser represents a user that has API access on the Stacki Frontend
type APIUser struct {
	Username string `json:"username"`
	Admin    bool   `json:"admin"`
	Groups   string `json:"groups"` // TODO: will this ever turn into a []string? may need to use custom unmarshal
}

// APIUserPerms represents a user and their permissions for the Stacki API
type APIUserPerms struct {
	User    string `json:"user"`
	Command string `json:"command"`
	Source  string `json:"source"`
}

/*
BlacklistCommand will list all the commands on the webservice blacklist. This shows the list of commands
that are not allowed to run by anyone. Not even the admin
*/
func (api *API) BlacklistCommand() ([]APICommand, error) {
	c := "list api blacklist command"
	b, err := cmd.RunCommand(c)
	if err != nil {
		return nil, err
	}
	commands := []APICommand{}
	err = json.Unmarshal(b, &commands)
	if err != nil {
		// it may have been just an empty output from the Frontend
		nullOutput := NullOutput{}
		err = json.Unmarshal(b, &nullOutput)
		if err != nil {
			// if we still can't recognize the output, return an error
			return nil, err
		}
		return commands, err
	}
	return commands, err
}

/*
Group will list app the users allowed to access the API framework
*/
func (api *API) Group() ([]APIGroup, error) {
	c := "list api group"
	b, err := cmd.RunCommand(c)
	if err != nil {
		return nil, err
	}
	groups := []APIGroup{}
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

/*
GroupPerms will list API group permissions
Arguments

	{group}

	List API permissions for one or more groups
*/
func (api *API) GroupPerms(group string) ([]APIGroupPerms, error) {
	c := fmt.Sprintf("list api group perms group='%s'", group)
	b, err := cmd.RunCommand(c)
	if err != nil {
		return nil, err
	}
	groupPerms := []APIGroupPerms{}
	err = json.Unmarshal(b, &groupPerms)
	if err != nil {
		// it may have been just an empty output from the Frontend
		nullOutput := NullOutput{}
		err = json.Unmarshal(b, &nullOutput)
		if err != nil {
			// if we still can't recognize the output, return an error
			return nil, err
		}
		return groupPerms, err
	}
	return groupPerms, err
}

/*
SudoCommand will list all commandss on the webservice sudo list. This shows the list of commands that will be run using sudo
*/
func (api *API) SudoCommand() ([]APICommand, error) {
	c := "list api sudo command"
	b, err := cmd.RunCommand(c)
	if err != nil {
		return nil, err
	}
	commands := []APICommand{}
	err = json.Unmarshal(b, &commands)
	if err != nil {
		// it may have been just an empty output from the Frontend
		nullOutput := NullOutput{}
		err = json.Unmarshal(b, &nullOutput)
		if err != nil {
			// if we still can't recognize the output, return an error
			return nil, err
		}
		return commands, err
	}
	return commands, err
}

/*
User will list all users registered to use the API and the groups they belong to
*/
func (api *API) User() ([]APIUser, error) {
	c := "list api user"
	b, err := cmd.RunCommand(c)
	if err != nil {
		return nil, err
	}
	users := []APIUser{}
	err = json.Unmarshal(b, &users)
	if err != nil {
		// it may have been just an empty output from the Frontend
		nullOutput := NullOutput{}
		err = json.Unmarshal(b, &nullOutput)
		if err != nil {
			// if we still can't recognize the output, return an error
			return nil, err
		}
		return users, err
	}
	return users, err
}

/*
UserPerms will list all user permissions. The permissions are classified as either permissions that
belong to the user, or to those that are inherited from the Group that the user belings to.
*/
func (api *API) UserPerms(user string) ([]APIUserPerms, error) {
	c := fmt.Sprintf("list api user perms user='%s'", user)
	b, err := cmd.RunCommand(c)
	if err != nil {
		return nil, err
	}
	perms := []APIUserPerms{}
	err = json.Unmarshal(b, &perms)
	if err != nil {
		// it may have been just an empty output from the Frontend
		nullOutput := NullOutput{}
		err = json.Unmarshal(b, &nullOutput)
		if err != nil {
			// if we still can't recognize the output, return an error
			return nil, err
		}
		return perms, err
	}
	return perms, err
}
