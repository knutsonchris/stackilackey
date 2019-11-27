package listcommand

import (
	"encoding/json"
	"fmt"

	"github.com/knutsonchris/stackilackey/cmd"
)

// Network represents a network in the Stacki database
type Network struct {
	NetworkName string `json:"network"`
	Address     string `json:"address"`
	Mask        string `json:"mask"`
	Gateway     string `json:"gateway"`
	MTU         int    `json:"mtu"`
	Zone        string `json:"zone"`
	DNS         bool   `json:"dns"`
	PXE         bool   `json:"pxe"`
}

/*
Network will list the defined networks for this system
Arguments

	[network ...]

	Zero, one or more network names. If no network names are supplied,
	info about all the known networks is listed.
*/
func (network *Network) Network(networkName string) ([]Network, error) {
	c := fmt.Sprintf("list network %s", networkName)
	b, err := cmd.RunCommand(c)
	if err != nil {
		return nil, err
	}
	networks := []Network{}
	err = json.Unmarshal(b, &networks)
	if err != nil {
		// it may have been just an empty output from the Frontend
		nullOutput := NullOutput{}
		err = json.Unmarshal(b, &nullOutput)
		if err != nil {
			// if we still can't recognize the output, return an error
			return nil, err
		}
		return networks, err
	}
	return networks, err
}
