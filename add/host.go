package add

import (
	"fmt"

	"github.td.teradata.com/ck250037/stacki-lackey-2/cmd"
)

type host struct {
}

/*
Host will add a new host to the cluster.
Arguments

	{host}

	A single host name.  If the hostname is of the standard form of
	basename-rack-rank the default values for the appliance, rack,
	and rank parameters are taken from the hostname.


Parameters

	[box=string]

	The box name for the host. The default is: "default".

	[environment=string]

	Name of the host environment.  For most users this is not specified.
	Environments allow you to partition hosts into logical groups.

	[rack=string]

	The number of the rack where the machine is located. The convention
	in Stacki is to start numbering at 0. If not provided and the host
	name is of the standard form the rack number is taken from the host
	name.

	[rank=string]

	The position of the machine in the rack. The convention in Stacki
	is to number from the bottom of the rack to the top starting at 0.
	If not provided and the host name is of the standard form the rank
	number is taken from the host name.
*/
func (host *host) Host(hostName, appliance, box, environment, rack, rank string) ([]byte, error) {

	argKeys := []string{"appliance", "box", "environment", "rack", "rank"}
	argValues := []interface{}{appliance, box, environment, rack, rank}
	baseCommand := fmt.Sprintf("add host %s", hostName)

	c, err := cmd.ArgsExpander(baseCommand, argKeys, argValues)
	if err != nil {
		return nil, err
	}

	return cmd.RunCommand(c)
}

/*
Alias adds an alias to a host.
Arguments

	{host}

	Host name of machine


Parameters

	{alias=string}

	Alias for the host.

	{interface=string}

	Interface of host.
*/
func (host *host) Alias(hostName, alias, interfaceName string) ([]byte, error) {
	args := []interface{}{hostName, alias, interfaceName}
	c := fmt.Sprintf("add host alias %s alais='%s' interface='%s'", args...)
	return cmd.RunCommand(c)
}

/*
Attrnwill add an attribute to a host and set the associated values.
Arguments

	[host ...]

	Host name of machine


Parameters

	{attr=string}

	Name of the attribute

	{value=string}

	Value of the attribute

	[shadow=boolean]

	If set to true, then set the 'shadow' value (only readable by root
	and apache).
*/
func (host *host) Attr(hostName, attr, value string, shadow bool) ([]byte, error) {
	var shadowstr string
	if shadow == true {
		shadowstr = "true"
	} else {
		shadowstr = "false"
	}
	c := fmt.Sprintf("add host attr %s attr='%s' value='%s' shadow='%s'", hostName, attr, value, shadowstr)
	return cmd.RunCommand(c)
}

/*
Bonded adds a channel bonded interface for a host.
Arguments

	{host}

	Host name of machine


Parameters

	[channel=string]

	The channel name (e.g., "bond0").

	[interfaces=string]

	The physical interfaces that will be bonded. The interfaces
	can be a comma-separated list (e.g., "eth0,eth1") or a space-separated
	list (e.g., "eth0 eth1").

	[ip=string]

	The IP address to assign to the bonded interface.

	[name=string]

	The host name associated with the bonded interface. If name is not
	specified, then the interface get the internal host name
	(e.g., backend-0-0).

	[network=string]

	The network to be assigned to this interface. This is a named network
	(e.g., 'private') and must be listable by the command
	'rocks list network'.

	[options=string]

	Bonding Options. These are applied to the bonding device
	as BONDING_OPTS in the ifcfg-bond* files.
*/
func (host *host) Bonded(hostName, channel, interfaces, ip, name, network, options string) ([]byte, error) {
	args := []interface{}{hostName, channel, interfaces, ip, name, network, options}
	c := fmt.Sprintf("add host bonded %s channel='%s' interfaces='%s' ip='%s' name='%s' network='%s' options='%s'", args...)
	return cmd.RunCommand(c)
}

/*
Bridge will add a bridge interface to a given host.
Arguments

	{host}

	Hostname


Parameters

	[interface=string]

	Physical interface to be bridged

	[name=string]

	Name for the bridge interface.

	[network=string]

	Name of the network on which the physical
	device to be bridged exists.
*/
func (host *host) Bridge(hostName, interfaceName, name, network string) ([]byte, error) {
	args := []interface{}{hostName, interfaceName, name, network}
	c := fmt.Sprintf("add host bridge %s interface='%s' name='%s' network='%s'", args...)
	return cmd.RunCommand(c)
}

/*
Firewall will add a firewall rule for the specified hosts.
Arguments

	{host ...}

	Host name of machine


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
func (host *host) Firewall(hostName, action, chain, protocol, service, comment, flags, network, outputNetwork, rulename, table string) ([]byte, error) {
	args := []interface{}{hostName, action, chain, protocol, service, comment, flags, network, outputNetwork, rulename, table}
	c := fmt.Sprintf("add host firewall %s action='%s' chain='%s' protocol='%s' service='%s' comment='%s' flags='%s' network='%s' output-network='%s' rulename='%s' table=%s", args...)
	return cmd.RunCommand(c)
}

/*
Group will add a group to one or more hosts.
Arguments

	{host ...}

	One or more host names.


Parameters

	{group=string}

	Group for the host.
*/
func (host *host) Group(hostName, group string) ([]byte, error) {
	c := fmt.Sprintf("stack add host group %s group='%s'", hostName, group)
	return cmd.RunCommand(c)
}

