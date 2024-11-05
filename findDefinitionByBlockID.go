package blockeditor

func (b *editor) findDefinitionByID(blockID string) *BlockDefinition {
	block := b.findBlockByID(blockID)

	if block == nil {
		return nil
	}

	return b.findDefinitionByType(block.Type)
}
