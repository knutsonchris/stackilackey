package stackilackey

import (
	"github.td.teradata.com/ck250037/stackilackey/add"
	"github.td.teradata.com/ck250037/stackilackey/listcommand"
	"github.td.teradata.com/ck250037/stackilackey/remove"
	"github.td.teradata.com/ck250037/stackilackey/set"
)

// StackCommand exposes stacki functionality
type StackCommand struct {
	Add    add.Add
	Set    set.Set
	List   listcommand.List
	Remove remove.Remove
}
