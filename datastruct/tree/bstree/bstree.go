package bstree

// 普通二叉搜索树

import (
	"dsalgingo/datastruct/tree"
	"reflect"
)

type Node struct {
	val   tree.Comparabler
	vals  []tree.Comparabler
	left  *Node
	right *Node
}

func (n *Node) Value() tree.Comparabler {
	return n.val
}

func (n *Node) Values() []tree.Comparabler {
	return n.vals
}

type Tree struct {
	count               int
	root                *Node
	tranFunc            func([]*Node) []tree.Noder
	_type               string
	_lazyPostOrderStack []*Node
	_lazyPreOrderStack  []*Node
	_lazyInOrderStack   []*Node
	err                 error
}

func NewTree() *Tree {
	return &Tree{
		count: 0,
		tranFunc: func(nodes []*Node) []tree.Noder {
			if len(nodes) == 0 {
				return nil
			}
			result := make([]tree.Noder, 0, len(nodes))
			for i := 0; i < len(nodes); i++ {
				result = append(result, nodes[i])
			}
			return result
		},
	}
}

func (tr *Tree) Search(v tree.Comparabler) tree.Noder {
	tr.resetErr()
	if tr.count == 0 {
		tr.err = tree.ErrEmptyTree
		return nil
	}
	if !tr.isSameType(v) {
		tr.err = tree.ErrTypeConflict
		return nil
	}
	res := tr.root.search(&Node{val: v})
	if res == nil {
		tr.err = tree.ErrTargetNotFound
	}
	return res
}
func (tr *Tree) Add(v tree.Comparabler) {
	tr.resetErr()
	if tr._type == "" {
		tr.initType(v)
	}
	if !tr.isSameType(v) {
		tr.err = tree.ErrTypeConflict
		return
	}
	if tr.count == 0 {
		tr.root = &Node{val: v, vals: []tree.Comparabler{v}}
		tr.count++
		return
	}
	if err := tr.root.add(&Node{val: v, vals: []tree.Comparabler{v}}); err == nil {
		tr.count++
		return
	} else {
		tr.err = err
	}
	return
}
func (tr *Tree) Del(v tree.Comparabler) {
	tr.resetErr()
	if v == nil {
		tr.err = tree.ErrIllegalParams
		return
	}
	if tr.count == 0 {
		tr.err = tree.ErrEmptyTree
		return
	}
	if !tr.isSameType(v) {
		tr.err = tree.ErrTypeConflict
		return
	}
	cmp := tr.root.val.Compare(v)
	if tr.count == 1 && cmp == tree.CompareResultEqual {
		tr.count--
		return
	}
	if tr.root.del(nil, &Node{val: v}) {
		tr.count--
		if tr.count == 0 {
			tr.root = nil
		}
		return
	}
	tr.err = tree.ErrTargetNotFound
	return
}
func (tr *Tree) Update(v tree.Comparabler) {
	tr.resetErr()
	if tr.count == 0 {
		tr.err = tree.ErrEmptyTree
		return
	}
	if !tr.isSameType(v) {
		tr.err = tree.ErrTypeConflict
		return
	}
	if tr.root.update(&Node{val: v}) {
		return
	}
	tr.err = tree.ErrTargetNotFound
	return
}

func (tr *Tree) PostOrderTraverse() []tree.Noder {
	if tr.count == 0 {
		return nil
	}
	return tr.tranFunc(tr.root.postOrderTraverse())
}
func (tr *Tree) PreOrderTraverse() []tree.Noder {
	if tr.count == 0 {
		return nil
	}
	return tr.tranFunc(tr.root.preOrderTraverse())
}
func (tr *Tree) InOrderTraverse() []tree.Noder {
	if tr.count == 0 {
		return nil
	}
	return tr.tranFunc(tr.root.inOrderTraverse())
}

