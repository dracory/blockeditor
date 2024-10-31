package blockeditor

import (
	"net/http"

	"github.com/gouniverse/bs"
	"github.com/gouniverse/form"
	"github.com/gouniverse/hb"
	"github.com/gouniverse/ui"
	"github.com/gouniverse/utils"
	"github.com/samber/lo"
)

// blockSettingsModal shows the block settings modal
func (b *editor) blockSettingsModal(r *http.Request) string {
	blockID := utils.Req(r, BLOCK_ID, "")

	if blockID == "" {
		return "no block id"
	}

	block := b.blockFindByID(blockID)

	if block == nil {
		return "no block found"
	}

	definition, found := lo.Find(b.blockDefinitions, func(d BlockDefinition) bool {
		return d.Type == block.Type()
	})

	blockForm := form.NewForm(form.FormOptions{
		Fields: lo.IfF(found, func() []form.Field {
			fields := lo.Map(definition.Fields, func(f form.Field, _ int) form.Field {
				f.Value = block.Parameter(f.Name)
				// we add the settings prefix to not conflict with other form fields (i.e. content)
				f.Name = SETTINGS_PREFIX + f.Name
				print(f.Name, f.Value)
				return f
			})
			return fields
		}).Else([]form.Field{}),
	})

	blocksJSON, err := ui.BlocksToJson(b.blocks)

	if err != nil {
		return err.Error()
	}

	blockForm.AddField(form.Field{
		Name:      b.name,
		Label:     "Editor Blocks",
		Type:      form.FORM_FIELD_TYPE_TEXTAREA,
		Value:     blocksJSON,
		Invisible: true,
	})

	modalCloseScript := `document.getElementById('ModalBlockUpdate').remove();document.getElementById('ModalBackdrop').remove();`

	buttonClose := hb.Button().
		Type("button").
		Child(hb.I().Class(`bi bi-chevron-left me-2`)).
		Text("Close").
		Class("btn btn-secondary").
		Data("bs-dismiss", "modal").
		OnClick(modalCloseScript)

	buttonUpdate := bs.ButtonLink().
		Class("btn btn-success").
		Child(hb.I().Class(`bi bi-check me-2`)).
		Text("Update").
		HxPost(b.url(map[string]string{
			EDITOR_ID:               b.id,
			EDITOR_NAME:             b.name,
			EDITOR_HANDLER_ENDPOINT: b.handleEndpoint,
			ACTION:                  ACTION_BLOCK_SETTINGS_UPDATE,
			BLOCK_ID:                blockID,
		})).
		HxInclude("#ModalBlockUpdate").
		HxTarget(`#` + b.id + `_wrapper`).
		HxSwap(`outerHTML`)
		// OnClick(modalCloseScript)

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
								Text("Update Block Settings"),
							hb.Button().
								Type("button").
								Class("btn-close").
								Data("bs-dismiss", "modal").
								OnClick(modalCloseScript),
						}),

						bs.ModalBody().
							Child(blockForm.Build()).
							Child(hb.NewDiv().
								Style("font-size: 0.8rem;margin-top: 1rem;").
								Child(hb.NewSpan().Text("Block ID: " + blockID))),

						bs.ModalFooter().
							Style("display:flex; justify-content: space-between;").
							Children([]hb.TagInterface{
								buttonClose,

								buttonUpdate,
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
