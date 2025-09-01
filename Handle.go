package blockeditor

import (
	"net/http"

	"github.com/dracory/req"
)

// Handle handles the block editor
//
// Business logic:
// - gets the action from the request
// - gets the editor id, name, content, and handler endpoint from the request
// - creates a new editor
// - if an action is provided, executes the action
// - if no action is provided, renders the editor
//
// Parameters:
// - w: the http.ResponseWriter
// - r: the http.Request
// - blockDefinitions: the block definitions
//
// Returns:
// - the response string
func Handle(w http.ResponseWriter, r *http.Request, blockDefinitions []BlockDefinition) string {
	action := req.GetString(r, ACTION)
	editorID := req.GetString(r, EDITOR_ID)
	editorName := req.GetString(r, EDITOR_NAME)
	editorContent := req.GetString(r, editorName)
	editorHandlerEndpoint := req.GetString(r, EDITOR_HANDLER_ENDPOINT)

	if editorID == "" {
		return "no editor id"
	}

	if editorName == "" {
		return "no editor name"
	}

	if editorContent == "" {
		return "no editor content"
	}

	blockEditor, err := NewEditor(NewEditorOptions{
		ID:               editorID,
		Name:             editorName,
		Value:            editorContent,
		HandleEndpoint:   editorHandlerEndpoint,
		BlockDefinitions: blockDefinitions,
	})

	if err != nil {
		return err.Error()
	}

	if action == ACTION_BLOCK_ADD {
		return blockEditor.blockAdd(r)
	}

	if action == ACTION_BLOCK_ADD_MODAL {
		return blockEditor.blockAddModal(r)
	}

	if action == ACTION_BLOCK_DELETE {
		return blockEditor.blockDelete(r)
	}

	if action == ACTION_BLOCK_DUPLICATE {
		return blockEditor.blockDuplicate(r)
	}

	if action == ACTION_BLOCK_MOVE_UP {
		return blockEditor.blockMoveUp(r)
	}

	if action == ACTION_BLOCK_MOVE_DOWN {
		return blockEditor.blockMoveDown(r)
	}

	if action == ACTION_BLOCK_MOVE_INTO {
		return blockEditor.blockMoveInto(r)
	}

	if action == ACTION_BLOCK_MOVE_OUT {
		return blockEditor.blockMoveOut(r)
	}

	if action == ACTION_BLOCK_SETTINGS {
		return blockEditor.blockSettingsModal(r)
	}

	if action == ACTION_BLOCK_SETTINGS_UPDATE {
		return blockEditor.blockSettingsUpdate(r)
	}

	return blockEditor.ToHTML()
}
