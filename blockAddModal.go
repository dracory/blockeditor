package blockeditor

import (
	"net/http"

	"github.com/gouniverse/bs"
	"github.com/gouniverse/hb"
	"github.com/gouniverse/ui"
	"github.com/gouniverse/utils"
	"github.com/samber/lo"
)

// blockAddModal shows the block add modal
func (e *editor) blockAddModal(r *http.Request) string {
	atPosition := utils.Req(r, "at_position", "0")
	parentID := utils.Req(r, "parent_id", "")

	modalCloseScript := `document.getElementById('ModalBlockUpdate').remove();document.getElementById('ModalBackdrop').remove();`

	blocksJSON, err := ui.BlocksToJson(e.blocks)

	if err != nil {
		return err.Error()
	}

	blockTiles := lo.Map(e.blockDefinitions, func(d BlockDefinition, _ int) hb.TagInterface {
		link := e.url(map[string]string{
			ACTION:                  ACTION_BLOCK_ADD,
			EDITOR_ID:               e.id,
			EDITOR_NAME:             e.name,
			EDITOR_HANDLER_ENDPOINT: e.handleEndpoint,
			BLOCK_TYPE:              d.Type,
			e.name:                  blocksJSON,
			"parent_id":             parentID,
			"at_position":           atPosition,
		})

		return bs.Card().
			HxPost(link).
			HxTarget("#" + e.id + "_wrapper").
			HxSwap(`outerHTML`).
			Class("w-100 h-100").
			Style(`cursor:pointer`).
			Child(bs.CardBody().
				Style(`text-align:center;`).
				Child(hb.Div().Child(d.Icon)).
				Child(hb.Heading6().
					Text(d.Type)))
	})

	buttonClose := hb.Button().
		Type("button").
		Child(hb.I().Class(`bi bi-chevron-left me-2`)).
		Text("Close").
		Class("btn btn-secondary").
		Data("bs-dismiss", "modal").
		OnClick(modalCloseScript)

	modal := bs.Modal().
		ID("ModalBlockUpdate").
		Class("fade show").
		Style(`display:block;position:fixed;top:50%;left:50%;transform:translate(-50%,-50%);z-index:1051;`).
		Children([]hb.TagInterface{
			bs.ModalDialog().
				Children([]hb.TagInterface{
					bs.ModalContent().Children([]hb.TagInterface{
						bs.ModalHeader().Children([]hb.TagInterface{
							hb.Heading5().
								Class("modal-title").
								Text("Add Block"),
							hb.Button().
								Type("button").
								Class("btn-close").
								Data("bs-dismiss", "modal").
								OnClick(modalCloseScript),
						}),

						bs.ModalBody().
							Child(bs.Row().
								Class("g-3").
								Children(lo.Map(blockTiles, func(tile hb.TagInterface, _ int) hb.TagInterface {
									return bs.Column(4).Child(tile)
								}))),

						bs.ModalFooter().
							Style("display:flex; justify-content: space-between;").
							Children([]hb.TagInterface{
								buttonClose,

								// buttonUpdate,
							}),
					}),
				}),
		})

	backdrop := hb.Div().
		ID("ModalBackdrop").
		Class("modal-backdrop fade show").
		Style("display:block;")

	return hb.Wrap().Children([]hb.TagInterface{
		modal,
		backdrop,
	}).ToHTML()
}
