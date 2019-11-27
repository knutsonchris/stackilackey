package remove

import (
	"fmt"

	"github.com/knutsonchris/stackilackey/cmd"
)

type host struct {
}

/* Host will remove a host from the database. This command will remove all related database rows for each specified host.
Arguments

	{host}

	Name of host to remove from the database
*/
func (host *host) Host(hostName string) ([]byte, error) {
	c := fmt.Sprintf("remove host %s", hostName)
	return cmd.RunCommand(c)
}

/*
Attr will remove an attribute for a host.
Arguments

	[host ...]

	One host


Parameters

	{attr=string}

	The attribute name that should be removed.
*/
func (host *host) Attr(hostName, attrName string) ([]byte, error) {
	c := fmt.Sprintf("remove host attr %s attr='%s'", hostName, attrName)
	return cmd.RunCommand(c)
}

/*
Firewall will remove a firewall rule for a host. To remove a rule, you must supply the name of the rule.
Arguments

	{host ...}

	Name of a host machine.
*/
func (host *host) Firewall(hostName, ruleName string) ([]byte, error) {
	c := fmt.Sprintf("remove host firewall %s rulename='%s'", hostName, ruleName)
	return cmd.RunCommand(c)
}

/*
StorageController will remove a storage controller configuration for the specified host
Arguments

	{host ...}

	Host name of machine


Parameters

	{slot=integer}

	Slot address(es). This can be a comma-separated list. If slot is '*',
	adapter/enclosure address applies to all slots.
*/
func (host *host) StorageController(hostName, slot string) ([]byte, error) {
	c := fmt.Sprintf("remove host storage controller %s slot='%s'", hostName, slot)
	return cmd.RunCommand(c)
}

/*
Route will remove a static route for a host
Arguments

	{host ...}

	Name of a host machine.


Parameters

	{address=string}

	The address of the static route to remove. This argument is required.

	[syncnow=string]

	If set to true, the routing table will be updated as well as the db.
*/
func (host *host) Route(hostName, address, syncnow string) ([]byte, error) {
	c := fmt.Sprintf("remove host route %s address='%s' syncnow='%s'", hostName, address, syncnow)
	return cmd.RunCommand(c)
}
