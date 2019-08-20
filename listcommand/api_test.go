package listcommand

import (
	"testing"
)

func TestAPI_BlacklistCommand(t *testing.T) {
	// run a list blacklist command and make sure we get back a standard one
	a := API{}
	commands, err := a.BlacklistCommand()
	if err != nil {
		t.Fatalf("list api blacklist command failed with error %s", err)
	}
	if commands[0].Owner != "None" {
		t.Errorf("list api blacklist command failed. expected the default owner to be `none`, got %s", commands[0].Owner)
	}
}

func TestAPI_Group(t *testing.T) {
	// run a list api group command and make sure the default group exists
	a := API{}
	groups, err := a.Group()
	if err != nil {
		t.Fatalf("list api group failed with error %s", err)
	}
	for _, group := range groups {
		if group.Group == "default" {
			return
		}
	}
	t.Errorf("list api group failed. Expected to find a `default` API group but none was found")
}

func TestAPI_GroupPerms(t *testing.T) {
	// run a list api group perms command and look for the default group
	a := API{}
	perms, err := a.GroupPerms("default")
	if err != nil {
		t.Fatalf("list api group perms default failed with error %s", err)
	}
	if perms[0].Group != "default" {
		t.Errorf("list api group perms failed. listed default and got %s", perms[0].Group)
	}
}

func TestAPI_SudoCommand(t *testing.T) {
	// run a list api sudo command and make sure the defaults have no owner
	a := API{}
	commands, err := a.SudoCommand()
	if err != nil {
		t.Fatalf("list api sudo command failed with error %s", err)
	}
	if commands[0].Owner != "None" {
		t.Errorf("list api sudo command failed. expected the commands to have no owner, got %s", commands[0].Owner)
	}
}

func TestAPI_User(t *testing.T) {
	// run a list api user command and make sure there is an admin in the list
	a := API{}
	users, err := a.User()
	if err != nil {
		t.Fatalf("list api user failed with error %s", err)
	}
	for _, user := range users {
		if user.Username == "admin" {
			return
		}
	}
	t.Errorf("list api user failed. expected to find `admin` user but none found")
}

func TestApi_UserPerms(t *testing.T) {
	// run a list api user perms command for a specific user and ensure that the user is in the list
	a := API{}
	perms, err := a.UserPerms("admin")
	if err != nil {
		t.Fatalf("list api user perms failed with error %s", err)
	}
	if perms[0].User != "admin" {
		t.Errorf("list api user perms failed. listed admin but got %s", perms[0].User)
	}
}
