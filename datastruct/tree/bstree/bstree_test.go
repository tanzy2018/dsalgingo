package bstree

import (
	"dsalgingo/datastruct/tree"
	"reflect"
	"testing"
)

type Int int

func (a Int) Compare(b tree.Comparabler) tree.CompareResult {
	_b, ok := b.(Int)
	if !ok {
		return tree.CompareResultUndefined
	}
	if a > _b {
		return tree.CompareResultGreater
	}
	if a < _b {
		return tree.CompareResultLess
	}
	return tree.CompareResultEqual
}

type Float float32

func (a Float) Compare(b tree.Comparabler) tree.CompareResult {
	_b, ok := b.(Float)
	if !ok {
		return tree.CompareResultUndefined
	}
	diff := a - _b
	if diff > 0.000001 {
		return tree.CompareResultGreater
	}
	if diff < -0.000001 {
		return tree.CompareResultLess
	}
	return tree.CompareResultEqual
}

func Test_node_Value(t *testing.T) {
	tests := []struct {
		name string
		n    *node
		want tree.Comparabler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.Value(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("node.Value() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_node_Values(t *testing.T) {
	tests := []struct {
		name string
		n    *node
		want []tree.Comparabler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.Values(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("node.Values() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTree(t *testing.T) {
	tests := []struct {
		name string
		want *Tree
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTree(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_Search(t *testing.T) {
	type args struct {
		v tree.Comparabler
	}
	tests := []struct {
		name string
		tr   *Tree
		args args
		want tree.Noder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.Search(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tree.Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_Add(t *testing.T) {
	type args struct {
		v tree.Comparabler
	}
	tests := []struct {
		name string
		tr   *Tree
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.tr.Add(tt.args.v)
		})
	}
}

func TestTree_Del(t *testing.T) {
	type args struct {
		v tree.Comparabler
	}
	tests := []struct {
		name string
		tr   *Tree
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.tr.Del(tt.args.v)
		})
	}
}

func TestTree_Update(t *testing.T) {
	type args struct {
		v tree.Comparabler
	}
	tests := []struct {
		name string
		tr   *Tree
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.tr.Update(tt.args.v)
		})
	}
}

func TestTree_PostOrderTraverse(t *testing.T) {
	tests := []struct {
		name string
		tr   *Tree
		want []tree.Noder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.PostOrderTraverse(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tree.PostOrderTraverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_PreOrderTraverse(t *testing.T) {
	tests := []struct {
		name string
		tr   *Tree
		want []tree.Noder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.PreOrderTraverse(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tree.PreOrderTraverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_InOrderTraverse(t *testing.T) {
	tests := []struct {
		name string
		tr   *Tree
		want []tree.Noder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.InOrderTraverse(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tree.InOrderTraverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_LazyPreOrderTraverse(t *testing.T) {
	tests := []struct {
		name string
		tr   *Tree
		want tree.Noder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.LazyPreOrderTraverse(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tree.LazyPreOrderTraverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_HasNextPreOrderTraverse(t *testing.T) {
	tests := []struct {
		name string
		tr   *Tree
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.HasNextPreOrderTraverse(); got != tt.want {
				t.Errorf("Tree.HasNextPreOrderTraverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_ResetPreOrderTraverse(t *testing.T) {
	tests := []struct {
		name string
		tr   *Tree
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.tr.ResetPreOrderTraverse()
		})
	}
}

func TestTree_LazyInOrderTraverse(t *testing.T) {
	tests := []struct {
		name string
		tr   *Tree
		want tree.Noder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.LazyInOrderTraverse(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tree.LazyInOrderTraverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_HasNextInOrderTraverse(t *testing.T) {
	tests := []struct {
		name string
		tr   *Tree
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.HasNextInOrderTraverse(); got != tt.want {
				t.Errorf("Tree.HasNextInOrderTraverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_ResetInOrderTraverse(t *testing.T) {
	tests := []struct {
		name string
		tr   *Tree
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.tr.ResetInOrderTraverse()
		})
	}
}

func TestTree_LazyPostOrderTraverse(t *testing.T) {
	tests := []struct {
		name string
		tr   *Tree
		want tree.Noder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.LazyPostOrderTraverse(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tree.LazyPostOrderTraverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_HasNextPostOrderTraverse(t *testing.T) {
	tests := []struct {
		name string
		tr   *Tree
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.HasNextPostOrderTraverse(); got != tt.want {
				t.Errorf("Tree.HasNextPostOrderTraverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_ResetPostOrderTraverse(t *testing.T) {
	tests := []struct {
		name string
		tr   *Tree
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.tr.ResetPostOrderTraverse()
		})
	}
}

func TestTree_initlazyPostOrderStack(t *testing.T) {
	type args struct {
		n *node
	}
	tests := []struct {
		name string
		tr   *Tree
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.tr.initlazyPostOrderStack(tt.args.n)
		})
	}
}

func TestTree_Count(t *testing.T) {
	tests := []struct {
		name string
		tr   *Tree
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.Count(); got != tt.want {
				t.Errorf("Tree.Count() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_Err(t *testing.T) {
	tests := []struct {
		name    string
		tr      *Tree
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.tr.Err(); (err != nil) != tt.wantErr {
				t.Errorf("Tree.Err() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTree_Type(t *testing.T) {
	tests := []struct {
		name string
		tr   *Tree
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.Type(); got != tt.want {
				t.Errorf("Tree.Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_resetErr(t *testing.T) {
	tests := []struct {
		name string
		tr   *Tree
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.tr.resetErr()
		})
	}
}

func TestTree_resetType(t *testing.T) {
	tests := []struct {
		name string
		tr   *Tree
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.tr.resetType()
		})
	}
}

func TestTree_isSameType(t *testing.T) {
	type args struct {
		v tree.Comparabler
	}
	tests := []struct {
		name string
		tr   *Tree
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.isSameType(tt.args.v); got != tt.want {
				t.Errorf("Tree.isSameType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_initType(t *testing.T) {
	type args struct {
		v tree.Comparabler
	}
	tests := []struct {
		name string
		tr   *Tree
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.tr.initType(tt.args.v)
		})
	}
}

func Test_node_search(t *testing.T) {
	type args struct {
		node *node
	}
	tests := []struct {
		name string
		root *node
		args args
		want *node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.root.search(tt.args.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("node.search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_node_add(t *testing.T) {
	type args struct {
		node *node
	}
	tests := []struct {
		name    string
		root    *node
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.root.add(tt.args.node); (err != nil) != tt.wantErr {
				t.Errorf("node.add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_node_del(t *testing.T) {
	type args struct {
		pre  *node
		node *node
	}
	tests := []struct {
		name string
		root *node
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.root.del(tt.args.pre, tt.args.node); got != tt.want {
				t.Errorf("node.del() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_node_delRoot(t *testing.T) {
	tests := []struct {
		name string
		root *node
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.root.delRoot(); got != tt.want {
				t.Errorf("node.delRoot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_node_popMinNode(t *testing.T) {
	type args struct {
		pre *node
	}
	tests := []struct {
		name string
		root *node
		args args
		want *node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.root.popMinNode(tt.args.pre); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("node.popMinNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_node_update(t *testing.T) {
	type args struct {
		node *node
	}
	tests := []struct {
		name string
		root *node
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.root.update(tt.args.node); got != tt.want {
				t.Errorf("node.update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_node_postOrderTraverse(t *testing.T) {
	tests := []struct {
		name string
		root *node
		want []*node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.root.postOrderTraverse(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("node.postOrderTraverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_node_preOrderTraverse(t *testing.T) {
	tests := []struct {
		name string
		root *node
		want []*node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.root.preOrderTraverse(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("node.preOrderTraverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_node_inOrderTraverse(t *testing.T) {
	tests := []struct {
		name string
		root *node
		want []*node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.root.inOrderTraverse(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("node.inOrderTraverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
