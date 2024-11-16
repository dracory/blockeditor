package blockeditor

import (
	"testing"

	"github.com/gouniverse/ui"
)

func Test_FlatTree_Children(t *testing.T) {
	blocks := ui.BlocksFromMap([]map[string]any{
		{
			"id": "1", "type": "test",
		},
		{
			"id": "2", "type": "test",
			"children": []map[string]any{
				{"id": "3", "type": "test"},
				{"id": "4", "type": "test"},
			},
		},
		{"id": "5", "type": "test"},
	})

	tree := NewFlatTree(blocks)

	chldren := tree.Children("2")

	if len(chldren) != 2 {
		t.Fatal("Expected 2 children, got:", len(chldren))
	}

	if chldren[0].ID != "3" {
		t.Fatal("Expected child ID '3', got:", chldren[0].ID)
	}

	if chldren[1].ID != "4" {
		t.Fatal("Expected child ID '4', got:", chldren[1].ID)
	}
}

func Test_FlatTree_Duplicate(t *testing.T) {
	tree := NewFlatTree(ui.BlocksFromMap([]map[string]any{
		{
			"id": "1", "type": "test",
		},
		{
			"id": "2", "type": "test",
			"children": []map[string]any{
				{"id": "3", "type": "test"},
				{"id": "4", "type": "test"},
			},
		},
		{"id": "5", "type": "test"},
	}))

	if len(tree.list) != 5 {
		t.Fatal("Expected 5 blocks, got:", len(tree.list))
	}

	tree.Duplicate("2")

	if len(tree.list) != 8 {
		t.Fatal("Expected 8 blocks, got:", len(tree.list))
	}

	duplicatedBlock := tree.list[5]

	if duplicatedBlock.ParentID != "" {
		t.Fatal("Expected parent ID '', got:", tree.list[5].ParentID)
	}

	if duplicatedBlock.Sequence != 4 {
		t.Fatal("Expected sequence 4, got:", tree.list[6].Sequence)
	}

	if len(tree.Children(duplicatedBlock.ID)) != 2 {
		t.Fatal("Expected 2 children, got:", len(tree.Children(duplicatedBlock.ID)))
	}
}

func Test_FlatTree_Find(t *testing.T) {
	tree := NewFlatTree(ui.BlocksFromMap([]map[string]any{
		{
			"id": "1", "type": "test",
		},
		{
			"id": "2", "type": "test",
			"children": []map[string]any{
				{"id": "3", "type": "test"},
				{"id": "4", "type": "test"},
			},
		},
		{"id": "5", "type": "test"},
	}))

	block := tree.Find("4")

	if block == nil {
		t.Fatal("Expected block, got nil")
	}

	if block.ID != "4" {
		t.Fatal("Expected block ID '4', got:", block.ID)
	}
}

func Test_FlatTree_FindNextSibling(t *testing.T) {
	tree := NewFlatTree(ui.BlocksFromMap([]map[string]any{
		{
			"id": "1", "type": "test",
		},
		{
			"id": "2", "type": "test",
			"children": []map[string]any{
				{"id": "3", "type": "test"},
				{"id": "4", "type": "test"},
				{"id": "5", "type": "test"},
			},
		},
		{"id": "6", "type": "test"},
	}))

	block := tree.FindNextSibling("4")

	if block == nil {
		t.Fatal("Expected block, got nil")
	}

	if block.ID != "5" {
		t.Fatal("Expected block ID '5', got:", block.ID)
	}
}

func Test_FlatTree_FindPreviousSibling(t *testing.T) {
	tree := NewFlatTree(ui.BlocksFromMap([]map[string]any{
		{
			"id": "1", "type": "test",
		},
		{
			"id": "2", "type": "test",
			"children": []map[string]any{
				{"id": "3", "type": "test"},
				{"id": "4", "type": "test"},
				{"id": "5", "type": "test"},
			},
		},
		{"id": "6", "type": "test"},
	}))

	block := tree.FindPreviousSibling("4")

	if block == nil {
		t.Fatal("Expected block, got nil")
	}

	if block.ID != "3" {
		t.Fatal("Expected block ID '3', got:", block.ID)
	}
}

