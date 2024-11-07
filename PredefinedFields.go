package blockeditor

import (
	"github.com/gouniverse/form"
	"github.com/gouniverse/hb"
	"github.com/gouniverse/ui"
)

func collapsibleStart(groupId string, groupName string, collapsed bool) string {
	collapsedClass := "collapsed"
	if collapsed {
		collapsedClass = "collapsed show"
	}
	// using card
	return `<div class="card mb-3">
		<div class="card-header" data-bs-target="#` + groupId + `" data-bs-toggle="collapse" style="cursor:pointer; background-color: #CCEBF8;">
		    <h5>
				<i class="bi bi-arrows-collapse"></i>
				` + groupName + `
			</h5>
		</div>
		<div class="card-body ` + collapsedClass + `" id="` + groupId + `" style="background-color: #E5F5FC;">`

	// // using fieldset
	// return `<fieldset data-bs-target="#` + groupId + `" data-bs-toggle="collapse"><legend><i class="bi bi-arrows-collapse"></i> ` + groupName + `</legend><div id="` + groupId + `" class="collapse">`
}

func collapsibleEnd() string {
	// using card
	return `</div></div>`

	// // using fieldset
	// return `</div></fieldset>`
}

func FieldGroupStart(groupId string, groupName string, collapsed bool) form.FieldInterface {
	return form.NewField(form.FieldOptions{
		Type:  form.FORM_FIELD_TYPE_RAW,
		Value: collapsibleStart(groupId, groupName, collapsed),
	})
}

func FieldGroupEnd() form.FieldInterface {
	return form.NewField(form.FieldOptions{
		Type:  form.FORM_FIELD_TYPE_RAW,
		Value: collapsibleEnd(),
	})
}

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
			Value: collapsibleStart("html_fields", "HTML Settings", true),
		}),
		// form.NewField(form.FieldOptions{
		// 	Type: form.FORM_FIELD_TYPE_RAW,
		// 	Value: `<div class="card pb-3">
		// 		<div class="card-header" data-bs-target="#html_fields" data-bs-toggle="collapse">
		// 			<i class="bi bi-arrows-collapse"></i>
		// 			HTML Settings
		// 		</div>
		// 		<div class="card-body collapse" id="html_fields">`,
		// }),
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
			Value: collapsibleEnd(),
		}),
	}
}

func ApplyHTMLParameters(block ui.BlockInterface, blockTag *hb.Tag) {
	id := block.Parameter("html_id")
	class := block.Parameter("html_class")
	style := block.Parameter("html_style")
	color := block.Parameter("text_color")

	if id != "" {
		blockTag.ID(id)
	}

	if class != "" {
		blockTag.Class(class)
	}

	if style != "" {
		blockTag.Style(style)
	}

	if color != "" {
		blockTag.Style("color:" + color)
	}
}