func (tr *Tree) LazyPreOrderTraverse() tree.Noder {
	_len := len(tr._lazyPreOrderStack)
	if _len == 0 {
		return nil
	}
	node := tr._lazyPreOrderStack[_len-1]
	tr._lazyPreOrderStack = tr._lazyPreOrderStack[:_len-1]
	if node.right != nil {
		tr._lazyPreOrderStack = append(tr._lazyPreOrderStack, node.right)
	}
	if node.left != nil {
		tr._lazyPreOrderStack = append(tr._lazyPreOrderStack, node.left)
	}
	return node
}

func (tr *Tree) HasNextPreOrderTraverse() bool {
	if len(tr._lazyPreOrderStack) > 0 {
		return true
	}
	return false
}

func (tr *Tree) ResetPreOrderTraverse() {
	tr._lazyPreOrderStack = make([]*Node, 0, tr.count+1)
	if tr.root != nil {
		tr._lazyPreOrderStack = append(tr._lazyPreOrderStack, tr.root)
	}
}

func (tr *Tree) LazyInOrderTraverse() tree.Noder {
	return nil
}

func (tr *Tree) HasNextInOrderTraverse() bool {
	if len(tr._lazyInOrderStack) > 0 {
		return true
	}
	return false
}

func (tr *Tree) ResetInOrderTraverse() {
	// TODO:: 这里需要补充
	tr._lazyInOrderStack = make([]*Node, 0, tr.count+1)
	if tr.root != nil {
		tr._lazyInOrderStack = append(tr._lazyInOrderStack, tr.root)
	}
}

func (tr *Tree) LazyPostOrderTraverse() tree.Noder {
	// TODO:: 这里需要补充
	return nil
}
func (tr *Tree) HasNextPostOrderTraverse() bool {
	if len(tr._lazyPostOrderStack) > 0 {
		return true
	}
	return false
}
func (tr *Tree) ResetPostOrderTraverse() {
	// TODO:: 这里需要补充
	tr._lazyPostOrderStack = make([]*Node, 0, tr.count+1)
	if tr.root != nil {
		tr._lazyPostOrderStack = append(tr._lazyPostOrderStack, tr.root)
	}
}

func (tr *Tree) initlazyPostOrderStack(node *Node) {
	// TODO:: 这里需要补充
	tr._lazyPostOrderStack = make([]*Node, 0, tr.count+1)
	if tr.root != nil {
		tr._lazyPostOrderStack = append(tr._lazyPostOrderStack, tr.root)
	}
}

func (tr *Tree) Count() int {
	return tr.count
}

func (tr *Tree) Err() error {
	return tr.err
}

func (tr *Tree) Type() string {
	return tr._type
}

func (tr *Tree) resetErr() {
	tr.err = nil
}

func (tr *Tree) resetType() {
	tr._type = ""
}

func (tr *Tree) isSameType(v tree.Comparabler) bool {
	if tr._type == "" {
		return false
	}
	_type := reflect.TypeOf(v)
	if tr._type != _type.PkgPath()+"."+_type.Name() {
		return false
	}
	return true
}

func (tr *Tree) initType(v tree.Comparabler) {
	_type := reflect.TypeOf(v)
	tr._type = _type.PkgPath() + "." + _type.Name()
}

func (root *Node) search(node *Node) *Node {
	if node == nil || root == nil {
		return nil
	}
	cmp := node.val.Compare(root.val)
	if cmp == tree.CompareResultEqual {
		return root
	}
	if root.left != nil && cmp == tree.CompareResultLess {
		if res := root.left.search(node); res != nil {
			return res
		}
	}
	if root.right != nil && cmp == tree.CompareResultGreater {
		if res := root.right.search(node); res != nil {
			return res
		}
	}
	return nil
}

func (root *Node) add(node *Node) error {
	if node == nil {
		return tree.ErrIllegalParams
	}
	if root == nil {
		root = node
		return nil
	}

	cmp := node.val.Compare(root.val)
	if cmp == tree.CompareResultEqual {
		if len(root.vals) == 0 {
			root.vals = make([]tree.Comparabler, 0, 1)
		}
		root.vals = append(root.vals, node.val)
		return nil
	}

	if cmp == tree.CompareResultLess {
		if root.left == nil {
			root.left = node
			return nil
		}
		return root.left.add(node)
	}

	if cmp == tree.CompareResultGreater {
		if root.right == nil {
			root.right = node
			return nil
		}
		return root.right.add(node)
	}
	return tree.ErrUndefinedCompareResult
}

