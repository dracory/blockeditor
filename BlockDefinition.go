package blockeditor

import (
	"github.com/gouniverse/form"
	"github.com/gouniverse/hb"
	"github.com/gouniverse/ui"
)

type BlockDefinition struct {
	Icon            hb.TagInterface
	Type            string
	Fields          []form.Field
	AllowedChildren []string
	ToHTML          func(block ui.BlockInterface) string // optional
}
