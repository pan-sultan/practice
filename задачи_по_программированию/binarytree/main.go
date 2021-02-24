package main

import (
	"fmt"
)

func main() {
	tree := new(BTree)
	tree.Add(&_Data{Value: 2})
	tree.Add(&_Data{Value: 7})
	tree.Add(&_Data{Value: 5})
	tree.Add(&_Data{Value: 1})
	tree.Add(&_Data{Value: 6})
	tree.Add(&_Data{Value: 5})
	tree.Add(&_Data{Value: 11})
	tree.Add(&_Data{Value: 9})
	tree.Add(&_Data{Value: 4})

	tree.Print()
}

type BTree struct {
	tree *_Tree
}

type _Tree struct {
	Left  *_Tree // smaller
	Right *_Tree // bigger
	Data  *_Data
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

func (b *BTree) MaxDepth() int {
	var deep func(*_Tree, int) int
	deep = func(t *_Tree, level int) int {
		if t == nil {
			return level
		}

		level++
		levelLeft := deep(t.Left, level)
		levelRight := deep(t.Right, level)

		if levelLeft < levelRight {
			return levelRight
		}
		return levelLeft
	}

	return deep(b.tree, 0)
}

func (b *BTree) Print() {
	printChar := func(b byte) {
		fmt.Printf("%c", b)
	}

	printSpaces := func(n int) {
		for ; n != 0; n-- {
			printChar(' ')
		}
	}

	calcBranchesLen := func(depth int) int {
		if depth <= 1 {
			return depth * 2
		}

		nodesCount := powSmart(depth-1, 2)
		return nodesCount*4 + nodesCount
	}

	calcLeavesLen := func(depth int) int {
		if depth <= 1 {
			return depth * 2
		}

		nodesCount := powSmart(depth-1, 2)
		return nodesCount*6 + nodesCount
	}

	maxDepth := b.MaxDepth()
	maxLen := calcLeavesLen(maxDepth)
	var printTree func(int, ...*_Tree)

	printTree = func(level int, nodes ...*_Tree) {
		level++
		allNil := true
		subNodes := make([]*_Tree, len(nodes)*2)
		for i, t := range nodes {
			if t != nil {
				allNil = false
			}

			var left, right *_Tree
			if t != nil {
				left = t.Left
				right = t.Right
			}

			subNodes[i*2] = left
			subNodes[i*2+1] = right
		}

		if level > 3 || allNil {
			return
		}

		if level != 1 {
			spacesCount := (maxLen - calcBranchesLen(level)) / 2
			printSpaces(spacesCount)
			for i := range nodes {
				if (i % 2) == 0 {
					printChar('/')
					if i != 0 {
						printSpaces(3)
					}
				} else {
					printSpaces(2)
					printChar('\\')
				}
			}
			fmt.Println()
		}

		spacesCount := (maxLen - calcLeavesLen(level)) / 2
		printSpaces(spacesCount)

		for i, t := range nodes {
			if i != 0 {
				printSpaces(2)
			}

			if t != nil {
				fmt.Printf("%02d", t.Data.Value)
			} else {
				printSpaces(2)
			}
		}

		fmt.Println()
		printTree(level, subNodes...)
	}

	printTree(0, b.tree)
}

func (d *_Data) Equal(n *_Data) bool {
	return d.Value == n.Value
}

type _Data struct {
	Value int
}

func (d *_Data) More(n *_Data) bool {
	return d.Value > n.Value
}

func (d *_Data) Less(n *_Data) bool {
	return d.Value < n.Value
}

func powSmart(a int, n int) int {
	if n == 0 {
		return 1
	}

	if n == 1 || a == 1 {
		return a
	}

	pows := make([]int, 0, 10)
	res := a

	for ; n > 1; n /= 2 {
		if (n % 2) == 1 {
			pows = append(pows, res)
		}
		res *= res
	}

	for _, p := range pows {
		res *= p
	}

	return res
}