// FieldsAnimation is a predefined set of animation fields
func FieldsAnimation() []form.FieldInterface {
	return []form.FieldInterface{
		form.NewField(form.FieldOptions{
			Type:  form.FORM_FIELD_TYPE_RAW,
			Value: collapsibleStart("animation_fields", "Animation Settings", true),
		}),
		form.NewField(form.FieldOptions{
			Name:  "animation_delay",
			Label: "Animation Delay",
			Type:  form.FORM_FIELD_TYPE_STRING,
			Help:  `Defines how long the animation should wait before it starts, i.e. 1s.`,
		}),
		form.NewField(form.FieldOptions{
			Name:  "animation_direction",
			Label: "Animation Direction",
			Type:  form.FORM_FIELD_TYPE_SELECT,
			Options: []form.FieldOption{
				{
					Value: "",
					Key:   "",
				},
				{
					Value: "alternate",
					Key:   "alternate",
				},
				{
					Value: "alternate-reverse",
					Key:   "alternate-reverse",
				},
				{
					Value: "normal",
					Key:   "normal",
				},
				{
					Value: "reverse",
					Key:   "reverse",
				},
			},
		}),
		form.NewField(form.FieldOptions{
			Name:  "animation_duration",
			Label: "Animation Duration",
			Type:  form.FORM_FIELD_TYPE_STRING,
			Help:  `Defines how long the animation should take to complete, i.e. 1s.`,
		}),
		form.NewField(form.FieldOptions{
			Name:  "animation-_ill_mode",
			Label: "Animation Fill Mode",
			Type:  form.FORM_FIELD_TYPE_SELECT,
			Options: []form.FieldOption{
				{
					Value: "",
					Key:   "",
				},
				{
					Value: "none",
					Key:   "none",
				},
				{
					Value: "forwards",
					Key:   "forwards",
				},
				{
					Value: "backwards",
					Key:   "backwards",
				},
				{
					Value: "both",
					Key:   "both",
				},
			},
		}),
		form.NewField(form.FieldOptions{
			Name:  "animation_iteration_count",
			Label: "Animation Iteration Count",
			Type:  form.FORM_FIELD_TYPE_STRING,
			Help:  `Defines how many times the animation should be played, i.e. infinite or number.`,
		}),
		form.NewField(form.FieldOptions{
			Name:  "animation_name",
			Label: "Animation Name",
			Type:  form.FORM_FIELD_TYPE_STRING,
			Help:  `Defines the name of the animation, i.e. bounce.`,
		}),
		form.NewField(form.FieldOptions{
			Name:  "animation_play_state",
			Label: "Animation Play State",
			Type:  form.FORM_FIELD_TYPE_SELECT,
			Options: []form.FieldOption{
				{
					Value: "",
					Key:   "",
				},
				{
					Value: "running",
					Key:   "running",
				},
				{
					Value: "paused",
					Key:   "paused",
				},
			},
		}),
		form.NewField(form.FieldOptions{
			Name:  "animation_timing_function",
			Label: "Animation Timing Function",
			Type:  form.FORM_FIELD_TYPE_SELECT,
			Options: []form.FieldOption{
				{
					Value: "",
					Key:   "",
				},
				{
					Value: "ease",
					Key:   "ease",
				},
				{
					Value: "linear",
					Key:   "linear",
				},
				{
					Value: "ease-in",
					Key:   "ease-in",
				},
				{
					Value: "ease-out",
					Key:   "ease-out",
				},
				{
					Value: "ease-in-out",
					Key:   "ease-in-out",
				},
				{
					Value: "step-start",
					Key:   "step-start",
				},
				{
					Value: "step-end",
					Key:   "step-end",
				},
			},
		}),
		form.NewField(form.FieldOptions{
			Type:  form.FORM_FIELD_TYPE_RAW,
			Value: collapsibleEnd(),
		}),
	}
}

func ApplyAnimationParameters(block ui.BlockInterface, blockTag *hb.Tag) {
	animationName := block.Parameter("animation_name")
	animationDelay := block.Parameter("animation_delay")
	animationDuration := block.Parameter("animation_duration")
	animationTimingFunction := block.Parameter("animation_timing_function")
	animationIterationCount := block.Parameter("animation_iteration_count")
	animationDirection := block.Parameter("animation_direction")
	animationFillMode := block.Parameter("animation-_ill_mode")
	animationPlayState := block.Parameter("animation_play_state")

	blockTag.StyleIf(animationName != "", "animation-name:"+animationName)
	blockTag.StyleIf(animationDelay != "", "animation-delay:"+animationDelay)
	blockTag.StyleIf(animationDuration != "", "animation-duration:"+animationDuration)
	blockTag.StyleIf(animationTimingFunction != "", "animation-timing-function:"+animationTimingFunction)
	blockTag.StyleIf(animationIterationCount != "", "animation-iteration-count:"+animationIterationCount)
	blockTag.StyleIf(animationDirection != "", "animation-direction:"+animationDirection)
	blockTag.StyleIf(animationFillMode != "", "animation-fill-mode:"+animationFillMode)
	blockTag.StyleIf(animationPlayState != "", "animation-play-state:"+animationPlayState)
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
			Value: collapsibleStart("align_fields", "Alignment Settings", true),
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
			Value: collapsibleEnd(),
		}),
	}
}

