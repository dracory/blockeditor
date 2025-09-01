package blockeditor

import (
	"github.com/dracory/hb"
	"github.com/dracory/ui"
	"github.com/samber/lo"
	"github.com/spf13/cast"
)

func (b *editor) cardButtonDropdown(block ui.BlockInterface) hb.TagInterface {
	flatBlock := NewFlatTree(b.blocks).Find(block.ID())

	if flatBlock == nil {
		return nil
	}

	definition := b.findDefinitionByType(flatBlock.Type)

	if definition == nil {
		return nil
	}

	areChildrenAllowed := lo.TernaryF(definition == nil, func() bool { return false }, func() bool { return definition.AllowChildren })

	position := lo.IndexOf(b.blocks, block)

	urlAddChild := b.url(map[string]string{
		ACTION:                  ACTION_BLOCK_ADD_MODAL,
		EDITOR_ID:               b.id,
		EDITOR_NAME:             b.name,
		EDITOR_HANDLER_ENDPOINT: b.handleEndpoint,
		"parent_id":             block.ID(),
	})

	linkAddChild := hb.Button().
		Child(hb.I().Class(`bi bi-plus-circle me-2`)).
		Text("Add Child").
		HxPost(urlAddChild).
		HxInclude("#" + b.id).
		HxTarget(`#` + b.id + `_wrapper`).
		HxSwap(`beforeend`)

	urlDuplicate := b.url(map[string]string{
		ACTION:                  ACTION_BLOCK_DUPLICATE,
		EDITOR_ID:               b.id,
		EDITOR_NAME:             b.name,
		EDITOR_HANDLER_ENDPOINT: b.handleEndpoint,
		BLOCK_ID:                block.ID(),
	})

	linkDuplicate := hb.Button().
		Child(hb.I().Class(`bi bi-stack me-2`)).
		Text("Duplicate").
		HxPost(urlDuplicate).
		HxInclude("#" + b.id).
		HxTarget(`#` + b.id + `_wrapper`).
		HxSwap(`outerHTML`)

	urlInsertBefore := b.url(map[string]string{
		ACTION:                  ACTION_BLOCK_ADD_MODAL,
		EDITOR_ID:               b.id,
		EDITOR_NAME:             b.name,
		EDITOR_HANDLER_ENDPOINT: b.handleEndpoint,
		BLOCK_ID:                block.ID(),
		"at_position":           cast.ToString(position),
	})

	linkInsertBefore := hb.Button().
		Child(hb.I().Class(`bi bi-arrow-90deg-right me-2`)).
		Text("Insert sibling before").
		HxPost(urlInsertBefore).
		HxInclude("#" + b.id).
		HxTarget(`#` + b.id + `_wrapper`).
		HxSwap(`beforeend`)

	urlInsertAfter := b.url(map[string]string{
		ACTION:                  ACTION_BLOCK_ADD_MODAL,
		EDITOR_ID:               b.id,
		EDITOR_NAME:             b.name,
		EDITOR_HANDLER_ENDPOINT: b.handleEndpoint,
		BLOCK_ID:                block.ID(),
		"at_position":           cast.ToString(position + 1),
	})

	linkInsertAfter := hb.Button().
		Child(hb.I().Class(`bi bi-arrow-return-right me-2`)).
		Text("Insert sibling after").
		HxPost(urlInsertAfter).
		HxInclude("#" + b.id).
		HxTarget(`#` + b.id + `_wrapper`).
		HxSwap(`beforeend`)

	urlMoveIntoPrevious := b.url(map[string]string{
		ACTION:                  ACTION_BLOCK_MOVE_INTO,
		EDITOR_ID:               b.id,
		EDITOR_NAME:             b.name,
		EDITOR_HANDLER_ENDPOINT: b.handleEndpoint,
		BLOCK_ID:                block.ID(),
		"in_sibling":            "previous",
	})

	linkMoveIntoPrevious := hb.Button().
		Child(hb.I().Class(`bi bi-arrow-up-right-square me-2`)).
		Text("Move into previous").
		HxPost(urlMoveIntoPrevious).
		HxInclude("#" + b.id).
		HxTarget(`#` + b.id + `_wrapper`).
		HxSwap(`outerHTML`)

	urlMoveIntoNext := b.url(map[string]string{
		ACTION:                  ACTION_BLOCK_MOVE_INTO,
		EDITOR_ID:               b.id,
		EDITOR_NAME:             b.name,
		EDITOR_HANDLER_ENDPOINT: b.handleEndpoint,
		BLOCK_ID:                block.ID(),
		"in_sibling":            "next",
	})

	linkMoveIntoNext := hb.Button().
		Child(hb.I().Class(`bi bi-arrow-down-right-square me-2`)).
		Text("Move into next").
		HxPost(urlMoveIntoNext).
		HxInclude("#" + b.id).
		HxTarget(`#` + b.id + `_wrapper`).
		HxSwap(`outerHTML`)

	urlMoveDown := b.url(map[string]string{
		ACTION:                  ACTION_BLOCK_MOVE_DOWN,
		EDITOR_ID:               b.id,
		EDITOR_NAME:             b.name,
		EDITOR_HANDLER_ENDPOINT: b.handleEndpoint,
		BLOCK_ID:                block.ID(),
	})

	linkMoveDown := hb.Button().
		Child(hb.I().Class(`bi bi-arrow-down me-2`)).
		Text("Move Down").
		HxPost(urlMoveDown).
		HxInclude("#" + b.id).
		HxTarget(`#` + b.id + `_wrapper`).
		HxSwap(`outerHTML`)

	urlMoveUp := b.url(map[string]string{
		ACTION:                  ACTION_BLOCK_MOVE_UP,
		EDITOR_ID:               b.id,
		EDITOR_NAME:             b.name,
		EDITOR_HANDLER_ENDPOINT: b.handleEndpoint,
		BLOCK_ID:                block.ID(),
	})

	linkMoveUp := hb.Button().
		Child(hb.I().Class(`bi bi-arrow-up me-2`)).
		Text("Move Up").
		HxPost(urlMoveUp).
		HxInclude("#" + b.id).
		HxTarget(`#` + b.id + `_wrapper`).
		HxSwap(`outerHTML`)

	urlMoveOutBefore := b.url(map[string]string{
		ACTION:                  ACTION_BLOCK_MOVE_OUT,
		EDITOR_ID:               b.id,
		EDITOR_NAME:             b.name,
		EDITOR_HANDLER_ENDPOINT: b.handleEndpoint,
		"block_id":              block.ID(),
		"to_position":           "before",
	})

	linkMoveOutBefore := hb.Button().
		Child(hb.I().Class(`bi bi-arrow-up-left-square me-2`)).
		Text("Move out before parent").
		HxPost(urlMoveOutBefore).
		HxInclude("#" + b.id).
		HxTarget(`#` + b.id + `_wrapper`).
		HxSwap(`outerHTML`)

	urlMoveOutAfter := b.url(map[string]string{
		ACTION:                  ACTION_BLOCK_MOVE_OUT,
		EDITOR_ID:               b.id,
		EDITOR_NAME:             b.name,
		EDITOR_HANDLER_ENDPOINT: b.handleEndpoint,
		"block_id":              block.ID(),
		"to_position":           "after",
	})

	linkMoveOutAfter := hb.Button().
		Child(hb.I().Class(`bi bi-arrow-down-left-square me-2`)).
		Text("Move out after parent").
		HxPost(urlMoveOutAfter).
		HxInclude("#" + b.id).
		HxTarget(`#` + b.id + `_wrapper`).
		HxSwap(`outerHTML`)

	urlDelete := b.url(map[string]string{
		ACTION:                  ACTION_BLOCK_DELETE,
		EDITOR_ID:               b.id,
		EDITOR_NAME:             b.name,
		EDITOR_HANDLER_ENDPOINT: b.handleEndpoint,
		BLOCK_ID:                block.ID(),
	})

	linkDelete := hb.Button().
		Child(hb.I().Class(`bi bi-trash me-2`)).
		Text("Delete").
		HxPost(urlDelete).
		HxInclude("#" + b.id).
		HxTarget(`#` + b.id + `_wrapper`).
		HxSwap(`outerHTML`)

	urlSettings := b.url(map[string]string{
		ACTION:                  ACTION_BLOCK_SETTINGS,
		EDITOR_ID:               b.id,
		EDITOR_NAME:             b.name,
		EDITOR_HANDLER_ENDPOINT: b.handleEndpoint,
		BLOCK_ID:                block.ID(),
	})

	linkSettings := hb.Button().
		Child(hb.I().Class(`bi bi-gear me-2`)).
		Text("Settings").
		HxPost(urlSettings).
		HxInclude("#" + b.id).
		HxTarget(`#` + b.id + `_wrapper`).
		HxSwap(`beforeend`)

	dropdown := hb.Div().
		Class(`BlockOptions dropdown float-start me-2`).
		Child(hb.Button().
			Class("btn btn-secondary btn-sm dropdown-toggle").
			Style(`border-radius: 30px;z-index: 100;font-size: 11px;padding: 2px;line-height: 1px;`).
			Type("button").
			Attr("data-bs-toggle", "dropdown").
			Attr("aria-expanded", "false").
			Child(hb.I().Class(`bi bi-three-dots-vertical`).Style(`font-size: 10px;`))).
		Child(hb.UL().
			Class("dropdown-menu").
			ChildIf(areChildrenAllowed, hb.LI().Child(linkAddChild.Class("dropdown-item"))).
			ChildIf(areChildrenAllowed, hb.Div().Class("dropdown-divider")).
			// Move up
			Child(hb.LI().Child(linkMoveUp.Class("dropdown-item"))).
			// Move down
			Child(hb.LI().Child(linkMoveDown.Class("dropdown-item"))).
			Child(hb.Div().Class("dropdown-divider")).
			// Insert before
			Child(hb.LI().Child(linkInsertBefore.Class("dropdown-item"))).
			// Insert after
			Child(hb.LI().Child(linkInsertAfter.Class("dropdown-item"))).
			Child(hb.Div().Class("dropdown-divider")).
			// Move into previous
			Child(hb.LI().Child(linkMoveIntoPrevious.Class("dropdown-item"))).
			// Move into next
			Child(hb.LI().Child(linkMoveIntoNext.Class("dropdown-item"))).
			// Only show the move out option if there is a parent
			ChildIf(flatBlock.ParentID != "", hb.Div().Class("dropdown-divider")).
			// Move out before
			ChildIf(flatBlock.ParentID != "", hb.LI().Child(linkMoveOutBefore.Class("dropdown-item"))).
			// Move out after
			ChildIf(flatBlock.ParentID != "", hb.LI().Child(linkMoveOutAfter.Class("dropdown-item"))).
			Child(hb.Div().Class("dropdown-divider")).
			// Duplicate
			Child(hb.LI().Child(linkDuplicate.Class("dropdown-item"))).
			Child(hb.Div().Class("dropdown-divider")).
			// Settings
			Child(hb.LI().Child(linkSettings.Class("dropdown-item"))).
			Child(hb.Div().Class("dropdown-divider")).
			// Delete
			Child(hb.LI().Child(linkDelete.Class("dropdown-item"))),
		)

	return dropdown
}
