package queue

import (
	"reflect"
	"testing"
)

func TestNewQueue(t *testing.T) {
	type args struct {
		cap int
	}
	tests := []struct {
		name string
		args args
		want *Queue
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewQueue(tt.args.cap); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_Enqueue(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		q       *Queue
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.q.Enqueue(tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("Queue.Enqueue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestQueue_Dequeue(t *testing.T) {
	tests := []struct {
		name    string
		q       *Queue
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.q.Dequeue()
			if (err != nil) != tt.wantErr {
				t.Errorf("Queue.Dequeue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Queue.Dequeue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_Clear(t *testing.T) {
	tests := []struct {
		name string
		q    *Queue
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.q.Clear()
		})
	}
}

func TestQueue_Cap(t *testing.T) {
	tests := []struct {
		name string
		q    *Queue
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.q.Cap(); got != tt.want {
				t.Errorf("Queue.Cap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_Len(t *testing.T) {
	tests := []struct {
		name string
		q    *Queue
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.q.Len(); got != tt.want {
				t.Errorf("Queue.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}
