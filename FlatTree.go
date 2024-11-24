package blockeditor

import (
	"sort"

	"github.com/gouniverse/ui"
	"github.com/gouniverse/uid"
	"github.com/samber/lo"
)

type FlatBlock struct {
	ID         string
	Type       string
	ParentID   string
	Sequence   int
	Parameters map[string]string
}

type FlatTree struct {
	list []FlatBlock
}

func NewFlatTree(blocks []ui.BlockInterface) *FlatTree {
	blocksExt := traverse(blocks, "")
	return &FlatTree{
		list: blocksExt,
	}
}

func (tree *FlatTree) Add(parentID string, flatBlock FlatBlock) {
	children := tree.Children(parentID)
	flatBlock.Sequence = len(children)
	flatBlock.ParentID = parentID
	tree.list = append(tree.list, flatBlock)

	tree.RecalculateSequences(parentID)
}

// AddBlock adds a new ui.BlockInterface to the FlatTree
func (tree *FlatTree) AddBlock(parentID string, block ui.BlockInterface) {
	children := tree.Children(parentID)

	flatBlock := FlatBlock{
		ID:         block.ID(),
		Type:       block.Type(),
		ParentID:   parentID,
		Sequence:   len(children),
		Parameters: block.Parameters(),
	}

	tree.list = append(tree.list, flatBlock)

	tree.RecalculateSequences(parentID)
}

// Children returns the children of the FlatBlock with the given parentID
func (tree *FlatTree) Children(parentID string) []FlatBlock {
	childrenExt := make([]FlatBlock, 0)

	sequence := []int{}
	for _, flatBlock := range tree.list {
		if flatBlock.ParentID == parentID {
			sequence = append(sequence, flatBlock.Sequence)
			childrenExt = append(childrenExt, flatBlock)
		}
	}

	sortedSequence := sort.IntSlice(sequence)
	sortedSequence.Sort()

	sortedChildren := make([]FlatBlock, 0)

	for _, seq := range sortedSequence {
		for _, flatBlock := range childrenExt {
			if flatBlock.Sequence == seq {
				sortedChildren = append(sortedChildren, flatBlock)
			}
		}
	}

	return sortedChildren
}

// Clone creates a shallow clone of a FlatBlock (no children)
//
// This is used to create a clone of a FlatBlock, so that the original FlatBlock
// is not modified, but we can modify the clone safely
//
// Remember to update the ID, Sequence, and ParentID of the copy with new values
func (tree *FlatTree) Clone(flatBlock FlatBlock) FlatBlock {
	return FlatBlock{
		ID:         flatBlock.ID,
		Type:       flatBlock.Type,
		ParentID:   flatBlock.ParentID,
		Sequence:   flatBlock.Sequence,
		Parameters: flatBlock.Parameters,
	}
}

// Duplicate creates a deep clone of a FlatBlock (with children)
// and adds it to the tree, under the same parent
//
// Business Logic:
// - travserses the tree to find all blocks to be duplicated
// - makes a map with current IDs as keys, newly generated IDs as values
// - clones each block, and replaces the ID with the new ID
// - assignes the correct mapped IDs and ParentIDs
// - adds the cloned blocks to the tree directly (using list)
// - moves the duplicated block under the block being duplicated
func (tree *FlatTree) Duplicate(blockID string) {
	block := tree.Find(blockID)

	if block == nil {
		return
	}

	blocks := tree.Traverse(blockID)

	if len(blocks) == 0 {
		return
	}

	mapIDs := make(map[string]string)

	for _, block := range blocks {
		mapIDs[block.ID] = uid.HumanUid()
	}

	clonedBlocks := make([]FlatBlock, 0)
	for _, block := range blocks {
		newID := lo.ValueOr(mapIDs, block.ID, block.ID)
		newParentID := lo.ValueOr(mapIDs, block.ParentID, block.ParentID)
		clonedBlock := tree.Clone(block)
		clonedBlock.ID = newID
		clonedBlock.ParentID = newParentID
		clonedBlocks = append(clonedBlocks, clonedBlock)
	}

	tree.list = append(tree.list, clonedBlocks...)

	newID := mapIDs[blockID]

	tree.MoveToPosition(newID, block.ParentID, block.Sequence+1)
}

func (tree *FlatTree) Exists(flatBlockID string) bool {
	for _, flatBlock := range tree.list {
		if flatBlock.ID == flatBlockID {
			return true
		}
	}
	return false
}

func (tree *FlatTree) Find(flatBlockID string) *FlatBlock {
	if flatBlockID == "" {
		return nil
	}

	for _, flatBlock := range tree.list {
		if flatBlock.ID == flatBlockID {
			return &flatBlock
		}
	}
	return nil
}

func (tree *FlatTree) FindNextSibling(flatBlockID string) *FlatBlock {
	block := tree.Find(flatBlockID)

	if block == nil {
		return nil
	}

	children := tree.Children(block.ParentID)

	_, index, found := lo.FindIndexOf(children, func(bExt FlatBlock) bool {
		return bExt.ID == flatBlockID
	})

	if !found {
		return nil
	}

	if index == len(children)-1 {
		return nil
	}

	return &children[index+1]
}

func (tree *FlatTree) FindPreviousSibling(flatBlockID string) *FlatBlock {
	block := tree.Find(flatBlockID)

	if block == nil {
		return nil
	}

	children := tree.Children(block.ParentID)

	_, index, found := lo.FindIndexOf(children, func(bExt FlatBlock) bool {
		return bExt.ID == flatBlockID
	})

	if !found {
		return nil
	}

	if index == 0 {
		return nil
	}

	return &children[index-1]
}

