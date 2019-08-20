package set

import (
	"fmt"

	"github.td.teradata.com/ck250037/stackilackey/cmd"
)

type host struct {
}

/*
interface channel {host ...} {channel=string} [interface=string] [mac=string]
	[network=string]
interface default {host ...} [default=boolean] [interface=string] [mac=string]
	[network=string]
interface interface {host ...} {interface=string} [mac=string] [network=string]
interface ip {host} {ip=string} [interface=string] [mac=string]
	[network=string]
interface mac {host} {mac=string} [interface=string] [network=string]
interface module {host ...} {module=string} [interface=string] [mac=string]
	[network=string]
interface name {host} {name=string} [interface=string] [mac=string]
	[network=string]
interface network {host ...} {network=string} [interface=string] [mac=string]
interface options {host ...} {options=string} [interface=string] [mac=string]
	[network=string]
interface vlan {host ...} {vlan=integer} [interface=string] [mac=string]
	[network=string]
metadata {host ...} {metadata=string}
name {host} {name=string}
power {host ...} {command=string} [debug=boolean]
rack {host ...} {rack=string}
rank {host ...} {rank=string}
storage [host ...] {action=string} {enclosure=string} {slot=string}
*/

/*
Appliance will set the appliance for a list of hosts.
Arguments

	{host ...}

	One or more host names.


Parameters

	{appliance=string}

	Appliance name (e.g. "backend").
*/
func (host *host) Appliance(hostName, applianceName string) ([]byte, error) {
	c := fmt.Sprintf("set host appliance %s appliance='%s'", hostName, applianceName)
	return cmd.RunCommand(c)
}

/*
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
	a := []interface{}{hostName, attr, value, shadowstr}
	c := fmt.Sprintf("set host attr %s attr='%s' value='%s' shadow='%s'", a...)
	return cmd.RunCommand(c)
}

/*
Boot will set a bootaction for a host. A hosts action can be set to 'install' or to 'os'.
Arguments

	{host ...}

	One or more host names.


Parameters

	{action=string}

	The label name for the bootaction. This must be one of: 'os',
	'install'.

	[nukecontroller=boolean]

	Set the host to overwrite the controller configuration on next
	install. Default: False

	[nukedisks=boolean]

	Set the host to erase all disks on next install. Default: False

	[sync=boolean]

	Controls if 'sync host boot' needs to be run after setting the
	bootaction. Default: True
*/
func (host *host) Boot(hostName, action string, nukecontroller, nukedisks, sync bool) ([]byte, error) {
	var nukecontrollerstr string
	if nukecontroller == true {
		nukecontrollerstr = "true"
	} else {
		nukecontrollerstr = "false"
	}
	var nukedisksstr string
	if nukedisks == true {
		nukedisksstr = "true"
	} else {
		nukedisksstr = "false"
	}
	var syncstr string
	if sync == true {
		syncstr = "true"
	} else {
		syncstr = "false"
	}
	a := []interface{}{hostName, action, nukecontrollerstr, nukedisksstr, syncstr}
	c := fmt.Sprintf("set host boot %s action='%s' nukecontroller='%s' nukedisks='%s' sync='%s'", a...)
	return cmd.RunCommand(c)
}

/*
Box will set the box for a host
Arguments

	{host ...}

	One host name.


Parameters

	{box=string}

	The name of the box (e.g. default)
*/
func (host *host) Box(hostName, box string) ([]byte, error) {
	c := fmt.Sprintf("set host box %s box='%s'", hostName, box)
	return cmd.RunCommand(c)
}

/*
Bootactionwill update a bootaction for a host.
Arguments

	{host ...}

	One or more host names.


Parameters

	{action=string}

	bootaction name. This should already exist via 'stack list bootaction'

	{type=string}

	type of bootaction. can be one of 'os' or 'install'

	[sync=boolean]

	controls if 'sync host boot' needs to be run after
	setting the bootaction.
*/
func (host *host) Bootaction(hostName, action, typeName string, sync bool) ([]byte, error) {
	var syncstr string
	if sync == true {
		syncstr = "true"
	} else {
		syncstr = "false"
	}
	a := []interface{}{hostName, action, typeName, syncstr}
	c := fmt.Sprintf("set host bootaction %s action='%s' type='%s' sync='%s'", a...)
	return cmd.RunCommand(c)
}

