package sort

import (
	"reflect"
	"testing"
)

func TestQsort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"nil", args{nil}, nil},
		{"1 ele", args{[]int{1}}, []int{1}},
		{"2 sorted ele", args{[]int{1, 2}}, []int{1, 2}},
		{"2 unsorted ele", args{[]int{2, 1}}, []int{1, 2}},
		{"ase sorted ele", args{[]int{1, 2, 3, 4}}, []int{1, 2, 3, 4}},
		{"desc sorted ele", args{[]int{4, 3, 2, 1}}, []int{1, 2, 3, 4}},
		{"unsorted ele", args{[]int{4, 2, 3, 1, 6, 5, 7}}, []int{1, 2, 3, 4, 5, 6, 7}},
		{"unsorted duplicated ele", args{[]int{4, 2, 3, 1, 6, 5, 5, 2, 3, 7}}, []int{1, 2, 2, 3, 3, 4, 5, 5, 6, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Qsort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Qsort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"nil", args{nil}, nil},
		{"1 ele", args{[]int{1}}, []int{1}},
		{"2 sorted ele", args{[]int{1, 2}}, []int{1, 2}},
		{"2 unsorted ele", args{[]int{2, 1}}, []int{1, 2}},
		{"ase sorted ele", args{[]int{1, 2, 3, 4}}, []int{1, 2, 3, 4}},
		{"desc sorted ele", args{[]int{4, 3, 2, 1}}, []int{1, 2, 3, 4}},
		{"unsorted ele", args{[]int{4, 2, 3, 1, 6, 5, 7}}, []int{1, 2, 3, 4, 5, 6, 7}},
		{"unsorted duplicated ele", args{[]int{4, 2, 3, 1, 6, 5, 5, 2, 3, 7}}, []int{1, 2, 2, 3, 3, 4, 5, 5, 6, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeSort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertionSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"nil", args{nil}, nil},
		{"1 ele", args{[]int{1}}, []int{1}},
		{"2 sorted ele", args{[]int{1, 2}}, []int{1, 2}},
		{"2 unsorted ele", args{[]int{2, 1}}, []int{1, 2}},
		{"ase sorted ele", args{[]int{1, 2, 3, 4}}, []int{1, 2, 3, 4}},
		{"desc sorted ele", args{[]int{4, 3, 2, 1}}, []int{1, 2, 3, 4}},
		{"unsorted ele", args{[]int{4, 2, 3, 1, 6, 5, 7}}, []int{1, 2, 3, 4, 5, 6, 7}},
		{"unsorted duplicated ele", args{[]int{4, 2, 3, 1, 6, 5, 5, 2, 3, 7}}, []int{1, 2, 2, 3, 3, 4, 5, 5, 6, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InsertionSort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertionSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCounterSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"nil", args{nil}, nil},
		{"1 ele", args{[]int{1}}, []int{1}},
		{"2 sorted ele", args{[]int{1, 2}}, []int{1, 2}},
		{"2 unsorted ele", args{[]int{2, 1}}, []int{1, 2}},
		{"ase sorted ele", args{[]int{1, 2, 3, 4}}, []int{1, 2, 3, 4}},
		{"desc sorted ele", args{[]int{4, 3, 2, 1}}, []int{1, 2, 3, 4}},
		{"unsorted ele", args{[]int{4, 2, 3, 1, 6, 5, 7}}, []int{1, 2, 3, 4, 5, 6, 7}},
		{"unsorted duplicated ele", args{[]int{4, 2, 3, 1, 6, 5, 5, 2, 3, 7}}, []int{1, 2, 2, 3, 3, 4, 5, 5, 6, 7}},
		{"neg 1 ele", args{[]int{-1}}, []int{-1}},
		{"neg 2 sorted ele", args{[]int{-1, -2}}, []int{-2, -1}},
		{"neg 2 unsorted ele", args{[]int{-2, -1}}, []int{-2, -1}},
		{"neg desc sorted ele", args{[]int{-1, -2, -3, -4}}, []int{-4, -3, -2, -1}},
		{"neg asc sorted ele", args{[]int{-4, -3, -2, -1}}, []int{-4, -3, -2, -1}},
		{"neg unsorted ele", args{[]int{-4, -2, -3, -1, -6, -5, -7}}, []int{-7, -6, -5, -4, -3, -2, -1}},
		{"neg unsorted duplicated ele", args{[]int{-4, -2, -3, -1, -6, -5, -5, -2, -3, -7}}, []int{-7, -6, -5, -5, -4, -3, -3, -2, -2, -1}},
		{"mixed sorted ele", args{[]int{-4, -3, -2, -1, 0, 1, 2, 5, 7}}, []int{-4, -3, -2, -1, 0, 1, 2, 5, 7}},
		{"mixed unsorted ele", args{[]int{4, 2, 3, 1, 6, 5, -5, -2, 3, -7}}, []int{-7, -5, -2, 1, 2, 3, 3, 4, 5, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CounterSort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CountSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
