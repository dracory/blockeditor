package blockeditor

func (b *editor) findBlockByID(blockID string) *FlatBlock {
	flatTree := NewFlatTree(b.blocks)
	blockExt := flatTree.Find(blockID)

	if blockExt == nil {
		return nil
	}

	return blockExt
}