/*
Comment will set the comment field for a list of hosts.
Arguments

	{host ...}

	One or more host names.


Parameters

	{comment=string}

	The string to assign to the comment field for each host.
*/
func (host *host) Comment(hostName, comment string) ([]byte, error) {
	c := fmt.Sprintf("set host comment %s comment='%s'", hostName, comment)
	return cmd.RunCommand(c)
}

/*
Environment will specify an Environment for the gives hosts.  Environments are
used to add another level to attribute resolution.  This is commonly
used to partition a single Frontend into managing multiple clusters.
Arguments

	{host ...}

	One or more host names.


Parameters

	{environment=string}

	The environment name to assign to each host.
*/
func (host *host) Environment(hostName, environmentName string) ([]byte, error) {
	c := fmt.Sprintf("set host environment %s environment='%s'", hostName, environmentName)
	return cmd.RunCommand(c)
}

/*
InterfaceChannel will set the channel for a named interface.
Arguments

	{host ...}

	One or more hosts.


Parameters

	{channel=string}

	The channel for an interface. Use channel=NULL to clear.

	[interface=string]

	Name of the interface.

	[mac=string]

	MAC address of the interface.

	[network=string]

	Network name of the interface.
*/
func (host *host) InterfaceChannel(hostName, channel, interfaceName, mac, network string) ([]byte, error) {
	a := []interface{}{hostName, channel, interfaceName, mac, network}
	c := fmt.Sprintf("set host interface channel %s channel='%s' interface='%s' mac='%s' network='%s'", a...)
	return cmd.RunCommand(c)
}

/*
InterfaceDefualt will designate one network as the default route for a set of hosts.
Either the interface, mac, or network paramater is required.
Arguments

	{host ...}

	One or more hosts.


Parameters

	[default=boolean]

	Can be used to set the value of default to False.

	[interface=string]

	Name of the interface.

	[mac=string]

	MAC address of the interface.

	[network=string]

	Network name of the interface.
*/
func (host *host) InterfaceDefault(hostName, interfaceName, mac, network string, defaultFlag bool) ([]byte, error) {
	var defaultstr string
	if defaultFlag == true {
		defaultstr = "true"
	} else {
		defaultstr = "false"
	}
	a := []interface{}{hostName, defaultstr, interfaceName, mac, network}
	c := fmt.Sprintf("set host interface default %s default='%s' interface='%s' mac='%s' network='%s'", a...)
	return cmd.RunCommand(c)
}

/*
InterfaceInterface will set the logical interface of a mac address for particular hosts.
Arguments

	{host ...}

	One or more hosts.


Parameters

	{interface=string}

	Name of the interface.

	[mac=string]

	MAC address of the interface.

	[network=string]

	Network name of the interface.
*/
func (host *host) InterfaceInterface(hostName, interfaceName, mac, network string) ([]byte, error) {
	a := []interface{}{hostName, interfaceName, mac, network}
	c := fmt.Sprintf("set host interface interface %s interface='%s' mac='%s' network='%s'", a...)
	return cmd.RunCommand(c)
}

/*
InterfaceIP will set the IP address for the named interface for one host.
Arguments

	{host}

	A single host.


Parameters

	{ip=string}

	The IP address to set. Use ip=AUTO to let the system pick one for
	you or ip=NULL to clear the IP address.

	[interface=string]

	Name of the interface.

	[mac=string]

	MAC address of the interface.

	[network=string]

	Network name of the interface.
*/
func (host *host) InterfaceIP(hostName, interfaceName, mac, network string) ([]byte, error) {
	a := []interface{}{hostName, interfaceName, mac, network}
	c := fmt.Sprintf("set host interface ip %s ip='%s' interface='%s' mac='%s' network='%s'", a...)
	return cmd.RunCommand(c)
}

