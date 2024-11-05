package blockeditor

import "github.com/gouniverse/form"

// FieldsHTML is a predefined set of fields common to all HTML elements
// - HTML ID (html_id)
// - HTML Class (html_class)
// - HTML Style (html_style)
//
// Parameters:
// - none
//
// Returns:
// - []form.Field - The fields
func FieldsHTML() []form.FieldInterface {
	return []form.FieldInterface{
		form.NewField(form.FieldOptions{
			Type:  form.FORM_FIELD_TYPE_RAW,
			Value: `<fieldset><legend>HTML Settings</legend>`,
		}),
		form.NewField(form.FieldOptions{
			Name:  "html_id",
			Label: "HTML ID",
			Type:  form.FORM_FIELD_TYPE_STRING,
			Help:  `Use this field to add an ID to the HTML element.`,
		}),
		form.NewField(form.FieldOptions{
			Name:  "html_class",
			Label: "HTML Class",
			Type:  form.FORM_FIELD_TYPE_STRING,
			Help:  `Use this field to add classes to the HTML element.`,
		}),
		form.NewField(form.FieldOptions{
			Name:  "html_style",
			Label: "HTML Style",
			Type:  form.FORM_FIELD_TYPE_TEXTAREA,
			Help:  `Use this field to add inline styles to the HTML element.`,
		}),
		form.NewField(form.FieldOptions{
			Type:  form.FORM_FIELD_TYPE_RAW,
			Value: `</fieldset>`,
		}),
	}
}

// FieldsAlign is a predefined set of alignment fields
// - Text Align (text_align)
// - Vertical Align (vertical_align)
//
// Parameters:
// - none
//
// Returns:
// - []form.Field - The fields
func FieldsAlign() []form.FieldInterface {
	return []form.FieldInterface{
		form.NewField(form.FieldOptions{
			Type:  form.FORM_FIELD_TYPE_RAW,
			Value: `<fieldset><legend>Alignment Settings</legend>`,
		}),
		form.NewField(form.FieldOptions{
			Name:  "text_align",
			Label: "Text Align",
			Type:  form.FORM_FIELD_TYPE_SELECT,
			Help:  `The text alignment of the block, i.e. center.`,
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
		}),
		form.NewField(form.FieldOptions{
			Name:  "vertical_align",
			Label: "Vertical Align",
			Type:  form.FORM_FIELD_TYPE_SELECT,
			Help:  `The vertical alignment of the block, i.e. baseline.`,
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
		}),
		form.NewField(form.FieldOptions{
			Type:  form.FORM_FIELD_TYPE_RAW,
			Value: `</fieldset>`,
		}),
	}
}

// FieldsBackground is a predefined set of background fields
// - Background Color (background_color)
// - Background Image URL (background_image_url)
// - Background Attachment (background_attachment)
// - Background Position (background_position)
// - Background Repeat (background_repeat)
// - Background Size (background_size)
//
// Parameters:
// - none
//
// Returns:
// - []form.Field - The fields
func FieldsBackground() []form.FieldInterface {
	return []form.FieldInterface{
		form.NewField(form.FieldOptions{
			Type:  form.FORM_FIELD_TYPE_RAW,
			Value: `<fieldset><legend>Background Settings</legend>`,
		}),
		form.NewField(form.FieldOptions{
			Name:  "background_color",
			Label: "Background Color",
			Type:  form.FORM_FIELD_TYPE_STRING,
			Help:  `The background color of the block, i.e. #000000.`,
		}),
		form.NewField(form.FieldOptions{
			Name:  "background_image_url",
			Label: "Background Image URL",
			Type:  form.FORM_FIELD_TYPE_STRING,
			Help:  `The background image URL of the block, i.e. https://example.com/image.png.`,
		}),
		form.NewField(form.FieldOptions{
			Name:  "background_attachment",
			Label: "Background Attachment",
			Type:  form.FORM_FIELD_TYPE_SELECT,
			Help:  `The background attachment of the block, i.e. fixed.`,
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
		}),
		form.NewField(form.FieldOptions{
			Name:  "background_repeat",
			Label: "Background Repeat",
			Type:  form.FORM_FIELD_TYPE_SELECT,
			Help:  `The background image repeat of the block, i.e. no-repeat.`,
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
		}),
		form.NewField(form.FieldOptions{
			Name:  "background_position",
			Label: "Background Position",
			Type:  form.FORM_FIELD_TYPE_SELECT,
			Help:  `The background image position of the block, i.e. center`,
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
		}),
		form.NewField(form.FieldOptions{
			Name:  "background_size",
			Label: "Background Size",
			Type:  form.FORM_FIELD_TYPE_SELECT,
			Help:  `The background image size of the block, i.e. cover.`,
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
		}),
		form.NewField(form.FieldOptions{
			Type:  form.FORM_FIELD_TYPE_RAW,
			Value: `</fieldset>`,
		}),
	}
}

