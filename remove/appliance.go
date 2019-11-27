package remove

import (
	"fmt"

	"github.com/knutsonchris/stackilackey/cmd"
)

type appliance struct {
}

/*
Appliance will remove an appliance definition from the system. This can be
called with just the appliance or it can be further
qualified by supplying the root XML node name and/or the
graph XML file name.
Arguments

	{name}

	The name of the appliance.
*/
func (appliance *appliance) Appliance(applianceName string) ([]byte, error) {
	c := fmt.Sprintf("remove appliance %s", applianceName)
	return cmd.RunCommand(c)
}

/*
Attr will remove an attribute for an appliance
Arguments

	[appliance ...]

	One or more appliances


Parameters

	{attr=string}

	The attribute name that should be removed.
*/
func (appliance *appliance) Attr(applianceName, attr string) ([]byte, error) {
	c := fmt.Sprintf("remove appliance attr %s attr='%s'", applianceName, attr)
	return cmd.RunCommand(c)
}

/*
Firewall will remove a firewall service rule for an appliance type. To remove the rule, you must supply the name of the rule.
Arguments

	{appliance ...}

	Name of an appliance type (e.g., "backend").


Parameters

	{rulename=string}

	Name of the Appliance-specific rule
*/
func (appliance *appliance) Firewall(applianceName, ruleName string) ([]byte, error) {
	c := fmt.Sprintf("remove appliance firewall %s rulename='%s'", applianceName, ruleName)
	return cmd.RunCommand(c)
}

/*
Route will remove a static route for an appliance type
Arguments

	{appliance ...}

	Appliance name. This argument is required.


Parameters

	{address=string}

	The address of the static route to remove.
*/
func (appliance *appliance) Route(applianceName, address string) ([]byte, error) {
	c := fmt.Sprintf("remove appliance route %s address='%s'", applianceName, address)
	return cmd.RunCommand(c)
}

/*
StorageController will remove a storage controller configuration for an appliance type
Arguments

	{appliance ...}

	Appliance type (e.g., "backend").


Parameters

	{slot=integer}

	Slot address(es). This can be a comma-separated list. If slot is '*',
	adapter/enclosure address applies to all slots.
*/
func (appliance *appliance) StorageController(applianceName, slot string) ([]byte, error) {
	c := fmt.Sprintf("remove appliance storage controller %s slot='%s'", applianceName, slot)
	return cmd.RunCommand(c)
}
