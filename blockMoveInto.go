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

	blockExt := flatTree.Find(blockID)

	if intoPrevious {
		previous := flatTree.FindPreviousSibling(blockExt.ID)

		if previous == nil {
			return b.ToHTML()
		}

		newBlockExt := flatTree.Clone(*blockExt)

		flatTree.Remove(blockExt.ID)
		flatTree.Add(previous.ID, newBlockExt)
	}

	if intoNext {
		next := flatTree.FindNextSibling(blockExt.ID)

		if next == nil {
			return b.ToHTML()
		}

		newBlockExt := flatTree.Clone(*blockExt)

		flatTree.Remove(blockExt.ID)
		flatTree.Add(next.ID, newBlockExt)
	}

	b.blocks = flatTree.ToBlocks()

	return b.ToHTML()
}
