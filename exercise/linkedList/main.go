package main

import "fmt"

//Node 추가할려면 맨마지막 노드를 알아서 거기를 추가해야한다. 혹은 tail을 가지고있어서 tail에 추가

type Node struct {
	next *Node
	val  int
}

func main() {
	rootNode := &Node{val: 0}
	tail := rootNode

	for i := 1; i < 10; i++ {
		tail = AddNode(tail, i)
	}

	PrintNodes(rootNode)
	rootNode, tail = DelNode(rootNode.next, rootNode, tail)
	PrintNodes(rootNode)
	rootNode, tail = DelNode(rootNode, rootNode, tail)
	PrintNodes(rootNode)
	rootNode, tail = DelNode(tail, rootNode, tail)
	PrintNodes(rootNode)
}

func AddNode(tail *Node, val int) *Node {
	node := &Node{val: val}
	tail.next = node
	return node
}

func DelNode(node *Node, root *Node, tail *Node) (*Node, *Node) {
	if node == root {
		root = root.next
		if root == nil {
			tail = nil
		}
		return root, tail
	}
	prev := root
	for prev.next != node {
		prev = prev.next
	}

	if node == tail {
		prev.next = nil
		tail = prev
	} else {
		prev.next = prev.next.next
	}

	return root, tail
}

func PrintNodes(root *Node) {
	node := root
	for node.next != nil {
		fmt.Printf("%d ->", node.val)
		node = node.next
	}
	fmt.Printf("%d\n", node.val)
}
