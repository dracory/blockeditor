package blockeditor

import "github.com/dracory/hb"

func (b *editor) cardButtonMoveUp(blockID string) *hb.Tag {
	icon := hb.I().Class(`bi bi-arrow-up`)
	buttonMoveUp := hb.Button().
		Class("btn btn-info text-white btn-sm ms-2 float-end").
		Style(`border-radius: 30px;font-size: 10px;`).
		Style(`padding: 2px 2px; line-height: 1;`).
		Type("button").
		Child(icon).
		Title("Move Up").
		HxPost(b.url(map[string]string{
			EDITOR_ID:               b.id,
			EDITOR_NAME:             b.name,
			EDITOR_HANDLER_ENDPOINT: b.handleEndpoint,
			ACTION:                  ACTION_BLOCK_MOVE_UP,
			BLOCK_ID:                blockID,
		})).
		HxInclude(`#` + b.id).
		HxTarget(`#` + b.id + `_wrapper`).
		HxSwap(`outerHTML`)

	return buttonMoveUp
}
