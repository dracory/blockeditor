package blockeditor

import (
	"net/http"

	"github.com/gouniverse/utils"
)

// blockMoveUp moves the block out of the parent block
func (b *editor) blockMoveOut(r *http.Request) string {
	blockID := utils.Req(r, BLOCK_ID, "")

	if blockID == "" {
		return "no block id"
	}

	flatTree := NewFlatTree(b.blocks)

	block := flatTree.Find(blockID)

	if block == nil {
		return b.ToHTML()
	}

	parent := flatTree.Parent(block.ID)

	if parent == nil {
		return b.ToHTML()
	}

	flatTree.MoveToParent(block.ID, parent.ParentID)

	b.blocks = flatTree.ToBlocks()

	return b.ToHTML()
}
