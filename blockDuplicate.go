package blockeditor

import (
	"net/http"

	"github.com/dracory/req"
	"github.com/gouniverse/hb"
)

// blockDelete removes a block from the editor
func (b *editor) blockDuplicate(r *http.Request) string {
	blockID := req.GetString(r, BLOCK_ID)

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
	flatTree.Duplicate(blockID)
	b.blocks = flatTree.ToBlocks()

	return b.ToHTML()
}