func ApplyAlignmentParameters(block ui.BlockInterface, blockTag *hb.Tag) {
	textAlign := block.Parameter("text_align")
	verticalAlign := block.Parameter("vertical_align")

	blockTag.StyleIf(textAlign != "", "text-align:"+textAlign)
	blockTag.StyleIf(verticalAlign != "", "vertical-align:"+verticalAlign)
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
			Value: collapsibleStart("background_fields", "Background Settings", true),
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
			Value: collapsibleEnd(),
		}),
	}
}

func ApplyBackgroundParameters(block ui.BlockInterface, blockTag *hb.Tag) {
	backgroundColor := block.Parameter("background_color")
	backgroundAttachment := block.Parameter("background_attachment")
	backgroundImageUrl := block.Parameter("background_image_url")
	backgroundRepeat := block.Parameter("background_repeat")
	backgroundPosition := block.Parameter("background_position")
	backgroundSize := block.Parameter("background_size")

	blockTag.StyleIf(backgroundImageUrl != "", "background-image:url("+backgroundImageUrl+")")
	blockTag.StyleIf(backgroundColor != "", "background-color:"+backgroundColor)
	blockTag.StyleIf(backgroundAttachment != "", "background-attachment:"+backgroundAttachment)
	blockTag.StyleIf(backgroundRepeat != "", "background-repeat:"+backgroundRepeat)
	blockTag.StyleIf(backgroundPosition != "", "background-position:"+backgroundPosition)
	blockTag.StyleIf(backgroundSize != "", "background-size:"+backgroundSize)
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
			Value: collapsibleStart("border_fields", "Border Settings", true),
		}),
		form.NewField(form.FieldOptions{
			Name:  "border",
			Label: "Border",
			Type:  form.FORM_FIELD_TYPE_STRING,
			Help:  `The border of the block, i.e. 1px solid #000000.`,
		}),
		form.NewField(form.FieldOptions{
			Type:  form.FORM_FIELD_TYPE_RAW,
			Value: collapsibleEnd(),
		}),
	}
}

func ApplyBorderParameters(block ui.BlockInterface, blockTag *hb.Tag) {
	border := block.Parameter("border")
	blockTag.StyleIf(border != "", "border:"+border)
}

