package blockeditor

import (
	"errors"

	"github.com/gouniverse/ui"
	"github.com/gouniverse/uid"
	"github.com/gouniverse/utils"
)

type NewEditorOptions struct {
	ID               string
	Name             string
	Value            string
	HandleEndpoint   string
	BlockDefinitions []BlockDefinition
}

func NewEditor(options NewEditorOptions) (*editor, error) {
	if options.ID == "" {
		options.ID = `id_` + uid.TimestampNano()
	}

	if options.HandleEndpoint == "" {
		return nil, errors.New("no handle endpoint")
	}

	if options.Name == "" {
		return nil, errors.New("no name")
	}

	if options.Value == "" {
		return nil, errors.New("no value")
	}

	if !utils.IsJSON(options.Value) {
		return nil, errors.New("value is not valid JSON")
	}

	blocks, err := ui.BlocksFromJson(options.Value)

	if err != nil {
		return nil, err
	}

	editor := &editor{}
	editor.handleEndpoint = options.HandleEndpoint
	editor.id = options.ID
	editor.name = options.Name
	editor.value = options.Value
	editor.blocks = blocks
	editor.blockDefinitions = options.BlockDefinitions
	return editor, nil
}
