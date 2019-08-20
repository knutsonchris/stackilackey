package add

import (
	"fmt"

	"github.td.teradata.com/ck250037/stacki-lackey-2/cmd"
)

type network struct {
}

/*
Add a network to the database. By default, the "private" network is already defined.
Arguments

	{name}

	Name of the new network.


Parameters

	{address=string}

	Network address.

	{mask=string}

	Network mask.

	[dns=boolean]

	If set to True this network will be included in the builtin DNS server.
	The default value is false.

	[gateway=string]

	Default gateway for the network. This is optional, not all networks
	require gateways.

	[mtu=string]

	The MTU for the new network. Default is 1500.

	[pxe=boolean]

	If set to True this network will be managed by the builtin DHCP/PXE
	server.
	The default is False.

	[zone=string]

	The Domain name or the DNS Zone name to use
	for all hosts of this particular subnet. Default
	is set to the name of the network.
*/
func (network *network) Network(networkName, address, mask, gateway, mtu, zone string, dns, pxe bool) ([]byte, error) {
	var dnsstr string
	if dns == true {
		dnsstr = "true"
	} else {
		dnsstr = "false"
	}
	var pxestr string
	if pxe == true {
		pxestr = "true"
	} else {
		pxestr = "false"
	}

	argKeys := []string{"address", "mask", "dns", "gateway", "mtu", "pxe", "zone"}
	argValues := []interface{}{address, mask, dnsstr, gateway, mtu, pxestr, zone}
	baseCommand := fmt.Sprintf("add network %s", networkName)

	c, err := cmd.ArgsExpander(baseCommand, argKeys, argValues)
	if err != nil {
		return nil, err
	}

	return cmd.RunCommand(c)
}
