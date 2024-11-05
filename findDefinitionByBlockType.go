package blockeditor

func (b *editor) findDefinitionByType(blockType string) *BlockDefinition {
	for _, blockDefinition := range b.blockDefinitions {
		if blockDefinition.Type == blockType {
			return &blockDefinition
		}
	}

	return nil
}
