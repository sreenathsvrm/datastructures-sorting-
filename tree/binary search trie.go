package main

import "fmt"

var count int

// Nodrepresents the components of a binary searc  B
type Nod struct {
	data  int
	left  *Nod
	right *Nod
}

type Bst struct {
	root *Nod
}

// insert will add a Nodto the  B
func (B *Bst) insert(k int) {

	temp := &Nod{data: k}

	if B.root == nil {
		B.root = temp
		return
	}

	currentNode := B.root

	for {
		if temp.data < currentNode.data {
			if currentNode.left == nil {
				currentNode.left = temp
				return
			}
			currentNode = currentNode.left
		} else {
			if currentNode.right == nil {
				currentNode.right = temp
				return
			}
			currentNode = currentNode.right
		}
	}

}

// search
// and return true if there is a Nodwith that value
func (n *Nod) search(k int) bool {
	count++
	if n == nil {
		return false
	}

	if n.data < k {
		return n.right.search(k)

	} else if n.data > k {
		return n.left.search(k)
	}
	return true
}

func (n *Nod) inorder(nod *Nod) {
	if nod == nil {
		return
	}

	n.inorder(nod.left)

	fmt.Println(nod.data)

	n.inorder(nod.right)
}

func (B *Bst) preorder() {

	Phelper(B.root)
}

func Phelper(node *Nod) {
	if node == nil {
		return
	}
	fmt.Println(node.data)
	Phelper(node.left)
	Phelper(node.right)
}

func (n *Nod) postorder(nod *Nod) {
	if nod == nil {
		return
	}

	n.postorder(nod.left)
	n.postorder(nod.right)
	fmt.Println(nod.data)
}

func minValued(root *Nod) *Nod {
	temp := root
	for nil != temp && temp.left != nil {
		temp = temp.left
	}
	return temp
}

// Delete node from binary search tree
func Delete(root *Nod, val int) *Nod {
	if nil == root {
		return root
	}
	if root.data > val {
		root.left = Delete(root.left, val)
	}
	if root.data < val {
		root.right = Delete(root.right, val)
	}
	if root.data == val {
		if root.left == nil && root.right == nil {
			root = nil
			return root
		}
		if root.left == nil && root.right != nil {
			temp := root.right
			root = nil
			root = temp
			return root
		}
		if root.left != nil && root.right == nil {
			temp := root.left
			root = nil
			root = temp
			return root
		}

		tempNode := minValued(root.right)
		root.data = tempNode.data
		root.right = Delete(root.right, tempNode.data)
	}
	return root
}
func main() {
	B := &Bst{}
	B.insert(50)
	B.insert(300)
	B.insert(54)
	B.insert(130)
	B.insert(70)
	B.insert(440)
	B.insert(43)
	B.insert(320)
	B.insert(75)
	B.insert(80)
	B.insert(260)
	B.insert(460)
	//fmt.Println  B.search(54))
	B.preorder()

}
