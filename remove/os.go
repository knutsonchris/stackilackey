package remove

import (
	"fmt"

	"github.td.teradata.com/ck250037/stacki-lackey-2/cmd"
)

type os struct {
}

/*
OS will remove an OS definition from the system
Arguments

	{os ...}

	The OS type (e.g., "linux", "sunos").
*/
func (o *os) OS(osName string) ([]byte, error) {
	c := fmt.Sprintf("remove os %s", osName)
	return cmd.RunCommand(c)
}

/*
Attr will remove an attribute for an OS
Arguments

	[os ...]

	One or more OS specifications (e.g., 'linux').


Parameters

	{attr=string}

	The attribute name that should be removed.
*/
func (o *os) Attr(osName, attr string) ([]byte, error) {

	argKeys := []string{"attr"}
	argValues := []interface{}{attr}
	baseCommand := fmt.Sprintf("remove os attr %s", osName)

	c, err := cmd.ArgsExpander(baseCommand, argKeys, argValues)
	if err != nil {
		return nil, err
	}

	return cmd.RunCommand(c)
}

/*
Firewall will remove a firewall rule for an OS type. To remove a rule, one must supply the name of the rule.
Arguments

	{os ...}

	Name of an OS type (e.g., "linux", "sunos").


Parameters

	{rulename=string}

	Name of the OS-specific rule
*/
func (o *os) Firewall(osName, rulename string) ([]byte, error) {

	argKeys := []string{"rulename"}
	argValues := []interface{}{rulename}
	baseCommand := fmt.Sprintf("remove os firewall %s", osName)

	c, err := cmd.ArgsExpander(baseCommand, argKeys, argValues)
	if err != nil {
		return nil, err
	}

	return cmd.RunCommand(c)
}

/*
Route will remove a static route for an OS type
Arguments

	{os ...}

	The OS type (e.g., 'linux', 'sunos', etc.).


Parameters

	{address=string}

	The address of the static route to remove.
*/
func (o *os) Route(osName, address string) ([]byte, error) {

	argKeys := []string{"address"}
	argValues := []interface{}{address}
	baseCommand := fmt.Sprintf("remove os route %s", osName)

	c, err := cmd.ArgsExpander(baseCommand, argKeys, argValues)
	if err != nil {
		return nil, err
	}

	return cmd.RunCommand(c)
}

/*
StorageController will remove a storage controller configuration for an OS type
Arguments

	{os ...}

	OS type (e.g., 'redhat', 'sles').


Parameters

	{slot=integer}

	Slot address(es). This can be a comma-separated list. If slot is '*',
	adapter/enclosure address applies to all slots.

	[adapter=integer]

	Adapter address. If adapter is '*', enclosure/slot address applies to
	all adapters.

	[enclosure=integer]

	Enclosure address. If enclosure is '*', adapter/slot address applies
	to all enclosures.
*/
func (o *os) StorageController(osName, slot, adapter, enclosure string) ([]byte, error) {

	argKeys := []string{"slot", "adapter", "enclosure"}
	argValues := []interface{}{slot, adapter, enclosure}
	baseCommand := fmt.Sprintf("remove os storage controller %s", osName)

	c, err := cmd.ArgsExpander(baseCommand, argKeys, argValues)
	if err != nil {
		return nil, err
	}

	return cmd.RunCommand(c)
}

/*
StoragePartition will remove a storage partition for an os type
Arguments

	[os]

	OS Name


Parameters

	[device=string]

	Device whose partition configuration needs to be removed from
	the database.

	[mountpoint=string]

	Mountpoint for the partition that needs to be removed from
	the database.

*/
func (o *os) StoragePartition(osName, device, mountpoint string) ([]byte, error) {

	argKeys := []string{"device", "mountpoint"}
	argValues := []interface{}{device, mountpoint}
	baseCommand := fmt.Sprintf("remove os storage partition %s", osName)

	c, err := cmd.ArgsExpander(baseCommand, argKeys, argValues)
	if err != nil {
		return nil, err
	}

	return cmd.RunCommand(c)
}
