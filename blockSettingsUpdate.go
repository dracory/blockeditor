package blockeditor

import (
	"net/http"

	"github.com/dracory/req"
	"github.com/gouniverse/hb"
)

// blockSettingsUpdate updates the block settings
func (b *editor) blockSettingsUpdate(r *http.Request) string {
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

	settings := b.findPostedSettings(r)

	if len(settings) > 0 {
		flatBlock.Parameters = settings
	}

	flatTree.Update(*flatBlock)

	b.blocks = flatTree.ToBlocks()

	modalCloseScript := `document.getElementById('ModalBlockUpdate').remove();document.getElementById('ModalBackdrop').remove();`
	return b.ToHTML() + hb.Script(modalCloseScript).ToHTML()
}

func (b *editor) findPostedSettings(r *http.Request) map[string]string {
	all := req.GetAll(r)

	if len(all) == 0 {
		return map[string]string{}
	}

	settings := map[string]string{}

	for key, values := range all {
		if len(values) == 0 {
			continue
		}

		value := values[0]
		if isPrefixedKey(key, SETTINGS_PREFIX) {
			origKey := unprefixKey(key, SETTINGS_PREFIX)
			settings[origKey] = value
		}
	}

	return settings
}
