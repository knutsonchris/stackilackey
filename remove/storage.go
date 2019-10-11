package remove

import (
	"fmt"

	"github.td.teradata.com/ck250037/stackilackey/cmd"
)

type controller struct {
}

type partition struct {
}

// Controller will remove a storage controller configuration from the database.
func (controller *controller) Controller(slot string) ([]byte, error) {
	c := fmt.Sprintf("remove storage controller slot='%s'", slot)
	return cmd.RunCommand(c)
}

// Partition will remove a storage partitioon from the database
func (paritition *partition) Partition(name, device, mountpoint string) ([]byte, error) {

	argKeys := []string{"device", "mountpoint"}
	argValues := []interface{}{device, mountpoint}
	baseCommand := fmt.Sprintf("remove storage partition %s", name)

	c, err := cmd.ArgsExpander(baseCommand, argKeys, argValues)
	if err != nil {
		return nil, err
	}
	return cmd.RunCommand(c)
}
