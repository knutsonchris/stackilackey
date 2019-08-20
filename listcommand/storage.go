package listcommand

import (
	"encoding/json"
	"strconv"

	"github.td.teradata.com/ck250037/stacki-lackey-2/cmd"
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

// Array represents a storage array in the Stacki database
type Array struct {
	Scope      string `json:"scope"`
	Device     string `json:"device"`
	PardID     int    `json:"partid"`
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
			//panic(string(b))
			return nil, err
		}
		return controllers, err
	}
	return controllers, err
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
