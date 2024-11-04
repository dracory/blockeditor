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
	flatTree.MoveUp(blockID)
	b.blocks = flatTree.ToBlocks()

	return b.ToHTML()
}
