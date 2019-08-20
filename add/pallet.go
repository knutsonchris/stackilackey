package add

import (
	"fmt"

	"github.td.teradata.com/ck250037/stacki-lackey-2/cmd"
)

type pallet struct {
}

/*
Pallet will add pallet ISO images to this machine's pallet directory. This command
copies all files in the ISOs to the local machine. The default location
is a directory under /export/stack/pallets.
Arguments

	[pallet ...]

	A list of pallet ISO images to add to the local machine. If no list is
	supplied, then if a pallet is mounted on /mnt/cdrom, it will be copied
	to the local machine. If the pallet is hosted on the internet, it will
	be downloaded and stored on the local machine.


Parameters

	[clean=bool]

	If set, then remove all files from any existing pallets of the same
	name, version, and architecture before copying the contents of the
	pallets onto the local disk.  This parameter should not be set
	when adding multi-CD pallets such as the OS pallet, but should be set
	when adding single pallet CDs such as the Grid pallet.

	[dir=string]

	The base directory to copy the pallet to.
	The default is: /export/stack/pallets.

	[password=string]

	If the pallet's download server requires authentication.

	[updatedb=string]

	Add the pallet info to the cluster database.
	The default is: true.

	[username=string]

	If the pallet's download server requires authentication.
*/
func (pallet *pallet) Pallet(palletName, dir, password, updatedb, username string, clean bool) ([]byte, error) {
	var cleanstr string
	if clean == true {
		cleanstr = "true"
	} else {
		cleanstr = "false"
	}
	args := []interface{}{palletName, cleanstr, dir, password, updatedb, username}
	c := fmt.Sprintf("add pallet %s clean='%s' dir='%s' password='%s' updatedb='%s' username='%s'", args...)
	return cmd.RunCommand(c)
}

/*
Tag will add a tag to one or more pallets if it does not already exist.
Arguments

	{pallet ...}

	Name of one or more pallets.


Parameters

	{tag=string}

	Name of the tag

	{value=string}

	Value of the attribute

	[arch=string]

	Arch of the pallet

	[os=string]

	OS of the pallet

	[release=string]

	Release of the pallet

	[version=string]

	Version of the pallet
*/
func (pallet *pallet) Tag(palletName, tag, value, arch, os, release, version string) ([]byte, error) {
	args := []interface{}{palletName, tag, value, arch, os, release, version}
	c := fmt.Sprintf("add pallet tag %s tag='%s' value='%s' arch='%s' os='%s' release='%s' version='%s'", args...)
	return cmd.RunCommand(c)
}
