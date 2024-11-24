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
		return hb.Wrap().
			Child(hb.Swal(hb.SwalOptions{
				Icon:  "error",
				Title: "Error",
				Text:  "No block id",
			})).
			ToHTML()
	}

	block := NewFlatTree(b.blocks).Find(blockID)

	if block == nil {
		return hb.Wrap().
			Child(hb.Swal(hb.SwalOptions{
				Icon:  "error",
				Title: "Error",
				Text:  "No block found",
			})).
			ToHTML()
	}

	definition := b.findDefinitionByType(block.Type)

	fields := lo.IfF(definition != nil, func() []form.FieldInterface { return definition.Fields }).
		Else([]form.FieldInterface{})

	// add the status field, if not already present
	hasStatus := lo.SomeBy(fields, func(f form.FieldInterface) bool { return f.GetName() == "status" })

	if !hasStatus {
		fields = append([]form.FieldInterface{form.NewField(form.FieldOptions{
			Name:  "status",
			Label: "Status",
			Type:  form.FORM_FIELD_TYPE_SELECT,
			Options: []form.FieldOption{
				{
					Value: "",
					Key:   "",
				},
				{
					Value: "Published",
					Key:   "published",
				},
				{
					Value: "Unpublished",
					Key:   "unpublished",
				},
			},
			Help: "The status of the block.",
		})}, fields...)
	}

	fieldsWithPrefix := lo.Map(fields, func(f form.FieldInterface, _ int) form.FieldInterface {
		fieldName := f.GetName()
		newField := form.NewField(form.FieldOptions{})

		// if the field is not a raw field, add the prefix
		if f.GetType() != form.FORM_FIELD_TYPE_RAW {
			settingsFieldName := prefixKey(fieldName, SETTINGS_PREFIX)
			newField.SetName(settingsFieldName)
		}

		newField.SetID(f.GetID())
		newField.SetLabel(f.GetLabel())
		newField.SetType(f.GetType())
		newField.SetHelp(f.GetHelp())

		// if the field is a raw field, set the original value
		if f.GetType() == form.FORM_FIELD_TYPE_RAW {
			newField.SetValue(f.GetValue())
		} else {
			fieldValue := block.Parameters[fieldName]
			newField.SetValue(fieldValue)
		}

		newField.SetOptions(f.GetOptions())
		newField.SetOptionsF(f.GetOptionsF())
		newField.SetRequired(f.GetRequired())

		if newField.GetType() == form.FORM_FIELD_TYPE_RAW {
			return newField
		}

		return newField
	})

	blockForm := form.NewForm(form.FormOptions{
		Fields: fieldsWithPrefix,
	})

	blocksJSON, err := ui.MarshalBlocksToJson(b.blocks)

	if err != nil {
		return err.Error()
	}

	blockForm.AddField(&form.Field{
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

	modal := bs.Modal().
		ID("ModalBlockUpdate").
		Class("fade show modal-xl").
		Style(`display:block;position:fixed;top:50%;left:50%;transform:translate(-50%,-50%);z-index:1051;`).
		Children([]hb.TagInterface{
			bs.ModalDialog().
				Class("modal-dialog").
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
								Style("height: 600px; overflow-y: auto; pading-right: 5px;").
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
		hb.Style(`
#ModalBlockUpdate fieldset {
	border: 1px solid #ced4da;
	border-radius: 5px;
	padding: 10px;
	margin-bottom: 20px;
	background-color: honeydew;
}

#ModalBlockUpdate fieldset legend {
	float: none;
	width: auto;
	font-weight: bold;
	font-size: 24px;
	padding: 0 10px;
	border: 1px solid #ced4da;
	border-radius: 10px;
	background-color: aliceblue;
}
`),
		modal,
		backdrop,
	}).ToHTML()
}
