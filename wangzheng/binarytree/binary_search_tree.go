package binarytree

type BinarySearchTree struct {
	*tree
	complete func(v interface{}, node *Node) int
}

func NewBinarySearchTree(v interface{}, complete func(v interface{}, node *Node) int) *BinarySearchTree {
	return &BinarySearchTree{&tree{
		&Node{val: v},
	}, complete}
}

// 二叉查找树中查找节点
func (b *BinarySearchTree) Find(v interface{}) *Node {
	p := b.root
	for p != nil {
		result := b.complete(v, p)
		if result == 0 && p.delete == false {
			return p
		} else if result == 1 {
			p = p.left
		} else if result == 2 {
			p = p.right
		}
	}
	return nil
}

func (b *BinarySearchTree) Insert(val interface{}) {
	p := b.root
	for p != nil {
		result := b.complete(val, p)
		if result == 1 {
			if p.left == nil  {
				p.left = &Node{val: val}
				return
			}
			p = p.left
		} else {
			if p.right == nil {
				p.right = &Node{val: val}
				return
			}
			p = p.right
		}
	}
}

func (b *BinarySearchTree) Delete(val interface{}) {

	p := b.Find(val)
	if p == nil {
		return
	}
	p.delete = true
}
