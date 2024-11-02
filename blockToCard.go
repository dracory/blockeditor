package blockeditor

import (
	"github.com/gouniverse/hb"
	"github.com/gouniverse/ui"
	"github.com/samber/lo"
)

// blockToCard creates a card for a block
func (b *editor) blockToCard(block ui.BlockInterface) *hb.Tag {
	buttonMoveUp := b.cardButtonMoveUp(block.ID())
	buttonMoveDown := b.cardButtonMoveDown(block.ID())
	buttonEdit := b.cardButtonSettings(block.ID())
	buttonDelete := b.cardButtonDelete(block.ID())
	buttonDropdown := b.cardButtonDropdown(block)

	definition, found := lo.Find(b.blockDefinitions, func(blockDefinition BlockDefinition) bool {
		return blockDefinition.Type == block.Type()
	})

	hasRenderer := false

	if found {
		hasRenderer = definition.ToHTML != nil
	}

	render := lo.IfF(hasRenderer, func() string {
		return definition.ToHTML(block)
	}).
		ElseF(func() string {
			// return definition.ToHTML(block)
			return hb.NewTag("center").Child(definition.Icon).Style("font-size: 40px;").ToHTML()
		})

	card := hb.Div().
		Class(`BlockCard card`).
		Child(
			hb.Div().
				Class(`card-header bg-info`).
				Style(`--bs-bg-opacity: 0.2;`).
				Style(`padding: 2px 10px;font-size: 11px;`).
				Child(buttonDropdown).
				Text(block.Type()).
				Child(buttonDelete).
				Child(buttonEdit).
				Child(buttonMoveUp).
				Child(buttonMoveDown),
		).
		Child(hb.Div().
			Class(`card-body bg-info`).
			ClassIf(block.Type() == "row", `row`).
			Style(`--bs-bg-opacity: 0.1;`).
			ChildrenIfF(len(block.Children()) > 0, func() []hb.TagInterface {
				return lo.Map(block.Children(), func(child ui.BlockInterface, _ int) hb.TagInterface {
					return b.blockToCard(child)
				})
			}).
			HTMLIf(len(block.Children()) < 1, render))

	if block.Type() == "column" {
		width := block.Parameter("width")

		if width == "" {
			width = "12"
		}

		return hb.Div().
			Class("col-" + width).
			Child(card)
	}

	return card

}
