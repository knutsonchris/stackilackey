package add

import (
	"fmt"

	"github.td.teradata.com/ck250037/stacki-lackey-2/cmd"
)

type route struct {
}

/*
Route will add a route for all hosts in the cluster.
Parameters

	{address=string}

	Host or network address

	{gateway=string}

	Network (e.g., IP address), subnet name (e.g., 'private', 'public'), or
	a device gateway (e.g., 'eth0').

	[interface=string]

	Specific interface to send traffic through. Should only be used if
	you need traffic to go through a VLAN interface (e.g., 'eth0.1').

	[netmask=string]

	Specifies the netmask for a network route.  For a host route
	this is not required and assumed to be 255.255.255.255
*/
func (route *route) Route(address, gateway, interfaceName, netmask string) ([]byte, error) {
	args := []interface{}{address, gateway, interfaceName, netmask}
	c := fmt.Sprintf("add route address='%s' gateway='%s' interface='%s' netmask='%s'", args...)
	return cmd.RunCommand(c)
}
