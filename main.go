package stackilackey

import (
	"github.com/knutsonchris/stackilackey/add"
	"github.com/knutsonchris/stackilackey/listcommand"
	"github.com/knutsonchris/stackilackey/remove"
	"github.com/knutsonchris/stackilackey/set"
)

// StackCommand exposes stacki functionality
type StackCommand struct {
	Add    add.Add
	Set    set.Set
	List   listcommand.List
	Remove remove.Remove
}
