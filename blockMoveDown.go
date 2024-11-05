package blockeditor

import (
	"net/http"

	"github.com/gouniverse/hb"
	"github.com/gouniverse/utils"
)

// blockMoveDown moves the block down
func (b *editor) blockMoveDown(r *http.Request) string {
	blockID := utils.Req(r, BLOCK_ID, "")

	if blockID == "" {
		return hb.Wrap().
			Child(hb.Swal(hb.SwalOptions{
				Icon:  "error",
				Title: "Error",
				Text:  "No block id",
			})).
			Child(b).
			ToHTML()
	}

	flatTree := NewFlatTree(b.blocks)
	flatTree.MoveDown(blockID)

	b.blocks = flatTree.ToBlocks()

	return b.ToHTML()
}
