package blockeditor

import (
	"net/http"

	"github.com/gouniverse/hb"
	"github.com/gouniverse/utils"
)

// blockMoveUp moves the block into the previous block
func (b *editor) blockMoveInto(r *http.Request) string {
	blockID := utils.Req(r, BLOCK_ID, "")
	inSibling := utils.Req(r, "in_sibling", "")

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

	flatTree := NewFlatTree(b.blocks)

	blockExt := flatTree.Find(blockID)

	if intoPrevious {
		previous := flatTree.FindPreviousSibling(blockExt.ID)

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

		newBlock := flatTree.Clone(*blockExt)

		flatTree.Remove(blockExt.ID)
		flatTree.Add(previous.ID, newBlock)
	}

	if intoNext {
		next := flatTree.FindNextSibling(blockExt.ID)

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

		newBlock := flatTree.Clone(*blockExt)

		flatTree.Remove(blockExt.ID)
		flatTree.Add(next.ID, newBlock)
	}

	b.blocks = flatTree.ToBlocks()

	return b.ToHTML()
}