/*
InterfaceMac will set the mac address for a named interface on the host.
Arguments

	{host}

	A single host.


Parameters

	{mac=string}

	MAC address of the interface. Usually of the form dd:dd:dd:dd:dd:dd
	where d is a hex digit. This format is not enforced. Use mac=NULL to
	clear the mac address.

	[interface=string]

	Name of the interface.

	[network=string]

	Network name of the interface.
*/
func (host *host) InterfaceMac(hostName, mac, interfaceName, network string) ([]byte, error) {
	a := []interface{}{hostName, mac, interfaceName, network}
	c := fmt.Sprintf("set host mac ip %s ip='%s' interface='%s' mac='%s' network='%s'", a...)
	return cmd.RunCommand(c)
}

/*
InterfaceModule will set the device module for a named interface. On Linux this will get
translated to an entry in /etc/modprobe.conf.
Arguments

	{host ...}

	One or more hosts.


Parameters

	{module=string}

	Module name.

	[interface=string]

	Name of the interface.

	[mac=string]

	MAC address of the interface.

	[network=string]

	Network name of the interface.
*/
func (host *host) InterfaceModule(hostName, module, interfaceName, mac, network string) ([]byte, error) {
	a := []interface{}{hostName, module, interfaceName, mac, network}
	c := fmt.Sprintf("set host interface module %s module='%s' interface='%s' mac='%s' network='%s'", a...)
	return cmd.RunCommand(c)
}

/*
InterfaceName will set the logical name of a network interface on a particular host.
Arguments

	{host}

	A single host.


Parameters

	{name=string}

	Name of this interface (e.g. newname). This is only the
	name associated with a certain interface. FQDNs are disallowed.
	To set the domain or zone for an interface, use the
	"stack add network" command, and then associate the interface
	with the network

	[interface=string]

	Name of the interface.

	[mac=string]

	MAC address of the interface.

	[network=string]

	Network name of the interface.
*/
func (host *host) InterfaceName(hostName, name, interfaceName, mac, network string) ([]byte, error) {
	a := []interface{}{hostName, name, interfaceName, mac, network}
	c := fmt.Sprintf("set host interface name %s name='%s' interface='%s' mac='%s' network='%s'", a...)
	return cmd.RunCommand(c)
}

/*
InterfaceNetwork will set the network for a named interface on one or more hosts.
Arguments

	{host ...}

	One or more hosts.


Parameters

	{network=string}

	Network name of the interface.

	[interface=string]

	Name of the interface.

	[mac=string]

	MAC address of the interface.
*/
func (host *host) InterfaceNetwork(hostName, network, interfaceName, mac string) ([]byte, error) {
	a := []interface{}{hostName, network, interfaceName, mac}
	c := fmt.Sprintf("set host interface network %s network='%s' interface='%s' mac='%s'", a...)
	return cmd.RunCommand(c)
}

/*
InterfaceOptions will set the options for a device module for a named interface. On Linux,
this will get translated to an entry in /etc/modprobe.conf.
Arguments

	{host ...}

	One or more hosts.


Parameters

	{options=string}

	The options for an interface. Use options=NULL to clear.
	options="dhcp", and options="noreport" have
	special meaning. options="bonding-opts=\"\"" sets up bonding
	options for bonded interfaces

	[interface=string]

	Name of the interface.

	[mac=string]

	MAC address of the interface.

	[network=string]

	Network name of the interface.
*/
func (host *host) InterfaceOptions(hostName, options, interfaceName, mac, network string) ([]byte, error) {
	a := []interface{}{hostName, options, interfaceName, mac, network}
	c := fmt.Sprintf("set host interface options %s options='%s' interface='%s' mac='%s' network='%s'", a...)
	return cmd.RunCommand(c)
}

