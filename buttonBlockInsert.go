package blockeditor

import (
	"github.com/gouniverse/hb"
	"github.com/spf13/cast"
)

// buttonBlockInsert creates a dropdown button for inserting a new block
//
// The dropdown will list all of the available block types
func (e *editor) buttonBlockInsert(parentID string, atPosition int, single bool) *hb.Tag {
	link := e.url(map[string]string{
		ACTION:                  ACTION_BLOCK_ADD_MODAL,
		EDITOR_ID:               e.id,
		EDITOR_NAME:             e.name,
		EDITOR_HANDLER_ENDPOINT: e.handleEndpoint,
		"parent_id":             parentID,
		"at_position":           cast.ToString(atPosition),
	})

	button := hb.Button().
		Class("ButtonBlockInsert btn btn-secondary btn-sm").
		Style(`border-radius: 30px;z-index: 100`).
		Type("button").
		Child(hb.I().Class(`bi bi-plus-circle me-2`)).
		TextIf(single, "add new block").
		TextIf(!single, "insert new block").
		HxPost(link).
		HxInclude("#" + e.id).
		HxTarget("#" + e.id + "_wrapper").
		HxSwap(`beforeend`)

	return button
}
