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
	flatTree.MoveDown(blockID)

	b.blocks = flatTree.ToBlocks()

	return b.ToHTML()
}
