package add

import (
	"fmt"

	"github.td.teradata.com/ck250037/stackilackey/cmd"
)

type api struct {
}

// BlacklistCommand will add a command to the webservice blacklist. This disallows the command from running, by anyone, including the admin. This has the
// granularity of a command. This means that you can only blacklist individual commands, and not entire verbs of commands.
func (api *api) BlacklistCommand(command string) ([]byte, error) {
	c := fmt.Sprintf("add api blacklist command command='%s'", command)
	return cmd.RunCommand(c)
}

// Group will add a group to the REST API Group list
func (api *api) Group(group string) ([]byte, error) {
	c := fmt.Sprintf("add api group %s", group)
	return cmd.RunCommand(c)
}

// GroupPerms will add permissions to an API group
func (api *api) GroupPerms(group string, perm string) ([]byte, error) {
	c := fmt.Sprintf("add api group perms %s perm='%s'", group, perm)
	return cmd.RunCommand(c)
}

// Sudo Command will add Add a command, or a set of commands, to the webservice sudo list. This allows the webservice to sudo up to
// root to run the commands. It can take a regular expression or an individual command.
func (api *api) SudoCommand(command string) ([]byte, error) {
	c := fmt.Sprintf("add api sudo command command='%s'", command)
	return cmd.RunCommand(c)
}

// User will Create a user to the REST API. This command will print out a JSON string
// that contains the Username, API Key, and Hostname of the API server.
func (api *api) User(username string, admin bool, group string) ([]byte, error) {
	var adminstr string
	if admin == true {
		adminstr = "true"
	} else {
		adminstr = "false"
	}
	c := fmt.Sprintf("add api user %s admin='%s' group='%s'", username, adminstr, group)
	return cmd.RunCommand(c)
}

// UserGroup will add users to an existing group
func (api *api) UserGroup(user string, group string) ([]byte, error) {
	c := fmt.Sprintf("add api user group %s group=%s", user, group)
	return cmd.RunCommand(c)
}

// UserPerms will set permissions for an API user
func (api *api) UserPerms(user string, perm string) ([]byte, error) {
	c := fmt.Sprintf("add api user perms %s perm='%s'", user, perm)
	return cmd.RunCommand(c)
}
