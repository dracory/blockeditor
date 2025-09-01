package blockeditor

import (
	"github.com/dracory/hb"
	"github.com/dracory/ui"
	"github.com/samber/lo"
)

// blockToCard creates a card for a block
func (b *editor) blockToCard(block ui.BlockInterface) *hb.Tag {
	buttonMoveUp := b.cardButtonMoveUp(block.ID())
	buttonMoveDown := b.cardButtonMoveDown(block.ID())
	buttonEdit := b.cardButtonSettings(block.ID())
	buttonDelete := b.cardButtonDelete(block.ID())
	buttonDropdown := b.cardButtonDropdown(block)

	definition := b.findDefinitionByType(block.Type())

	if definition == nil {
		return hb.Div().
			Class("alert alert-warning").
			Text("Block " + block.Type() + " renderer does not exist")
	}

	status := block.Parameter("status")
	cardBackgroundColor := "bg-info"
	if status != "published" {
		cardBackgroundColor = "bg-danger"
	}

	tag := hb.NewTag("center").
		Child(definition.Icon).
		Style("font-size: 40px;")

	if definition.ToTag != nil {
		tag = definition.ToTag(block)
	}

	if len(block.Children()) > 0 {
		if definition.ToTag == nil {
			tag = hb.NewWrap()
		}
		tag.Children(lo.Map(block.Children(), func(child ui.BlockInterface, _ int) hb.TagInterface {
			return b.blockToCard(child)
		}))
	}

	card := hb.Div().
		Class(`BlockCard card`).
		Child(
			hb.Div().
				Class(`card-header`).
				Class(cardBackgroundColor).
				Style(`--bs-bg-opacity: 0.2;`).
				Style(`padding: 8px 10px; font-size: 12px; text-transform: uppercase; font-weight: bold;letter-spacing: 2px;`).
				Child(buttonDropdown).
				Text(block.Type()).
				Child(buttonDelete).
				Child(buttonEdit).
				Child(buttonMoveUp).
				Child(buttonMoveDown),
		).
		Child(hb.Div().
			Class(`card-body`).
			Class(cardBackgroundColor).
			Style(`--bs-bg-opacity: 0.1;`).
			Child(tag))
		// ChildIf(len(block.Children()) < 1, b.blockDivider().Child(b.buttonBlockInsert(blockExt.ID, 0, false))).
		// ChildrenIfF(len(block.Children()) > 0, func() []hb.TagInterface {
		// 	return lo.Map(block.Children(), func(child ui.BlockInterface, position int) hb.TagInterface {
		// 		return hb.Wrap().
		// 			// Child(b.blockDivider().Child(b.buttonBlockInsert(blockExt.ID, position, false))).
		// 			Child(b.blockToCard(child))
		// 	})
		// }).
		// ChildIf(len(block.Children()) > 0, b.blockDivider().Child(b.buttonBlockInsert(blockExt.ID, len(block.Children()), false))).
		// HTMLIf(len(block.Children()) < 1, render))

	if definition.Wrapper != nil {
		return definition.Wrapper(block).Child(card)
	}

	return card

}
