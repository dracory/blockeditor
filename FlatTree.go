package blockeditor

import (
	"sort"

	"github.com/gouniverse/ui"
	"github.com/samber/lo"
)

type BlockExt struct {
	ID         string
	Type       string
	ParentID   string
	Sequence   int
	Parameters map[string]string
}

type FlatTree struct {
	list []BlockExt
}

func (f *FlatTree) Children(parentID string) []BlockExt {
	childrenExt := make([]BlockExt, 0)

	sequence := []int{}
	for _, blockExt := range f.list {
		if blockExt.ParentID == parentID {
			sequence = append(sequence, blockExt.Sequence)
			childrenExt = append(childrenExt, blockExt)
		}
	}

	sortedSequence := sort.IntSlice(sequence)
	sortedSequence.Sort()

	sortedChildren := make([]BlockExt, 0)

	for _, seq := range sortedSequence {
		for _, blockExt := range childrenExt {
			if blockExt.Sequence == seq {
				sortedChildren = append(sortedChildren, blockExt)
			}
		}
	}

	return sortedChildren
}

func NewFlatTree(blocks []ui.BlockInterface) *FlatTree {
	blocksExt := traverse(blocks, "")
	return &FlatTree{
		list: blocksExt,
	}
}

func (f *FlatTree) FindBlockExt(blockID string) *BlockExt {
	for _, blockExt := range f.list {
		if blockExt.ID == blockID {
			return &blockExt
		}
	}

	return nil
}

func (f *FlatTree) FindPreviousBlockExt(blockExt BlockExt) *BlockExt {
	children := f.Children(blockExt.ParentID)

	_, index, found := lo.FindIndexOf(children, func(bExt BlockExt) bool {
		return bExt.ID == blockExt.ID
	})

	if !found {
		return nil
	}

	if index == 0 {
		return nil
	}

	return &children[index-1]
}

func (f *FlatTree) FindNextBlockExt(blockExt BlockExt) *BlockExt {
	children := f.Children(blockExt.ParentID)

	_, index, found := lo.FindIndexOf(children, func(bExt BlockExt) bool {
		return bExt.ID == blockExt.ID
	})

	if !found {
		return nil
	}

	if index == len(children)-1 {
		return nil
	}

	return &children[index+1]
}

func (f *FlatTree) AddBlockExt(parentID string, blockExt BlockExt) {
	children := f.Children(parentID)
	blockExt.Sequence = len(children)
	blockExt.ParentID = parentID
	f.list = append(f.list, blockExt)
}

func (f *FlatTree) RemoveBlockExt(blockExt BlockExt) {
	for i, ext := range f.list {
		if ext.ID == blockExt.ID {
			f.list = append(f.list[:i], f.list[i+1:]...)
		}
	}
}

func (f *FlatTree) UpdateBlockExt(blockExt BlockExt) {
	for i, ext := range f.list {
		if ext.ID == blockExt.ID {
			f.list[i] = blockExt
		}
	}
}

func (f *FlatTree) AddBlock(parentID string, block ui.BlockInterface) {
	children := f.Children(parentID)

	blockExt := BlockExt{
		ID:         block.ID(),
		Type:       block.Type(),
		ParentID:   parentID,
		Sequence:   len(children),
		Parameters: block.Parameters(),
	}

	f.list = append(f.list, blockExt)
}

func (f *FlatTree) ToBlocks() []ui.BlockInterface {
	parentBlocks := f.Children("")

	blocks := make([]ui.BlockInterface, 0)

	for _, blockExt := range parentBlocks {
		blocks = append(blocks, f.blockExtToBlock(blockExt))
	}

	return blocks
}

func (f *FlatTree) blockExtToBlock(blockExt BlockExt) ui.BlockInterface {
	childrenExt := f.Children(blockExt.ID)

	children := []ui.BlockInterface{}
	for _, childExt := range childrenExt {
		children = append(children, f.blockExtToBlock(childExt))
	}

	block := ui.NewFromMap(map[string]interface{}{
		"id":         blockExt.ID,
		"type":       blockExt.Type,
		"parameters": blockExt.Parameters,
		"children":   children,
	})

	return block
}

func traverse(blocks []ui.BlockInterface, parentID string) []BlockExt {
	list := make([]BlockExt, 0)

	for index, block := range blocks {
		blockExt := BlockExt{
			ID:         block.ID(),
			Type:       block.Type(),
			ParentID:   parentID,
			Sequence:   index,
			Parameters: block.Parameters(),
		}
		list = append(list, blockExt)
		list = append(list, traverse(block.Children(), block.ID())...)
	}

	return list
}
