package blockeditor

import (
	"net/http"

	"github.com/gouniverse/hb"
	"github.com/gouniverse/utils"
)

// blockMoveUp moves the block out of the parent block, before or after
func (b *editor) blockMoveOut(r *http.Request) string {
	blockID := utils.Req(r, BLOCK_ID, "")
	toPosition := utils.Req(r, "to_position", "")

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

	tree := NewFlatTree(b.blocks)

	block := tree.Find(blockID)

	if block == nil {
		return b.ToHTML()
	}

	parent := tree.Parent(block.ID)

	if parent == nil {
		return b.ToHTML()
	}

	if toPosition == "before" {
		tree.MoveToPosition(block.ID, parent.ParentID, parent.Sequence)
	} else if toPosition == "after" {
		tree.MoveToPosition(block.ID, parent.ParentID, parent.Sequence+1)
	} else {
		tree.MoveToParent(block.ID, parent.ParentID)
	}

	b.blocks = tree.ToBlocks()

	return b.ToHTML()
}
