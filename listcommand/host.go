package listcommand

import (
	"encoding/json"
	"fmt"

	"github.td.teradata.com/ck250037/stacki-lackey-2/cmd"
)

// Host represents a host stored in the stacki database
type Host struct {
	Name          string `json:"host"`
	Rack          string `json:"rack"`
	Rank          string `json:"rank"`
	Appliance     string `json:"appliance"`
	OS            string `json:"os"`
	Box           string `json:"box"`
	Environemnt   string `json:"environment"`
	OSAction      string `json:"osaction"`
	InstallAction string `json:"installaction"`
	Comment       string `json:"comment"`
}

// HostInterface represents a host interface in the Stacki database
type HostInterface struct {
	Host      string `json:"host"`
	Interface string `json:"interface"`
	Default   bool   `json:"default"`
	Network   string `json:"network"`
	Mac       string `json:"mac"`
	IP        string `json:"ip"`
	Name      string `json:"name"`
	Module    string `json:"module"`
	Vlan      string `json:"vlan"`
	Options   string `json:"options"`
	Channel   string `json:"channel"`
}

/*
Host will list the Appliance, and physical position info for a list of hosts.
Arguments

	[host ...]

	Zero, one or more host names. If no host names are supplied, info about
	all the known hosts is listed.
*/
// TODO: do we want to include the parameter `hash`?
func (host *Host) Host(hostName string) ([]Host, error) {
	c := fmt.Sprintf("list host %s", hostName)
	b, err := cmd.RunCommand(c)
	if err != nil {
		return nil, err
	}
	hosts := []Host{}
	err = json.Unmarshal(b, &hosts)
	if err != nil {
		// it may have been just an empty output from the Frontend
		nullOutput := NullOutput{}
		err = json.Unmarshal(b, &nullOutput)
		if err != nil {
			// if we still can't recognize the output, return an error
			return nil, err
		}
		return hosts, err
	}
	return hosts, err
}

/*
Attr will list the set of sttributes for hosts
Arguments

	[host]

	Host name of machine
*/
// TODO: implement attr and display parameters
func (host *Host) Attr(hostName string) ([]Attr, error) {
	c := fmt.Sprintf("list host attr %s", hostName)
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
Firewall will list the current firewall files for the named host
Arguments

	[host]

	One host
*/
func (host *Host) Firewall(hostName string) ([]Firewall, error) {
	c := fmt.Sprintf("list host firewall %s", hostName)
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
StorageController will list the storage controller configuration for a given host(s).
Arguments

	[host ...]

	Zero, one or more host names. If no host names are supplied,
	the routes for all the hosts are listed.
*/
func (host *Host) StorageController(hostName string) ([]Controller, error) {
	c := fmt.Sprintf("list host storage controller %s", hostName)
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
Route will list the static routes that are assigned to a host.
Arguments

	[host]

	Host name of machine
*/
func (host *Host) Route(hostName string) ([]Route, error) {
	c := fmt.Sprintf("list host route %s", hostName)
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
Interface will list the the interface definitions for hosts. For each host supplied on the command line, this command prints the hostname and interface definitions for that host.
Arguments

	[host ...]

	Zero, one or more host names. If no host names are supplied, info about
	all the known hosts is listed.
*/
func (host *Host) Interface(hostName string) ([]HostInterface, error) {
	c := fmt.Sprintf("list host interface %s", hostName)
	b, err := cmd.RunCommand(c)
	if err != nil {
		return nil, err
	}
	ifaces := []HostInterface{}
	err = json.Unmarshal(b, &ifaces)
	if err != nil {
		// it may have been just an empty output from the Frontend
		nullOutput := NullOutput{}
		err = json.Unmarshal(b, &nullOutput)
		if err != nil {
			// if we still can't recognize the output, return an error
			return nil, err
		}
		return ifaces, err
	}
	return ifaces, err
}

// TODO: list host alais
// TODO: list host bonded
// TODO: list host bridge
// TODO: list host key
// TODO: list host message
// TODO: list host partition