func FieldsFlexBox() []form.FieldInterface {
	return []form.FieldInterface{
		form.NewField(form.FieldOptions{
			Type:  form.FORM_FIELD_TYPE_RAW,
			Value: collapsibleStart("flexbox_fields", "Flexbox Settings", true),
		}),
		// align-content
		form.NewField(form.FieldOptions{
			Name:  "align_content",
			Label: "Align Content",
			Type:  form.FORM_FIELD_TYPE_SELECT,
			Help:  `Defines how each line is aligned in the container, i.e. center. Applies only when the flex-wrap is set to wrap.`,
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
					Value: "flex-end",
					Key:   "flex-end",
				},
				{
					Value: "flex-start",
					Key:   "flex-start",
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
					Value: "normal",
					Key:   "normal",
				},
				{
					Value: "space-around",
					Key:   "space-around",
				},
				{
					Value: "space-between",
					Key:   "space-between",
				},
				{
					Value: "space-evenly",
					Key:   "space-evenly",
				},
				{
					Value: "stretch",
					Key:   "stretch",
				},
			},
		}),
		// align-items
		form.NewField(form.FieldOptions{
			Name:  "align_items",
			Label: "Align Items",
			Type:  form.FORM_FIELD_TYPE_SELECT,
			Help:  `Defines how the flex items are aligned in the cross axis, i.e. center.`,
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
					Value: "flex-end",
					Key:   "flex-end",
				},
				{
					Value: "flex-start",
					Key:   "flex-start",
				},
				{
					Value: "initial",
					Key:   "initial",
				},
				{
					Value: "inherit",
					Key:   "inherit",
				},
				{
					Value: "stretch",
					Key:   "stretch",
				},
			},
		}),
		// align-self
		form.NewField(form.FieldOptions{
			Name:  "align_self",
			Label: "Align Self",
			Type:  form.FORM_FIELD_TYPE_SELECT,
			Help:  `Defines how the flex item is aligned in the cross axis, but apllies only to that item, i.e. center.`,
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
					Value: "flex-end",
					Key:   "flex-end",
				},
				{
					Value: "flex-start",
					Key:   "flex-start",
				},
				{
					Value: "initial",
					Key:   "initial",
				},
				{
					Value: "inherit",
					Key:   "inherit",
				},
				{
					Value: "stretch",
					Key:   "stretch",
				},
			},
		}),
		// flex-basis
		form.NewField(form.FieldOptions{
			Name:  "flex_basis",
			Label: "Flex Basis",
			Type:  form.FORM_FIELD_TYPE_STRING,
			Help:  `Defines the initial size of the flex item, i.e. auto or 40px.`,
		}),
		// flex-direction
		form.NewField(form.FieldOptions{
			Name:  "flex_direction",
			Label: "Flex Direction",
			Type:  form.FORM_FIELD_TYPE_SELECT,
			Help:  `Defines how the flex items are ordered in the flex container, i.e. row.`,
			Options: []form.FieldOption{
				{
					Value: "",
					Key:   "",
				},
				{
					Value: "column",
					Key:   "column",
				},
				{
					Value: "column-reverse",
					Key:   "column-reverse",
				},
				{
					Value: "row",
					Key:   "row",
				},
				{
					Value: "row-reverse",
					Key:   "row-reverse",
				},
			},
		}),
		// flex-flow
		form.NewField(form.FieldOptions{
			Name:  "flex_flow",
			Label: "Flex Flow",
			Type:  form.FORM_FIELD_TYPE_STRING,
			Help:  `Shorthand property for flex-direction and flex-wrap.`,
		}),
		// flex-grow
		form.NewField(form.FieldOptions{
			Name:  "flex_grow",
			Label: "Flex Grow",
			Type:  form.FORM_FIELD_TYPE_NUMBER,
			Help:  `Defines the flex grow factor of the flex item, if there is space available, i.e. 2.`,
		}),
		// flex-shrink
		form.NewField(form.FieldOptions{
			Name:  "flex_shrink",
			Label: "Flex Shrink",
			Type:  form.FORM_FIELD_TYPE_NUMBER,
			Help:  `Defines the flex shrink factor of the flex item, if there is not enough space, i.e. 2.`,
		}),
		// flex-wrap
		form.NewField(form.FieldOptions{
			Name:  "flex_wrap",
			Label: "Flex Wrap",
			Type:  form.FORM_FIELD_TYPE_SELECT,
			Help:  `Defines whether the flex items are forced onto one line or can wrap onto multiple lines, i.e. nowrap.`,
			Options: []form.FieldOption{
				{
					Value: "",
					Key:   "",
				},
				{
					Value: "nowrap",
					Key:   "nowrap",
				},
				{
					Value: "wrap",
					Key:   "wrap",
				},
				{
					Value: "wrap-reverse",
					Key:   "wrap-reverse",
				},
			},
		}),
		// justify-content
		form.NewField(form.FieldOptions{
			Name:  "justify_content",
			Label: "Justify Content",
			Type:  form.FORM_FIELD_TYPE_SELECT,
			Help:  `Defines how the flex items are positioned accross the main axis in the flex container, i.e. center.`,
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
					Value: "flex-end",
					Key:   "flex-end",
				},
				{
					Value: "flex-start",
					Key:   "flex-start",
				},
				{
					Value: "initial",
					Key:   "initial",
				},
				{
					Value: "inherit",
					Key:   "inherit",
				},
				{
					Value: "space-around",
					Key:   "space-around",
				},
				{
					Value: "space-between",
					Key:   "space-between",
				},
				{
					Value: "space-evenly",
					Key:   "space-evenly",
				},
			},
		}),
		// order
		form.NewField(form.FieldOptions{
			Type:  form.FORM_FIELD_TYPE_RAW,
			Value: collapsibleEnd(),
		}),
	}
}