func Test_FlatTree_FlatBlockToBlock(t *testing.T) {
	tree := NewFlatTree(ui.BlocksFromMap([]map[string]any{
		{"id": "1", "type": "test"},
		{
			"id": "2", "type": "test",
			"children": []map[string]any{
				{"id": "3", "type": "test"},
				{"id": "4", "type": "test"},
			},
		},
		{"id": "5", "type": "test"},
	}))

	block := tree.flatBlockToBlock(tree.list[1])

	if block == nil {
		t.Fatal("Expected block, got nil")
	}

	if block.ID() != "2" {
		t.Fatal("Expected block ID '2', got:", block.ID())
	}

	if len(block.Children()) != 2 {
		t.Fatal("Expected 2 child, got:", len(block.Children()))
	}
}

func Test_FlatTree_MoveDown(t *testing.T) {
	tree := NewFlatTree(ui.BlocksFromMap([]map[string]any{
		{"id": "1", "type": "test"},
		{
			"id": "2", "type": "test",
			"children": []map[string]any{
				{"id": "3", "type": "test"},
				{"id": "4", "type": "test"},
				{"id": "5", "type": "test"},
			},
		},
		{"id": "6", "type": "test"},
	}))

	tree.MoveDown("4")

	if len(tree.Children("2")) != 3 {
		t.Fatal("Expected 3 child, got:", len(tree.Children("2")))
	}

	if tree.Children("2")[2].ID != "4" {
		t.Fatal("Expected child ID '4', got:", tree.Children("2")[2].ID)
	}

	if tree.Children("2")[1].ID != "5" {
		t.Fatal("Expected child ID '5', got:", tree.Children("2")[1].ID)
	}

	if tree.Children("2")[0].ID != "3" {
		t.Fatal("Expected child ID '3', got:", tree.Children("2")[0].ID)
	}
}

func Test_FlatTree_MoveToParent(t *testing.T) {
	tree := NewFlatTree(ui.BlocksFromMap([]map[string]any{
		{"id": "1", "type": "test"},
		{
			"id": "2", "type": "test",
			"children": []map[string]any{
				{"id": "3", "type": "test"},
				{"id": "4", "type": "test"},
				{"id": "5", "type": "test"},
			},
		},
		{"id": "6", "type": "test"},
	}))

	tree.MoveToParent("6", "2")

	if len(tree.Children("2")) != 4 {
		t.Fatal("Expected 4 child, got:", len(tree.Children("2")))
	}

	if tree.Children("2")[0].ID != "3" {
		t.Fatal("Expected child ID '3', got:", tree.Children("2")[0].ID)
	}

	if tree.Children("2")[1].ID != "4" {
		t.Fatal("Expected child ID '4', got:", tree.Children("2")[1].ID)
	}

	if tree.Children("2")[2].ID != "5" {
		t.Fatal("Expected child ID '5', got:", tree.Children("2")[1].ID)
	}

	if tree.Children("2")[3].ID != "6" {
		t.Fatal("Expected child ID '6', got:", tree.Children("2")[3].ID)
	}
}

func Test_FlatTree_MoveToPosition(t *testing.T) {
	tree := NewFlatTree(ui.BlocksFromMap([]map[string]any{
		{"id": "1", "type": "test"},
		{
			"id": "2", "type": "test",
			"children": []map[string]any{
				{"id": "3", "type": "test"},
				{"id": "4", "type": "test"},
				{"id": "5", "type": "test"},
			},
		},
		{"id": "6", "type": "test"},
	}))

	tree.MoveToPosition("6", "2", 2)

	if len(tree.Children("2")) != 4 {
		t.Fatal("Expected 4 child, got:", len(tree.Children("2")))
	}

	if tree.Children("2")[0].ID != "3" {
		t.Fatal("Expected child ID '3', got:", tree.Children("2")[0].ID)
	}

	if tree.Children("2")[1].ID != "4" {
		t.Fatal("Expected child ID '4', got:", tree.Children("2")[1].ID)
	}

	if tree.Children("2")[2].ID != "6" {
		t.Fatal("Expected child ID '6', got:", tree.Children("2")[2].ID)
	}

	if tree.Children("2")[3].ID != "5" {
		t.Fatal("Expected child ID '5', got:", tree.Children("2")[3].ID)
	}

	if len(tree.Children("")) != 2 {
		t.Fatal("Expected 2 child, got:", len(tree.Children("")))
	}
}

