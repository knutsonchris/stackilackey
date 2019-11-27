package listcommand

import (
	"encoding/json"
	"fmt"

	"github.com/knutsonchris/stackilackey/cmd"
)

// Environment represents a environment in the Stacki Database
type Environment struct {
	EnvironmentName string `json:"environment"`
}

/*
Environment will return a list of environments
Arguments

	[environment ...]

	Optional list of environment names.
*/
func (e *Environment) Environment(environmentName string) ([]Environment, error) {
	c := "list environment"
	b, err := cmd.RunCommand(c)
	if err != nil {
		return nil, err
	}
	environments := []Environment{}
	err = json.Unmarshal(b, &environments)
	if err != nil {
		// it may have been just an empty output from the Frontend
		nullOutput := NullOutput{}
		err = json.Unmarshal(b, &nullOutput)
		if err != nil {
			// if we still can't recognize the output, return an error
			return nil, err
		}
		return environments, err
	}
	return environments, err
}

/*
Attr will list the set of attributes for environments
Arguments

	[environment]

	Name of environment (e.g. "test")
*/
func (e *Environment) Attr(environmentName, attr string) ([]Attr, error) {

	argkeys := []string{"attr"}
	argvalues := []interface{}{attr}
	baseCommand := fmt.Sprintf("list environment attr %s", environmentName)

	c, err := cmd.ArgsExpander(baseCommand, argkeys, argvalues)
	if err != nil {
		return nil, err
	}

	b, err := cmd.RunCommand(c)
	if err != nil {
		return nil, err
	}
	attrs := []Attr{}
	err = json.Unmarshal(b, &attrs)
	if err != nil {
		// it may have been just an empty output from the Frontend
		nullOutput := NullOutput{}
		err = json.Unmarshal(b, &nullOutput)
		if err != nil {
			// if we still can't recognize the output, return an error
			return nil, err
		}
		return attrs, err
	}
	return attrs, err
}

/*
Firewall will list the firewall rules for a given environment
Arguments

	[environment ...]

	Zero or more environments. If no environments are supplied,
	the firewall rules for all environments are listed.
*/
func (e *Environment) Firewall(environmentName string) ([]Firewall, error) {
	c := fmt.Sprintf("list environment firewall %s", environmentName)
	b, err := cmd.RunCommand(c)
	if err != nil {
		return nil, err
	}
	firewalls := []Firewall{}
	err = json.Unmarshal(b, &firewalls)
	if err != nil {
		// it may have been just an empty output from the Frontend
		nullOutput := NullOutput{}
		err = json.Unmarshal(b, &nullOutput)
		if err != nil {
			// if we still can't recognize the output, return an error
			return nil, err
		}
		return firewalls, err
	}
	return firewalls, err
}

/*
Route will list the routes for one or more environments
Arguments

	[environment ...]

	Zero, one or more environments.
*/
func (e *Environment) Route(environmentName string) ([]Route, error) {
	c := fmt.Sprintf("list environment route %s", environmentName)
	b, err := cmd.RunCommand(c)
	if err != nil {
		return nil, err
	}
	routes := []Route{}
	err = json.Unmarshal(b, &routes)
	if err != nil {
		// it may have been just an empty output from the Frontend
		nullOutput := NullOutput{}
		err = json.Unmarshal(b, &nullOutput)
		if err != nil {
			// if we still can't recognize the output, return an error
			return nil, err
		}
		return routes, err
	}
	return routes, err
}

/*
StorageController will list the storage controller configuration for one or more environments
*/
func (e *Environment) StorageController(environmentName string) ([]Controller, error) {
	c := fmt.Sprintf("list environment storage controller %s", environmentName)
	b, err := cmd.RunCommand(c)
	if err != nil {
		return nil, err
	}
	controllers := []Controller{}
	err = json.Unmarshal(b, &controllers)
	if err != nil {
		// it may have been just an empty output from the Frontend
		nullOutput := NullOutput{}
		err = json.Unmarshal(b, &nullOutput)
		if err != nil {
			// if we still can't recognize the output, return an error
			return nil, err
		}
		return controllers, err
	}
	return controllers, err
}
