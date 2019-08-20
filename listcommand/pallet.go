package listcommand

import (
	"encoding/json"
	"fmt"

	"github.td.teradata.com/ck250037/stackilackey/cmd"
)

// Pallet represents a
type Pallet struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Release string `json:"release"`
	Arch    string `json:"arch"`
	OS      string `json:"OS"`
	Boxes   string `json:"boxes"`
	URL     string `json:"url"`
}

/*
Pallet will list the status of available pallets
Arguments

	[pallet ...]

	List of pallets. This should be the pallet base name (e.g., base, hpc,
	kernel). If no pallets are listed, then status for all the pallets are
	listed.


Parameters

	{expanded=bool}

	Displays an additional column containing the url of the pallet.

	[arch=string]

	The architecture of the pallet to be listed. If no architecture is
	supplied, then all architectures of a pallet will be listed.

	[os=string]

	The OS of the pallet to be listed. If no OS is supplied, then all OS
	versions of a pallet will be listed.

	[release=string]

	The release number of the pallet to be listed. If no release number is
	supplied, then all releases of a pallet will be listed.

	[version=string]

	The version number of the pallets to list. If no version number is
	supplied, then all versions of a pallet will be listed.
*/
func (pallet *Pallet) Pallet(palletName, arch, os, release, version string) ([]Pallet, error) {
	// to grab the url field every time, hard code expanded=true
	argKeys := []string{"expanded", "arch", "os", "release", "version"}
	argValues := []interface{}{true, arch, os, release, version}

	baseCommand := fmt.Sprintf("list pallet %s", palletName)
	c, err := cmd.ArgsExpander(baseCommand, argKeys, argValues)
	if err != nil {
		return nil, err
	}
	b, err := cmd.RunCommand(c)
	if err != nil {
		return nil, err
	}
	pallets := []Pallet{}
	err = json.Unmarshal(b, &pallets)
	if err != nil {
		// it may have been just an empty output from the Frontend
		nullOutput := NullOutput{}
		err = json.Unmarshal(b, &nullOutput)
		if err != nil {
			// if we still can't recognize the output, return an error
			return nil, err
		}
		return pallets, err
	}
	return pallets, err
}
