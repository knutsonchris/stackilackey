package remove

import (
	"testing"

	"github.td.teradata.com/ck250037/stackilackey/add"
	"github.td.teradata.com/ck250037/stackilackey/listcommand"
)

func TestGroup_Group(t *testing.T) {
	// before we can test the removal of a group we need to add one
	add := add.Add{}
	_, err := add.Group.Group("testremovegroup")
	if err != nil {
		t.Fatalf("remove group set up failed. add group failed with error %s", err)
	}

	// make sure it was added correctly
	list := listcommand.List{}
	groups, err := list.Group.Group()
	if err != nil {
		t.Fatalf("remove group set up failed. list group failed with error %s", err)
	}
	var found bool
	for _, group := range groups {
		if group.GroupName == "testremovegroup" {
			found = true
		}
	}
	if !found {
		t.Fatal("remove group set up failed. unable to find test group after add")
	}

	// now we can remove it
	g := group{}
	_, err = g.Group("testremovegroup")
	if err != nil {
		t.Fatalf("remove group failed with error %s", err)
	}

	// make sure it was removed correctly
	groups, err = list.Group.Group()
	if err != nil {
		t.Fatalf("remove group check failed. list group failed with error %s", err)
	}
	for _, group := range groups {
		if group.GroupName == "testremovegroup" {
			t.Fatal("remove group check failed. did not expect to find test group after remove")
		}
	}
}
