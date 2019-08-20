package listcommand

import (
	"encoding/json"
	"fmt"

	"github.td.teradata.com/ck250037/stacki-lackey-2/cmd"
)

// Appliance represents an appliance stored in the stacki database
type Appliance struct {
	ApplianceName string `json:"appliance"`
	// Public    string `json:"public"` NOTE: this is un-necessary. default to public=true
}

// NullOutput can absorb an empty response from the stacki frontend
type NullOutput struct {
	Output interface{} `json:"Output"`
}

// XML represents the json output of stack list appliance xml <applianceName>
type XML struct {
	Col0 string `json:"col-0"`
	Col1 string `json:"col-1"`
}

/*
Appliance will list the appliances defined in the cluster database.
Arguments

	[appliance ...]

	Optional list of appliance names.
*/
func (a *Appliance) Appliance(applianceName string) ([]Appliance, error) {
	c := fmt.Sprintf("list appliance %s", applianceName)
	b, err := cmd.RunCommand(c)
	if err != nil {
		return nil, err
	}
	appliances := []Appliance{}
	err = json.Unmarshal(b, &appliances)
	if err != nil {
		// it may have been just an empty output from the Frontend
		nullOutput := NullOutput{}
		err = json.Unmarshal(b, &nullOutput)
		if err != nil {
			// if we still can't recognize the output, return an error
			return nil, err
		}
		return appliances, err
	}
	return appliances, err
}

/*
Attr will list the set of attributes for the appliance.
Arguments

	[appliance]

	Name of appliance

*/
func (a *Appliance) Attr(applianceName, attr string) ([]Attr, error) {
	c := fmt.Sprintf("list appliance attr %s attr='%s'", applianceName, attr)
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
Firewall will list the firewall rules for a given appliance type.
Arguments

	[appliance ...]

	Zero, one or more appliance names. If no appliance names are supplied,x
	the firewall rules for all the appliances are listed.
*/
func (a *Appliance) Firewall(applianceName string) ([]Firewall, error) {
	c := fmt.Sprintf("list appliance firewall %s", applianceName)
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
Route will list the routes for a given appliance type.
Arguments

	[appliance ...]

	Zero, one or more appliance names. If no appliance names are supplied,
	the routes for all the appliances are listed.
*/
func (a *Appliance) Route(applianceName string) ([]Route, error) {
	c := fmt.Sprintf("list appliance route %s", applianceName)
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
StorageController will list the storage controller configuration for a given appliance.
Arguments

	[appliance ...]

	Zero, one or more appliance names. If no appliance names are supplied,
	the routes for all the appliances are listed.
*/
func (a *Appliance) StorageController(applianceName string) ([]Controller, error) {
	c := fmt.Sprintf("list appliance storage controller %s", applianceName)
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

/*
XML will list the XML profile for a given appliance type. This is useful
for high level debugging but will be missing any host specific
variables. It cannot be used to pass into 'rocks list host profile'
to create a complete Kickstart/Jumpstart profile.
Arguments

	[appliance ...]

	Optional list of appliance names.
*/
func (a *Appliance) XML(applianceName string) ([]XML, error) {
	c := fmt.Sprintf("list appliance xml %s", applianceName)
	b, err := cmd.RunCommand(c)
	if err != nil {
		return nil, err
	}
	xmls := []XML{}
	err = json.Unmarshal(b, &xmls)
	if err != nil {
		// it may have been just an empty output from the Frontend
		nullOutput := NullOutput{}
		err = json.Unmarshal(b, &nullOutput)
		if err != nil {
			// if we still can't recognize the output, return an error
			return nil, err
		}
		return xmls, err
	}
	return xmls, err
}
