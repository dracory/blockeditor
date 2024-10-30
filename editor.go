package blockeditor

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gouniverse/hb"
	"github.com/gouniverse/ui"
	"github.com/gouniverse/utils"
	"github.com/samber/lo"
	"github.com/spf13/cast"
)

var _ hb.TagInterface = (*editor)(nil)

// == TYPE ====================================================================

type editor struct {
	hb.Tag
	id               string
	name             string
	value            string
	handleEndpoint   string
	blocks           []ui.BlockInterface
	blockDefinitions []BlockDefinition
}

// func (b *editor) availableTypes() []string {
// 	return lo.Map(b.blockDefinitions, func(blockDefinition BlockDefinition, _ int) string {
// 		return blockDefinition.Type
// 	})
// }

func (b *editor) ToHTML() string {
	blocksJSON, err := ui.BlocksToJson(b.blocks)

	if err != nil {
		return err.Error()
	}

	editor := hb.Div().
		Style("border:1px solid silver; padding:10px").
		ChildIf(len(b.blocks) < 1, b.buttonBlockInsert(0, true)).
		HTMLIf(len(b.blocks) > 0, b.blocksToCards(b.blocks))

	// buttonSeeSource := hb.Span().
	// 	ID(b.id + "_button_see_source").
	// 	Text("See source").
	// 	Style("margin-top:10px; margin-bottom:10px; cursor:pointer; color:blue; text-decoration:underline").
	// 	OnClick(`var display = document.getElementById("` + b.id + `").style.display; if (display == "none") { document.getElementById("` + b.id + `").style.display = "block"; } else { document.getElementById("` + b.id + `").style.display = "none"; }`)

	textarea := hb.TextArea().
		ID(b.id).
		Name(b.name).
		Style("width:100%;height:300px;display:none;").
		Text(blocksJSON)

	wrapperID := b.id + "_wrapper"

	style := `
 #` + wrapperID + ` .BlockOptions .dropdown-toggle:after {
   content: none;
 }

 #` + wrapperID + ` .BlockSeparator {
   border-top: 1px dashed #999;
 }

 #` + wrapperID + ` .BlockSeparator .ButtonBlockInsert {
   display: none;
 }

 #` + wrapperID + ` .BlockSeparator:hover {
   border-top: 5px solid #999;
 }

 #` + wrapperID + ` .BlockSeparator:hover .ButtonBlockInsert {
   display: inline-block;
 }
	`

	return hb.Wrap().
		Child(hb.Style(style)).
		Child(hb.Div().
			ID(wrapperID).
			Child(editor).
			// Child(buttonSeeSource).
			Child(textarea)).
		ToHTML()
}

// buttonBlockInsert creates a dropdown button for inserting a new block
//
// The dropdown will list all of the available block types
func (e *editor) buttonBlockInsert(atPosition int, single bool) *hb.Tag {
	link := e.url(map[string]string{
		ACTION:                  ACTION_BLOCK_ADD_MODAL,
		EDITOR_ID:               e.id,
		EDITOR_NAME:             e.name,
		EDITOR_HANDLER_ENDPOINT: e.handleEndpoint,
		"at_position":           cast.ToString(atPosition),
	})

	button := hb.Button().
		Class("ButtonBlockInsert btn btn-secondary btn-sm").
		Style(`border-radius: 30px;z-index: 100`).
		Type("button").
		Child(hb.I().Class(`bi bi-plus-circle me-2`)).
		TextIf(single, "add new block").
		TextIf(!single, "insert new block").
		HxPost(link).
		HxTarget("#" + e.id + "_wrapper").
		HxSwap(`beforeend`)

	return button

	// button := hb.Button().
	// 	Class("btn btn-secondary btn-sm dropdown-toggle").
	// 	Style(`border-radius: 30px;z-index: 100`).
	// 	Type("button").
	// 	Attr("data-bs-toggle", "dropdown").
	// 	Attr("aria-expanded", "false").
	// 	Child(hb.I().Class(`bi bi-plus-circle me-2`)).
	// 	TextIf(single, " add new block ").
	// 	TextIf(!single, " insert new block ")

	// dropdown := hb.Div().
	// 	Class(`dropdown`).
	// 	Child(button).
	// 	Child(hb.UL().
	// 		Class("dropdown-menu").
	// 		Children(lo.Map(b.availableTypes(), func(blockType string, _ int) hb.TagInterface {
	// 			link := b.url(map[string]string{
	// 				ACTION:                  ACTION_BLOCK_ADD,
	// 				EDITOR_ID:               b.id,
	// 				EDITOR_NAME:             b.name,
	// 				EDITOR_HANDLER_ENDPOINT: b.handleEndpoint,
	// 				BLOCK_TYPE:              blockType,
	// 				"at_position":           strconv.Itoa(atPosition),
	// 			})

	// 			dropdownItem := hb.Hyperlink().
	// 				Class("dropdown-item").
	// 				Href(link).
	// 				Child(hb.Text(blockType)).
	// 				HxPost(link).
	// 				HxTarget("#" + b.id + "_wrapper").
	// 				HxSwap(`outerHTML`)

	// 			return hb.LI().
	// 				Child(dropdownItem)
	// 		})))

	// return hb.NewWrap().
	// 	Child(dropdown)
}