/*
Interface will add an interface to a host and set the associated values.
Arguments

	{host}

	Host name of machine


Parameters

	[channel=string]

	The channel for an interface.

	[default=boolean]

	If true, the name associated with this interface becomes the hostname
	and the interface's gateway becomes the default gateway.

	[interface=string]

	The interface name on the host (e.g., 'eth0', 'eth1')

	[ip=string]

	The IP address to assign to the interface (e.g., '192.168.1.254')

	[mac=string]

	The MAC address of the interface (e.g., '00:11:22:33:44:55')

	[module=string]

	The device driver name (or module) of the interface (e.g., 'e1000')

	[name=string]

	The name to assign to the interface

	[network=string]

	The name of the network to assign to the interface (e.g., 'private')

	[vlan=string]

	The VLAN ID to assign the interface
*/
func (host *host) Interface(hostName, channel, interfaceName, ip, mac, module, name, network, vlan string, defaultFlag bool) ([]byte, error) {
	var defaultstr string
	if defaultFlag == true {
		defaultstr = "true"
	} else {
		defaultstr = "false"
	}
	args := []interface{}{hostName, channel, defaultstr, interfaceName, ip, mac, module, name, network, vlan}
	c := fmt.Sprintf("add host interface %s channel='%s' default='%s' interface='%s' ip='%s' mac='%s' module='%s' name='%s' network='%s' vlan='%s", args...)
	return cmd.RunCommand(c)
}

/*
Key will add a public key for a host. One use of this public key is to authenticate messages sent from remote services.
Arguments

	{host}

	Host name of machine


Parameters

	[key=string]

	A public key. This can be the actual key or it can be a path name to
	a file that contains a public key (e.g., /tmp/public.key).
*/
func (host *host) Key(hostName, key string) ([]byte, error) {
	c := fmt.Sprintf("add host key %s key='%s", hostName, key)
	return cmd.RunCommand(c)
}

/*
Message will add a message to one or more host message queues.
Arguments

	{host ...}

	Zero, one or more host names. If no host names are supplied, the
	message is sent to all hosts.


Parameters

	{message=string}

	Message text

	[channel=string]

	Name of the channel
*/
func (host *host) Message(hostName, message, channel string) ([]byte, error) {
	args := []interface{}{hostName, message, channel}
	c := fmt.Sprintf("add host message %s message='%s' channel='%s'", args...)
	return cmd.RunCommand(c)
}

/*
Partition will add partitioning information to the database.
Arguments

	{host ...}

	Hostname


Parameters

	[device=string]

	Device to be added. For example, sdb, sdb1, etc.

	[formatflags=string]

	Flags used for formatting the partition

	[fs=string]

	File system type of partition. For example, "ext4", "xfs"

	[mountpoint=string]

	Mount point for the device. For example, "/state/partition1", "/hadoop01".

	[partid=string]

	ID of partition

	[partitionflags=string]

	Flags used for partitioning

	[sectorstart=string]

	Starting sector of partition

	[size=string]

	Size of partition

	[uuid=string]

	UUID for the partition
*/
func (host *host) Partition(hostName, device, formatFlags, fs, mountPoint, partID, partitionFlags, sectorStart, size, uuid string) ([]byte, error) {
	args := []interface{}{hostName, device, formatFlags, fs, mountPoint, partID, partitionFlags, sectorStart, size, uuid}
	c := fmt.Sprintf("add host partition %s device='%s' formatflags='%s' fs='%s' mountpoint='%s' partid='%s' partitionflags='%s' sectorstart='%s' size='%s' uuid='%s'", args...)
	return cmd.RunCommand(c)
}

/*
Route will add a route for a host.
Arguments

	{host ...}

	Host name of machine


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

	[syncnow=string]

	Add route to the routing table immediately
*/
func (host *host) Route(hostName, address, gateway, interfaceName, netmask, syncnow string) ([]byte, error) {
	args := []interface{}{hostName, address, gateway, interfaceName, netmask, syncnow}
	c := fmt.Sprintf("add host route %s address='%s' gateway='%s' interface='%s' netmask='%s' syncnow='%s'", args...)
	return cmd.RunCommand(c)
}

/*
StorageController will add a storage controller configuration for the specified hosts.
Arguments

	{host ...}

	Host name of machine


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
func (host *host) StorageController(hostName string, arrayID, adapter, enclosure, raidlevel, slot int) ([]byte, error) {
	args := []interface{}{hostName, arrayID, adapter, enclosure, raidlevel, slot}
	c := fmt.Sprintf("add host storage controller %s arrayid='%d' adapter='%d' enclosure='%d' raidlevel='%d' slot='%d'", args...)
	return cmd.RunCommand(c)
}
