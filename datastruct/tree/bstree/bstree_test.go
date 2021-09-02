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

func TestNode_Value(t *testing.T) {
	tests := []struct {
		name string
		tr   *Node
		want tree.Comparabler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.Value(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Node.Value() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_Values(t *testing.T) {
	tests := []struct {
		name string
		tr   *Node
		want []tree.Comparabler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.Values(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Node.Values() = %v, want %v", got, tt.want)
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
		g    *Tree
		args args
		want tree.Noder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.Search(tt.args.v); !reflect.DeepEqual(got, tt.want) {
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
		g    *Tree
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.g.Add(tt.args.v)
		})
	}
}

func TestTree_Del(t *testing.T) {
	type args struct {
		v tree.Comparabler
	}
	tests := []struct {
		name string
		g    *Tree
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.g.Del(tt.args.v)
		})
	}
}

func TestTree_Update(t *testing.T) {
	type args struct {
		v tree.Comparabler
	}
	tests := []struct {
		name string
		g    *Tree
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.g.Update(tt.args.v)
		})
	}
}

func TestTree_PostOrderTraverse(t *testing.T) {
	tests := []struct {
		name string
		g    *Tree
		want []tree.Noder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.PostOrderTraverse(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tree.PostOrderTraverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_PreOrderTraverse(t *testing.T) {
	tests := []struct {
		name string
		g    *Tree
		want []tree.Noder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.PreOrderTraverse(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tree.PreOrderTraverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_InOrderTraverse(t *testing.T) {
	tests := []struct {
		name string
		g    *Tree
		want []tree.Noder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.InOrderTraverse(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tree.InOrderTraverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_Count(t *testing.T) {
	tests := []struct {
		name string
		g    *Tree
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.Count(); got != tt.want {
				t.Errorf("Tree.Count() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_Err(t *testing.T) {
	tests := []struct {
		name    string
		g       *Tree
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.g.Err(); (err != nil) != tt.wantErr {
				t.Errorf("Tree.Err() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTree_Type(t *testing.T) {
	tests := []struct {
		name string
		g    *Tree
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.Type(); got != tt.want {
				t.Errorf("Tree.Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_resetErr(t *testing.T) {
	tests := []struct {
		name string
		g    *Tree
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.g.resetErr()
		})
	}
}

func TestTree_resetType(t *testing.T) {
	tests := []struct {
		name string
		g    *Tree
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.g.resetType()
		})
	}
}

func TestTree_isSameType(t *testing.T) {
	type args struct {
		v tree.Comparabler
	}
	tests := []struct {
		name string
		g    *Tree
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.isSameType(tt.args.v); got != tt.want {
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
		g    *Tree
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.g.initType(tt.args.v)
		})
	}
}

func TestNode_search(t *testing.T) {
	type args struct {
		node *Node
	}
	tests := []struct {
		name string
		root *Node
		args args
		want *Node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.root.search(tt.args.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Node.search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_add(t *testing.T) {
	type args struct {
		node *Node
	}
	tests := []struct {
		name    string
		root    *Node
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.root.add(tt.args.node); (err != nil) != tt.wantErr {
				t.Errorf("Node.add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNode_del(t *testing.T) {
	type args struct {
		pre  *Node
		node *Node
	}
	tests := []struct {
		name string
		root *Node
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.root.del(tt.args.pre, tt.args.node); got != tt.want {
				t.Errorf("Node.del() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_delRoot(t *testing.T) {
	tests := []struct {
		name string
		root *Node
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.root.delRoot(); got != tt.want {
				t.Errorf("Node.delRoot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_popMinNode(t *testing.T) {
	type args struct {
		pre *Node
	}
	tests := []struct {
		name string
		root *Node
		args args
		want *Node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.root.popMinNode(tt.args.pre); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Node.popMinNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_update(t *testing.T) {
	type args struct {
		node *Node
	}
	tests := []struct {
		name string
		root *Node
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.root.update(tt.args.node); got != tt.want {
				t.Errorf("Node.update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_postOrderTraverse(t *testing.T) {
	tests := []struct {
		name string
		root *Node
		want []*Node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.root.postOrderTraverse(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Node.postOrderTraverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_preOrderTraverse(t *testing.T) {
	tests := []struct {
		name string
		root *Node
		want []*Node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.root.preOrderTraverse(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Node.preOrderTraverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_inOrderTraverse(t *testing.T) {
	tests := []struct {
		name string
		root *Node
		want []*Node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.root.inOrderTraverse(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Node.inOrderTraverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
