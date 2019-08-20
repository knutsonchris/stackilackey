package remove

import (
	"fmt"

	"github.td.teradata.com/ck250037/stacki-lackey-2/cmd"
)

type environment struct {
}

/*
Environment will remove an environment. If the environment is currently being used, (has attributes, or hosts) an error is raised.
Arguments

	{environment ...}

	One or more Environment specifications (e.g., 'test').
*/
func (e *environment) Environment(environmentName string) ([]byte, error) {
	c := fmt.Sprintf("remove environment %s", environmentName)
	return cmd.RunCommand(c)
}

/*
Attr will remove an attribute from an environment
Arguments

	[environment ...]

	One or more Environment specifications (e.g., 'test').


Parameters

	{attr=string}

	The attribute name that should be removed.
*/
func (e *environment) Attr(environmentName, attrName string) ([]byte, error) {

	argkeys := []string{"attr"}
	argvalues := []interface{}{attrName}
	baseCommand := fmt.Sprintf("remove environment attr %s", environmentName)

	c, err := cmd.ArgsExpander(baseCommand, argkeys, argvalues)
	if err != nil {
		return nil, err
	}

	return cmd.RunCommand(c)
}

/*
Firewall will remove a firewall service rule for an environment. To remove the rule you must supply the name of the rule
Arguments

	{environment ...}

	An environment name.


Parameters

	{rulename=string}

	Name of the environment-specific rule
*/
func (e *environment) Firewall(environmentName, ruleName string) ([]byte, error) {

	argkeys := []string{"rulename"}
	argvalues := []interface{}{ruleName}
	baseCommand := fmt.Sprintf("remove environment firewall %s", environmentName)

	c, err := cmd.ArgsExpander(baseCommand, argkeys, argvalues)
	if err != nil {
		return nil, err
	}

	return cmd.RunCommand(c)
}

/*
Route will remove an environment static route
Arguments

	{environment ...}

	Environment name


Parameters

	{address=string}

	The address of the route to remove.
*/
func (e *environment) Route(environmentName, address string) ([]byte, error) {

	argkeys := []string{"address"}
	argvalues := []interface{}{address}
	baseCommand := fmt.Sprintf("remove environment route %s", environmentName)

	c, err := cmd.ArgsExpander(baseCommand, argkeys, argvalues)
	if err != nil {
		return nil, err
	}

	return cmd.RunCommand(c)
}

/*
StorageController will remove a storage controller configuration for an environment
Arguments

	{environment ...}

	An environment name.


Parameters

	{slot=integer}

	Slot address(es). This can be a comma-separated list. If slot is '*',
	adapter/enclosure address applies to all slots.
*/
func (e *environment) StorageController(environmentName, slot string) ([]byte, error) {

	argkeys := []string{"slot"}
	argvalues := []interface{}{slot}
	baseCommand := fmt.Sprintf("remove environment storage controller %s", environmentName)

	c, err := cmd.ArgsExpander(baseCommand, argkeys, argvalues)
	if err != nil {
		return nil, err
	}

	return cmd.RunCommand(c)
}
