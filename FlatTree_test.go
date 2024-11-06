package blockeditor

import (
	"testing"

	"github.com/gouniverse/ui"
)

func Test_FlatTree_Children(t *testing.T) {
	tree := NewFlatTree([]ui.BlockInterface{
		ui.NewBlock().SetType("test").SetID("1").SetChildren([]ui.BlockInterface{}),
		ui.NewBlock().SetType("test").SetID("2").SetChildren([]ui.BlockInterface{
			ui.NewBlock().SetType("test").SetID("3").SetChildren([]ui.BlockInterface{}),
			ui.NewBlock().SetType("test").SetID("4").SetChildren([]ui.BlockInterface{}),
		}),
		ui.NewBlock().SetType("test").SetID("5").SetChildren([]ui.BlockInterface{}),
	})

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

func Test_FlatTree_Find(t *testing.T) {
	tree := NewFlatTree([]ui.BlockInterface{
		ui.NewBlock().SetType("test").SetID("1").SetChildren([]ui.BlockInterface{}),
		ui.NewBlock().SetType("test").SetID("2").SetChildren([]ui.BlockInterface{
			ui.NewBlock().SetType("test").SetID("3").SetChildren([]ui.BlockInterface{}),
			ui.NewBlock().SetType("test").SetID("4").SetChildren([]ui.BlockInterface{}),
		}),
		ui.NewBlock().SetType("test").SetID("5").SetChildren([]ui.BlockInterface{}),
	})

	block := tree.Find("4")

	if block == nil {
		t.Fatal("Expected block, got nil")
	}

	if block.ID != "4" {
		t.Fatal("Expected block ID '4', got:", block.ID)
	}
}

func Test_FlatTree_FindNextSibling(t *testing.T) {
	tree := NewFlatTree([]ui.BlockInterface{
		ui.NewBlock().SetType("test").SetID("1").SetChildren([]ui.BlockInterface{}),
		ui.NewBlock().SetType("test").SetID("2").SetChildren([]ui.BlockInterface{
			ui.NewBlock().SetType("test").SetID("3").SetChildren([]ui.BlockInterface{}),
			ui.NewBlock().SetType("test").SetID("4").SetChildren([]ui.BlockInterface{}),
			ui.NewBlock().SetType("test").SetID("5").SetChildren([]ui.BlockInterface{}),
		}),
		ui.NewBlock().SetType("test").SetID("6").SetChildren([]ui.BlockInterface{}),
	})

	block := tree.FindNextSibling("4")

	if block == nil {
		t.Fatal("Expected block, got nil")
	}

	if block.ID != "5" {
		t.Fatal("Expected block ID '5', got:", block.ID)
	}
}

func Test_FlatTree_FindPreviousSibling(t *testing.T) {
	tree := NewFlatTree([]ui.BlockInterface{
		ui.NewBlock().SetType("test").SetID("1").SetChildren([]ui.BlockInterface{}),
		ui.NewBlock().SetType("test").SetID("2").SetChildren([]ui.BlockInterface{
			ui.NewBlock().SetType("test").SetID("3").SetChildren([]ui.BlockInterface{}),
			ui.NewBlock().SetType("test").SetID("4").SetChildren([]ui.BlockInterface{}),
			ui.NewBlock().SetType("test").SetID("5").SetChildren([]ui.BlockInterface{}),
		}),
		ui.NewBlock().SetType("test").SetID("6").SetChildren([]ui.BlockInterface{}),
	})

	block := tree.FindPreviousSibling("4")

	if block == nil {
		t.Fatal("Expected block, got nil")
	}

	if block.ID != "3" {
		t.Fatal("Expected block ID '3', got:", block.ID)
	}
}

func Test_FlatTree_FlatBlockToBlock(t *testing.T) {
	tree := NewFlatTree([]ui.BlockInterface{
		ui.NewBlock().SetType("test").SetID("1").SetChildren([]ui.BlockInterface{}),
		ui.NewBlock().SetType("test").SetID("2").SetChildren([]ui.BlockInterface{
			ui.NewBlock().SetType("test").SetID("3").SetChildren([]ui.BlockInterface{}),
			ui.NewBlock().SetType("test").SetID("4").SetChildren([]ui.BlockInterface{}),
		}),
		ui.NewBlock().SetType("test").SetID("5").SetChildren([]ui.BlockInterface{}),
	})

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
	tree := NewFlatTree([]ui.BlockInterface{
		ui.NewBlock().SetType("test").SetID("1").SetChildren([]ui.BlockInterface{}),
		ui.NewBlock().SetType("test").SetID("2").SetChildren([]ui.BlockInterface{
			ui.NewBlock().SetType("test").SetID("3").SetChildren([]ui.BlockInterface{}),
			ui.NewBlock().SetType("test").SetID("4").SetChildren([]ui.BlockInterface{}),
			ui.NewBlock().SetType("test").SetID("5").SetChildren([]ui.BlockInterface{}),
		}),
		ui.NewBlock().SetType("test").SetID("6").SetChildren([]ui.BlockInterface{}),
	})

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
	tree := NewFlatTree([]ui.BlockInterface{
		ui.NewBlock().SetType("test").SetID("1").SetChildren([]ui.BlockInterface{}),
		ui.NewBlock().SetType("test").SetID("2").SetChildren([]ui.BlockInterface{
			ui.NewBlock().SetType("test").SetID("3").SetChildren([]ui.BlockInterface{}),
			ui.NewBlock().SetType("test").SetID("4").SetChildren([]ui.BlockInterface{}),
			ui.NewBlock().SetType("test").SetID("5").SetChildren([]ui.BlockInterface{}),
		}),
		ui.NewBlock().SetType("test").SetID("6").SetChildren([]ui.BlockInterface{}),
	})

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
	tree := NewFlatTree([]ui.BlockInterface{
		ui.NewBlock().SetType("test").SetID("1").SetChildren([]ui.BlockInterface{}),
		ui.NewBlock().SetType("test").SetID("2").SetChildren([]ui.BlockInterface{
			ui.NewBlock().SetType("test").SetID("3").SetChildren([]ui.BlockInterface{}),
			ui.NewBlock().SetType("test").SetID("4").SetChildren([]ui.BlockInterface{}),
			ui.NewBlock().SetType("test").SetID("5").SetChildren([]ui.BlockInterface{}),
		}),
		ui.NewBlock().SetType("test").SetID("6").SetChildren([]ui.BlockInterface{}),
	})

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
	tree := NewFlatTree([]ui.BlockInterface{
		ui.NewBlock().SetType("test").SetID("1").SetChildren([]ui.BlockInterface{}),
		ui.NewBlock().SetType("test").SetID("2").SetChildren([]ui.BlockInterface{
			ui.NewBlock().SetType("test").SetID("3").SetChildren([]ui.BlockInterface{}),
			ui.NewBlock().SetType("test").SetID("4").SetChildren([]ui.BlockInterface{}),
			ui.NewBlock().SetType("test").SetID("5").SetChildren([]ui.BlockInterface{}),
		}),
		ui.NewBlock().SetType("test").SetID("6").SetChildren([]ui.BlockInterface{}),
	})

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
	tree := NewFlatTree([]ui.BlockInterface{
		ui.NewBlock().SetType("test").SetID("1").SetChildren([]ui.BlockInterface{}),
		ui.NewBlock().SetType("test").SetID("2").SetChildren([]ui.BlockInterface{
			ui.NewBlock().SetType("test").SetID("3").SetChildren([]ui.BlockInterface{}),
			ui.NewBlock().SetType("test").SetID("4").SetChildren([]ui.BlockInterface{}),
		}),
		ui.NewBlock().SetType("test").SetID("5").SetChildren([]ui.BlockInterface{}),
	})

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
	tree := NewFlatTree([]ui.BlockInterface{
		ui.NewBlock().SetType("test").SetID("1").SetChildren([]ui.BlockInterface{}),
		ui.NewBlock().SetType("test").SetID("2").SetChildren([]ui.BlockInterface{
			ui.NewBlock().SetType("test").SetID("3").SetChildren([]ui.BlockInterface{}),
			ui.NewBlock().SetType("test").SetID("4").SetChildren([]ui.BlockInterface{}),
			ui.NewBlock().SetType("test").SetID("5").SetChildren([]ui.BlockInterface{}),
		}),
		ui.NewBlock().SetType("test").SetID("6").SetChildren([]ui.BlockInterface{}),
	})

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

func Test_FlatTree_RecalculateSequences(t *testing.T) {
	tree := NewFlatTree([]ui.BlockInterface{
		ui.NewBlock().SetType("test").SetID("1").SetChildren([]ui.BlockInterface{}),
		ui.NewBlock().SetType("test").SetID("2").SetChildren([]ui.BlockInterface{
			ui.NewBlock().SetType("test").SetID("3").SetChildren([]ui.BlockInterface{}),
			ui.NewBlock().SetType("test").SetID("4").SetChildren([]ui.BlockInterface{}),
		}),
		ui.NewBlock().SetType("test").SetID("5").SetChildren([]ui.BlockInterface{}),
	})

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
