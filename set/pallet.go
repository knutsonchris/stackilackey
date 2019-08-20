package set

import (
	"fmt"

	"github.td.teradata.com/ck250037/stackilackey/cmd"
)

type pallet struct {
}

/*
Tag will set a tag for one or more pallets.
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
	c := fmt.Sprintf("set pallet tag %s tag='%s' value='%s' arch='%s' os='%s' release='%s' version='%s'", args...)
	return cmd.RunCommand(c)
}
