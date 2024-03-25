package main

type Node struct {
	key  string
	next *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) insert(k string) {
	newNode := &Node{key: k}
	newNode.next = l.head
	l.head = newNode
}

func (l *LinkedList) search(k string) bool {
	curr := l.head
	for curr != nil {
		if curr.key == k {
			return true
		}
		curr = curr.next
	}
	return false
}

func (l *LinkedList) delete(k string) {
	if l.head.key == k {
		l.head = l.head.next
		return
	}

	curr := l.head
	for curr.next != nil {
		if curr.next.key == k {
			curr.next = curr.next.next
			return
		}
		curr = curr.next
	}
}
