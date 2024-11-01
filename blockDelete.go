package blockeditor

import (
	"net/http"

	"github.com/gouniverse/utils"
)

// blockDelete removes a block from the editor
func (b *editor) blockDelete(r *http.Request) string {
	blockID := utils.Req(r, BLOCK_ID, "")

	if blockID == "" {
		return "no block id"
	}

	flatTree := NewFlatTree(b.blocks)
	blockExt := flatTree.FindBlockExt(blockID)

	if blockExt == nil {
		return b.ToHTML()
	}

	flatTree.RemoveBlockExt(*blockExt)
	b.blocks = flatTree.ToBlocks()

	return b.ToHTML()
}
