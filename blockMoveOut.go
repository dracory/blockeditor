package blockeditor

import (
	"net/http"

	"github.com/gouniverse/hb"
	"github.com/gouniverse/utils"
)

// blockMoveUp moves the block out of the parent block
func (b *editor) blockMoveOut(r *http.Request) string {
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
