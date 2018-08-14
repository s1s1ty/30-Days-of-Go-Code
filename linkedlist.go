package main

import "fmt"

// Node ...
type Node struct {
	num      int
	nextNode *Node
}

func (l *Node) appendFirst(n int) {
	nw := &Node{
		num:      n,
		nextNode: l.nextNode,
	}
	l.nextNode = nw
}

func (l *Node) appendLast(n int) {
	if l.nextNode == nil {
		l.nextNode = &Node{
			num:      n,
			nextNode: l.nextNode,
		}
	} else {
		for l.nextNode != nil {
			l = l.nextNode
		}
		l.nextNode = &Node{
			num:      n,
			nextNode: l.nextNode,
		}
	}
}

func (l *Node) show() {
	_show(l.nextNode)
}

func _show(l *Node) {
	if l == nil {
		return
	}
	fmt.Println(l.num)
	_show(l.nextNode)
}

func main() {
	var nl Node
	nl.appendLast(7)
	nl.appendFirst(4)
	nl.appendFirst(5)
	nl.appendLast(3)
	nl.show()
}
