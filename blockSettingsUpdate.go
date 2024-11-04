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
		return "no block id"
	}

	flatTree := NewFlatTree(b.blocks)
	flatBlock := flatTree.Find(blockID)

	if flatBlock == nil {
		return "no block found"
	}

	// block := b.blockFindByID(blockID)

	// if block == nil {
	// 	return "no block found"
	// }

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

	// content := lo.ValueOr(settings, "content", "")

	// block.SetContent(content)

	// for i := 0; i < len(b.blocks); i++ {
	// 	if b.blocks[i].ID() == blockID {
	// 		b.blocks[i] = block
	// 	}
	// }

	modalCloseScript := `document.getElementById('ModalBlockUpdate').remove();document.getElementById('ModalBackdrop').remove();`
	return b.ToHTML() + hb.Script(modalCloseScript).ToHTML()
}