// FieldsBorder is a predefined set of border fields
// - Border (border)
//
// Parameters:
// - none
//
// Returns:
// - []form.Field - The fields
func FieldsBorder() []form.FieldInterface {
	return []form.FieldInterface{
		form.NewField(form.FieldOptions{
			Type:  form.FORM_FIELD_TYPE_RAW,
			Value: `<fieldset><legend>Border Settings</legend>`,
		}),
		form.NewField(form.FieldOptions{
			Name:  "border",
			Label: "Border",
			Type:  form.FORM_FIELD_TYPE_STRING,
			Help:  `The border of the block, i.e. 1px solid #000000.`,
		}),
		form.NewField(form.FieldOptions{
			Type:  form.FORM_FIELD_TYPE_RAW,
			Value: `</fieldset>`,
		}),
	}
}

// FieldsFont is a predefined set of font fields
// - Font Family (font_family)
// - Font Size (font_size)
// - Font Weight (font_weight)
// - Font Style (font_style)
//
// Parameters:
// - none
//
// Returns:
// - []form.Field - The fields
func FieldsFont() []form.FieldInterface {
	return []form.FieldInterface{
		form.NewField(form.FieldOptions{
			Type:  form.FORM_FIELD_TYPE_RAW,
			Value: `<fieldset><legend>Font Settings</legend>`,
		}),
		form.NewField(form.FieldOptions{
			Name:  "font_family",
			Label: "Font Family",
			Type:  form.FORM_FIELD_TYPE_STRING,
			Help:  `The font family of the text.`,
		}),
		form.NewField(form.FieldOptions{
			Name:  "font_size",
			Label: "Font Size",
			Type:  form.FORM_FIELD_TYPE_STRING,
			Help:  `The font size of the text i.e. 16px.`,
		}),
		form.NewField(form.FieldOptions{
			Name:  "font_style",
			Label: "Font Style",
			Type:  form.FORM_FIELD_TYPE_SELECT,
			Options: []form.FieldOption{
				{
					Value: "",
					Key:   "",
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
					Value: "italic",
					Key:   "italic",
				},
				{
					Value: "normal",
					Key:   "normal",
				},
				{
					Value: "oblique",
					Key:   "oblique",
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
		}),
		form.NewField(form.FieldOptions{
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
			Help: `The font weight of the text i.e. normal, bold, etc.`,
		}),
		form.NewField(form.FieldOptions{
			Type:  form.FORM_FIELD_TYPE_RAW,
			Value: `</fieldset>`,
		}),
	}
}

// FieldsMargin is a predefined set of margin fields
// - Margin Top (margin_top)
// - Margin Bottom (margin_bottom)
// - Margin Left (margin_left)
// - Margin Right (margin_right)
//
// Parameters:
// - none
//
// Returns:
// - []form.Field - The fields
func FieldsMargin() []form.FieldInterface {
	return []form.FieldInterface{
		form.NewField(form.FieldOptions{
			Type:  form.FORM_FIELD_TYPE_RAW,
			Value: `<fieldset><legend>Margin Settings</legend>`,
		}),
		form.NewField(form.FieldOptions{
			Name:  "margin_top",
			Label: "Margin Top",
			Type:  form.FORM_FIELD_TYPE_STRING,
			Help:  `The margin top of the block, i.e. 20px.`,
		}),
		form.NewField(form.FieldOptions{
			Name:  "margin_bottom",
			Label: "Margin Bottom",
			Type:  form.FORM_FIELD_TYPE_STRING,
			Help:  `The margin bottom of the block, i.e. 20px.`,
		}),
		form.NewField(form.FieldOptions{
			Name:  "margin_left",
			Label: "Margin Left",
			Type:  form.FORM_FIELD_TYPE_STRING,
			Help:  `The margin left of the block, i.e. 20px.`,
		}),
		form.NewField(form.FieldOptions{
			Name:  "margin_right",
			Label: "Margin Right",
			Type:  form.FORM_FIELD_TYPE_STRING,
			Help:  `The margin right of the block, i.e. 20px.`,
		}),
		form.NewField(form.FieldOptions{
			Type:  form.FORM_FIELD_TYPE_RAW,
			Value: `</fieldset>`,
		}),
	}
}

// FieldsPadding is a predefined set of padding fields
// - Padding Top (padding_top)
// - Padding Bottom (padding_bottom)
// - Padding Left (padding_left)
// - Padding Right (padding_right)
//
// Parameters:
// - none
//
// Returns:
// - []form.Field - The fields
func FieldsPadding() []form.FieldInterface {
	return []form.FieldInterface{
		form.NewField(form.FieldOptions{
			Type:  form.FORM_FIELD_TYPE_RAW,
			Value: `<fieldset><legend>Padding Settings</legend>`,
		}),
		form.NewField(form.FieldOptions{
			Name:  "padding_top",
			Label: "Padding Top",
			Type:  form.FORM_FIELD_TYPE_STRING,
			Help:  `The padding top of the block, i.e. 20px.`,
		}),
		form.NewField(form.FieldOptions{
			Name:  "padding_bottom",
			Label: "Padding Bottom",
			Type:  form.FORM_FIELD_TYPE_STRING,
			Help:  `The padding bottom of the block, i.e. 20px.`,
		}),
		form.NewField(form.FieldOptions{
			Name:  "padding_left",
			Label: "Padding Left",
			Type:  form.FORM_FIELD_TYPE_STRING,
			Help:  `The padding left of the block, i.e. 20px.`,
		}),
		form.NewField(form.FieldOptions{
			Name:  "padding_right",
			Label: "Padding Right",
			Type:  form.FORM_FIELD_TYPE_STRING,
			Help:  `The padding right of the block, i.e. 20px.`,
		}),
		form.NewField(form.FieldOptions{
			Type:  form.FORM_FIELD_TYPE_RAW,
			Value: `</fieldset>`,
		}),
	}
}

// FieldsText is a predefined set of text fields
// - Text Color (text_color)
// - Line Height (line_height)
//
// Parameters:
// - none
//
// Returns:
// - []form.Field - The fields
func FieldsText() []form.FieldInterface {
	return []form.FieldInterface{
		form.NewField(form.FieldOptions{
			Type:  form.FORM_FIELD_TYPE_RAW,
			Value: `<fieldset><legend>Text Settings</legend>`,
		}),
		form.NewField(form.FieldOptions{
			Name:  "text_color",
			Label: "Text Color",
			Type:  form.FORM_FIELD_TYPE_STRING,
			Help:  `The color of the text.`,
		}),
		form.NewField(form.FieldOptions{
			Name:  "line_height",
			Label: "Line Height",
			Type:  form.FORM_FIELD_TYPE_STRING,
			Help:  `The line height of the text i.e. 16px.`,
		}),
		form.NewField(form.FieldOptions{
			Type:  form.FORM_FIELD_TYPE_RAW,
			Value: `</fieldset>`,
		}),
	}
}
