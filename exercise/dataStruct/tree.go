package dataStruct

import "fmt"

type TreeNode struct {
	Val    int
	Childs []*TreeNode
}

type Tree struct {
	Root *TreeNode
}

func (t *TreeNode) AddNode(val int) {
	t.Childs = append(t.Childs, &TreeNode{Val: val})
}

func (t *Tree) AddNode(val int) {
	if t.Root == nil {
		t.Root = &TreeNode{Val: val}
	} else {
		t.Root.Childs = append(t.Root.Childs, &TreeNode{Val: val})
	}
}

func (t *Tree) DFS_Recursion() {
	DFS_Recursion(t.Root)
}

func DFS_Recursion(node *TreeNode) {
	fmt.Printf("%d->", node.Val)
	for i := 0; i < len(node.Childs); i++ {
		DFS_Recursion(node.Childs[i])
	}
}

func (t *Tree) DFS_Stack() {
	s := []*TreeNode{}
	s = append(s, t.Root)

	for len(s) > 0 {
		var last *TreeNode
		last, s = s[len(s)-1], s[:len(s)-1]

		fmt.Printf("%d->", last.Val)

		for i := 0; i < len(last.Childs); i++ {
			s = append(s, last.Childs[i])
		}
	}

}
