package main

import "fmt"

type node struct {
	data int
	next *node
	prev *node
}
type doublylinkedList struct {
	head   *node
	length int
	tail   *node
}

func (l *doublylinkedList) endInsert(val int) {
	var n node
	n.data = val
	if l.length == 0 {
		l.head = &n
		l.tail = &n
		l.length++
		return
	}
	ptr := l.tail
	ptr.next = &n
	n.prev = ptr
	l.length++
	l.tail = &n
	return
}
func (l *doublylinkedList) frontInsert(val int) {
	var n node
	n.data = val
	if l.length == 0 {
		l.head = &n
		l.tail = &n
		l.length++
		return
	}
	ptr := l.head
	ptr.prev = &n
	n.next = ptr
	l.head = &n
	l.length++
	return
}
func (l *doublylinkedList) insert(val, pos int) {
	var n node
	n.data = val
	ptr := l.head
	for i := 1; i < l.length; i++ {
		if i == pos {
			ptr.prev.next = &n
			n.prev = ptr.prev
			n.next = ptr
			ptr.prev = &n
			l.length++
			return
		}
		ptr = ptr.next
	}
}
func (l *doublylinkedList) delNode(val int) {
	ptr := l.head
	len := l.length
	if len != 0 {
		for len != 0 {
			if ptr.data == val {
				if ptr.prev == nil {
					l.head = ptr.next
					l.head.prev = nil
					l.length--
					return
				} else if ptr.next == nil {
					ptr1 := ptr.prev
					ptr1.next = nil
					l.tail = ptr1
					l.length--
					return
				} else {
					ptr1 := ptr.next
					ptr1.prev = ptr.prev
					ptr.prev.next = ptr1
					l.length--
					return
				}
			}
			ptr = ptr.next
			len--
		}
	} else {
		fmt.Println("Linkedlist is empty")
		return
	}
	fmt.Println("The element is not in the linked list")
	return
}
func (l doublylinkedList) display() {
	ptr := l.head
	for i := 0; i < l.length; i++ {
		fmt.Print(ptr.data, "->")
		ptr = ptr.next
	}
	fmt.Println("")
	ptr1 := l.tail
	for i := 0; i < l.length; i++ {
		fmt.Print(ptr1.data, "->")
		ptr1 = ptr1.prev
	}
	fmt.Println()
}
func main() {
	var n doublylinkedList
	n.delNode(3)
	n.frontInsert(3)
	n.endInsert(4)
	n.endInsert(5)
	n.frontInsert(1)
	n.frontInsert(2)
	n.endInsert(8)
	n.endInsert(7)
	n.display()
	n.delNode(10)
	n.delNode(3)
	n.delNode(2)
	n.delNode(7)
	n.display()
	n.insert(10, 3)
	n.display()
}
