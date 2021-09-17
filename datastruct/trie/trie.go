package trie

// TODO::前缀树
type Trier interface {
	// 添加
	Insert(v interface{}) bool
	// 查找
	Search(v interface{}) bool
	// 前缀匹配
	StartWith(v interface{}) []interface{}
}