func Test_FlatTree_MoveUp(t *testing.T) {
	tree := NewFlatTree(ui.BlocksFromMap([]map[string]any{
		{"id": "1", "type": "test"},
		{
			"id": "2", "type": "test",
			"children": []map[string]any{
				{"id": "3", "type": "test"},
				{"id": "4", "type": "test"},
				{"id": "5", "type": "test"},
			},
		},
		{"id": "6", "type": "test"},
	}))

	tree.MoveUp("4")

	if len(tree.Children("2")) != 3 {
		t.Fatal("Expected 3 child, got:", len(tree.Children("2")))
	}

	if tree.Children("2")[0].ID != "4" {
		t.Fatal("Expected child ID '4', got:", tree.Children("2")[0].ID)
	}

	if tree.Children("2")[1].ID != "3" {
		t.Fatal("Expected child ID '3', got:", tree.Children("2")[1].ID)
	}

	if tree.Children("2")[2].ID != "5" {
		t.Fatal("Expected child ID '5', got:", tree.Children("2")[1].ID)
	}
}

func Test_FlatTree_Parent(t *testing.T) {
	tree := NewFlatTree(ui.BlocksFromMap([]map[string]any{
		{"id": "1", "type": "test"},
		{
			"id": "2", "type": "test",
			"children": []map[string]any{
				{"id": "3", "type": "test"},
				{"id": "4", "type": "test"},
			},
		},
		{"id": "5", "type": "test"},
	}))

	block := tree.Parent("4")

	if block == nil {
		t.Fatal("Expected block, got nil")
	}

	if block.ID != "2" {
		t.Fatal("Expected block ID '2', got:", block.ID)
	}

	if len(tree.Children("2")) != 2 {
		t.Fatal("Expected 2 child, got:", len(tree.Children("2")))
	}
}

func Test_FlatTree_Remove(t *testing.T) {
	tree := NewFlatTree(ui.BlocksFromMap([]map[string]any{
		{"id": "1", "type": "test"},
		{
			"id": "2", "type": "test",
			"children": []map[string]any{
				{"id": "3", "type": "test"},
				{"id": "4", "type": "test"},
				{"id": "5", "type": "test"},
			},
		},
		{"id": "6", "type": "test"},
	}))

	tree.Remove("4")

	if len(tree.Children("2")) != 2 {
		t.Fatal("Expected 2 child, got:", len(tree.Children("2")))
	}

	if tree.Children("2")[0].ID != "3" {
		t.Fatal("Expected child ID '3', got:", tree.Children("2")[0].ID)
	}

	if tree.Children("2")[0].Sequence != 0 {
		t.Fatal("Expected child sequence 0, got:", tree.Children("2")[0].Sequence)
	}

	if tree.Children("2")[1].ID != "5" {
		t.Fatal("Expected child ID '5', got:", tree.Children("2")[1].ID)
	}

	if tree.Children("2")[1].Sequence != 1 {
		t.Fatal("Expected child sequence 1, got:", tree.Children("2")[1].Sequence)
	}
}

func Test_FlatTree_RemoveOrphans(t *testing.T) {
	tree := NewFlatTree(ui.BlocksFromMap([]map[string]any{
		{"id": "1", "type": "test"},
		{
			"id": "2", "type": "test",
			"children": []map[string]any{
				{"id": "3", "type": "test"},
				{"id": "4", "type": "test"},
				{"id": "5", "type": "test"},
			},
		},
		{"id": "6", "type": "test"},
	}))

	// append orphan blocks
	tree.list = append(tree.list, FlatBlock{ID: "77", Type: "orphan", ParentID: "43", Sequence: 1})
	tree.list = append(tree.list, FlatBlock{ID: "84", Type: "orphan", ParentID: "43", Sequence: 0})
	// prepend orphan blocks
	tree.list = append([]FlatBlock{{ID: "73", Type: "orphan", ParentID: "43", Sequence: 1}}, tree.list...)
	tree.list = append([]FlatBlock{{ID: "86", Type: "orphan", ParentID: "43", Sequence: 0}}, tree.list...)
	// insert orphan blocks at the middle
	tree.list = append(tree.list[:2], append([]FlatBlock{{ID: "99", Type: "orphan", ParentID: "43", Sequence: 1}}, tree.list[2:]...)...)

	orphanIDs := []string{"77", "84", "73", "86", "99"}

	for _, orphanID := range orphanIDs {
		if !tree.Exists(orphanID) {
			t.Fatal("Expected orphan block, got nil, ID:", orphanID)
		}
	}

	if len(tree.list) != 11 {
		t.Fatal("Expected 11 blocks, got:", len(tree.list))
	}

	tree.RemoveOrphans()

	for _, orphanID := range orphanIDs {
		if tree.Exists(orphanID) {
			t.Fatal("Expected orphan block to be removed, ID:", orphanID)
		}
	}

	if len(tree.list) != 6 {
		t.Fatal("Expected 6 blocks, got:", len(tree.list))
	}
}