// blocksToCards creates a card for each block
func (b *editor) blocksToCards(blocks []ui.BlockInterface) string {
	wrap := hb.Wrap()

	wrap.Child(b.blockDivider().Child(b.buttonBlockInsert(0, false)))

	for index, block := range blocks {
		position := index + 1

		wrap.Child(b.blockToCard(block))

		wrap.Child(b.blockDivider().Child(b.buttonBlockInsert(position, false)))
	}

	return wrap.ToHTML()
}

// blockToCard creates a card for a block
func (b *editor) blockToCard(block ui.BlockInterface) *hb.Tag {
	buttonMoveUp := b.cardButtonMoveUp(block.ID())
	buttonMoveDown := b.cardButtonMoveDown(block.ID())
	buttonEdit := b.cardButtonSettings(block.ID())
	buttonDelete := b.cardButtonDelete(block.ID())

	definition, found := lo.Find(b.blockDefinitions, func(blockDefinition BlockDefinition) bool {
		return blockDefinition.Type == block.Type()
	})

	render := lo.If(!found, "No renderer for type: "+block.Type()).
		ElseF(func() string { return definition.ToHTML(block) })

	position := lo.IndexOf(b.blocks, block)

	dropdown := hb.Div().
		Class(`BlockOptions dropdown float-start me-2`).
		Child(hb.Button().
			Class("btn btn-secondary btn-sm dropdown-toggle").
			Style(`border-radius: 30px;z-index: 100;font-size: 11px;padding: 5px;line-height: 1px;`).
			Type("button").
			Attr("data-bs-toggle", "dropdown").
			Attr("aria-expanded", "false").
			Child(hb.I().Class(`bi bi-three-dots-vertical`))).
		Child(hb.UL().
			Class("dropdown-menu").
			Child(func() hb.TagInterface {
				link := b.url(map[string]string{
					ACTION:                  ACTION_BLOCK_ADD_MODAL,
					BLOCK_ID:                block.ID(),
					EDITOR_ID:               b.id,
					EDITOR_NAME:             b.name,
					EDITOR_HANDLER_ENDPOINT: b.handleEndpoint,
					BLOCK_TYPE:              "image",
					"at_position":           cast.ToString(position),
				})
				dropdownItem := hb.Hyperlink().
					Class("dropdown-item").
					Href(link).
					Text("Insert before").
					HxPost(link).
					HxTarget("#" + b.id + "_wrapper").
					HxSwap(`beforeend`)

				return hb.LI().
					Child(dropdownItem)
			}()).
			Child(func() hb.TagInterface {
				link := b.url(map[string]string{
					ACTION:                  ACTION_BLOCK_ADD_MODAL,
					BLOCK_ID:                block.ID(),
					EDITOR_ID:               b.id,
					EDITOR_NAME:             b.name,
					EDITOR_HANDLER_ENDPOINT: b.handleEndpoint,
					BLOCK_TYPE:              "image",
					"at_position":           cast.ToString(position + 1),
				})
				dropdownItem := hb.Hyperlink().
					Class("dropdown-item").
					Href(link).
					Text("Insert after").
					HxPost(link).
					HxTarget("#" + b.id + "_wrapper").
					HxSwap(`beforeend`)

				return hb.LI().
					Child(dropdownItem)
			}()))

	card := hb.Div().
		Class(`card`).
		Child(
			hb.Div().Class(`card-header`).
				Child(dropdown).
				Text(block.Type()).
				Child(buttonDelete).
				Child(buttonEdit).
				Child(buttonMoveUp).
				Child(buttonMoveDown),
		).
		Child(hb.Div().
			Class(`card-body`).
			HTML(render))

	return card

}

