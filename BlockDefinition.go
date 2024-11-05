package blockeditor

import (
	"github.com/gouniverse/form"
	"github.com/gouniverse/hb"
	"github.com/gouniverse/ui"
)

type BlockDefinition struct {
	// Icon is the icon of the block
	Icon hb.TagInterface

	// Type is the type of the block
	Type string

	// Fields are the fields for the paameters of the block which can be edited
	Fields []form.Field

	// AllowChildren determines if the block can have children
	AllowChildren bool

	// AllowedChildTypes is a list of block types that can be added as children
	AllowedChildTypes []string

	// ToHTML is a function that converts a block to HTML to be displayed
	// in the block card in the editor
	ToHTML func(block ui.BlockInterface) string // optional
}