/*
Arguments

	{host ...}

	One or more hosts.


Parameters

	{vlan=integer}

	The VLAN ID that should be updated. This must be an integer and the
	pair 'subnet/vlan' must be defined in the VLANs table.

	[interface=string]

	Name of the interface.

	[mac=string]

	MAC address of the interface.

	[network=string]

	Network name of the interface.
*/
func (host *host) InterfaceVlan(hostName, interfaceName, mac, network string, vlan int) ([]byte, error) {
	a := []interface{}{hostName, vlan, interfaceName, mac, network}
	c := fmt.Sprintf("set host interface vlan %s vlan='%d', interface='%s' mac='%s' network='%s'", a...)
	return cmd.RunCommand(c)
}

/*
Metadata will set the metadata for a list of hosts. The metadata is
reserved for the user and is not used internally by Stacki. The
intention is to provide a mechanism similar to the AWS
meta-data to allow arbitrary data to be attached to a host.

It is recommended that the metadata (if used) should be a JSON
document but there are no assumptions from Stacki on the
structure of the data.

Metadata should be accessed using the "metadata" read-only attribute.
Arguments

	{host ...}

	One or more host names.


Parameters

	{metadata=string}

	The metadata document
*/
func (host *host) Metadata(hostName, metadata string) ([]byte, error) {
	c := fmt.Sprintf("set host metadata %s metadata='%s'", hostName, metadata)
	return cmd.RunCommand(c)
}

/*
Name will rename a host.
Arguments

	{host}

	The current name of the host.


Parameters

	{name=string}

	The new name for the host.
*/
func (host *host) Name(hostName, name string) ([]byte, error) {
	c := fmt.Sprintf("set host name %s name='%s'", hostName, name)
	return cmd.RunCommand(c)
}

/*
Power will send a "power" command to a host. Valid power commands are: on, off and reset. This
command uses IPMI to change the power setting on a host.
Arguments

	{host ...}

	One or more host names.


Parameters

	{command=string}

	The power command to execute. Valid power commands are: "on", "off" and "reset".

	[debug=boolean]

	Print debug output from the ipmitool command.
*/
func (host *host) Power(hostName, command string, debug bool) ([]byte, error) {
	var debugstr string
	if debug == true {
		debugstr = "true"
	} else {
		debugstr = "false"
	}
	args := []interface{}{hostName, command, debugstr}
	c := fmt.Sprintf("set host power %s command='%s' debug='%s'", args...)
	return cmd.RunCommand(c)
}

/*
Rack will set the rack number for a list of hosts.
Arguments

	{host ...}

	One or more host names.


Parameters

	{rack=string}

	The rack name (usually a number) to assign to each host.
*/
func (host *host) Rack(hostName, rack string) ([]byte, error) {
	c := fmt.Sprintf("set host rack %s rack='%s'", hostName, rack)
	return cmd.RunCommand(c)
}

/*
Rank will set the rank number for a list of hosts.
Arguments

	{host ...}

	One or more host names.


Parameters

	{rank=string}

	The rank number to assign to each host.
*/
func (host *host) Rank(hostName, rank string) ([]byte, error) {
	c := fmt.Sprintf("set host rank %s rank='%s'", hostName, rank)
	return cmd.RunCommand(c)
}

/*
Storage will set the state for a storage device for hosts (e.g., to change the state
of a disk from 'offline' to 'online').
Arguments

	[host ...]

	Zero, one or more host names. If no host names are supplied, this
	command will apply the state change to all known hosts.


Parameters

	{action=string}

	The action to perform on the device. Valid actions are: online,
	offline, configure, locate-on and locate-off.

	{enclosure=string}

	An integer id for the enclosure that contains the storage device.

	{slot=string}

	An integer id for the slot that contains the storage device.
*/
func (host *host) Storage(hostName, action, enclosure, slot string) ([]byte, error) {
	args := []interface{}{hostName, action, enclosure, slot}
	c := fmt.Sprintf("set host storage %s action='%s' enclosure='%s' slot='%s'", args...)
	return cmd.RunCommand(c)
}
