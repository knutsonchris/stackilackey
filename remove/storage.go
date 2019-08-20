package remove

import (
	"fmt"

	"github.td.teradata.com/ck250037/stacki-lackey-2/cmd"
)

type controller struct {
}

// Remove will remove a storage controller configuration from the database.
func (controller *controller) Controller(slot string) ([]byte, error) {
	c := fmt.Sprintf("remove storage controller slot='%s'", slot)
	return cmd.RunCommand(c)
}
