package main

func main() {
}

type BTree struct {
	tree *_Tree
}

type _Tree struct {
	Left  *_Tree // smaller
	Right *_Tree // bigger
	Data  *_Data
}

type _Data struct {
	N int
}

func (b *BTree) Add(d *_Data) {
	ntree := &_Tree{Data: d}
	if b.tree == nil {
		b.tree = ntree
		return
	}

	for tree := b.tree; ; {
		if tree.Data.Equal(d) {
			return
		}

		if tree.Data.More(d) {
			if tree.Right == nil {
				tree.Right = ntree
				return
			}
			tree = tree.Right
		} else if tree.Data.Less(d) {
			if tree.Left == nil {
				tree.Left = ntree
				return
			}
			tree = tree.Left
		}
	}
}

func (b *BTree) Exists(d *_Data) bool {
	for tree := b.tree; tree != nil; {
		if tree.Data.Equal(d) {
			return true
		}

		if tree.Data.More(d) {
			tree = tree.Right
		} else if tree.Data.Less(d) {
			tree = tree.Left
		}
	}

	return false
}

func (d *_Data) Equal(n *_Data) bool {
	return d.N == n.N
}

func (d *_Data) More(n *_Data) bool {
	return d.N > n.N
}

func (d *_Data) Less(n *_Data) bool {
	return d.N < n.N
}
