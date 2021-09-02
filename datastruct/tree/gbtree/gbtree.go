package gbtree

// 普通二叉树

import (
	"dsalgingo/datastruct/tree"
	"math/rand"
	"time"
)

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

type Node struct {
	val   tree.Comparabler
	vals  []tree.Comparabler
	left  *Node
	right *Node
}

func (tr *Node) Value() tree.Comparabler {
	return tr.val
}

func (tr *Node) Values() []tree.Comparabler {
	return tr.vals
}

type Tree struct {
	count int
	root  *Node
}

func NewTree() *Tree {
	return &Tree{
		count: 0,
	}
}

func (g *Tree) Search(_ tree.Comparabler) tree.Noder {
	panic("not implemented") // TODO: Implement
}

func (g *Tree) Add(_ tree.Comparabler) {
	panic("not implemented") // TODO: Implement
}

func (g *Tree) Del(_ tree.Comparabler) {
	panic("not implemented") // TODO: Implement
}

func (g *Tree) Update(_ tree.Comparabler) {
	panic("not implemented") // TODO: Implement
}

func (g *Tree) PostOrderTraverse() []tree.Noder {
	panic("not implemented") // TODO: Implement
}

func (g *Tree) PreOrderTraverse() []tree.Noder {
	panic("not implemented") // TODO: Implement
}

func (g *Tree) InOrderTraverse() []tree.Noder {
	panic("not implemented") // TODO: Implement
}

func (g *Tree) LazyPostOrderTraverse() tree.Noder {
	panic("not implemented") // TODO: Implement
}

func (g *Tree) HasNextPostOrderTraverse() bool {
	panic("not implemented") // TODO: Implement
}

func (g *Tree) ResetPostOrderTraverse() {
	panic("not implemented") // TODO: Implement
}

func (g *Tree) LazyPreOrderTraverse() tree.Noder {
	panic("not implemented") // TODO: Implement
}

func (g *Tree) HasNextPreOrderTraverse() bool {
	panic("not implemented") // TODO: Implement
}

func (g *Tree) ResetPreOrderTraverse() {
	panic("not implemented") // TODO: Implement
}

func (g *Tree) LazyInOrderTraverse() tree.Noder {
	panic("not implemented") // TODO: Implement
}

func (g *Tree) HasNextInOrderTraverse() bool {
	panic("not implemented") // TODO: Implement
}

func (g *Tree) ResetInOrderTraverse() {
	panic("not implemented") // TODO: Implement
}

func (g *Tree) Count() int {
	panic("not implemented") // TODO: Implement
}

func (g *Tree) Err() error {
	panic("not implemented") // TODO: Implement
}

func (g *Tree) Type() string {
	panic("not implemented") // TODO: Implement
}

func (root *Node) search(node *Node) *Node {
	if node == nil || root == nil {
		return nil
	}
	if root.val == node.val {
		return root
	}
	if root.left != nil {
		if res := root.left.search(node); res != nil {
			return res
		}
	}
	if root.right != nil {
		if res := root.right.search(node); res != nil {
			return res
		}
	}
	return nil
}

func (root *Node) add(node *Node) bool {
	if node == nil {
		return false
	}
	if root == nil {
		root = node
		return true
	}

	if root.left == nil {
		root.left = node
		return true
	}

	if root.right == nil {
		root.right = node
		return true
	}

	// 简单点 通过奇偶来确定放哪一边
	if rng.Intn(1024)%2 == 0 {
		return root.left.add(node)
	} else {
		return root.right.add(node)
	}
}

func (root *Node) del(node *Node) bool {
	if root == nil || node == nil {
		return false
	}
	tNode := root.search(node)
	if tNode == nil {
		return false
	}
	if tNode.left == nil && tNode.right == nil {
		root.delLeaf(tNode)
		return true
	}

	if tNode.left == nil && tNode.right != nil {
		tNode.val = tNode.right.val
		tNode.left = tNode.right.left
		tNode.right = tNode.right.right
		return true
	}

	if tNode.left != nil && tNode.right == nil {
		tNode.val = tNode.left.val
		tNode.left = tNode.left.left
		tNode.right = tNode.left.right
		return true
	}

	// 简单点 通过奇偶来确定放哪一边
	if rng.Intn(1024)%2 == 0 {
		left := tNode.left
		tNode.val = left.val
		tNode.left = left.left
		if tNode.left == nil {
			tNode.left = tNode.right
			tNode.right = nil
			return true
		}
		if left.right != nil {
			right := tNode.right
			tNode.right = left.right
			tNode.add(right)
			return true
		}
	} else {
		right := tNode.right
		tNode.val = right.val
		tNode.right = right.right
		if tNode.right == nil {
			tNode.right = tNode.left
			tNode.left = nil
			return true
		}
		if right.left != nil {
			left := tNode.left
			tNode.left = right.left
			tNode.add(left)
			return true
		}
	}
	return false
}

func (root *Node) delLeaf(node *Node) bool {
	if root == nil || node == nil || !(node.left == nil && node.right == nil) {
		return false
	}

	if root.val == node.val {
		return true
	}
	if root.left != nil {
		if res := root.left.delLeaf(node); res {
			root.left = nil
			return true
		}
	}
	if root.right != nil {
		if res := root.right.delLeaf(node); res {
			root.right = nil
			return true
		}
	}
	return false
}

func (root *Node) update(node *Node) bool {
	if node == nil || root == nil {
		return false
	}
	if res := root.search(node); res != nil {
		res.val = node.val
		return true
	}
	return false
}

func (root *Node) postOrderTraverse() []*Node {
	if root == nil {
		return nil
	}
	result := []*Node{}
	if root.left != nil {
		result = append(result, root.left.preOrderTraverse()...)
	}
	if root.left != nil {
		result = append(result, root.right.preOrderTraverse()...)
	}
	result = append(result, root)
	return result
}

func (root *Node) preOrderTraverse() []*Node {
	if root == nil {
		return nil
	}
	result := []*Node{root}
	if root.left != nil {
		result = append(result, root.left.preOrderTraverse()...)
	}
	if root.left != nil {
		result = append(result, root.right.preOrderTraverse()...)
	}
	return result
}

func (root *Node) inOrderTraverse() []*Node {
	if root == nil {
		return nil
	}
	result := []*Node{}
	if root.left != nil {
		result = append(result, root.left.preOrderTraverse()...)
	}
	result = append(result, root)
	if root.left != nil {
		result = append(result, root.right.preOrderTraverse()...)
	}
	return result
}