func (root *Node) del(pre, node *Node) bool {
	if root == nil || node == nil {
		return false
	}
	cmp := node.val.Compare(root.val)
	if cmp == tree.CompareResultEqual {
		// 如果节点的元素不止一个,则摘除一个元素直接返回
		if len(root.vals) > 1 {
			root.vals = root.vals[:len(root.vals)-1]
			return true
		}
		// 如果是根节点
		if pre == nil {
			return root.delRoot()
		}
		// 该节点是叶子节点
		if root.left == nil && root.right == nil {
			if pre.left != nil && pre.left.val.Compare(pre.left.val) == 0 {
				pre.left = nil
			} else {
				pre.right = nil
			}
			return true
		}
		// 该节点有左右两个子节点
		if root.left != nil && root.right != nil {
			if root.right.left == nil {
				root.right.left = root.left
				if pre.left != nil && pre.left.val.Compare(node.val) == 0 {
					pre.left = root.right
				} else {
					pre.right = root.right
				}
				return true
			}
			minNode := root.right.left.popMinNode(root.right)
			minNode.left = root.left
			minNode.right = root.right
			if pre.left != nil && pre.left.val.Compare(node.val) == 0 {
				pre.left = minNode
			} else {
				pre.right = minNode
			}
			return true
		}
		// 只有左节点
		if root.left != nil {
			root.val = root.left.val
			root.vals = root.left.vals
			root.left = root.left.left
			root.right = root.left.right
			return true
		}
		// 只有右节点
		if root.right != nil {
			root.val = root.right.val
			root.vals = root.right.vals
			root.left = root.right.left
			root.right = root.right.right
			return true
		}
	}
	if cmp == tree.CompareResultLess && root.left != nil {
		return root.left.del(root, node)
	}
	if cmp == tree.CompareResultGreater && root.right != nil {
		return root.right.del(root, node)
	}

	return false
}

func (root *Node) delRoot() bool {
	// 没有子树
	if root.left == nil && root.right == nil {
		root = nil
		return true
	}
	// 有左右子树
	if root.left != nil && root.right != nil {
		minNode := root.right.popMinNode(root)
		root.val = minNode.val
		root.vals = minNode.vals
		return true
	}
	// 只有左子树
	if root.left != nil {
		root.val = root.left.val
		root.vals = root.left.vals
		root.left = root.left.left
		root.right = root.left.right
		return true
	}
	// 只有右子树
	if root.right != nil {
		root.val = root.right.val
		root.vals = root.right.vals
		root.left = root.right.left
		root.right = root.right.right
		return true
	}
	return false
}

func (root *Node) popMinNode(pre *Node) *Node {
	if root == nil || pre == nil {
		return nil
	}
	if root.left != nil {
		return root.left.popMinNode(root)
	}
	if root.right == nil {
		pre.left = nil
	}

	if root.right != nil {
		pre.left = root.right
	}
	return root
}

func (root *Node) update(node *Node) bool {
	// 目前看起来没有什么用的更新
	if node == nil || root == nil {
		return false
	}
	if res := root.search(node); res != nil {
		if res.val.Compare(node.val) != tree.CompareResultEqual {
			res.val = node.val
			res.vals = []tree.Comparabler{node.val}
			return true
		}
	}
	return false
}

func (root *Node) postOrderTraverse() []*Node {
	if root == nil {
		return nil
	}
	result := []*Node{}
	if root.left != nil {
		result = append(result, root.left.postOrderTraverse()...)
	}
	if root.right != nil {
		result = append(result, root.right.postOrderTraverse()...)
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
	if root.right != nil {
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
		result = append(result, root.left.inOrderTraverse()...)
	}
	result = append(result, root)
	if root.right != nil {
		result = append(result, root.right.inOrderTraverse()...)
	}
	return result
}