// blockAdd creates a new block and inserts it at the specified position
func (b *editor) blockAdd(r *http.Request) string {
	blockType := utils.Req(r, BLOCK_TYPE, "")
	atPosition := utils.Req(r, "at_position", "")

	if blockType == "" {
		return "no block type"
	}

	if atPosition == "" {
		return "no position"
	}

	atPositionInt, err := strconv.Atoi(atPosition)

	if err != nil {
		return err.Error()
	}

	blockNew := ui.NewBlock()
	blockNew.SetType(blockType)

	b.blocks = append(b.blocks[:atPositionInt], append([]ui.BlockInterface{blockNew}, b.blocks[atPositionInt:]...)...)

	return b.ToHTML()
}

// blockDelete removes a block from the editor
func (b *editor) blockDelete(r *http.Request) string {
	blockID := utils.Req(r, BLOCK_ID, "")

	if blockID == "" {
		return "no block id"
	}

	for i, block := range b.blocks {
		if block.ID() == blockID {
			b.blocks = append(b.blocks[:i], b.blocks[i+1:]...)
			break
		}
	}

	return b.ToHTML()
}

// blockDivider creates a divider
func (b *editor) blockDivider() *hb.Tag {
	return hb.Div().
		Class(`BlockSeparator`).
		Style(`margin: 3px 0px;`).
		Style(`clear:both; height:1px; position: relative;`).
		Style(`display: flex; justify-content: center; align-items: center;`)
}

// blockMoveDown moves the block down
func (b *editor) blockMoveDown(r *http.Request) string {
	blockID := utils.Req(r, BLOCK_ID, "")

	if blockID == "" {
		return "no block id"
	}

	for i, block := range b.blocks {
		if block.ID() == blockID {
			if i < len(b.blocks)-1 {
				b.blocks[i], b.blocks[i+1] = b.blocks[i+1], b.blocks[i]
			}
			break
		}
	}

	return b.ToHTML()
}

// blockMoveUp moves the block up
func (b *editor) blockMoveUp(r *http.Request) string {
	blockID := utils.Req(r, BLOCK_ID, "")

	if blockID == "" {
		return "no block id"
	}

	for i, block := range b.blocks {
		if block.ID() == blockID {
			if i > 0 {
				b.blocks[i], b.blocks[i-1] = b.blocks[i-1], b.blocks[i]
			}
			break
		}
	}

	return b.ToHTML()
}

// blockFindByID finds a block by its ID
func (b *editor) blockFindByID(id string) ui.BlockInterface {

	for _, block := range b.blocks {
		if block.ID() == id {
			return block
		}
	}

	return nil
}

// blockSettingsUpdate updates the block settings
func (b *editor) blockSettingsUpdate(r *http.Request) string {
	blockID := utils.Req(r, BLOCK_ID, "")

	if blockID == "" {
		return "no block id"
	}

	block := b.blockFindByID(blockID)

	if block == nil {
		return "no block found"
	}

	all := utils.ReqAll(r)
	settings := map[string]string{}
	for key, values := range all {
		value := values[0]
		if strings.HasPrefix(key, SETTINGS_PREFIX) {
			key, _ := strings.CutPrefix(key, SETTINGS_PREFIX)
			settings[key] = value
		}
	}

	if len(settings) > 0 {
		block.SetParameters(settings)
	}

	// content := lo.ValueOr(settings, "content", "")

	// block.SetContent(content)

	for i := 0; i < len(b.blocks); i++ {
		if b.blocks[i].ID() == blockID {
			b.blocks[i] = block
		}
	}

	modalCloseScript := `document.getElementById('ModalBlockUpdate').remove();document.getElementById('ModalBackdrop').remove();`
	return b.ToHTML() + hb.Script(modalCloseScript).ToHTML()
}