func (tree *FlatTree) MoveDown(flatBlockID string) {
	block := tree.Find(flatBlockID)

	if block == nil {
		return
	}

	next := tree.FindNextSibling(block.ID)

	if next == nil {
		return
	}

	nextSequence := next.Sequence
	sequence := block.Sequence

	block.Sequence = nextSequence
	next.Sequence = sequence

	tree.Update(*block)
	tree.Update(*next)

	tree.RecalculateSequences(block.ParentID)
}

func (tree *FlatTree) MoveToParent(flatBlockID, parentID string) {
	block := tree.Find(flatBlockID)

	if block == nil {
		return
	}

	if block.ParentID == parentID {
		return
	}

	children := tree.Children(block.ParentID)

	block.ParentID = parentID
	block.Sequence = len(children)

	tree.Update(*block)

	tree.RecalculateSequences(parentID)
}

func (tree *FlatTree) MoveToPosition(flatBlockID, parentID string, position int) {
	tree.MoveToParent(flatBlockID, parentID)

	block := tree.Find(flatBlockID)

	if block == nil {
		return
	}

	if block.Sequence == position {
		return
	}

	if position < 0 {
		return // position already at the top
	}

	if position > len(tree.Children(parentID)) {
		return // position already at the bottom
	}

	if block.Sequence < position {
		// move down
		for i := block.Sequence; i < position; i++ {
			tree.MoveDown(flatBlockID)
		}
	} else {
		// move up
		for i := block.Sequence; i > position; i-- {
			tree.MoveUp(flatBlockID)
		}
	}
}

func (tree *FlatTree) MoveUp(flatBlockID string) {
	block := tree.Find(flatBlockID)

	if block == nil {
		return
	}

	previous := tree.FindPreviousSibling(block.ID)

	if previous == nil {
		return
	}

	previousSequence := previous.Sequence
	sequence := block.Sequence

	block.Sequence = previousSequence
	previous.Sequence = sequence

	tree.Update(*block)
	tree.Update(*previous)

	tree.RecalculateSequences(block.ParentID)
}

func (tree *FlatTree) Parent(flatBlockID string) *FlatBlock {
	block := tree.Find(flatBlockID)

	if block == nil {
		return nil
	}

	return tree.Find(block.ParentID)
}

func (tree *FlatTree) RecalculateSequences(blockID string) {
	children := tree.Children(blockID)

	for i, child := range children {
		child.Sequence = i
		tree.Update(child)
	}
}

// Remove removes the block with the given id
//
// Buisiness Logic:
// - checks if the block exists, if not, do nothing
// - removes the block from the list
// - recalculates the sequences of the parent's children
func (tree *FlatTree) Remove(flatBlockID string) {
	flatBlock := tree.Find(flatBlockID)

	if flatBlock == nil {
		return
	}

	parentID := flatBlock.ParentID
	for i, ext := range tree.list {
		if ext.ID == flatBlockID {
			tree.list = append(tree.list[:i], tree.list[i+1:]...)
		}
	}

	tree.RemoveOrphans()

	tree.RecalculateSequences(parentID)
}

// RemoveOrphans removes all orphaned blocks that have no parent
//
// Buisiness Logic:
// - finds and creates a new list without orphaned blocks
// - non orphaned blocks are the ones that have a parent or root blocks
// - updates the list with the new list
//
// Parameters:
// - none
//
// Returns:
// - none
func (tree *FlatTree) RemoveOrphans() {
	nonOrphans := make([]FlatBlock, 0)

	for _, block := range tree.list {
		if block.ParentID == "" {
			nonOrphans = append(nonOrphans, block)
			continue
		}

		parent := tree.Parent(block.ID)

		if parent != nil {
			nonOrphans = append(nonOrphans, block)
		}
	}

	tree.list = nonOrphans
}

func (tree *FlatTree) Traverse(blockID string) []FlatBlock {
	block := tree.Find(blockID)

	if block == nil {
		return []FlatBlock{}
	}

	travsersed := make([]FlatBlock, 0)
	travsersed = append(travsersed, *block)

	for _, child := range tree.Children(blockID) {
		travsersed = append(travsersed, tree.Traverse(child.ID)...)
	}

	return travsersed
}

func (tree *FlatTree) Update(flatBlock FlatBlock) {
	for i, ext := range tree.list {
		if ext.ID == flatBlock.ID {
			tree.list[i] = flatBlock
		}
	}
}

func (tree *FlatTree) ToBlocks() []ui.BlockInterface {
	parentBlocks := tree.Children("")

	blocks := make([]ui.BlockInterface, 0)

	for _, flatBlock := range parentBlocks {
		blocks = append(blocks, tree.flatBlockToBlock(flatBlock))
	}

	return blocks
}

func (tree *FlatTree) flatBlockToBlock(flatBlock FlatBlock) ui.BlockInterface {
	childrenExt := tree.Children(flatBlock.ID)

	children := []ui.BlockInterface{}
	for _, childExt := range childrenExt {
		children = append(children, tree.flatBlockToBlock(childExt))
	}

	block := ui.NewBlockFromMap(map[string]interface{}{
		"id":         flatBlock.ID,
		"type":       flatBlock.Type,
		"parameters": flatBlock.Parameters,
		"children":   children,
	})

	return block
}

func traverse(blocks []ui.BlockInterface, parentID string) []FlatBlock {
	list := make([]FlatBlock, 0)

	for index, block := range blocks {
		flatBlock := blockToFlatBlock(block, parentID, index)
		list = append(list, flatBlock)
		list = append(list, traverse(block.Children(), block.ID())...)
	}

	return list
}

func blockToFlatBlock(block ui.BlockInterface, parentID string, sequence int) FlatBlock {
	return FlatBlock{
		ID:         block.ID(),
		Type:       block.Type(),
		ParentID:   parentID,
		Sequence:   sequence,
		Parameters: block.Parameters(),
	}
}
