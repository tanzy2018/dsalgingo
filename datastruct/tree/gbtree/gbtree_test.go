package gbtree

import (
	"dsalgingo/datastruct/tree"
	"reflect"
	"testing"
)

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
		in0 tree.Comparabler
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
			if got := tt.g.Search(tt.args.in0); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tree.Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_Add(t *testing.T) {
	type args struct {
		in0 tree.Comparabler
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
			tt.g.Add(tt.args.in0)
		})
	}
}

func TestTree_Del(t *testing.T) {
	type args struct {
		in0 tree.Comparabler
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
			tt.g.Del(tt.args.in0)
		})
	}
}

func TestTree_Update(t *testing.T) {
	type args struct {
		in0 tree.Comparabler
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
			tt.g.Update(tt.args.in0)
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

func TestTree_LazyPostOrderTraverse(t *testing.T) {
	tests := []struct {
		name string
		g    *Tree
		want tree.Noder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.LazyPostOrderTraverse(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tree.LazyPostOrderTraverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_HasNextPostOrderTraverse(t *testing.T) {
	tests := []struct {
		name string
		g    *Tree
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.HasNextPostOrderTraverse(); got != tt.want {
				t.Errorf("Tree.HasNextPostOrderTraverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_ResetPostOrderTraverse(t *testing.T) {
	tests := []struct {
		name string
		g    *Tree
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.g.ResetPostOrderTraverse()
		})
	}
}

func TestTree_LazyPreOrderTraverse(t *testing.T) {
	tests := []struct {
		name string
		g    *Tree
		want tree.Noder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.LazyPreOrderTraverse(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tree.LazyPreOrderTraverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_HasNextPreOrderTraverse(t *testing.T) {
	tests := []struct {
		name string
		g    *Tree
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.HasNextPreOrderTraverse(); got != tt.want {
				t.Errorf("Tree.HasNextPreOrderTraverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_ResetPreOrderTraverse(t *testing.T) {
	tests := []struct {
		name string
		g    *Tree
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.g.ResetPreOrderTraverse()
		})
	}
}

func TestTree_LazyInOrderTraverse(t *testing.T) {
	tests := []struct {
		name string
		g    *Tree
		want tree.Noder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.LazyInOrderTraverse(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tree.LazyInOrderTraverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_HasNextInOrderTraverse(t *testing.T) {
	tests := []struct {
		name string
		g    *Tree
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.HasNextInOrderTraverse(); got != tt.want {
				t.Errorf("Tree.HasNextInOrderTraverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_ResetInOrderTraverse(t *testing.T) {
	tests := []struct {
		name string
		g    *Tree
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.g.ResetInOrderTraverse()
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
		name string
		root *Node
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.root.add(tt.args.node); got != tt.want {
				t.Errorf("Node.add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_del(t *testing.T) {
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
			if got := tt.root.del(tt.args.node); got != tt.want {
				t.Errorf("Node.del() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_delLeaf(t *testing.T) {
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
			if got := tt.root.delLeaf(tt.args.node); got != tt.want {
				t.Errorf("Node.delLeaf() = %v, want %v", got, tt.want)
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
