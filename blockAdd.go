package blockeditor

import (
	"net/http"
	"strconv"

	"github.com/dracory/hb"
	"github.com/dracory/req"
	"github.com/dracory/ui"
)

// blockAdd creates a new block and inserts it at the specified position
func (b *editor) blockAdd(r *http.Request) string {
	blockType := req.GetString(r, BLOCK_TYPE)
	atPosition := req.GetString(r, "at_position")
	parentID := req.GetString(r, "parent_id")

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

	blockNew := ui.NewBlock()
	blockNew.SetType(blockType)

	tree := NewFlatTree(b.blocks)

	tree.AddBlock(parentID, blockNew)

	tree.MoveToPosition(blockNew.ID(), parentID, atPositionInt)

	b.blocks = tree.ToBlocks()

	return b.ToHTML()
}
