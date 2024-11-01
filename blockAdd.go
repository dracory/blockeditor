package blockeditor

import (
	"net/http"
	"strconv"

	"github.com/gouniverse/ui"
	"github.com/gouniverse/utils"
)

// blockAdd creates a new block and inserts it at the specified position
func (b *editor) blockAdd(r *http.Request) string {
	blockType := utils.Req(r, BLOCK_TYPE, "")
	atPosition := utils.Req(r, "at_position", "")
	parentID := utils.Req(r, "parent_id", "")

	if blockType == "" {
		return "no block type"
	}

	if atPosition == "" {
		return "no position"
	}

	atPositionInt, err := strconv.Atoi(atPosition)

	if err != nil {
		return err.Error()
	}

	blockNew := ui.NewBlock()
	blockNew.SetType(blockType)

	if parentID != "" {
		flatTree := NewFlatTree(b.blocks)
		flatTree.AddBlock(parentID, blockNew)
		b.blocks = flatTree.ToBlocks()
	} else {
		b.blocks = append(b.blocks[:atPositionInt], append([]ui.BlockInterface{blockNew}, b.blocks[atPositionInt:]...)...)
	}

	return b.ToHTML()
}
