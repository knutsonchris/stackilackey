package listcommand

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/knutsonchris/stackilackey/cmd"
)

// FlexString exists because the Stacki API will occasionally give us an int
type FlexString string

// Controller represents a storage controller in the Stacki database
type Controller struct {
	Enclosure FlexString `json:"enclosure"`
	Adapter   FlexString `json:"adapter"`
	Slot      FlexString `json:"slot"`
	RaidLevel FlexString `json:"raidlevel"`
	ArrayID   FlexString `json:"arrayid"`
	Options   string     `json:"options"`
}

// Partition represents a storage array in the Stacki database
type Partition struct {
	Scope      string `json:"scope"`
	Device     string `json:"device"`
	PartID     int    `json:"partid"`
	Mountpoint string `json:"mountpoint"`
	Size       int    `json:"size"`
	FStype     string `json:"fstype"`
	Options    string `json:"options"`
}

// Controller will list the global storage controller connfiguration
func (controller *Controller) Controller() ([]Controller, error) {
	c := "list storage controller"
	b, err := cmd.RunCommand(c)
	if err != nil {
		return nil, err
	}
	controllers := []Controller{}
	err = json.Unmarshal(b, &controllers)
	if err != nil {
		// it may have been just an empty output from the Frontend
		nullOutput := NullOutput{}
		err2 := json.Unmarshal(b, &nullOutput)
		if err2 != nil {
			// if we still can't recognize the output, return an error
			return nil, err
		}
		return controllers, err
	}
	return controllers, err
}

// Partition will list storage partition information, with an optional scope host, os, or appliance name
// leave scope blank for global scope
func (partition *Partition) Partition(scope string) ([]Partition, error) {
	c := fmt.Sprintf("list storage partition %s", scope)
	b, err := cmd.RunCommand(c)
	if err != nil {
		return nil, err
	}
	partitions := []Partition{}
	err = json.Unmarshal(b, &partitions)
	if err != nil {
		// it may have been just an empty output from the Frontend
		nullOutput := NullOutput{}
		err2 := json.Unmarshal(b, &nullOutput)
		if err2 != nil {
			// if we still can't recognize the output, return an error
			return nil, err
		}
		return partitions, err
	}
	return partitions, err
}

// UnmarshalJSON is a FlexString reciever, to call out the Stacki API on its lies
func (f *FlexString) UnmarshalJSON(b []byte) error {
	if b[0] != '"' {
		var i int
		err := json.Unmarshal(b, &i)
		if err != nil {
			return err
		}
		*f = FlexString(strconv.Itoa(i))
		return nil
	}
	return json.Unmarshal(b, (*string)(f))
}
