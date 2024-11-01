package blockeditor

import (
	"github.com/gouniverse/hb"
	"github.com/gouniverse/ui"
)

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
