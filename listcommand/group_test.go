package listcommand

import (
	"testing"

	"github.td.teradata.com/ck250037/stackilackey/add"
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
	for _, group := range groups {
		if group.GroupName == "testGroup" {
			return
		}
	}
	t.Errorf("list group failed. unable to find  expected group testGroup")
}