func Test_FlatTree_RecalculateSequences(t *testing.T) {
	tree := NewFlatTree(ui.BlocksFromMap([]map[string]any{
		{"id": "1", "type": "test"},
		{
			"id": "2", "type": "test",
			"children": []map[string]any{
				{"id": "3", "type": "test"},
				{"id": "4", "type": "test"},
			},
		},
		{"id": "5", "type": "test"},
	}))

	tree.RecalculateSequences("2")

	if len(tree.Children("2")) != 2 {
		t.Fatal("Expected 2 children, got:", len(tree.Children("2")))
	}

	if tree.Children("2")[0].Sequence != 0 {
		t.Fatal("Expected sequence 0, got:", tree.Children("2")[0].Sequence)
	}

	if tree.Children("2")[1].Sequence != 1 {
		t.Fatal("Expected sequence 1, got:", tree.Children("2")[1].Sequence)
	}

	if tree.Children("2")[0].ID != "3" {
		t.Fatal("Expected child ID '3', got:", tree.Children("2")[0].ID)
	}

	if tree.Children("2")[1].ID != "4" {
		t.Fatal("Expected child ID '4', got:", tree.Children("2")[1].ID)
	}
}

func Test_FlatTree_Traverse(t *testing.T) {
	tree := NewFlatTree(ui.BlocksFromMap([]map[string]any{
		{"id": "1", "type": "test"},
		{
			"id": "2", "type": "test",
			"children": []map[string]any{
				{
					"id": "2.1", "type": "test",
					"children": []map[string]any{
						{"id": "2.1.1", "type": "test"},
						{"id": "2.1.2", "type": "test"},
					},
				},
				{
					"id": "2.2", "type": "test",
					"children": []map[string]any{
						{"id": "2.2.1", "type": "test"},
						{"id": "2.2.2", "type": "test"},
					},
				},
			},
		},
		{"id": "3", "type": "test"},
	}))

	blocks := tree.Traverse("2")

	if len(blocks) != 7 {
		t.Fatal("Expected 2 blocks, got:", len(blocks))
	}

	if blocks[0].ID != "2" {
		t.Fatal("Expected block ID '2.1', got:", blocks[0].ID)
	}

	if blocks[1].ID != "2.1" {
		t.Fatal("Expected block ID '2.1', got:", blocks[1].ID)
	}

	if blocks[2].ID != "2.1.1" {
		t.Fatal("Expected block ID '2.1.1', got:", blocks[2].ID)
	}

	if blocks[3].ID != "2.1.2" {
		t.Fatal("Expected block ID '2.1.2', got:", blocks[3].ID)
	}

	if blocks[4].ID != "2.2" {
		t.Fatal("Expected block ID '2.2', got:", blocks[4].ID)
	}

	if blocks[5].ID != "2.2.1" {
		t.Fatal("Expected block ID '2.2.1', got:", blocks[5].ID)
	}

	if blocks[6].ID != "2.2.2" {
		t.Fatal("Expected block ID '2.2.2', got:", blocks[6].ID)
	}
}

func Test_FlatTree_Update(t *testing.T) {
	// tree := NewFlatTree([]ui.BlockInterface{
	// 	ui.Block().SetType("test").SetID("1").SetChildren([]ui.BlockInterface{}),
	// })

	tree := NewFlatTree(ui.BlocksFromMap([]map[string]any{
		{"id": "1", "type": "test"},
	}))

	block := tree.Find(`1`)

	if block == nil {
		t.Fatal("Expected block, got nil")
	}

	block.Type = "updated"

	tree.Update(*block)

	foundBlock := tree.Find(`1`)

	if foundBlock == nil {
		t.Fatal("Expected block, got nil")
	}

	if foundBlock.Type != "updated" {
		t.Fatal("Expected type 'updated', got:", foundBlock.Type)
	}
}
