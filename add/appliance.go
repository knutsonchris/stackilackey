package add

import (
	"fmt"

	"github.com/knutsonchris/stackilackey/cmd"
)

type appliance struct {
}

/*
Appliance will add an appliance specification to the database.
Arguments

	{appliance}

	The appliance name (e.g., 'backend', 'frontend', 'nas').


Parameters

	[node=string]

	The name of the root XML node (e.g., 'backend', 'nas'). If
	not supplied, the node name is set to the appliance name.

	[public=bool]

	True means this appliance will be displayed by 'insert-ethers' in
	the Appliance menu. The default is 'yes'.
*/
func (appliance *appliance) Appliance(applianceName, node string, public bool) ([]byte, error) {
	var publicstr string
	if public == true {
		publicstr = "true"
	} else {
		publicstr = "false"
	}
	c := fmt.Sprintf("add appliance %s node='%s' public='%s'", applianceName, node, publicstr)
	return cmd.RunCommand(c)
}

/*
Attr will add an attribute to an appliance and set the associated values.
Arguments

	[appliance ...]

	Name of appliance


Parameters

	{attr=string}

	Name of the attribute

	{value=string}

	Value of the attribute

	[shadow=boolean]

	If set to true, then set the 'shadow' value (only readable by root
	and apache).
*/
func (appliance *appliance) Attr(applianceName, attr, value string, shadow bool) ([]byte, error) {
	var shadowstr string
	if shadow == true {
		shadowstr = "true"
	} else {
		shadowstr = "false"
	}
	c := fmt.Sprintf("add appliance attr %s attr='%s' value='%s' shadow='%s'", applianceName, attr, value, shadowstr)
	return cmd.RunCommand(c)
}

/*
Firewall will add a firewall rule for an appliance type
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
func (appliance *appliance) Firewall(applianceName, action, chain, protocol, service, comment, flags, network, outputNetwork, rulename, table string) ([]byte, error) {
	args := []interface{}{applianceName, action, chain, protocol, service, comment, flags, network, outputNetwork, rulename, table}
	c := fmt.Sprintf("add appliance firewall %s action='%s' chain='%s' protocol='%s' service='%s' comment='%s' flags='%s' network='%s' output-network='%s' rulename='%s' table='%s'", args...)
	return cmd.RunCommand(c)
}

/*
Route will add a route for an appliance type in the cluster.
Arguments

	{appliance ...}

	The appliance type (e.g., 'backend').


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
func (appliance *appliance) Route(applianceName, address, gateway, interfaceName, netmask string) ([]byte, error) {
	args := []interface{}{applianceName, address, gateway, interfaceName, netmask}
	c := fmt.Sprintf("add appliance route %s address='%s' gateway='%s' interface='%s' netmask='%s'", args...)
	return cmd.RunCommand(c)
}

/*
StorageController will add a storage controller configuration for an appliance type.
Arguments

	{appliance ...}

	Appliance type (e.g., "backend").


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
// TODO: the api refuses to accept an integer, must be a string
func (appliance *appliance) StorageController(applianceName string, arrayID, adapter, enclosure, raidlevel int, slot string) ([]byte, error) {
	args := []interface{}{applianceName, arrayID, adapter, enclosure, raidlevel, slot}
	c := fmt.Sprintf("add appliance storage controller %s arrayid='%d' adapter='%d' enclosure='%d' raidlevel='%d' slot='%s'", args...)
	return cmd.RunCommand(c)
}
