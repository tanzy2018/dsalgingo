package trie

import "strings"

// 简单英语字典树
type simpleEnglishTrieNode struct {
	words    []string
	isWord   bool
	children [26]*simpleEnglishTrieNode
}

type SimpleEnglishTrie struct {
	words           int
	isCaseSensitive bool
	root            *simpleEnglishTrieNode
}

func isAlphabet(s string) bool {
	for i, _ := range s {
		if !((s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z')) {
			return false
		}
	}
	return true
}

func strings2interfaces(arr []string) []interface{} {
	res := make([]interface{}, len(arr))
	for i, _ := range arr {
		res[i] = arr[i]
	}
	return res
}

func preHandleInput(v interface{}) (string, bool) {
	raw, ok := v.(string)
	if !ok {
		return "", false
	}
	raw = strings.Trim(raw, " ")
	if len(raw) == 0 {
		return "", false
	}

	if !isAlphabet(raw) {
		return "", false
	}
	return raw, true
}

func NewSimpleEnglishTrie(isCaseSensitive bool) *SimpleEnglishTrie {
	return &SimpleEnglishTrie{
		isCaseSensitive: isCaseSensitive,
		root:            &simpleEnglishTrieNode{},
	}
}

func (s *SimpleEnglishTrie) Insert(v interface{}) bool {
	raw, ok := preHandleInput(v)
	if !ok {
		return false
	}
	word := []byte(strings.ToLower(raw))
	index := word[0] - 'a'
	if s.root.children[index] == nil {
		s.root.children[index] = &simpleEnglishTrieNode{}
	}
	if s.root.children[index].insert(word, 0, raw, s.isCaseSensitive) {
		s.words++
	}
	return true
}

func (s *SimpleEnglishTrie) Search(v interface{}) bool {
	raw, ok := preHandleInput(v)
	if !ok {
		return false
	}
	word := []byte(strings.ToLower(raw))
	index := word[0] - 'a'
	if s.root.children[index] == nil {
		return false
	}
	return s.root.children[index].search(word, 0, raw, s.isCaseSensitive)
}

func (s *SimpleEnglishTrie) StartWith(v interface{}) []interface{} {
	raw, ok := preHandleInput(v)
	if !ok {
		return nil
	}
	word := []byte(strings.ToLower(raw))
	index := word[0] - 'a'
	if s.root.children[index] == nil {
		return nil
	}
	res := s.root.children[index].startWith(word, 0, raw)
	if len(res) > 0 {
		return strings2interfaces(res)
	}
	return nil
}

func (s *SimpleEnglishTrie) Len() int {
	return s.words
}

func (s *simpleEnglishTrieNode) insert(word []byte, i int, raw string, isCaseSensitive bool) bool {
	if i == len(word)-1 {
		// 大小写不敏感
		if s.isWord && !isCaseSensitive {
			return false
		}
		// 大小写敏感
		if s.isWord {
			for _, v := range s.words {
				if v == raw {
					return false
				}
			}
			s.words = append(s.words, raw)
			return true
		}
		s.words = []string{raw}
		s.isWord = true
		return true
	}
	index := word[i+1] - 'a'
	if s.children[index] == nil {
		s.children[index] = &simpleEnglishTrieNode{}
	}
	return s.children[index].insert(word, i+1, raw, isCaseSensitive)
}

func (s *simpleEnglishTrieNode) search(word []byte, i int, raw string, isCaseSensitive bool) bool {
	if i == len(word)-1 {
		if !s.isWord {
			return false
		}
		if isCaseSensitive {
			for _, v := range s.words {
				if v == raw {
					return true
				}
			}
			return false
		}
		return true
	}
	index := word[i+1] - 'a'
	if s.children[index] == nil {
		return false
	}
	return s.children[index].search(word, i+1, raw, isCaseSensitive)
}

func (s *simpleEnglishTrieNode) startWith(word []byte, i int, raw string) []string {
	if i == len(word)-1 {
		return s.traverseWords()
	}
	index := word[i+1] - 'a'
	if s.children[index] == nil {
		return nil
	}
	return s.children[index].startWith(word, i+1, raw)
}

func (s *simpleEnglishTrieNode) traverseWords() []string {
	res := make([]string, 0, 1)
	if s.isWord {
		res = append(res, s.words...)
	}
	for i, _ := range s.children {
		if s.children[i] != nil {
			res = append(res, s.children[i].traverseWords()...)
		}
	}
	return res
}
