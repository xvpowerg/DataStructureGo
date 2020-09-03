package main

import "fmt"

type Node interface {
	GetData() int
	GetLeft() Node
	GetRight() Node
	SetLeft(n Node)
	SetRight(n Node)
	String() string
}
type TreeNode struct {
	data      int
	leftNode  Node
	rightNode Node
}

func (tn *TreeNode) GetData() int {
	return tn.data
}
func (tn *TreeNode) GetLeft() Node {
	return tn.leftNode
}
func (tn *TreeNode) GetRight() Node {
	return tn.rightNode
}

func (tn *TreeNode) SetLeft(n Node) {
	tn.leftNode = n
}
func (tn *TreeNode) SetRight(n Node) {
	tn.rightNode = n
}

func (tn *TreeNode) String() string {
	return fmt.Sprintf("%d", tn.GetData())
}
func NewNode(data int) *TreeNode {
	tn := new(TreeNode)
	tn.data = data
	return tn
}

type Tree interface {
	Insert(n Node)
	GetRoot() Node
	Find(n Node) bool
}

type BinaryTree struct {
	root Node
}

func (bt *BinaryTree) GetRoot() Node {
	return bt.root
}
func (bt *BinaryTree) Find(n Node) bool {

	currentNode := bt.GetRoot()
	findValue := n.GetData()

	for {
		if currentNode == nil {
			return false
		}

		if findValue == currentNode.GetData() {
			return true
		} else if findValue < currentNode.GetData() {
			currentNode = currentNode.GetLeft()
		} else if findValue > currentNode.GetData() {
			currentNode = currentNode.GetRight()
		}
	}
}

func NewBinaryTree(root *TreeNode) *BinaryTree {
	bt := new(BinaryTree)
	bt.root = root
	return bt
}

func (bt *BinaryTree) Insert(n Node) {
	currentNode := bt.root
	var tmpNode Node
	if currentNode == nil {
		bt.root = n
		return
	}
	for {
		if n.GetData() < currentNode.GetData() {
			tmpNode = currentNode.GetLeft()
			if tmpNode == nil {
				currentNode.SetLeft(n)
				return
			}
		} else if n.GetData() > currentNode.GetData() {
			tmpNode = currentNode.GetRight()
			if tmpNode == nil {
				currentNode.SetRight(n)
				return
			}
		}
		currentNode = tmpNode
	}

}
func main() {
	nod1 := NewNode(20)
	nod2 := NewNode(10)
	nod3 := NewNode(30)
	nod4 := NewNode(8)
	nod5 := NewNode(29)

	nod6 := NewNode(10)
	btree := NewBinaryTree(nod1)
	btree.Insert(nod2)
	btree.Insert(nod3)
	btree.Insert(nod4)
	btree.Insert(nod5)

	fmt.Println(btree.Find(nod6))
}
