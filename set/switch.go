package set

import (
	"fmt"

	"github.td.teradata.com/ck250037/stackilackey/cmd"
)

type netSwitch struct {
}

/*
HostInterface will edit the switch host relation that stacki keeps in its database. This command
changes the name of a host's interface that is associated with a specific port on a switch.
Arguments

	{switch}

	One switch name.


Parameters

	{host=string}

	One host name.

	{interface=string}

	Name of the interface.

	{port=string}

	The port number on the switch.
*/
func (netSwitch *netSwitch) HostInterface(switchName, host, interfaceName, port string) ([]byte, error) {
	args := []interface{}{switchName, host, interfaceName, port}
	c := fmt.Sprintf("set switch host interface %s host='%s' interface='%s' switch='%s'", args...)
	return cmd.RunCommand(c)
}

/*
Host port will change port association on the switch that this host's interface maps to.
Arguments

	{switch}

	One switch name.


Parameters

	{host=string}

	One host name.

	{interface=string}

	Name of the interface.

	{port=string}

	The port number on the switch.
*/
func (netSwitch *netSwitch) HostPort(switchName, host, interfaceName, port string) ([]byte, error) {
	args := []interface{}{switchName, host, interfaceName, port}
	c := fmt.Sprintf("set switch host port %s host='%s' interface='%s' switch='%s'", args...)
	return cmd.RunCommand(c)
}

/*
PartitionMembership will set membership state on an infiniband partition in the Stacki database for a switch.
Arguments

	{switch}

	The name of the switches to add partition members to.


Parameters

	{name=string}

	The name of the partition to set membership status on the switch.

	[enforce_sm=boolean]

	If a switch is not an infiniband subnet manager an error will be raised.

	[guid=string]
*/
func (netSwitch *netSwitch) PartitionMembership(switchName, name, guid, interfaceName, member, membership string, enforceSM bool) ([]byte, error) {
	var enforceSMstr string
	if enforceSM == true {
		enforceSMstr = "true"
	} else {
		enforceSMstr = "false"
	}
	args := []interface{}{switchName, name, enforceSMstr, guid, interfaceName, member, membership}
	c := fmt.Sprintf("set switch partition membership %s name='%s' enforce_sm='%s' guid='%s' interface='%s' member='%s' membership='%s'", args...)
	return cmd.RunCommand(c)
}

/*
PartitionOptions will set the infiniband partition flags in the Stacki database.
Note that a sync is still required to enact this change on the switch.
Arguments

	{switch}

	The name of the switches on which to set these options.


Parameters

	{name=string}

	The name of the partition to set this flag on.  Must either be 'Default'
	or a hex value between 0x0001-0x7ffe (1 and 32,766).  The name will be
	normalized to the lower-case, leading-zero hex format: 'aaa' -> '0x0aaa'

	[enforce_sm=boolean]

	If a switch is not an infiniband subnet manager an error will be raised.

	[options=string]

	A list of options to set on the partition.  The format is
	'flag=value flag2=value2'.  Currently supported are 'ipoib=True|False'
	and 'defmember=limited|full'.  Unless explicitly specified, 'ipoib' and
	'defmember' are not set.
*/
func (netSwitch *netSwitch) PartitionOptions(switchName, name, options string, enforceSM bool) ([]byte, error) {
	var enforceSMstr string
	if enforceSM == true {
		enforceSMstr = "true"
	} else {
		enforceSMstr = "false"
	}
	args := []interface{}{switchName, name, enforceSMstr, options}
	c := fmt.Sprintf("set switch partition options %s name='%s', enforce_sm='%s' options='%s'", args...)
	return cmd.RunCommand(c)
}

/*
SM will enable to subnet manager for the given switch.
Arguments

	{switch}

	Exactly one infiniband switch name which will become the subnet manager for
	the fabric it is on.  All other infiniband switches on the same fabric will
	have their subnet manager status disabled.  Fabric is determined soley based
	on the 'ibfabric' attribute.


Parameters

	[disable=boolean]

	When set to True, will disable subnet manager status on that switch only.
	Defaults to False.
*/
func (netSwitch *netSwitch) SM(switchName string, disable bool) ([]byte, error) {
	var disablestr string
	if disable == true {
		disablestr = "true"
	} else {
		disablestr = "false"
	}
	c := fmt.Sprintf("set switch sm %s disable='%s'", switchName, disablestr)
	return cmd.RunCommand(c)
}