func ApplyFlexBoxParameters(block ui.BlockInterface, blockTag *hb.Tag) {
	alignContent := block.Parameter("align_content")
	blockTag.StyleIf(alignContent != "", "align-content:"+alignContent)

	alignItems := block.Parameter("align_items")
	blockTag.StyleIf(alignItems != "", "align-items:"+alignItems)

	alignSelf := block.Parameter("align_self")
	blockTag.StyleIf(alignSelf != "", "align-self:"+alignSelf)

	flexBasis := block.Parameter("flex_basis")
	blockTag.StyleIf(flexBasis != "", "flex-basis:"+flexBasis)

	flexDirection := block.Parameter("flex_direction")
	blockTag.StyleIf(flexDirection != "", "flex-direction:"+flexDirection)

	flexFlow := block.Parameter("flex_flow")
	blockTag.StyleIf(flexFlow != "", "flex-flow:"+flexFlow)

	flexGrow := block.Parameter("flex_grow")
	blockTag.StyleIf(flexGrow != "", "flex-grow:"+flexGrow)

	flexShrink := block.Parameter("flex_shrink")
	blockTag.StyleIf(flexShrink != "", "flex-shrink:"+flexShrink)

	flexWrap := block.Parameter("flex_wrap")
	blockTag.StyleIf(flexWrap != "", "flex-wrap:"+flexWrap)

	justifyContent := block.Parameter("justify_content")
	blockTag.StyleIf(justifyContent != "", "justify-content:"+justifyContent)

	order := block.Parameter("order")
	blockTag.StyleIf(order != "", "order:"+order)
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
			Value: collapsibleStart("font_fields", "Font Settings", true),
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
			Value: collapsibleEnd(),
		}),
	}
}

func ApplyFontParameters(block ui.BlockInterface, blockTag *hb.Tag) {
	fontFamily := block.Parameter("font_family")

	if fontFamily != "" {
		blockTag.Style("font-family:" + fontFamily)
	}

	fontSize := block.Parameter("font_size")

	if fontSize != "" {
		blockTag.Style("font-size:" + fontSize)
	}

	fontWeight := block.Parameter("font_weight")

	if fontWeight != "" {
		blockTag.Style("font-weight:" + fontWeight)
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
			Value: collapsibleStart("margin_fields", "Margin Settings", true),
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
			Value: collapsibleEnd(),
		}),
	}
}

func ApplyMarginParameters(block ui.BlockInterface, blockTag *hb.Tag) {
	marginTop := block.Parameter("margin_top")

	if marginTop != "" {
		blockTag.Style("margin-top:" + marginTop)
	}

	marginBottom := block.Parameter("margin_bottom")

	if marginBottom != "" {
		blockTag.Style("margin-bottom:" + marginBottom)
	}

	marginLeft := block.Parameter("margin_left")

	if marginLeft != "" {
		blockTag.Style("margin-left:" + marginLeft)
	}

	marginRight := block.Parameter("margin_right")

	if marginRight != "" {
		blockTag.Style("margin-right:" + marginRight)
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
			Value: collapsibleStart("padding_fields", "Padding Settings", true),
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
			Value: collapsibleEnd(),
		}),
	}
}

func ApplyPaddingParameters(block ui.BlockInterface, blockTag *hb.Tag) {
	paddingTop := block.Parameter("padding_top")

	if paddingTop != "" {
		blockTag.Style("padding-top:" + paddingTop)
	}

	paddingBottom := block.Parameter("padding_bottom")

	if paddingBottom != "" {
		blockTag.Style("padding-bottom:" + paddingBottom)
	}

	paddingLeft := block.Parameter("padding_left")

	if paddingLeft != "" {
		blockTag.Style("padding-left:" + paddingLeft)
	}

	paddingRight := block.Parameter("padding_right")

	if paddingRight != "" {
		blockTag.Style("padding-right:" + paddingRight)
	}
}

