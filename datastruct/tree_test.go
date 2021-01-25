package datastruct

import (
	"testing"
)

func TestNewLinkedTree(t *testing.T) {
	tree := NewLinkedTree(10)
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
	tree := NewLinkedTree(10)
	root := tree.Root()
	root.AddChildren(12)
	root.AddChildren(7)
	firstChildren := root.FirstChildren()
	firstChildren.AddChildren(30)
	tree.PreorderTraversal()
	tree.PostorderTraversal()
}

func TestNewLinkedBinaryTree(t *testing.T) {
	tree := NewLinkedBinaryTree(20)
	tree.Insert(10)
	tree.Insert(50)
	tree.Insert(5)
	tree.Insert(15)
	tree.Insert(25)
	tree.Insert(75)

	if tree.IsEmpty() {
		t.Errorf("actual: %t, want: %t", tree.IsEmpty(), false)
	}

	contains := tree.Contains(10)
	if !contains {
		t.Errorf("actual: %t, want: %t", contains, true)
	}
	nonContains := tree.Contains(17)
	if nonContains {
		t.Errorf("actual: %t, want: %t", nonContains, true)
	}

	min, err := tree.FindMin()
	if err != nil {
		t.Errorf("error: %e", err)
	}
	if min != 5 {
		t.Errorf("actual: %d, want: %d", min, 5)
	}

	max, err := tree.FindMax()
	if err != nil {
		t.Errorf("error: %e", err)
	}
	if max != 75 {
		t.Errorf("actual: %d, want: %d", max, 75)
	}

	tree.Remove(5)
	treeNode := tree.Find(10)
	if treeNode.Left() != nil {
		t.Errorf("actual: %t, want: %t", treeNode.Left() == nil, true)
	}

	tree.Remove(10)
	root := tree.Root()
	if root.Left() == nil || root.Left().Val() != 15 {
		var actual int
		if root.Left() == nil {
			actual = 0
		} else {
			actual = root.Left().Val()
		}
		t.Errorf("actual: %d, want: %d", actual, 15)
	}

	tree.Remove(50)
	if root.Right() == nil || root.Right().Val() != 75 {
		var actual int
		if root.Right() == nil {
			actual = 0
		} else {
			actual = root.Right().Val()
		}
		t.Errorf("actual: %d, want: %d", actual, 75)
	}

}
