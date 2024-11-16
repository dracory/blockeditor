package blockeditor

import (
	"net/http"
	"strconv"

	"github.com/gouniverse/hb"
	"github.com/gouniverse/ui"
	"github.com/gouniverse/utils"
)

// blockAdd creates a new block and inserts it at the specified position
func (b *editor) blockAdd(r *http.Request) string {
	blockType := utils.Req(r, BLOCK_TYPE, "")
	atPosition := utils.Req(r, "at_position", "")
	parentID := utils.Req(r, "parent_id", "")

	if blockType == "" {
		return hb.Wrap().
			Child(hb.Swal(hb.SwalOptions{
				Icon:  "error",
				Title: "Error",
				Text:  "No block type",
			})).
			Child(b).
			ToHTML()
	}

	if atPosition == "" {
		return hb.Wrap().
			Child(hb.Swal(hb.SwalOptions{
				Icon:  "error",
				Title: "Error",
				Text:  "No position",
			})).
			Child(b).
			ToHTML()
	}

	atPositionInt, err := strconv.Atoi(atPosition)

	if err != nil {
		return hb.Wrap().
			Child(hb.Swal(hb.SwalOptions{
				Icon:  "error",
				Title: "Error",
				Text:  err.Error(),
			})).
			Child(b).
			ToHTML()
	}

	blockNew := ui.Block()
	blockNew.SetType(blockType)

	tree := NewFlatTree(b.blocks)

	tree.AddBlock(parentID, blockNew)

	tree.MoveToPosition(blockNew.ID(), parentID, atPositionInt)

	b.blocks = tree.ToBlocks()

	return b.ToHTML()
}