// FieldsPositioning is a predefined set of positioning fields
// - Position (position)
// - Top (top)
// - Right (right)
// - Bottom (bottom)
// - Left (left)
// - Z-Index (zindex)
//
// Parameters:
// - none
//
// Returns:
// - []form.Field - The fields
func FieldsPositioning() []form.FieldInterface {
	return []form.FieldInterface{
		form.NewField(form.FieldOptions{
			Type:  form.FORM_FIELD_TYPE_RAW,
			Value: collapsibleStart("positioning_fields", "Positioning Settings", true),
		}),
		form.NewField(form.FieldOptions{
			Name:  "position",
			Label: "Position",
			Type:  form.FORM_FIELD_TYPE_SELECT,
			Options: []form.FieldOption{
				{
					Value: "",
					Key:   "",
				},
				{
					Value: "relative",
					Key:   "relative",
				},
				{
					Value: "absolute",
					Key:   "absolute",
				},
				{
					Value: "fixed",
					Key:   "fixed",
				},
				{
					Value: "sticky",
					Key:   "sticky",
				},
				{
					Value: "static",
					Key:   "static",
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
					Value: "unset",
					Key:   "unset",
				},
			},
		}),
		form.NewField(form.FieldOptions{
			Name:  "top",
			Label: "Top",
			Type:  form.FORM_FIELD_TYPE_STRING,
			Help:  `The top position of the block, i.e. 20px.`,
		}),
		form.NewField(form.FieldOptions{
			Name:  "right",
			Label: "Right",
			Type:  form.FORM_FIELD_TYPE_STRING,
			Help:  `The right position of the block, i.e. 20px.`,
		}),
		form.NewField(form.FieldOptions{
			Name:  "botom",
			Label: "Bottom",
			Type:  form.FORM_FIELD_TYPE_STRING,
			Help:  `The bottom position of the block, i.e. 20px.`,
		}),
		form.NewField(form.FieldOptions{
			Name:  "left",
			Label: "Left",
			Type:  form.FORM_FIELD_TYPE_STRING,
			Help:  `The left position of the block, i.e. 20px.`,
		}),
		form.NewField(form.FieldOptions{
			Name:  "z_index",
			Label: "Z-Index",
			Type:  form.FORM_FIELD_TYPE_STRING,
			Help:  `The z-index of the block, i.e. auto or a number.`,
		}),
		form.NewField(form.FieldOptions{
			Type:  form.FORM_FIELD_TYPE_RAW,
			Value: collapsibleEnd(),
		}),
	}
}

