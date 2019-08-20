package remove

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.td.teradata.com/ck250037/stackilackey/add"
	"github.td.teradata.com/ck250037/stackilackey/listcommand"
)

func TestController_Controller(t *testing.T) {
	// before we can test the removal of contoller we first need to add one
	add := add.Add{}
	_, err := add.Storage.Controller(99, 98, 97, 5, "96")
	if err != nil {
		t.Fatalf("remove storage controller set up failed. add storage controller failed with error %s", err)
	}
	// make sure it was added correctly
	list := listcommand.List{}
	controllers, err := list.StorageController.Controller()
	if err != nil {
		t.Fatalf("remove storage controller set up failed. list storage controller failed with error %s", err)
	}
	expected := listcommand.Controller{
		ArrayID:   "99",
		Adapter:   "98",
		Enclosure: "97",
		RaidLevel: "5",
		Slot:      "96",
	}
	var found bool
	for _, controller := range controllers {
		if cmp.Equal(controller, expected) {
			found = true
		}
	}
	if !found {
		t.Fatalf("remove storage controller set up failed. unable to find expected\n%+v\n controller in list of controllers\n%+v", expected, controllers)
	}

	// now we can remove it
	c := controller{}
	_, err = c.Controller("96")
	if err != nil {
		t.Fatalf("remove storage controller failed with error %s", err)
	}

	// make sure it was removed correctly
	controllers, err = list.StorageController.Controller()
	if err != nil {
		t.Fatalf("remove storage controller check failed. list storage controller failed with error %s", err)
	}
	for _, controller := range controllers {
		if cmp.Equal(controller, expected) {
			t.Fatal("remove storage controller faield. did not expect to find test controller after remove")
		}
	}
}
