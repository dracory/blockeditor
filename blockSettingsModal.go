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

	block := NewFlatTree(b.blocks).Find(blockID)

	if block == nil {
		return "no block found"
	}

	definition, found := lo.Find(b.blockDefinitions, func(d BlockDefinition) bool {
		return d.Type == block.Type
	})

	fields := lo.If(found, definition.Fields).Else([]form.Field{})

	// fields = b.settingFields(fields)

	fieldsWithPrefix := lo.Map(fields, func(f form.Field, _ int) form.Field {
		if f.Type == form.FORM_FIELD_TYPE_RAW {
			return f
		}

		// calculate the value before adding the prefix
		f.Value = block.Parameters[f.Name]

		// add prefix to not conflict with other form fields (i.e. content)
		f.Name = SETTINGS_PREFIX + f.Name
		return f
	})

	blockForm := form.NewForm(form.FormOptions{
		Fields: fieldsWithPrefix,
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
		Class("fade show modal-lg").
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
							Child(hb.NewDiv().
								Style("height: 300px; overflow-y: auto; pading-right: 5px;").
								Child(hb.NewDiv().
									Class(`alert alert-info`).
									Style("font-size: 16px;").
									Child(hb.Span().Text("Block: ")).
									Child(hb.Span().Text(block.Type).Style(`font-weight: bold;`)).
									Child(hb.Sup().
										Style("float:right;font-size: 11px;").
										Text("ID: " + block.ID))).
								Child(blockForm.Build())),

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
