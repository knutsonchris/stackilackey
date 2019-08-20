package listcommand

import (
	"encoding/json"
	"fmt"

	"github.td.teradata.com/ck250037/stackilackey/cmd"
)

// OS represents an OS defined in the Stacki database
type OS struct {
	OSName string `json:"os"`
}

/*
OS will list the OSes defined
Arguments

	[os ...]

	Optional list of os names.
*/
func (os *OS) OS(OSName string) ([]OS, error) {
	c := fmt.Sprintf("list os %s", OSName)
	b, err := cmd.RunCommand(c)
	if err != nil {
		return nil, err
	}
	OSList := []OS{}
	err = json.Unmarshal(b, &OSList)
	if err != nil {
		// it may have been just an empty output from the Frontend
		nullOutput := NullOutput{}
		err = json.Unmarshal(b, &nullOutput)
		if err != nil {
			// if we still can't recognize the output, return an error
			return nil, err
		}
		return OSList, err
	}
	return OSList, err
}

/*
Attr will list the set of attributes for oses
Arguments

	[os]

	Name of OS (e.g. "linux", "sunos")
*/
func (os *OS) Attr(osName string) ([]Attr, error) {
	c := fmt.Sprintf("list os attr %s", osName)
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
Firewall will list the firewall rules for an OS
Arguments

	[os ...]

	Zero, one or more OS names. If no OS names are supplied, the firewall
	rules for all OSes are listed.
*/
func (os *OS) Firewall(osName string) ([]Firewall, error) {
	c := fmt.Sprintf("list os firewall %s", osName)
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
Route will list the routes for one or more OS
Arguments

	[os ...]

	Zero, one or more os.
*/
func (os *OS) Route(osName string) ([]Route, error) {
	c := fmt.Sprintf("list os route %s", osName)
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
StorageController will list the storage controller configuration for a given OS
Arguments

	[os ...]

	Zero, one or more OS names. If no OS names are supplied,
	the routes for all the OSes are listed.
*/
func (os *OS) StorageController(osName string) ([]Controller, error) {
	c := fmt.Sprintf("list os storage controller %s", osName)
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
