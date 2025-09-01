package blockeditor

import (
	"net/http"

	"github.com/dracory/hb"
	"github.com/dracory/req"
)

// blockMoveUp moves the block into the previous block
func (b *editor) blockMoveInto(r *http.Request) string {
	blockID := req.GetString(r, BLOCK_ID)
	inSibling := req.GetString(r, "in_sibling")

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

	if inSibling == "" {
		return hb.Wrap().
			Child(hb.Swal(hb.SwalOptions{
				Icon:  "error",
				Title: "Error",
				Text:  "No sibling specifier",
			})).
			Child(b).
			ToHTML()
	}

	intoNext := inSibling == "next"
	intoPrevious := inSibling == "previous"

	tree := NewFlatTree(b.blocks)

	block := tree.Find(blockID)

	if block == nil {
		return b.ToHTML()
	}

	if intoPrevious {
		previous := tree.FindPreviousSibling(block.ID)

		if previous == nil {
			return hb.Wrap().
				Child(hb.Swal(hb.SwalOptions{
					Icon:  "error",
					Title: "Error",
					Text:  "No previous sibling found",
				})).
				Child(b).
				ToHTML()
		}

		definition := b.findDefinitionByType(previous.Type)

		if definition == nil {
			return hb.Wrap().
				Child(hb.Swal(hb.SwalOptions{
					Icon:  "error",
					Title: "Error",
					Text:  "Previous block definition not found",
				})).
				Child(b).
				ToHTML()
		}

		if !definition.AllowChildren {
			return hb.Wrap().
				Child(hb.Swal(hb.SwalOptions{
					Icon:  "error",
					Title: "Error",
					Text:  "Previous block definition does not allow children",
				})).
				Child(b).
				ToHTML()
		}

		tree.MoveToParent(block.ID, previous.ID)
	}

	if intoNext {
		next := tree.FindNextSibling(block.ID)

		if next == nil {
			return hb.Wrap().
				Child(hb.Swal(hb.SwalOptions{
					Icon:  "error",
					Title: "Error",
					Text:  "No next sibling found",
				})).
				Child(b).
				ToHTML()
		}

		definition := b.findDefinitionByType(next.Type)

		if definition == nil {
			return hb.Wrap().
				Child(hb.Swal(hb.SwalOptions{
					Icon:  "error",
					Title: "Error",
					Text:  "Next block definition not found",
				})).
				Child(b).
				ToHTML()
		}

		if !definition.AllowChildren {
			return hb.Wrap().
				Child(hb.Swal(hb.SwalOptions{
					Icon:  "error",
					Title: "Error",
					Text:  "Next block definition does not allow children",
				})).
				Child(b).
				ToHTML()
		}

		tree.MoveToParent(block.ID, next.ID)
	}

	b.blocks = tree.ToBlocks()

	return b.ToHTML()
}