func ApplyPositionParameters(block ui.BlockInterface, blockTag *hb.Tag) {
	position := block.Parameter("position")
	zIndex := block.Parameter("z_index")
	top := block.Parameter("top")
	right := block.Parameter("right")
	bottom := block.Parameter("bottom")
	left := block.Parameter("left")

	if position != "" {
		blockTag.Style("position:" + position)
	}

	if zIndex != "" {
		blockTag.Style("z-index:" + zIndex)
	}

	if top != "" {
		blockTag.Style("top:" + top)
	}

	if right != "" {
		blockTag.Style("right:" + right)
	}

	if bottom != "" {
		blockTag.Style("bottom:" + bottom)
	}

	if left != "" {
		blockTag.Style("left:" + left)
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
			Value: collapsibleStart("text_fields", "Text Settings", true),
		}),
		form.NewField(form.FieldOptions{
			Name:  "text_color",
			Label: "Text Color",
			Type:  form.FORM_FIELD_TYPE_STRING,
			Help:  `The color of the text.`,
		}),
		form.NewField(form.FieldOptions{
			Name:  "text_decoration",
			Label: "Text Decoration",
			Type:  form.FORM_FIELD_TYPE_SELECT,
			Options: []form.FieldOption{
				{
					Value: "",
					Key:   "",
				},
				{
					Value: "underline",
					Key:   "underline",
				},
				{
					Value: "overline",
					Key:   "overline",
				},
				{
					Value: "line-through",
					Key:   "line-through",
				},
				{
					Value: "blink",
					Key:   "blink",
				},
			},
		}),
		form.NewField(form.FieldOptions{
			Name:  "text_indent",
			Label: "Text Indent",
			Type:  form.FORM_FIELD_TYPE_STRING,
			Help:  `The indent of the text, i.e. 20px.`,
		}),
		// overflow
		form.NewField(form.FieldOptions{
			Name:  "text_overflow",
			Label: "Text Overflow",
			Type:  form.FORM_FIELD_TYPE_SELECT,
			Options: []form.FieldOption{
				{
					Value: "",
					Key:   "",
				},
				{
					Value: "clip",
					Key:   "clip",
				},
				{
					Value: "ellipsis",
					Key:   "ellipsis",
				},
			},
		}),
		// shadow
		form.NewField(form.FieldOptions{
			Name:  "text_shadow",
			Label: "Text Shadow",
			Type:  form.FORM_FIELD_TYPE_STRING,
			Help:  `The shadow of the text.`,
		}),
		// transform
		form.NewField(form.FieldOptions{
			Name:  "text_transform",
			Label: "Text Transform",
			Type:  form.FORM_FIELD_TYPE_SELECT,
			Options: []form.FieldOption{
				{
					Value: "",
					Key:   "",
				},
				{
					Value: "uppercase",
					Key:   "uppercase",
				},
				{
					Value: "lowercase",
					Key:   "lowercase",
				},
			},
		}),
		// line height
		form.NewField(form.FieldOptions{
			Name:  "line_height",
			Label: "Line Height",
			Type:  form.FORM_FIELD_TYPE_STRING,
			Help:  `The line height of the text i.e. 16px.`,
		}),
		// letter spacing
		form.NewField(form.FieldOptions{
			Name:  "letter_spacing",
			Label: "Letter Spacing",
			Type:  form.FORM_FIELD_TYPE_STRING,
			Help:  `The letter spacing of the text, i.e. 1px.`,
		}),
		// whitespace
		form.NewField(form.FieldOptions{
			Name:  "white_space",
			Label: "White Space",
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
					Value: "nowrap",
					Key:   "nowrap",
				},
				{
					Value: "pre",
					Key:   "pre",
				},
				{
					Value: "pre-wrap",
					Key:   "pre-wrap",
				},
				{
					Value: "pre-line",
					Key:   "pre-line",
				},
			},
		}),
		// word-break
		form.NewField(form.FieldOptions{
			Name:  "word_break",
			Label: "Word Break",
			Type:  form.FORM_FIELD_TYPE_SELECT,
			Options: []form.FieldOption{
				{
					Value: "",
					Key:   "",
				},
				{
					Value: "break-all",
					Key:   "break-all",
				},
				{
					Value: "keep-all",
					Key:   "keep-all",
				},
			},
		}),
		// word-spacing
		form.NewField(form.FieldOptions{
			Name:  "word_spacing",
			Label: "Word Spacing",
			Type:  form.FORM_FIELD_TYPE_STRING,
			Help:  `The word spacing of the text, i.e. 1px.`,
		}),
		form.NewField(form.FieldOptions{
			Type:  form.FORM_FIELD_TYPE_RAW,
			Value: collapsibleEnd(),
		}),
	}
}

func ApplyTextParameters(block ui.BlockInterface, blockTag *hb.Tag) {
	textColor := block.Parameter("text_color")
	textDecoration := block.Parameter("text_decoration")
	textIndent := block.Parameter("text_indent")
	textOverflow := block.Parameter("text_overflow")
	textShadow := block.Parameter("text_shadow")
	textTransform := block.Parameter("text_transform")
	lineHeight := block.Parameter("line_height")
	letterSpacing := block.Parameter("letter_spacing")
	whiteSpace := block.Parameter("white_space")
	wordBreak := block.Parameter("word_break")
	wordSpacing := block.Parameter("word_spacing")

	if textColor != "" {
		blockTag.Style("color:" + textColor)
	}

	if textDecoration != "" {
		blockTag.Style("text-decoration:" + textDecoration)
	}

	if textIndent != "" {
		blockTag.Style("text-indent:" + textIndent)
	}

	if textOverflow != "" {
		blockTag.Style("text-overflow:" + textOverflow)
	}

	if textShadow != "" {
		blockTag.Style("text-shadow:" + textShadow)
	}

	if textTransform != "" {
		blockTag.Style("text-transform:" + textTransform)
	}

	if lineHeight != "" {
		blockTag.Style("line-height:" + lineHeight)
	}

	if letterSpacing != "" {
		blockTag.Style("letter-spacing:" + letterSpacing)
	}

	if whiteSpace != "" {
		blockTag.Style("white-space:" + whiteSpace)
	}

	if wordBreak != "" {
		blockTag.Style("word-break:" + wordBreak)
	}

	if wordSpacing != "" {
		blockTag.Style("word-spacing:" + wordSpacing)
	}
}

