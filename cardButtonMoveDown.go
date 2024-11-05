package blockeditor

import "github.com/gouniverse/hb"

func (b *editor) cardButtonMoveDown(blockID string) *hb.Tag {
	icon := hb.I().Class(`bi bi-arrow-down`)

	buttonMoveDown := hb.Button().
		Class("btn btn-info text-white btn-sm ms-2 float-end").
		Style(`border-radius: 30px;font-size: 10px;`).
		Style(`padding: 2px 2px; line-height: 1;`).
		Type("button").
		Child(icon).
		HxPost(b.url(map[string]string{
			EDITOR_ID:               b.id,
			EDITOR_NAME:             b.name,
			EDITOR_HANDLER_ENDPOINT: b.handleEndpoint,
			ACTION:                  ACTION_BLOCK_MOVE_DOWN,
			BLOCK_ID:                blockID,
		})).
		HxInclude(`#` + b.id).
		HxTarget(`#` + b.id + `_wrapper`).
		HxSwap(`outerHTML`)

	return buttonMoveDown
}
