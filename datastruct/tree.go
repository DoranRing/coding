package datastruct

import (
	"errors"
	"fmt"
)

// Tree 树型ADT
type Tree interface {
	Root() TreeNode

	// PreorderTraversal 前序遍历(深度优先遍历之一)
	PreorderTraversal()

	// PostorderTraversal 后序遍历(深度优先遍历之一)
	PostorderTraversal()

	// LevelOrderTraversal 层序遍历(广度优先遍历之一)
	LevelOrderTraversal()
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
	l.preorderTraversal(l.root)
	fmt.Printf("\n")
}

func (l *LinkedTree) preorderTraversal(treeNode TreeNode) {
	if treeNode == nil {
		return
	}
	fmt.Printf("%d,", treeNode.Val())
	if treeNode.Children() != nil && len(treeNode.Children()) > 0 {
		for _, val := range treeNode.Children() {
			l.preorderTraversal(val)
		}
	}
}

func (l *LinkedTree) PostorderTraversal() {
	l.postorderTraversal(l.root)
	fmt.Printf("\n")
}

func (l *LinkedTree) postorderTraversal(treeNode TreeNode) {
	if treeNode == nil {
		return
	}
	if treeNode.Children() != nil && len(treeNode.Children()) > 0 {
		for _, val := range treeNode.Children() {
			l.preorderTraversal(val)
		}
	}
	fmt.Printf("%d,", treeNode.Val())
}

func (l *LinkedTree) LevelOrderTraversal() {
	queue := make([]TreeNode, 0, 0)
	if l.root != nil {
		queue = append(queue, l.root)
	}
	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]
		fmt.Printf("%d,", top.Val())
		if top.Children() == nil {
			continue
		}
		for _, treeNode := range top.Children() {
			queue = append(queue, treeNode)
		}
	}
	fmt.Printf("\n")

}