func FieldsTransition() []*form.Field {
	return []*form.Field{
		form.NewField(form.FieldOptions{
			Type:  form.FORM_FIELD_TYPE_RAW,
			Value: collapsibleStart("transition_fields", "Transition Settings", true),
		}),
		// transition
		form.NewField(form.FieldOptions{
			Name:  "transition",
			Label: "Transition",
			Type:  form.FORM_FIELD_TYPE_STRING,
			Help:  `The transition of the block.`,
		}),
		// transition-delay
		form.NewField(form.FieldOptions{
			Name:  "transition_delay",
			Label: "Transition Delay",
			Type:  form.FORM_FIELD_TYPE_STRING,
			Help:  `The transition delay of the block, i.e. 1s.`,
		}),
		// transition-duration
		form.NewField(form.FieldOptions{
			Name:  "transition_duration",
			Label: "Transition Duration",
			Type:  form.FORM_FIELD_TYPE_STRING,
			Help:  `The transition duration of the block, i.e. 1s.`,
		}),
		// transition-timing-function
		form.NewField(form.FieldOptions{
			Name:  "transition_timing_function",
			Label: "Transition Timing Function",
			Type:  form.FORM_FIELD_TYPE_SELECT,
			Options: []form.FieldOption{
				{
					Value: "",
					Key:   "",
				},
				{
					Value: "ease",
					Key:   "ease",
				},
				{
					Value: "linear",
					Key:   "linear",
				},
				{
					Value: "ease-in",
					Key:   "ease-in",
				},
				{
					Value: "ease-out",
					Key:   "ease-out",
				},
				{
					Value: "ease-in-out",
					Key:   "ease-in-out",
				},
			},
		}),
		// transition-property
		form.NewField(form.FieldOptions{
			Name:  "transition_property",
			Label: "Transition Property",
			Type:  form.FORM_FIELD_TYPE_SELECT,
			Help:  `Defines which properties should be transitioned.`,
			Options: []form.FieldOption{
				{
					Value: "",
					Key:   "",
				},
				{
					Value: "all",
					Key:   "all",
				},
				{
					Value: "none",
					Key:   "none",
				},
				{
					Value: "color",
					Key:   "color",
				},
				{
					Value: "background",
					Key:   "backgroundr",
				},
				{
					Value: "transform",
					Key:   "transform",
				},
			},
		}),
		form.NewField(form.FieldOptions{
			Type:  form.FORM_FIELD_TYPE_RAW,
			Value: collapsibleEnd(),
		}),
	}
}

func ApplyTransitionParameters(block ui.BlockInterface, blockTag *hb.Tag) {
	transition := block.Parameter("transition")
	transitionDelay := block.Parameter("transition_delay")
	transitionDuration := block.Parameter("transition_duration")
	transitionTimingFunction := block.Parameter("transition_timing_function")
	transitionProperty := block.Parameter("transition_property")

	if transition != "" {
		blockTag.Style("transition:" + transition)
	}

	if transitionDelay != "" {
		blockTag.Style("transition-delay:" + transitionDelay)
	}

	if transitionDuration != "" {
		blockTag.Style("transition-duration:" + transitionDuration)
	}

	if transitionTimingFunction != "" {
		blockTag.Style("transition-timing-function:" + transitionTimingFunction)
	}

	if transitionProperty != "" {
		blockTag.Style("transition-property:" + transitionProperty)
	}
}
