package blockeditor

import (
	"strings"

	"github.com/gouniverse/hb"
	"github.com/gouniverse/ui"
	"github.com/samber/lo"
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

func (b *editor) ToHTML() string {
	blocksJSON, err := ui.MarshalBlocksToJson(b.blocks)

	if err != nil {
		return err.Error()
	}

	editor := hb.Div().
		// Style("border:1px solid silver; padding:10px").
		ChildIf(len(b.blocks) < 1, b.buttonBlockInsert("", 0, true)).
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
#` + wrapperID + ` {
   text-align: left;
 }

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

 #` + wrapperID + ` #ModalBlockAdd .card:hover {
   background-color: cornsilk;
 }
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
