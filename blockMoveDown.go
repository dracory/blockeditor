package blockeditor

import (
	"net/http"

	"github.com/gouniverse/utils"
)

// blockMoveDown moves the block down
func (b *editor) blockMoveDown(r *http.Request) string {
	blockID := utils.Req(r, BLOCK_ID, "")

	if blockID == "" {
		return "no block id"
	}

	flatTree := NewFlatTree(b.blocks)
	blockExt := flatTree.FindBlockExt(blockID)
	nextBlockExt := flatTree.FindNextBlockExt(*blockExt)

	if nextBlockExt == nil {
		return b.ToHTML()
	}

	nextSequence := nextBlockExt.Sequence
	sequence := blockExt.Sequence

	blockExt.Sequence = nextSequence
	nextBlockExt.Sequence = sequence

	flatTree.UpdateBlockExt(*blockExt)
	flatTree.UpdateBlockExt(*nextBlockExt)

	b.blocks = flatTree.ToBlocks()

	return b.ToHTML()
}
