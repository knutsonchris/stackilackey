package listcommand

import "testing"

func TestFirewall_Firewall(t *testing.T) {
	// run a list command for global firewalls and make sure that there are some returned
	f := Firewall{}
	firewalls, err := f.Firewall()
	if err != nil {
		t.Fatalf("list firewall failed with error %s", err)
	}
	for _, firewall := range firewalls {
		if firewall.Name == "SSH" {
			return
		}
	}
	t.Errorf("list firewall failed. unable to find expected SSH firewall rule")
}
