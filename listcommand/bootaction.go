package listcommand

import (
	"encoding/json"

	"github.td.teradata.com/ck250037/stacki-lackey-2/cmd"
)

// Bootaction represents a PXE OOS install target
type Bootaction struct {
	BootactionName string `json:"bootaction"`
	TypeName       string `json:"type"`
	OS             string `json:"os"`
	Kernel         string `json:"kernel"`
	Ramdisk        string `json:"ramdisk"`
	Args           string `json:"args"`
}

/*
Bootaction will list the available PXE OS Install targets.
*/
func (bootaction *Bootaction) Bootaction() ([]Bootaction, error) {
	c := "list bootaction"
	b, err := cmd.RunCommand(c)
	if err != nil {
		return nil, err
	}
	bootactions := []Bootaction{}
	err = json.Unmarshal(b, &bootactions)
	if err != nil {
		// it may have been just an empty output from the Frontend
		nullOutput := NullOutput{}
		err = json.Unmarshal(b, &nullOutput)
		if err != nil {
			// if we still can't recognize the output, return an error
			return nil, err
		}
		return bootactions, err
	}
	return bootactions, err
}
