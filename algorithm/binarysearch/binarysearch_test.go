package binarysearch

import (
	"testing"
)

func TestBinarySearchEqual(t *testing.T) {
	type args struct {
		arr []int
		num int
	}
	arr := []int{1, 2, 4, 4, 5, 6}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"not found", args{arr, 3}, -1},
		{"less than min", args{arr, 0}, -1},
		{"greater than max", args{arr, 7}, -1},
		{"the first", args{arr, 1}, 0},
		{"the second", args{arr, 2}, 1},
		{"more equal elements", args{arr, 4}, 2},
		{"the last", args{arr, 6}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearchEqual(tt.args.arr, tt.args.num); got != tt.want {
				t.Errorf("BinarySearchEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinarySearchEqualFirst(t *testing.T) {
	type args struct {
		arr []int
		num int
	}
	arr := []int{1, 2, 4, 4, 4, 4, 4, 5, 6}
	tests := []struct {
		name      string
		args      args
		wantIndex int
	}{
		{"not found", args{arr, 3}, -1},
		{"less than min", args{arr, 0}, -1},
		{"greater than max", args{arr, 7}, -1},
		{"the first", args{arr, 1}, 0},
		{"the second", args{arr, 2}, 1},
		{"more equal elements", args{arr, 4}, 2},
		{"the last", args{arr, 6}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotIndex := BinarySearchEqualFirst(tt.args.arr, tt.args.num); gotIndex != tt.wantIndex {
				t.Errorf("BinarySearchEqualFirst() = %v, want %v", gotIndex, tt.wantIndex)
			}
		})
	}
}

func TestBinarySearchEqualLast(t *testing.T) {
	type args struct {
		arr []int
		num int
	}
	arr := []int{1, 2, 4, 4, 4, 4, 4, 5, 6}
	tests := []struct {
		name      string
		args      args
		wantIndex int
	}{
		{"not found", args{arr, 3}, -1},
		{"less than min", args{arr, 0}, -1},
		{"greater than max", args{arr, 7}, -1},
		{"the first", args{arr, 1}, 0},
		{"the second", args{arr, 2}, 1},
		{"more equal elements", args{arr, 4}, 6},
		{"the last", args{arr, 6}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotIndex := BinarySearchEqualLast(tt.args.arr, tt.args.num); gotIndex != tt.wantIndex {
				t.Errorf("BinarySearchEqualLast() = %v, want %v", gotIndex, tt.wantIndex)
			}
		})
	}
}

