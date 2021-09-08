package arithmetic

import (
	"reflect"
	"testing"
)

// 四则运算

func Test_binaryhalfAdder(t *testing.T) {
	type args struct {
		a bInt
		b bInt
	}
	tests := []struct {
		name         string
		args         args
		wantRes      bInt
		wantOverflow bInt
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, gotOverflow := binaryhalfAdder(tt.args.a, tt.args.b)
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("binaryhalfAdder() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
			if !reflect.DeepEqual(gotOverflow, tt.wantOverflow) {
				t.Errorf("binaryhalfAdder() gotOverflow = %v, want %v", gotOverflow, tt.wantOverflow)
			}
		})
	}
}

func TestBinaryFullAdder8(t *testing.T) {
	type args struct {
		a [8]bInt
		b [8]bInt
	}
	tests := []struct {
		name         string
		args         args
		wantRes      [8]bInt
		wantOverflow bInt
	}{
		// TODO: Add test cases.
		{"0+0", args{a: [8]bInt{}, b: [8]bInt{}}, [8]bInt{}, bIntZero},
		{"0+1", args{a: [8]bInt{}, b: [8]bInt{0, 0, 0, 0, 0, 0, 0, 1}}, [8]bInt{0, 0, 0, 0, 0, 0, 0, 1}, bIntZero},
		{"overflow", args{a: [8]bInt{1, 0, 0, 0, 0, 0, 0, 0}, b: [8]bInt{1, 0, 0, 0, 0, 0, 0, 0}}, [8]bInt{0, 0, 0, 0, 0, 0, 0, 0}, bIntOne},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, gotOverflow := BinaryFullAdder8(tt.args.a, tt.args.b)
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("BinaryFullAdder8() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
			if !reflect.DeepEqual(gotOverflow, tt.wantOverflow) {
				t.Errorf("BinaryFullAdder8() gotOverflow = %v, want %v", gotOverflow, tt.wantOverflow)
			}
		})
	}
}
