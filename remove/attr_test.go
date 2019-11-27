package remove

import (
	"testing"

	"github.com/knutsonchris/stackilackey/add"
	"github.com/knutsonchris/stackilackey/listcommand"
)

func TestAttr_Attr(t *testing.T) {
	// lets add a new global to not affect other tests
	add := add.Add{}
	_, err := add.Attr.Attr("testRemoveAttr", "testRemoveAttrValue", false)
	if err != nil {
		t.Fatalf("remove attr set up failed. add attr failed with error %s", err)
	}

	// make sure it was added correctly
	list := listcommand.List{}
	attrs, err := list.Attr.Attr("testRemoveAttr", false)
	if err != nil {
		t.Fatalf("remove attr set up failed. list attr failed with error %s", err)
	}
	var found bool
	for _, attr := range attrs {
		if attr.AttrName == "testRemoveAttr" {
			found = true
		}
	}

	if !found {
		t.Fatalf("remove attr set up failed. unable to find test attr after add")
	}

	// now we can remove it
	a := attr{}
	_, err = a.Attr("testRemoveAttr")
	if err != nil {
		t.Fatalf("remove attr failed with error %s", err)
	}

	// make sure it was removed correctly
	attrs, err = list.Attr.Attr("", false)
	if err != nil {
		t.Fatalf("remove attr check failed. list attr failed with error %s", err)
	}
	for _, attr := range attrs {
		if attr.AttrName == "testRemoveAttr" {
			t.Fatal("remove attr check failed. did not expect to find test attr  after remove")
		}
	}
}
