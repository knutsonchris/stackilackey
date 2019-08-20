package remove

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.td.teradata.com/ck250037/stackilackey/add"
	"github.td.teradata.com/ck250037/stackilackey/listcommand"
)

func TestBootaction_Bootaction(t *testing.T) {
	// before we can test the removal of a bootaction we need to add one
	add := add.Add{}
	_, err := add.Bootaction.Bootaction("testremovebootaction", "test_arg=yada", "examplekernel", "redhat", "exampleramdisk", "install")
	if err != nil {
		t.Fatalf("remove bootaction set up failed. add bootaction failed with error %s", err)
	}

	// make sure it was added correctly
	list := listcommand.List{}
	actions, err := list.Bootaction.Bootaction()
	if err != nil {
		t.Fatalf("remove bootaction set up failed. list bootaction failed with error %s", err)
	}
	expected := listcommand.Bootaction{
		BootactionName: "testremovebootaction",
		TypeName:       "install",
		OS:             "redhat",
		Kernel:         "examplekernel",
		Ramdisk:        "exampleramdisk",
		Args:           "test_arg=yada",
	}
	var found bool
	for _, action := range actions {
		if cmp.Equal(action, expected) {
			found = true
		}
	}
	if !found {
		t.Fatal("remove bootaction set up failed. unable to find test bootaction after add")
	}

	// now we can remove it
	b := bootaction{}
	_, err = b.Bootaction("testremovebootaction", "redhat", "install")
	if err != nil {
		t.Fatalf("remove bootaction failed with error %s", err)
	}

	// make sure it was removed correctly
	actions, err = list.Bootaction.Bootaction()
	if err != nil {
		t.Fatalf("remove bootaction check failed. list bootaction failed with error %s", err)
	}
	for _, action := range actions {
		if cmp.Equal(action, expected) {
			t.Fatal("remove bootaction check failed. did not expect to find test bootaction after remove")
		}
	}
}
