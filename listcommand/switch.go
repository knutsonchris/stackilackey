package listcommand

import (
	"encoding/json"
	"fmt"

	"github.com/knutsonchris/stackilackey/cmd"
)

/*
TODO:
switch [switch ...] [expanded=boolean]
switch config [switch ...] [raw=string]
switch host [switch ...]
switch mac [switch ...] [pinghosts=bool]
switch partition {switch} [enforce_sm=boolean] [name=string]
switch partition member {switch} [enforce_sm=boolean] [expanded=boolean]
	[name=string]
switch status [switch ...]
switch support
*/

// HostSwitch represents a switch in the Stacki database
type HostSwitch struct {
	SwitchName      string `json:"switch"`
	Rack            string `json:"rack"`
	Rank            string `json:"rank"`
	Appliance       string `json:"appliance"`
	Make            string `json:"make"`
	Model           string `json:"model"`
	IBSubnetManager string `json:"ib subnet manager"`
	IBFabric        string `json:"ib fabric"`
}

/*
Switch will list Appliance, physical position, and model of any hosts with appliance type of `switch`.
Arguments

	[switch ...]

	Zero, one or more switch names. If no switch names are supplies, info about
	all the known switches is listed.
*/
func (h *HostSwitch) Switch(switchName string) ([]HostSwitch, error) {
	c := fmt.Sprintf("list switch %s expanded=true", switchName)
	b, err := cmd.RunCommand(c)
	if err != nil {
		return nil, err
	}
	switches := []HostSwitch{}
	err = json.Unmarshal(b, &switches)
	if err != nil {
		// it may have been just an empty output from the Frontend
		nullOutput := NullOutput{}
		err = json.Unmarshal(b, &nullOutput)
		if err != nil {
			// if we still can't recognize the output, return an error
			return nil, err
		}
		return switches, err
	}
	return switches, err
}
