package listcommand

import "testing"

func TestOS_OS(t *testing.T) {
	// run a list OS command and make sure there is at least one
	o := OS{}
	OSList, err := o.OS("")
	if err != nil {
		t.Fatalf("list os failed with error %s", err)
	}
	if len(OSList) == 0 {
		t.Error("list os failed. unable  to find any OSes on the frontend")
	}
}