func NewLinkedTree(val int) *LinkedTree {
	return &LinkedTree{root: NewLinkedTreeNode(val)}
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

// BinaryTree 二叉树ADT
type BinaryTree interface {
	Root() BinaryTreeNode

	// MakeEmpty 清空
	MakeEmpty()

	// IsEmpty 检查为空
	IsEmpty() bool

	// Contains 包含值
	Contains(val int) bool

	Find(val int) BinaryTreeNode

	// FindMax 查询最大值
	FindMax() (int, error)

	// FindMin 查询最小值
	FindMin() (int, error)

	// Insert 新增值
	Insert(val int)

	// 删除值
	Remove(val int)

	// Contains 包含值
	ContainsByNode(val int, treeNode BinaryTreeNode) bool

	// FindMax 查询最大值
	FindMaxByNode(treeNode BinaryTreeNode) (int, error)

	// FindMin 查询最小值
	FindMinByNode(treeNode BinaryTreeNode) (int, error)

	// Insert 新增值
	InsertByNode(val int, treeNode BinaryTreeNode)

	// 删除值
	RemoveByNode(val int, treeNode BinaryTreeNode)

	// PreorderTraversal 前序遍历
	PreorderTraversal()

	// InorderTraversal 中序遍历
	InorderTraversal()

	// PostorderTraversal 后序遍历
	PostorderTraversal()

	// LevelOrderTraversal 层序遍历
	LevelOrderTraversal()
}

// BinaryTreeNode 二叉树节点
type BinaryTreeNode interface {
	Val() int

	SetVal(val int)

	Left() BinaryTreeNode

	SetLeft(treeNode BinaryTreeNode)

	Right() BinaryTreeNode

	SetRight(treeNode BinaryTreeNode)
}

type LinkedBinaryTree struct {
	root BinaryTreeNode
}

func NewLinkedBinaryTree(val int) *LinkedBinaryTree {
	return &LinkedBinaryTree{root: NewLinkedBinaryTreeNode(val)}
}

func (l *LinkedBinaryTree) Root() BinaryTreeNode {
	return l.root
}

func (l *LinkedBinaryTree) MakeEmpty() {
	l.root = nil
}

func (l *LinkedBinaryTree) IsEmpty() bool {
	return l.root == nil
}

func (l *LinkedBinaryTree) Contains(val int) bool {
	return l.find(l.Root(), val) != nil
}

func (l *LinkedBinaryTree) Find(val int) BinaryTreeNode {
	return l.find(l.root, val)
}

func (l *LinkedBinaryTree) find(treeNode BinaryTreeNode, val int) BinaryTreeNode {
	if treeNode == nil {
		return nil
	}
	if treeNode.Val() == val {
		return treeNode
	} else if treeNode.Val() > val {
		return l.find(treeNode.Left(), val)
	} else {
		return l.find(treeNode.Right(), val)
	}
}

func (l *LinkedBinaryTree) FindMax() (int, error) {
	if l.root == nil {
		return 0, errors.New("tree is null")
	}
	return l.findMax(l.root), nil
}

func (l *LinkedBinaryTree) findMax(treeNode BinaryTreeNode) int {
	if treeNode.Right() == nil {
		return treeNode.Val()
	}
	return l.findMax(treeNode.Right())
}

func (l *LinkedBinaryTree) FindMin() (int, error) {
	if l.IsEmpty() {
		return 0, errors.New("tree is null")
	}
	return l.findMin(l.root), nil
}

func (l *LinkedBinaryTree) findMin(treeNode BinaryTreeNode) int {
	if treeNode.Left() == nil {
		return treeNode.Val()
	}
	return l.findMin(treeNode.Left())
}

func (l *LinkedBinaryTree) Insert(val int) {
	if l.root == nil {
		l.root = NewLinkedBinaryTreeNode(val)
	} else {
		l.insert(l.root, val)
	}
}

func (l *LinkedBinaryTree) insert(treeNode BinaryTreeNode, val int) BinaryTreeNode {
	if treeNode == nil {
		return NewLinkedBinaryTreeNode(val)
	}
	if treeNode.Val() > val {
		treeNode.SetLeft(l.insert(treeNode.Left(), val))
	} else if treeNode.Val() < val {
		treeNode.SetRight(l.insert(treeNode.Right(), val))
	}
	return treeNode
}

func (l *LinkedBinaryTree) Remove(val int) {
	if l.root == nil {
		return
	}

	l.root = l.remove(l.root, val)
}

func (l *LinkedBinaryTree) remove(treeNode BinaryTreeNode, val int) BinaryTreeNode {
	if treeNode == nil {
		return treeNode
	}
	//
	if treeNode.Val() > val {
		treeNode.SetLeft(l.remove(treeNode.Left(), val))
	} else if treeNode.Val() < val {
		treeNode.SetRight(l.remove(treeNode.Right(), val))
	} else if treeNode.Left() != nil && treeNode.Right() != nil {
		treeNode.SetVal(l.findMin(treeNode.Right()))
		treeNode.SetRight(l.remove(treeNode.Right(), treeNode.Val()))
	} else {
		if treeNode.Left() != nil {
			return treeNode.Left()
		} else {
			return treeNode.Right()
		}
	}
	return treeNode
}

func (l *LinkedBinaryTree) ContainsByNode(val int, treeNode BinaryTreeNode) bool {
	return l.find(treeNode, val) != nil
}

func (l *LinkedBinaryTree) FindMaxByNode(treeNode BinaryTreeNode) (int, error) {
	if treeNode == nil {
		return 0, errors.New("element is null")
	}
	return l.findMax(treeNode), nil
}

func (l *LinkedBinaryTree) FindMinByNode(treeNode BinaryTreeNode) (int, error) {
	if treeNode == nil {
		return 0, errors.New("element is null")
	}
	return l.findMin(treeNode), nil
}

func (l *LinkedBinaryTree) InsertByNode(val int, treeNode BinaryTreeNode) {
	if treeNode == nil {
		return
	}
	l.insert(treeNode, val)
}

func (l *LinkedBinaryTree) RemoveByNode(val int, treeNode BinaryTreeNode) {
	l.remove(treeNode, val)
}

func (l *LinkedBinaryTree) PreorderTraversal() {
	l.preorderTraversalBinary(l.root)
}

func (l *LinkedBinaryTree) preorderTraversalBinary(treeNode BinaryTreeNode) {
	if treeNode == nil {
		return
	}
	fmt.Printf("%d,", treeNode.Val())
	l.preorderTraversalBinary(treeNode.Left())
	l.preorderTraversalBinary(treeNode.Right())
}

func (l *LinkedBinaryTree) InorderTraversal() {
	l.inorderTraversalBinary(l.root)
}

func (l *LinkedBinaryTree) inorderTraversalBinary(treeNode BinaryTreeNode) {
	if treeNode == nil {
		return
	}
	l.inorderTraversalBinary(treeNode.Left())
	fmt.Printf("%d,", treeNode.Val())
	l.inorderTraversalBinary(treeNode.Right())
}

func (l *LinkedBinaryTree) PostorderTraversal() {
	l.postorderTraversalBinary(l.root)
}

func (l *LinkedBinaryTree) postorderTraversalBinary(treeNode BinaryTreeNode) {
	if treeNode == nil {
		return
	}
	l.postorderTraversalBinary(treeNode.Left())
	fmt.Printf("%d,", treeNode.Val())
	l.postorderTraversalBinary(treeNode.Right())
}

func (l *LinkedBinaryTree) LevelOrderTraversal() {
	queue := make([]BinaryTreeNode, 0, 0)
	if l.root != nil {
		queue = append(queue, l.root)
	}
	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]
		fmt.Printf("%d,", top.Val())
		if top.Left() != nil {
			queue = append(queue, top.Left())
		}
		if top.Right() != nil {
			queue = append(queue, top.Right())
		}
	}
	fmt.Printf("\n")
}

type LinkedBinaryTreeNode struct {
	val   int
	left  BinaryTreeNode
	right BinaryTreeNode
}

func (l *LinkedBinaryTreeNode) SetVal(val int) {
	l.val = val
}

func (l *LinkedBinaryTreeNode) Val() int {
	return l.val
}

func (l *LinkedBinaryTreeNode) Left() BinaryTreeNode {
	return l.left
}

func (l *LinkedBinaryTreeNode) SetLeft(treeNode BinaryTreeNode) {
	l.left = treeNode
}

func (l *LinkedBinaryTreeNode) Right() BinaryTreeNode {
	return l.right
}

func (l *LinkedBinaryTreeNode) SetRight(treeNode BinaryTreeNode) {
	l.right = treeNode
}

func NewLinkedBinaryTreeNode(val int) *LinkedBinaryTreeNode {
	return &LinkedBinaryTreeNode{val: val}
}
