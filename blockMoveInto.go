package blockeditor

import (
	"net/http"

	"github.com/gouniverse/utils"
)

// blockMoveUp moves the block into the previous block
func (b *editor) blockMoveInto(r *http.Request) string {
	blockID := utils.Req(r, BLOCK_ID, "")
	inSibling := utils.Req(r, "in_sibling", "")

	if blockID == "" {
		return "no block id"
	}

	if inSibling == "" {
		return "no in_sibling"
	}

	intoNext := inSibling == "next"
	intoPrevious := inSibling == "previous"

	flatTree := NewFlatTree(b.blocks)

	blockExt := flatTree.FindBlockExt(blockID)

	if intoPrevious {
		previousBlockExt := flatTree.FindPreviousBlockExt(*blockExt)

		if previousBlockExt == nil {
			return b.ToHTML()
		}

		blockExt.ParentID = previousBlockExt.ID
	}

	if intoNext {
		nextBlockExt := flatTree.FindNextBlockExt(*blockExt)

		if nextBlockExt == nil {
			return b.ToHTML()
		}

		blockExt.ParentID = nextBlockExt.ID
	}

	flatTree.UpdateBlockExt(*blockExt)

	b.blocks = flatTree.ToBlocks()

	return b.ToHTML()
}
