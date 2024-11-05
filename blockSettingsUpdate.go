package blockeditor

import (
	"net/http"
	"strings"

	"github.com/gouniverse/hb"
	"github.com/gouniverse/utils"
)

// blockSettingsUpdate updates the block settings
func (b *editor) blockSettingsUpdate(r *http.Request) string {
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
	flatBlock := flatTree.Find(blockID)

	if flatBlock == nil {
		return hb.Wrap().
			Child(hb.Swal(hb.SwalOptions{
				Icon:  "error",
				Title: "Error",
				Text:  "No block found",
			})).
			Child(b).
			ToHTML()
	}

	all := utils.ReqAll(r)
	settings := map[string]string{}
	for key, values := range all {
		value := values[0]
		if strings.HasPrefix(key, SETTINGS_PREFIX) {
			key, _ := strings.CutPrefix(key, SETTINGS_PREFIX)
			settings[key] = value
		}
	}

	if len(settings) > 0 {
		flatBlock.Parameters = settings
	}

	flatTree.Update(*flatBlock)
	b.blocks = flatTree.ToBlocks()

	modalCloseScript := `document.getElementById('ModalBlockUpdate').remove();document.getElementById('ModalBackdrop').remove();`
	return b.ToHTML() + hb.Script(modalCloseScript).ToHTML()
}
