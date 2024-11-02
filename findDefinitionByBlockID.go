package blockeditor

func (b *editor) findDefinitionByID(blockID string) *BlockDefinition {
	block := b.findBlockByID(blockID)

	if block == nil {
		return nil
	}

	for _, blockDefinition := range b.blockDefinitions {
		if blockDefinition.Type == block.Type {
			return &blockDefinition
		}
	}

	return nil
}
