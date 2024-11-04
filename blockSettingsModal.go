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

	flatBlock := NewFlatTree(b.blocks).Find(blockID)

	if flatBlock == nil {
		return "no block found"
	}

	definition, found := lo.Find(b.blockDefinitions, func(d BlockDefinition) bool {
		return d.Type == flatBlock.Type
	})

	fields := lo.If(found, definition.Fields).Else([]form.Field{})

	fields = b.settingFields(fields)

	fieldsWithPrefix := lo.Map(fields, func(f form.Field, _ int) form.Field {
		if f.Type == form.FORM_FIELD_TYPE_RAW {
			return f
		}

		// calculate the value before adding the prefix
		f.Value = flatBlock.Parameters[f.Name]

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
								Child(blockForm.Build()).
								Child(hb.NewDiv().
									Style("font-size: 0.8rem;margin-top: 1rem;").
									Child(hb.NewSpan().Text("Block ID: " + blockID)))),

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

func (b *editor) settingFields(blockFields []form.Field) []form.Field {
	fields := []form.Field{}

	if len(blockFields) > 0 {
		fields = append(fields, form.Field{
			Type:  form.FORM_FIELD_TYPE_RAW,
			Value: `<fieldset><legend>Block Settings</legend>`,
		})
	}

	fields = append(fields, blockFields...)

	if len(blockFields) > 0 {
		fields = append(fields, form.Field{
			Type:  form.FORM_FIELD_TYPE_RAW,
			Value: `</fieldset>`,
		})
	}

	commonFields := []form.Field{
		{
			Name:  "html_id",
			Label: "HTML ID",
			Type:  form.FORM_FIELD_TYPE_STRING,
		},
		{
			Name:  "html_class",
			Label: "HTML Class",
			Type:  form.FORM_FIELD_TYPE_STRING,
		},
		{
			Name:  "text_color",
			Label: "Text Color",
			Type:  form.FORM_FIELD_TYPE_STRING,
		},
		{
			Name:  "text_align",
			Label: "Text Align",
			Type:  form.FORM_FIELD_TYPE_SELECT,
			Options: []form.FieldOption{
				{
					Value: "left",
					Key:   "left",
				},
				{
					Value: "center",
					Key:   "center",
				},
				{
					Value: "right",
					Key:   "right",
				},
			},
		},
		{
			Name:  "vertical_align",
			Label: "Vertical Align",
			Type:  form.FORM_FIELD_TYPE_SELECT,
			Options: []form.FieldOption{
				{
					Value: "top",
					Key:   "top",
				},
				{
					Value: "middle",
					Key:   "middle",
				},
				{
					Value: "bottom",
					Key:   "bottom",
				},
			},
		},
		{
			Name:  "font_size",
			Label: "Font Size",
			Type:  form.FORM_FIELD_TYPE_STRING,
		},
		{
			Name:  "font_family",
			Label: "Font Family",
			Type:  form.FORM_FIELD_TYPE_STRING,
		},
		{
			Name:  "font_weight",
			Label: "Font Weight",
			Type:  form.FORM_FIELD_TYPE_SELECT,
			Options: []form.FieldOption{
				{
					Value: "",
					Key:   "",
				},
				{
					Value: "normal",
					Key:   "normal",
				},
				{
					Value: "bold",
					Key:   "bold",
				},
				{
					Value: "bolder",
					Key:   "bolder",
				},
				{
					Value: "light",
					Key:   "light",
				},
				{
					Value: "lighter",
					Key:   "lighter",
				},
				{
					Value: "100",
					Key:   "100",
				},
				{
					Value: "200",
					Key:   "200",
				},
				{
					Value: "300",
					Key:   "300",
				},
				{
					Value: "400",
					Key:   "400",
				},
				{
					Value: "500",
					Key:   "500",
				},
				{
					Value: "600",
					Key:   "600",
				},
				{
					Value: "700",
					Key:   "700",
				},
				{
					Value: "800",
					Key:   "800",
				},
				{
					Value: "900",
					Key:   "900",
				},
			},
		},
	}

	fields = append(fields, form.Field{
		Type:  form.FORM_FIELD_TYPE_RAW,
		Value: `<fieldset><legend>Common Settings</legend>`,
	})

	fields = append(fields, commonFields...)

	fields = append(fields, form.Field{
		Type:  form.FORM_FIELD_TYPE_RAW,
		Value: `</fieldset>`,
	})

	marginFields := []form.Field{
		{
			Name:  "margin_top",
			Label: "Margin Top",
			Type:  form.FORM_FIELD_TYPE_STRING,
		},
		{
			Name:  "margin_bottom",
			Label: "Margin Bottom",
			Type:  form.FORM_FIELD_TYPE_STRING,
		},
		{
			Name:  "margin_left",
			Label: "Margin Left",
			Type:  form.FORM_FIELD_TYPE_STRING,
		},
		{
			Name:  "margin_right",
			Label: "Margin Right",
			Type:  form.FORM_FIELD_TYPE_STRING,
		},
	}

	fields = append(fields, form.Field{
		Type:  form.FORM_FIELD_TYPE_RAW,
		Value: `<fieldset><legend>Margin Settings</legend>`,
	})

	fields = append(fields, marginFields...)

	fields = append(fields, form.Field{
		Type:  form.FORM_FIELD_TYPE_RAW,
		Value: `</fieldset>`,
	})

	paddingFields := []form.Field{
		{
			Name:  "padding_top",
			Label: "Padding Top",
			Type:  form.FORM_FIELD_TYPE_STRING,
		},
		{
			Name:  "padding_bottom",
			Label: "Padding Bottom",
			Type:  form.FORM_FIELD_TYPE_STRING,
		},
		{
			Name:  "padding_left",
			Label: "Padding Left",
			Type:  form.FORM_FIELD_TYPE_STRING,
		},
		{
			Name:  "padding_right",
			Label: "Padding Right",
			Type:  form.FORM_FIELD_TYPE_STRING,
		},
	}

	fields = append(fields, form.Field{
		Type:  form.FORM_FIELD_TYPE_RAW,
		Value: `<fieldset><legend>Padding Settings</legend>`,
	})

	fields = append(fields, paddingFields...)

	fields = append(fields, form.Field{
		Type:  form.FORM_FIELD_TYPE_RAW,
		Value: `</fieldset>`,
	})

	backgroundFields := []form.Field{
		{
			Name:  "background_color",
			Label: "Background Color",
			Type:  form.FORM_FIELD_TYPE_STRING,
		},
		{
			Name:  "background_image_url",
			Label: "Background Image URL",
			Type:  form.FORM_FIELD_TYPE_STRING,
		},

		{
			Name:  "background_attachment",
			Label: "Background Attachment",
			Type:  form.FORM_FIELD_TYPE_SELECT,
			Options: []form.FieldOption{
				{
					Value: "",
					Key:   "",
				},
				{
					Value: "fixed",
					Key:   "fixed",
				},
				{
					Value: "inherit",
					Key:   "inherit",
				},
				{
					Value: "initial",
					Key:   "initial",
				},
				{
					Value: "local",
					Key:   "local",
				},
				{
					Value: "revert",
					Key:   "revert",
				},
				{
					Value: "scroll",
					Key:   "scroll",
				},
				{
					Value: "unset",
					Key:   "unset",
				},
			},
		},
		{
			Name:  "background_repeat",
			Label: "Background Repeat",
			Type:  form.FORM_FIELD_TYPE_SELECT,
			Options: []form.FieldOption{
				{
					Value: "",
					Key:   "",
				},
				{
					Value: "repeat",
					Key:   "repeat",
				},
				{
					Value: "repeat-x",
					Key:   "repeat-x",
				},
				{
					Value: "repeat-y",
					Key:   "repeat-y",
				},
				{
					Value: "no-repeat",
					Key:   "no-repeat",
				},
			},
		},
		{
			Name:  "background_position",
			Label: "Background Position",
			Type:  form.FORM_FIELD_TYPE_SELECT,
			Options: []form.FieldOption{
				{
					Value: "",
					Key:   "",
				},
				{
					Value: "bottom",
					Key:   "bottom",
				},
				{
					Value: "center",
					Key:   "center",
				},
				{
					Value: "inherit",
					Key:   "inherit",
				},
				{
					Value: "initial",
					Key:   "initial",
				},
				{
					Value: "left",
					Key:   "left",
				},
				{
					Value: "revert",
					Key:   "revert",
				},
				{
					Value: "right",
					Key:   "right",
				},
				{
					Value: "top",
					Key:   "top",
				},
				{
					Value: "unset",
					Key:   "unset",
				},
			},
		},
		{
			Name:  "background_size",
			Label: "Background Size",
			Type:  form.FORM_FIELD_TYPE_SELECT,
			Options: []form.FieldOption{
				{
					Value: "",
					Key:   "",
				},
				{
					Value: "auto",
					Key:   "auto",
				},
				{
					Value: "cover",
					Key:   "cover",
				},
				{
					Value: "contain",
					Key:   "contain",
				},
				{
					Value: "inherit",
					Key:   "inherit",
				},
				{
					Value: "initial",
					Key:   "initial",
				},
				{
					Value: "revert",
					Key:   "revert",
				},
				{
					Value: "unset",
					Key:   "unset",
				},
			},
		},
	}

	fields = append(fields, form.Field{
		Type:  form.FORM_FIELD_TYPE_RAW,
		Value: `<fieldset><legend>Background Settings</legend>`,
	})

	fields = append(fields, backgroundFields...)

	fields = append(fields, form.Field{
		Type:  form.FORM_FIELD_TYPE_RAW,
		Value: `</fieldset>`,
	})

	alignFields := []form.Field{
		{
			Name:  "text_align",
			Label: "Text Align",
			Type:  form.FORM_FIELD_TYPE_SELECT,
			Options: []form.FieldOption{
				{
					Value: "",
					Key:   "",
				},
				{
					Value: "center",
					Key:   "center",
				},
				{
					Value: "inherit",
					Key:   "inherit",
				},
				{
					Value: "initial",
					Key:   "initial",
				},
				{
					Value: "justify",
					Key:   "justify",
				},
				{
					Value: "left",
					Key:   "left",
				},
				{
					Value: "right",
					Key:   "right",
				},
				{
					Value: "revert",
					Key:   "revert",
				},
				{
					Value: "unset",
					Key:   "unset",
				},
			},
		},
		{
			Name:  "vertical_align",
			Label: "Vertical Align",
			Type:  form.FORM_FIELD_TYPE_SELECT,
			Options: []form.FieldOption{
				{
					Value: "",
					Key:   "",
				},
				{
					Value: "baseline",
					Key:   "baseline",
				},
				{
					Value: "sub",
					Key:   "sub",
				},
				{
					Value: "super",
					Key:   "super",
				},
				{
					Value: "revert",
					Key:   "revert",
				},
				{
					Value: "unset",
					Key:   "unset",
				},
			},
		},
	}

	fields = append(fields, form.Field{
		Type:  form.FORM_FIELD_TYPE_RAW,
		Value: `<fieldset><legend>Alignment Settings</legend>`,
	})

	fields = append(fields, alignFields...)

	fields = append(fields, form.Field{
		Type:  form.FORM_FIELD_TYPE_RAW,
		Value: `</fieldset>`,
	})

	fontFields := []form.Field{
		{
			Name:  "font_family",
			Label: "Font Family",
			Type:  form.FORM_FIELD_TYPE_STRING,
		},
		{
			Name:  "font_size",
			Label: "Font Size",
			Type:  form.FORM_FIELD_TYPE_STRING,
		},
	}

	fields = append(fields, form.Field{
		Type:  form.FORM_FIELD_TYPE_RAW,
		Value: `<fieldset><legend>Font Settings</legend>`,
	})

	fields = append(fields, fontFields...)

	fields = append(fields, form.Field{
		Type:  form.FORM_FIELD_TYPE_RAW,
		Value: `</fieldset>`,
	})

	advancedFields := []form.Field{
		{
			Name:  "html_style",
			Label: "HTML Style",
			Type:  form.FORM_FIELD_TYPE_TEXTAREA,
		},
	}

	fields = append(fields, form.Field{
		Type:  form.FORM_FIELD_TYPE_RAW,
		Value: `<fieldset><legend>Advanced Settings</legend>`,
	})

	fields = append(fields, advancedFields...)

	fields = append(fields, form.Field{
		Type:  form.FORM_FIELD_TYPE_RAW,
		Value: `</fieldset>`,
	})

	return fields
}