func TestBinarySearchLessLast(t *testing.T) {
	type args struct {
		arr []int
		num int
	}
	arr := []int{1, 2, 4, 4, 4, 4, 4, 5, 6}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"not found", args{arr, 3}, 1},
		{"less than min", args{arr, 0}, -1},
		{"the first", args{arr, 1}, -1},
		{"the second", args{arr, 2}, 0},
		{"more equal elements", args{arr, 4}, 1},
		{"the last", args{arr, 6}, 7},
		{"greater than max", args{arr, 7}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearchLessLast(tt.args.arr, tt.args.num); got != tt.want {
				t.Errorf("BinarySearchLessLast() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinarySearchGreaterFirst(t *testing.T) {
	type args struct {
		arr []int
		num int
	}
	arr := []int{1, 2, 4, 4, 4, 4, 4, 5, 6}
	tests := []struct {
		name      string
		args      args
		wantIndex int
	}{
		{"not found", args{arr, 3}, 2},
		{"less than min", args{arr, 0}, 0},
		{"the first", args{arr, 1}, 1},
		{"the second", args{arr, 2}, 2},
		{"more equal elements", args{arr, 4}, 7},
		{"the last", args{arr, 6}, -1},
		{"greater than max", args{arr, 7}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotIndex := BinarySearchGreaterFirst(tt.args.arr, tt.args.num); gotIndex != tt.wantIndex {
				t.Errorf("BinarySearchGreaterFirst() = %v, want %v", gotIndex, tt.wantIndex)
			}
		})
	}
}

func TestBinarySearchLess(t *testing.T) {
	type args struct {
		arr []int
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearchLess(tt.args.arr, tt.args.num); got != tt.want {
				t.Errorf("BinarySearchLess() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinarySearchLessLast_0(t *testing.T) {
	type args struct {
		arr []int
		num int
	}
	arr := []int{1, 2, 4, 4, 4, 4, 4, 5, 6}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"not found", args{arr, 3}, 1},
		{"less than min", args{arr, 0}, -1},
		{"the first", args{arr, 1}, -1},
		{"the second", args{arr, 2}, 0},
		{"more equal elements", args{arr, 4}, 1},
		{"the last", args{arr, 6}, 7},
		{"greater than max", args{arr, 7}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearchLessLast_0(tt.args.arr, tt.args.num); got != tt.want {
				t.Errorf("BinarySearchLessLast_0() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinarySearchGreater(t *testing.T) {
	type args struct {
		arr []int
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearchGreater(tt.args.arr, tt.args.num); got != tt.want {
				t.Errorf("BinarySearchGreater() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinarySearchGreaterFirst_0(t *testing.T) {
	type args struct {
		arr []int
		num int
	}
	arr := []int{1, 2, 4, 4, 4, 4, 4, 5, 6}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"not found", args{arr, 3}, 2},
		{"less than min", args{arr, 0}, 0},
		{"the first", args{arr, 1}, 1},
		{"the second", args{arr, 2}, 2},
		{"more equal elements", args{arr, 4}, 7},
		{"the last", args{arr, 6}, -1},
		{"greater than max", args{arr, 7}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearchGreaterFirst_0(tt.args.arr, tt.args.num); got != tt.want {
				t.Errorf("BinarySearchGreaterFirst_0() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinarySearchLessOrEqual(t *testing.T) {
	type args struct {
		arr []int
		num int
	}
	arr := []int{1, 2, 4, 4, 4, 4, 4, 5, 6}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"not found", args{arr, 3}, 1},
		{"less than min", args{arr, 0}, -1},
		{"the first", args{arr, 1}, 0},
		{"the second", args{arr, 2}, 1},
		{"more equal elements", args{arr, 4}, 4},
		{"the last", args{arr, 6}, 4},
		{"greater than max", args{arr, 7}, 4},
		{"test-7", args{arr, 0}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearchLessOrEqual(tt.args.arr, tt.args.num); got != tt.want {
				t.Errorf("BinarySearchLessOrEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinarySearchLessOrEqualLast(t *testing.T) {
	type args struct {
		arr []int
		num int
	}
	arr := []int{1, 2, 4, 4, 4, 4, 4, 5, 6}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"not found", args{arr, 3}, 1},
		{"less than min", args{arr, 0}, -1},
		{"the first", args{arr, 1}, 0},
		{"the second", args{arr, 2}, 1},
		{"more equal elements", args{arr, 4}, 6},
		{"the last", args{arr, 6}, 8},
		{"greater than max", args{arr, 7}, 8},
		{"test-7", args{arr, 0}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearchLessOrEqualLast(tt.args.arr, tt.args.num); got != tt.want {
				t.Errorf("BinarySearchLessOrEqualLast() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinarySearchGreaterOrEqual(t *testing.T) {
	type args struct {
		arr []int
		num int
	}
	arr := []int{1, 2, 4, 4, 4, 4, 4, 5, 6}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"not found", args{arr, 3}, 4},
		{"less than min", args{arr, 0}, 4},
		{"the first", args{arr, 1}, 4},
		{"the second", args{arr, 2}, 4},
		{"more equal elements", args{arr, 4}, 4},
		{"the last", args{arr, 6}, 8},
		{"greater than max", args{arr, 7}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearchGreaterOrEqual(tt.args.arr, tt.args.num); got != tt.want {
				t.Errorf("BinarySearchGreaterOrEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinarySearchGreaterOrEqualFirst(t *testing.T) {
	type args struct {
		arr []int
		num int
	}
	arr := []int{1, 2, 4, 4, 4, 4, 4, 5, 6}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"not found", args{arr, 3}, 2},
		{"less than min", args{arr, 0}, 0},
		{"the first", args{arr, 1}, 0},
		{"the second", args{arr, 2}, 1},
		{"more equal elements", args{arr, 4}, 2},
		{"the last", args{arr, 6}, 8},
		{"greater than max", args{arr, 7}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearchGreaterOrEqualFirst(tt.args.arr, tt.args.num); got != tt.want {
				t.Errorf("BinarySearchGreaterOrEqualFirst() = %v, want %v", got, tt.want)
			}
		})
	}
}
