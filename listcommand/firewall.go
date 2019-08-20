package listcommand

import (
	"encoding/json"
	"fmt"

	"github.td.teradata.com/ck250037/stackilackey/cmd"
)

// Firewall representus a firewall rule stored in the stacki database
type Firewall struct {
	Name          string `json:"name"`
	Table         string `json:"table"`
	Service       string `json:"service"`
	Protocol      string `json:"protocol"`
	Chain         string `json:"chain"`
	Action        string `json:"action"`
	Network       string `json:"network"`
	OutputNetwork string `json:"output-network"`
	Flags         string `json:"flags"`
	Comment       string `json:"comment"`
	Source        string `json:"source"`
	Type          string `json:"type"`
}

/*
Firewall will list the set of global firewall rules
*/
func (firewall *Firewall) Firewall() ([]Firewall, error) {
	c := fmt.Sprintf("list firewall")
	b, err := cmd.RunCommand(c)
	if err != nil {
		return nil, err
	}
	firewalls := []Firewall{}
	err = json.Unmarshal(b, &firewalls)
	if err != nil {
		// it may have been just an empty output from the Frontend
		nullOutput := NullOutput{}
		err = json.Unmarshal(b, &nullOutput)
		if err != nil {
			// if we still can't recognize the output, return an error
			return nil, err
		}
		return firewalls, err
	}
	return firewalls, err
}
