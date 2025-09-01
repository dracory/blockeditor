package blockeditor

import (
	"github.com/dracory/form"
	"github.com/dracory/hb"
	"github.com/dracory/ui"
)

type BlockDefinition struct {
	// Icon is the icon of the block
	Icon hb.TagInterface

	// Type is the type of the block
	Type string

	// Fields are the fields for the paameters of the block which can be edited
	Fields []form.FieldInterface

	// AllowChildren determines if the block can have children
	AllowChildren bool

	// AllowedChildTypes is a list of block types that can be added as children
	AllowedChildTypes []string

	// Wrapper is a function that wraps the block's card in a custom tag
	// suitable for a column that need to apply its width so that the
	// card is displayed with the correct width
	Wrapper func(block ui.BlockInterface) *hb.Tag

	// ToTag is a function that converts a block to a tag to be displayed
	// in the block's card body in the editor
	ToTag func(block ui.BlockInterface) *hb.Tag

	// ToHTML is a function that converts a block to HTML to be displayed
	// in the block card in the editor
	// Deprecated
	// ToHTML func(block ui.BlockInterface) string // optional
}
