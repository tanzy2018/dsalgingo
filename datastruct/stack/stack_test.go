package stack

import (
	"reflect"
	"testing"
)

func TestNewStack(t *testing.T) {
	type args struct {
		cap int
	}
	tests := []struct {
		name string
		args args
		want *Stack
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStack(tt.args.cap); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Pop(t *testing.T) {
	tests := []struct {
		name    string
		s       *Stack
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Pop()
			if (err != nil) != tt.wantErr {
				t.Errorf("Stack.Pop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Stack.Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Push(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		s       *Stack
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.Push(tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("Stack.Push() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStack_Clear(t *testing.T) {
	tests := []struct {
		name string
		s    *Stack
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Clear()
		})
	}
}

func TestStack_Len(t *testing.T) {
	tests := []struct {
		name string
		s    *Stack
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Len(); got != tt.want {
				t.Errorf("Stack.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Cap(t *testing.T) {
	tests := []struct {
		name string
		s    *Stack
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Cap(); got != tt.want {
				t.Errorf("Stack.Cap() = %v, want %v", got, tt.want)
			}
		})
	}
}
