package stackilackey

import (
	"github.td.teradata.com/ck250037/stacki-lackey-2/add"
	"github.td.teradata.com/ck250037/stacki-lackey-2/listcommand"
	"github.td.teradata.com/ck250037/stacki-lackey-2/remove"
	"github.td.teradata.com/ck250037/stacki-lackey-2/set"
)

// StackCommand exposes stacki functionality
type StackCommand struct {
	Add    add.Add
	Set    set.Set
	List   listcommand.List
	Remove remove.Remove
}
