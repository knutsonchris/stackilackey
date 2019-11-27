package add

import (
	"fmt"

	"github.com/knutsonchris/stackilackey/cmd"
)

type netswitch struct {
}

/*
Host will add a new host to a switch.
Arguments

	{switch}

	Name of the switch


Parameters

	{host=string}

	Name of the host being assigned a vlan id

	{port=string}

	Port the host is connected to the switch on

	[interface=string]

	Name of the interface you want to use to connect to the switch.
	Default: The first interface that is found that shares the
	same network as the switch.
*/
func (netswitch *netswitch) Host(switchName, host, port, interfaceName string) ([]byte, error) {
	args := []interface{}{switchName, host, port, interfaceName}
	c := fmt.Sprintf("add switch host %s host='%s' port='%s' interface='%s'", args...)
	return cmd.RunCommand(c)
}

/*
Partition will add a partition for an Infiniband switch to the Stacki database.
Note that a sync is still required to enact this change on the switch.
Arguments

	{switch}

	The name of the switches on which to create this partition.


Parameters

	{name=string}

	The name of the partition to set this flag on.  Must either be 'Default'
	or a hex value between 0x0001-0x7ffe (1 and 32,766).  The name will be
	normalized to the lower-case, leading-zero hex format: 'aaa' -> '0x0aaa'

	[enforce_sm=boolean]

	If a switch is not an infiniband subnet manager an error will be raised.

	[options=string]

	A set of options to create the partition with.  The format is
	'flag=value flag2=value2'.  Currently supported are 'ipoib=True|False'
	and 'defmember=limited|full'.  Unless explicitly specified, 'ipoib' and
	'defmember' are not set.
*/
func (netswitch *netswitch) Partition(switchName, name, options string, enforceSM bool) ([]byte, error) {
	var enforceSMstr string
	if enforceSM == true {
		enforceSMstr = "true"
	} else {
		enforceSMstr = "false"
	}
	args := []interface{}{switchName, name, enforceSMstr, options}
	c := fmt.Sprintf("add switch partition %s name='%s' enforce_sm='%s' options='%s'", args...)
	return cmd.RunCommand(c)

}

/*
PartitionMember will add members to an infiniband partition in the Stacki database for one or more switches.
Arguments

	{switch}

	The name of the switches to add partition members to.


Parameters

	{name=string}

	The name of the partition to add members to on the switch(es).

	[enforce_sm=boolean]

	If a switch is not an infiniband subnet manager an error will be raised.

	[guid=string]

	The GUID of the host's infiniband interface to use.

	[interface=string]

	The name of an infiniband interface to add as a member.  Must be specified
	with the name of the host the interface belongs to.

	[member=string]

	The hostname with an infiniband interface to add as a member.  Must be
	specified with the name of the interface to use.

	[membership=string]

	The membership state to use for this member on the partition.  Must be 'both',
	or 'limited'.  Defaults to 'limited'.
*/
func (netswitch *netswitch) PartitionMember(switchName, name, guid, interfaceName, member, membership string, enforceSM bool) ([]byte, error) {
	var enforceSMstr string
	if enforceSM == true {
		enforceSMstr = "true"
	} else {
		enforceSMstr = "false"
	}
	args := []interface{}{switchName, name, enforceSMstr, guid, interfaceName, member, membership}
	c := fmt.Sprintf("add switch partition member %s name='%s' enforce_sm='%s' guid='%s' interface='%s' member='%s' membership'%s'", args...)
	return cmd.RunCommand(c)
}
