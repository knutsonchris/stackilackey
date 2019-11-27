package remove

import (
	"fmt"

	"github.com/knutsonchris/stackilackey/cmd"
)

type firewall struct {
}

/*
Firewall will remove a global firewall rule. To remove a rule, you must supply the name of a rule.
Parameters

	{rulename=string}

	Name of the rule
*/
func (firewall *firewall) Firewall(ruleName string) ([]byte, error) {
	c := fmt.Sprintf("remove firewall rulename='%s'", ruleName)
	return cmd.RunCommand(c)
}
