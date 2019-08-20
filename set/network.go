package set

import (
	"fmt"

	"github.td.teradata.com/ck250037/stackilackey/cmd"
)

type network struct {
}

/*
Description

	Sets the network address of a network.


Arguments

	{network}

	The name of the network.


Parameters

	{address=string}

	Address that the named network should have.
*/
func (network *network) Address(networkName, address string) ([]byte, error) {
	c := fmt.Sprintf("set network address %s address='%s'", networkName, address)
	return cmd.RunCommand(c)
}

/*
DNS enables or Disables DNS for one of more networks.

If DNS is enabled for a network then all known hosts on that network
will have their hostnames and IP addresses in a DNS server running
on the Frontend.  This will serve both forward and reverse lookups.
Arguments

	{network ...}

	The names of one or more networks.


Parameters

	{dns=boolean}

	Set to True to enable DNS for the given networks.
*/
func (network *network) DNS(networkName string, dns bool) ([]byte, error) {
	var dnsstr string
	if dns == true {
		dnsstr = "true"
	} else {
		dnsstr = "false"
	}
	c := fmt.Sprintf("set network dns %s dns='%s'", networkName, dnsstr)
	return cmd.RunCommand(c)
}

/*
Gateway will set the network gateway of a network.
Arguments

	{network}

	The name of the network.


Parameters

	{gateway=string}

	Gateway that the named network should have.
*/
func (network *network) Gateway(networkName, gateway string) ([]byte, error) {
	c := fmt.Sprintf("set network gateway %s gateway='%s'", networkName, gateway)
	return cmd.RunCommand(c)
}

/*
Mask will se the network mask for one or more networks.
Arguments

	{network ...}

	The names of one or more networks.


Parameters

	{mask=string}

	Mask that the named network should have.
*/
func (network *network) Mask(networkName, mask string) ([]byte, error) {
	c := fmt.Sprintf("set network mask %s mask='%s'", networkName, mask)
	return cmd.RunCommand(c)
}

/*
MTU will set the MTU for one or more networks.
Arguments

	{network ...}

	The names of one or more networks.


Parameters

	{mtu=string}

	MTU value the networks should have.
*/
func (network *network) MTU(networkName, mtu string) ([]byte, error) {
	c := fmt.Sprintf("set network mtu %s mtu='%s'", networkName, mtu)
	return cmd.RunCommand(c)
}

/*
Name will set the network name of a network.
Arguments

	{network}

	The name of the network.


Parameters

	{name=string}

	Name that the named network should have.
*/
func (network *network) Name(networkName, name string) ([]byte, error) {
	c := fmt.Sprintf("set network name %s name='%s'", networkName, name)
	return cmd.RunCommand(c)
}

/*
PXE will enable or diable PXE for one or more networks. All hosts must be connected
to at least one network that had PXE enabled.
Arguments

	{network ...}

	The names of one or more networks.


Parameters

	{pxe=boolean}

	Set to True to enable PXE for the given networks.
*/
func (network *network) PXE(networkName string, pxe bool) ([]byte, error) {
	var pxestr string
	if pxe == true {
		pxestr = "true"
	} else {
		pxestr = "false"
	}
	c := fmt.Sprintf("set network pxe %s pxe='%s'", networkName, pxestr)
	return cmd.RunCommand(c)
}

/*
Zone will set the DNS zone (domain name) for a network.
Arguments

	{network}

	The name of the network.


Parameters

	{zone=string}

	Zone that the named network should have.
*/
func (network *network) Zone(networkName, zone string) ([]byte, error) {
	c := fmt.Sprintf("set network zone %s zone='%s'", networkName, zone)
	return cmd.RunCommand(c)
}