func (b *editor) cardButtonMoveDown(blockID string) *hb.Tag {
	icon := hb.I().Class(`bi bi-arrow-down`)

	buttonMoveDown := hb.Button().
		Class("btn btn-info text-white btn-sm ms-2 float-end").
		Style(`border-radius: 30px;font-size: 11px;`).
		Type("button").
		Child(icon).
		HxPost(b.url(map[string]string{
			EDITOR_ID:               b.id,
			EDITOR_NAME:             b.name,
			EDITOR_HANDLER_ENDPOINT: b.handleEndpoint,
			ACTION:                  ACTION_BLOCK_MOVE_DOWN,
			BLOCK_ID:                blockID,
		})).
		HxTarget(`#` + b.id + `_wrapper`).
		HxSwap(`outerHTML`)

	return buttonMoveDown
}

func (b *editor) cardButtonMoveUp(blockID string) *hb.Tag {
	icon := hb.I().Class(`bi bi-arrow-up`)
	buttonMoveUp := hb.Button().
		Class("btn btn-info text-white btn-sm ms-2 float-end").
		Style(`border-radius: 30px;font-size: 11px;`).
		Type("button").
		Child(icon).
		HxPost(b.url(map[string]string{
			EDITOR_ID:               b.id,
			EDITOR_NAME:             b.name,
			EDITOR_HANDLER_ENDPOINT: b.handleEndpoint,
			ACTION:                  ACTION_BLOCK_MOVE_UP,
			BLOCK_ID:                blockID,
		})).
		HxTarget(`#` + b.id + `_wrapper`).
		HxSwap(`outerHTML`)

	return buttonMoveUp
}

func (b *editor) cardButtonDelete(blockID string) *hb.Tag {
	icon := hb.I().Class(`bi bi-trash`)

	buttonDelete := hb.Button().
		Class("btn btn-danger text-white btn-sm ms-2 float-end").
		Style(`border-radius: 30px;font-size: 11px;`).
		Type("button").
		Child(icon).
		HxPost(b.url(map[string]string{
			EDITOR_ID:               b.id,
			EDITOR_NAME:             b.name,
			EDITOR_HANDLER_ENDPOINT: b.handleEndpoint,
			ACTION:                  ACTION_BLOCK_DELETE,
			BLOCK_ID:                blockID,
		})).
		HxTarget(`#` + b.id + `_wrapper`).
		HxSwap(`outerHTML`)

	return buttonDelete
}

func (b *editor) cardButtonSettings(blockID string) *hb.Tag {
	icon := hb.I().Class(`bi bi-gear`)

	buttonSettings := hb.Button().
		Class("btn btn-warning text-white btn-sm ms-2 float-end").
		Style(`border-radius: 30px;font-size: 11px;`).
		Type("button").
		Child(icon).
		HxPost(b.url(map[string]string{
			EDITOR_ID:               b.id,
			EDITOR_NAME:             b.name,
			EDITOR_HANDLER_ENDPOINT: b.handleEndpoint,
			ACTION:                  ACTION_BLOCK_SETTINGS,
			BLOCK_ID:                blockID,
		})).
		HxTarget(`body`).
		HxSwap(`beforeend`)

	return buttonSettings
}

// url returns the url for the editor handler
//
// Business logic:
// - adds all editor parameters, so that we can reconstruct the editor
// - adds the parameters
func (b *editor) url(params map[string]string) string {
	params[EDITOR_ID] = b.id
	params[EDITOR_NAME] = b.name
	params[EDITOR_HANDLER_ENDPOINT] = b.handleEndpoint

	separator := lo.Ternary(strings.Contains(b.handleEndpoint, "?"), "&", "?")
	url := b.handleEndpoint + separator + query(params)
	return url
}
