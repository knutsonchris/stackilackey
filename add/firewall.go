package add

import (
	"fmt"

	"github.com/knutsonchris/stackilackey/cmd"
)

type firewall struct {
}

/*
Firewall will add a global firewall rule for all the hosts in the cluster
Description

	Add a global firewall rule for the all hosts in the cluster.


Parameters

	{action=string}

	The iptables 'action' this rule should be applied to (e.g.,
	ACCEPT, REJECT, DROP).

	{chain=string}

	The iptables 'chain' this rule should be applied to (e.g.,
	INPUT, OUTPUT, FORWARD).

	{protocol=string}

	The protocol associated with the rule. For example, "tcp" or "udp".

	To have this firewall rule apply to all protocols, specify the
	keyword 'all'.

	{service=string}

	A comma seperated list of service identifier, port number or port range.

	For example "www", 8080, 0:1024, or "1:1024,8080".

	To have this firewall rule apply to all services, specify the keyword 'all'.

	[comment=string]

	A comment associated with this rule. The comment will be printed
	directly above the rule in the firewall configuration file.

	[flags=string]

	Optional flags associated with this rule. An example flag is:
	"-m state --state RELATED,ESTABLISHED".

	[network=string]

	The network this rule should be applied to. This is a named network
	(e.g., 'private') and must be one listed by the command
	'stack list network'.

	By default, the rule will apply to all networks.

	[output-network=string]

	The output network this rule should be applied to. This is a named
	network (e.g., 'private') and must be one listed by the command
	'stack list network'.

	By default, the rule will apply to all networks.

	[rulename=string]

	The rule name for the rule to add. This is the handle by
	which the admin can remove or override the rule.

	[table=string]

	The table to add the rule to. Valid values are 'filter',
	'nat', 'mangle', and 'raw'. If this parameter is not
	specified, it defaults to 'filter'
*/
func (firewall *firewall) Firewall(action, chain, protocol, service, comment, flags, network, outputNetwork, rulename, table string) ([]byte, error) {
	args := []interface{}{action, chain, protocol, service, comment, flags, network, outputNetwork, rulename, table}
	c := fmt.Sprintf("add firewall action='%s' chain='%s' protocol='%s' service='%s' comment='%s' flags='%s' network='%s' output-network='%s' rulename='%s' table='%s'", args...)
	return cmd.RunCommand(c)
}
