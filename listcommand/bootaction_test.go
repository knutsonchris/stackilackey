package listcommand

import "testing"

func TestBootaction_Bootaction(t *testing.T) {
	// run a list bootaction command and make sure that the `default` bootaction is there
	b := Bootaction{}
	bootactions, err := b.Bootaction()
	if err != nil {
		t.Errorf("list bootaction failed with error %s", err)
	}
	for _, action := range bootactions {
		if action.BootactionName == "default" {
			return
		}
	}
	t.Errorf("List bootaction failed. Expected a `default` bootaction but it was not found")
}
