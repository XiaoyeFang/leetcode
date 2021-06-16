package binarytree

type TrieNode struct {
	data      byte
	children  [26]*TrieNode // 更改为hash表
	isEndChar bool
}

func NewTrieTree() *TrieNode {
	return &TrieNode{data: '/'}
}

func (t *TrieNode) Insert(text []byte) {
	p := t
	for _, v := range text {
		index := v - 'a'
		if p.children[index] == nil {
			p.children[index] = &TrieNode{data: v}
		}
		p = p.children[index]
	}
	p.isEndChar = true
}

func (t *TrieNode) Search(text []byte) bool {
	p := t
	for _, v := range text {
		index :=v-'a'
		if p.children[index] != nil {
          p =p.children[index]
		} else {
			return false
		}
	}
	return p.isEndChar
}

