package blockeditor

import (
	"net/http"

	"github.com/gouniverse/utils"
)

// blockMoveUp moves the block up
func (b *editor) blockMoveUp(r *http.Request) string {
	blockID := utils.Req(r, BLOCK_ID, "")

	if blockID == "" {
		return "no block id"
	}

	flatTree := NewFlatTree(b.blocks)

	blockExt := flatTree.FindBlockExt(blockID)
	previousBlockExt := flatTree.FindPreviousBlockExt(*blockExt)

	if previousBlockExt == nil {
		return b.ToHTML()
	}

	previousSequence := previousBlockExt.Sequence
	sequence := blockExt.Sequence

	blockExt.Sequence = previousSequence
	previousBlockExt.Sequence = sequence

	flatTree.UpdateBlockExt(*blockExt)
	flatTree.UpdateBlockExt(*previousBlockExt)

	b.blocks = flatTree.ToBlocks()

	return b.ToHTML()
}
