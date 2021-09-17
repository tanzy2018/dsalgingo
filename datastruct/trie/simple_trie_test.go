package trie

import (
	"reflect"
	"testing"
)

func TestSimpleEnglishTrie_Insert(t *testing.T) {
	type args struct {
		v interface{}
	}
	insensitiveTr := NewSimpleEnglishTrie(false)
	sensitiveTr := NewSimpleEnglishTrie(true)
	tests := []struct {
		name string
		s    *SimpleEnglishTrie
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"test-0", insensitiveTr, args{"string"}, true},
		{"test-1", insensitiveTr, args{"String"}, true},
		{"test-2", sensitiveTr, args{"string"}, true},
		{"test-3", sensitiveTr, args{"String"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Insert(tt.args.v); got != tt.want {
				t.Errorf("SimpleEnglishTrie.Insert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSimpleEnglishTrie_Search(t *testing.T) {

	insensitiveTr := NewSimpleEnglishTrie(false)
	sensitiveTr := NewSimpleEnglishTrie(true)
	insensitiveTr1 := NewSimpleEnglishTrie(false)
	sensitiveTr1 := NewSimpleEnglishTrie(true)
	insensitiveTr.Insert("goodbye")
	insensitiveTr.Insert("Goodbye")
	insensitiveTr1.Insert("goodbye")
	insensitiveTr1.Insert("good")

	sensitiveTr.Insert("goodbye")
	sensitiveTr.Insert("Goodbye")
	sensitiveTr1.Insert("goodbye")
	sensitiveTr1.Insert("Good")
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		s    *SimpleEnglishTrie
		args args
		want bool
	}{
		// TODO: Add test cases.
		// 大小写不敏感
		{"test-insensitive-0", insensitiveTr, args{"goodbye"}, true},
		{"test-insensitive-1", insensitiveTr, args{"Goodbye"}, true},
		{"test-insensitive-2", insensitiveTr, args{"goodBye"}, true},
		{"test-insensitive-3", insensitiveTr, args{"good"}, false},
		{"test-insensitive-4", insensitiveTr1, args{"good"}, true},
		// 大小写敏感
		{"test-sensitive-0", sensitiveTr, args{"goodbye"}, true},
		{"test-sensitive-1", sensitiveTr, args{"Goodbye"}, true},
		{"test-sensitive-2", sensitiveTr, args{"goodBye"}, false},
		{"test-sensitive-3", sensitiveTr, args{"good"}, false},
		{"test-sensitive-4", sensitiveTr1, args{"good"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Search(tt.args.v); got != tt.want {
				t.Errorf("SimpleEnglishTrie.Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSimpleEnglishTrie_StartWith(t *testing.T) {
	type args struct {
		v interface{}
	}
	// 大小写敏感
	sensitiveTr := NewSimpleEnglishTrie(true)
	sensitiveTr.Insert("golden")
	sensitiveTr.Insert("good")
	sensitiveTr.Insert("Good")
	sensitiveTr.Insert("goodbye")
	sensitiveTr.Insert("goose")
	sensitiveTr.Insert("grow")

	// 大小写不敏感
	insensitiveTr := NewSimpleEnglishTrie(false)
	insensitiveTr.Insert("golden")
	insensitiveTr.Insert("good")
	insensitiveTr.Insert("Good")
	insensitiveTr.Insert("goodbye")
	insensitiveTr.Insert("goose")
	insensitiveTr.Insert("grow")
	tests := []struct {
		name string
		s    *SimpleEnglishTrie
		args args
		want []interface{}
	}{
		// TODO: Add test cases.
		{"test-sensitive-0", sensitiveTr, args{"go"}, []interface{}{"golden", "good", "Good", "goodbye", "goose"}},
		{"test-sensitive-1", sensitiveTr, args{"Go"}, []interface{}{"golden", "good", "Good", "goodbye", "goose"}},
		{"test-sensitive-2", sensitiveTr, args{"Home"}, nil},
		{"test-insensitive-0", insensitiveTr, args{"go"}, []interface{}{"golden", "good", "goodbye", "goose"}},
		{"test-insensitive-0", insensitiveTr, args{"Go"}, []interface{}{"golden", "good", "goodbye", "goose"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.StartWith(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SimpleEnglishTrie.StartWith() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSimpleEnglishTrie_Len(t *testing.T) {
	insensitiveTr := NewSimpleEnglishTrie(false)
	sensitiveTr := NewSimpleEnglishTrie(true)
	insensitiveTr.Insert("string")
	insensitiveTr.Insert("string")
	insensitiveTr.Insert("String")

	sensitiveTr.Insert("string")
	sensitiveTr.Insert("string")
	sensitiveTr.Insert("String")
	tests := []struct {
		name string
		s    *SimpleEnglishTrie
		want int
	}{
		// TODO: Add test cases.
		{"test-0", insensitiveTr, 1},
		{"test-0", sensitiveTr, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Len(); got != tt.want {
				t.Errorf("SimpleEnglishTrie.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}
