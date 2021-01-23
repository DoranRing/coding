package datastruct

import (
	"testing"
)

func TestNewLinkedTree(t *testing.T) {
	tree := NewLinkedTree(NewLinkedTreeNode(10))
	root := tree.Root()
	root.AddChildren(12)
	root.AddChildren(7)
	firstChildren := root.FirstChildren()
	if len(root.Children()) != 2 {
		t.Errorf("actual: %d, want: %d", len(root.Children()), 2)
	}
	if firstChildren.Val() != 12 {
		t.Errorf("actual:%d, want:%d", firstChildren.Val(), 12)
	}

	treeNode := root.FindChildren(1)
	if treeNode.Val() != 7 {
		t.Errorf("actual:%d, want:%d", treeNode.Val(), 7)
	}

	root.DeleteChildByIdx(0)
	if len(root.Children()) != 1 {
		t.Errorf("actual: %d, want: %d", len(root.Children()), 1)
	}

	root.DeleteChild(treeNode)
	if len(root.Children()) != 0 {
		t.Errorf("actual: %d, want: %d", len(root.Children()), 0)
	}
}

func TestLinkedTree_Traversal(t *testing.T) {
	tree := NewLinkedTree(NewLinkedTreeNode(10))
	root := tree.Root()
	root.AddChildren(12)
	root.AddChildren(7)
	firstChildren := root.FirstChildren()
	firstChildren.AddChildren(30)
	tree.PreorderTraversal()
	tree.PostorderTraversal()
}
