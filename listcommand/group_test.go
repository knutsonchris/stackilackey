package listcommand

import (
	"testing"

	"github.com/knutsonchris/stackilackey/add"
	"github.com/knutsonchris/stackilackey/remove"
)

func TestGroup_Group(t *testing.T) {
	// before we run the list group test we need to add one
	a := add.Add{}
	_, err := a.Group.Group("testGroup")
	if err != nil {
		t.Fatalf("list group set up failed. add group failed with error %s", err)
	}
	g := Group{}
	groups, err := g.Group()
	if err != nil {
		t.Fatalf("list group failed with error %s", err)
	}
	found := false
	for _, group := range groups {
		if group.GroupName == "testGroup" {
			found = true
		}
	}
	if !found {
		t.Fatal("list group failed. unable to find  expected group testGroup")
	}

	// to clean up after ourselves, remove the test group
	remove := remove.Remove{}
	removeGroup := remove.Group

	_, err = removeGroup.Group("testGroup")
	if err != nil {
		t.Errorf("list group failed. unable to remove test group after list test")
	}
}
