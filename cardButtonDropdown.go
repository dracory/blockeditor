package blockeditor

import (
	"github.com/gouniverse/hb"
	"github.com/gouniverse/ui"
	"github.com/samber/lo"
	"github.com/spf13/cast"
)

func (b *editor) cardButtonDropdown(block ui.BlockInterface) hb.TagInterface {
	blockExt := NewFlatTree(b.blocks).Find(block.ID())

	position := lo.IndexOf(b.blocks, block)

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
			Child(func() hb.TagInterface {
				// Add child
				link := b.url(map[string]string{
					ACTION:                  ACTION_BLOCK_ADD_MODAL,
					EDITOR_ID:               b.id,
					EDITOR_NAME:             b.name,
					EDITOR_HANDLER_ENDPOINT: b.handleEndpoint,
					"parent_id":             block.ID(),
					"at_position":           cast.ToString(position),
				})
				dropdownItem := hb.Hyperlink().
					Class("dropdown-item").
					Href(link).
					Child(hb.I().Class(`bi bi-plus-circle me-2`)).
					Text("Add child").
					HxPost(link).
					HxTarget("#" + b.id + "_wrapper").
					HxSwap(`beforeend`)

				return hb.LI().
					Child(dropdownItem)
			}()).
			Child(hb.Div().Class("dropdown-divider")).
			Child(func() hb.TagInterface {
				// Insert sibling before
				link := b.url(map[string]string{
					ACTION:                  ACTION_BLOCK_ADD_MODAL,
					EDITOR_ID:               b.id,
					EDITOR_NAME:             b.name,
					EDITOR_HANDLER_ENDPOINT: b.handleEndpoint,
					"parent_id":             "",
					"at_position":           cast.ToString(position),
				})
				dropdownItem := hb.Hyperlink().
					Class("dropdown-item").
					Href(link).
					Child(hb.I().Class(`bi bi-arrow-90deg-right me-2`)).
					Text("Insert sibling before").
					HxPost(link).
					HxTarget("#" + b.id + "_wrapper").
					HxSwap(`beforeend`)

				return hb.LI().
					Child(dropdownItem)
			}()).
			Child(func() hb.TagInterface {
				// Insert sibling after
				link := b.url(map[string]string{
					ACTION:                  ACTION_BLOCK_ADD_MODAL,
					EDITOR_ID:               b.id,
					EDITOR_NAME:             b.name,
					EDITOR_HANDLER_ENDPOINT: b.handleEndpoint,
					"parent_id":             "",
					"at_position":           cast.ToString(position + 1),
				})
				dropdownItem := hb.Hyperlink().
					Class("dropdown-item").
					Href(link).
					Child(hb.I().Class(`bi bi-arrow-return-right me-2`)).
					Text("Insert sibling after").
					HxPost(link).
					HxTarget("#" + b.id + "_wrapper").
					HxSwap(`beforeend`)

				return hb.LI().
					Child(dropdownItem)
			}()).
			Child(hb.Div().Class("dropdown-divider")).
			Child(func() hb.TagInterface {
				// Move into previous sibling
				link := b.url(map[string]string{
					ACTION:                  ACTION_BLOCK_MOVE_INTO,
					EDITOR_ID:               b.id,
					EDITOR_NAME:             b.name,
					EDITOR_HANDLER_ENDPOINT: b.handleEndpoint,
					"block_id":              block.ID(),
					"in_sibling":            "previous",
				})
				dropdownItem := hb.Hyperlink().
					Class("dropdown-item").
					Href(link).
					Child(hb.I().Class(`bi bi-arrow-up-right-square me-2`)).
					Text("Move into previous sibling").
					HxPost(link).
					HxTarget("#" + b.id + "_wrapper").
					HxSwap(`outerHTML`)

				return hb.LI().
					Child(dropdownItem)
			}()).
			Child(func() hb.TagInterface {
				// Move into next sibling
				link := b.url(map[string]string{
					ACTION:                  ACTION_BLOCK_MOVE_INTO,
					EDITOR_ID:               b.id,
					EDITOR_NAME:             b.name,
					EDITOR_HANDLER_ENDPOINT: b.handleEndpoint,
					"block_id":              block.ID(),
					"in_sibling":            "next",
				})
				dropdownItem := hb.Hyperlink().
					Class("dropdown-item").
					Href(link).
					Child(hb.I().Class(`bi bi-arrow-down-right-square me-2`)).
					Text("Move into next sibling").
					HxPost(link).
					HxTarget("#" + b.id + "_wrapper").
					HxSwap(`outerHTML`)

				return hb.LI().
					Child(dropdownItem)
			}()).
			ChildIf(blockExt.ParentID != "", hb.Div().Class("dropdown-divider")).
			ChildIf(blockExt.ParentID != "", func() hb.TagInterface {
				// Move into previous sibling
				link := b.url(map[string]string{
					ACTION:                  ACTION_BLOCK_MOVE_OUT,
					EDITOR_ID:               b.id,
					EDITOR_NAME:             b.name,
					EDITOR_HANDLER_ENDPOINT: b.handleEndpoint,
					"block_id":              block.ID(),
				})
				dropdownItem := hb.Hyperlink().
					Class("dropdown-item").
					Href(link).
					Child(hb.I().Class(`bi bi-arrow-down-left-square me-2`)).
					Text("Move out of parent block").
					HxPost(link).
					HxTarget("#" + b.id + "_wrapper").
					HxSwap(`outerHTML`)

				return hb.LI().
					Child(dropdownItem)
			}()))

	return dropdown
}
