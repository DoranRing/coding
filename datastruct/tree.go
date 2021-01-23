package datastruct

import "fmt"

// Tree 树型ADT
type Tree interface {
	Root() TreeNode

	// PreorderTraversal 前序遍历
	PreorderTraversal()

	// PostorderTraversal 后序遍历
	PostorderTraversal()
}

// TreeNode 树型节点ADT
type TreeNode interface {
	Val() int

	Children() []TreeNode

	FirstChildren() TreeNode

	FindChildren(idx int) TreeNode

	AddChildren(val int)

	DeleteChildByIdx(idx int)

	DeleteChild(treeNode TreeNode)
}

// LinkedTree 链式结构的树型实现
type LinkedTree struct {
	root TreeNode
}

func (l *LinkedTree) Root() TreeNode {
	return l.root
}

func (l *LinkedTree) PreorderTraversal() {
	preorderTraversal(l.root)
	fmt.Printf("\n")
}

func preorderTraversal(treeNode TreeNode) {
	fmt.Printf("%d,", treeNode.Val())
	if treeNode.Children() != nil && len(treeNode.Children()) > 0 {
		for _, val := range treeNode.Children() {
			preorderTraversal(val)
		}
	}
}

func (l *LinkedTree) PostorderTraversal() {
	postorderTraversal(l.root)
	fmt.Printf("\n")
}

func postorderTraversal(treeNode TreeNode) {
	if treeNode.Children() != nil && len(treeNode.Children()) > 0 {
		for _, val := range treeNode.Children() {
			preorderTraversal(val)
		}
	}
	fmt.Printf("%d,", treeNode.Val())
}

func NewLinkedTree(root TreeNode) *LinkedTree {
	return &LinkedTree{root: root}
}

// LinkedTreeNode 链式结构的树型节点实现
type LinkedTreeNode struct {
	val      int
	children []TreeNode
}

func NewLinkedTreeNode(val int) *LinkedTreeNode {
	return &LinkedTreeNode{val: val}
}

func (l *LinkedTreeNode) Val() int {
	return l.val
}

func (l *LinkedTreeNode) Children() []TreeNode {
	return l.children
}

func (l *LinkedTreeNode) FirstChildren() TreeNode {
	if l.children != nil && len(l.children) > 0 {
		return l.children[0]
	}
	return nil
}

func (l *LinkedTreeNode) FindChildren(idx int) TreeNode {
	if l.children != nil && idx < len(l.children) {
		return l.children[idx]
	}
	return nil
}

func (l *LinkedTreeNode) AddChildren(val int) {
	if l.children == nil {
		l.children = []TreeNode{NewLinkedTreeNode(val)}
	} else {
		l.children = append(l.children, NewLinkedTreeNode(val))
	}
}

func (l *LinkedTreeNode) DeleteChildByIdx(idx int) {
	if l.children == nil || len(l.children) <= idx {
		return
	}
	l.children = append(l.children[:idx], l.children[idx+1:]...)
}

func (l *LinkedTreeNode) DeleteChild(treeNode TreeNode) {
	if l.children == nil {
		return
	}
	for k, v := range l.children {
		if v == treeNode {
			l.children = append(l.children[:k], l.children[k+1:]...)
		}
	}
}
