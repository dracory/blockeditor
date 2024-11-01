package blockeditor

import (
	"strings"

	"github.com/gouniverse/hb"
	"github.com/gouniverse/ui"
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

	buttonSeeSource := hb.Span().
		ID(b.id + "_button_see_source").
		Text("See source").
		Style("margin-top:10px; margin-bottom:10px; cursor:pointer; color:blue; text-decoration:underline").
		OnClick(`var display = document.getElementById("` + b.id + `").style.display; if (display == "none") { document.getElementById("` + b.id + `").style.display = "block"; } else { document.getElementById("` + b.id + `").style.display = "none"; }`)

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

//  #` + wrapperID + ` .BlockCard .card-header {
//    display: none;
//  }

//  #` + wrapperID + ` .BlockCard:hover .card-header {
//    display: block;
//  }
	`

	return hb.Wrap().
		Child(hb.Style(style)).
		Child(hb.Div().
			ID(wrapperID).
			Child(editor).
			Child(buttonSeeSource).
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
}

// blockDivider creates a divider
func (b *editor) blockDivider() *hb.Tag {
	return hb.Div().
		Class(`BlockSeparator`).
		Style(`margin: 3px 0px;`).
		Style(`clear:both; height:1px; position: relative;`).
		Style(`display: flex; justify-content: center; align-items: center;`)
}

// blockFindByID finds a block by its ID
// func (b *editor) blockFindByID(id string) ui.BlockInterface {

// 	for _, block := range b.blocks {
// 		if block.ID() == id {
// 			return block
// 		}
// 	}

// 	return nil
// }

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
