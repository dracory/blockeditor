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

	blockExt := flatTree.FindBlockExt(blockID)
	parentBlockExt := flatTree.FindBlockExt(blockExt.ParentID)

	if parentBlockExt == nil {
		return b.ToHTML()
	}

	children := flatTree.Children(parentBlockExt.ParentID)

	blockExt.ParentID = parentBlockExt.ParentID
	blockExt.Sequence = len(children)

	flatTree.UpdateBlockExt(*blockExt)

	b.blocks = flatTree.ToBlocks()

	return b.ToHTML()
}
