package blockeditor

func (b *editor) findBlockByID(blockID string) *BlockExt {
	flatTree := NewFlatTree(b.blocks)
	blockExt := flatTree.FindBlockExt(blockID)

	if blockExt == nil {
		return nil
	}

	return blockExt
}
