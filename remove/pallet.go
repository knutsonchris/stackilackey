package remove

import (
	"fmt"

	"github.com/knutsonchris/stackilackey/cmd"
)

type pallet struct {
}

/*
Pallet will remove a pallet from both the database and the filesystem
Arguments

	{pallet ...}

	List of pallets. This should be the pallet base name (e.g., base, hpc,
	kernel).


Parameters

	[arch=string]

	The architecture of the pallet to be removed. If no architecture is
	supplied, then all architectures will be removed.

	[os=string]

	The OS of the pallet to be removed. If no OS is
	supplied, then all OSes will be removed.

	[release=string]

	The release id of the pallet to be removed. If no release id is
	supplied, then all releases of a pallet will be removed.

	[version=string]

	The version number of the pallet to be removed. If no version number is
	supplied, then all versions of a pallet will be removed.
*/
func (p *pallet) Pallet(palletName, arch, os, release, version string) ([]byte, error) {

	argKeys := []string{"arch", "os", "release", "version"}
	argValues := []interface{}{arch, os, release, version}
	baseCommand := fmt.Sprintf("remove pallet %s", palletName)

	c, err := cmd.ArgsExpander(baseCommand, argKeys, argValues)
	if err != nil {
		return nil, err
	}

	return cmd.RunCommand(c)
}
