package add

import (
	"fmt"

	"github.com/knutsonchris/stackilackey/cmd"
)

type os struct {
}

/*
Attr will add an attribute to an os and set the associated values.
Arguments

	[os ...]

	Name of os


Parameters

	{attr=string}

	Name of the attribute

	{value=string}

	Value of the attribute

	[shadow=boolean]

	If set to true, then set the 'shadow' value (only readable by root
	and apache).
*/
func (os *os) Attr(osName, attr, value string, shadow bool) ([]byte, error) {
	var shadowstr string
	if shadow == true {
		shadowstr = "true"
	} else {
		shadowstr = "false"
	}

	argKeys := []string{"attr", "value", "shadow"}
	argValues := []interface{}{attr, value, shadowstr}
	baseCommand := fmt.Sprintf("add os attr %s", osName)

	c, err := cmd.ArgsExpander(baseCommand, argKeys, argValues)
	if err != nil {
		return nil, err
	}

	return cmd.RunCommand(c)
}

/*
Firewall will add a firewall sule fot an OS type.
Arguments

	{os ...}

	OS type (e.g., 'linux', 'sunos').


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
func (os *os) Firewall(osName, action, chain, protocol, service, comment, flags, network, outputNetwork, rulename, table string) ([]byte, error) {

	argKeys := []string{"action", "chain", "protocol", "service", "comment", "flags", "network", "output-network", "rulename", "table"}
	argValues := []interface{}{action, chain, protocol, service, comment, flags, network, outputNetwork, rulename, table}
	baseCommand := fmt.Sprintf("add os firewall %s", osName)

	c, err := cmd.ArgsExpander(baseCommand, argKeys, argValues)
	if err != nil {
		return nil, err
	}

	return cmd.RunCommand(c)
}

/*
Route will add a route for an OS type.
Arguments

	{os ...}

	The OS type (e.g., 'linux', 'sunos', etc.). This argument is required.


Parameters

	{address=string}

	Host or network address

	{gateway=string}

	Network (e.g., IP address), subnet name (e.g., 'private', 'public'), or
	a device gateway (e.g., 'eth0').

	[interface=string]

	Specific interface to send traffic through. Should only be used if
	you need traffic to go through a VLAN interface (e.g., 'eth0.1').

	[netmask=string]

	Specifies the netmask for a network route.  For a host route
	this is not required and assumed to be 255.255.255.255
*/
func (os *os) Route(osName, address, gateway, interfaceName, netmask string) ([]byte, error) {

	argKeys := []string{"address", "gateway", "interface", "netmask"}
	argValues := []interface{}{address, gateway, interfaceName, netmask}
	baseCommand := fmt.Sprintf("add os route %s", osName)

	c, err := cmd.ArgsExpander(baseCommand, argKeys, argValues)
	if err != nil {
		return nil, err
	}

	return cmd.RunCommand(c)
}

/*
StorageController will add a storage controller configuration for an OS type.
Arguments

	{os ...}

	OS type (e.g., 'redhat', 'sles').


Parameters

	{arrayid=string}

	The 'arrayid' is used to determine which disks are grouped as part
	of the same array. For example, all the disks with arrayid of '1' will
	be part of the same array. Arrayids must be integers starting at 1
	or greater. If the arrayid is 'global', then 'hotspare' must
	have at least one slot definition (this is how one specifies a global
	hotspare).

	In addition, the arrays will be created in arrayid order, that is,
	the array with arrayid equal to 1 will be created first, arrayid
	equal to 2 will be created second, etc.

	[adapter=integer]

	Adapter address.

	[enclosure=integer]

	Enclosure address.

	[hotspare=integer]

	Slot address(es) of the hotspares associated with this array id. This
	can be a comma-separated list (like the 'slot' parameter). If the
	'arrayid' is 'global', then the specified slots are global hotspares.

	[raidlevel=integer]

	RAID level. Raid 0, 1, 5, 6, 10, 50, 60 are currently supported.

	[slot=integer]

	Slot address(es). This can be a comma-separated list meaning all disks
	in the specified slots will be associated with the same array
*/
// TODO: allow slot to be a list of ints for multiple slots
// TODO: figure out what to do with hotspare
func (os *os) StorageController(osName, arrayID string, adapter, enclosure, raidlevel, slot int) ([]byte, error) {

	argKeys := []string{"arrayid", "adapter", "enclosure", "raidlevel", "slot"}
	argValues := []interface{}{arrayID, adapter, enclosure, raidlevel, slot}
	baseCommand := fmt.Sprintf("add os storage controller %s", osName)

	c, err := cmd.ArgsExpander(baseCommand, argKeys, argValues)
	if err != nil {
		return nil, err
	}

	return cmd.RunCommand(c)
}
