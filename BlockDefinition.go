package blockeditor

import (
	"github.com/gouniverse/form"
	"github.com/gouniverse/hb"
	"github.com/gouniverse/ui"
)

type BlockDefinition struct {
	Icon   hb.TagInterface
	Type   string
	Fields []form.Field
	ToHTML func(block ui.BlockInterface) string
}
