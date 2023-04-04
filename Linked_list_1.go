package main

import "fmt"

type node struct {
	data int
	next *node
}
type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) insert(val int) {
	var n node
	n.data = val
	if l.length == 0 {
		l.head = &n
		l.length++
		return
	}
	ptr := l.head
	for i := 0; i < l.length; i++ {
		if ptr.next == nil {
			ptr.next = &n
			n.next = nil
			l.length++
			return
		}
		ptr = ptr.next
	}
}
func (l linkedList) display() {
	ptr := l.head
	for i := 0; i < l.length; i++ {
		fmt.Print(ptr.data, "->")
		ptr = ptr.next
	}
	fmt.Println("")
}
func (l *linkedList) sort() {
	if l.head == nil {
		return
	}
	ptr := l.head
	for i := 0; i < l.length-1; i++ {
		ptr1 := ptr.next
		for j := 0; j < l.length-1-i; j++ {
			if ptr1.data < ptr.data {
				temp := ptr1.data
				ptr1.data = ptr.data
				ptr.data = temp
			}
			ptr1 = ptr1.next
		}
		ptr = ptr.next
	}
}
func main() {
	var n linkedList
	n.insert(10)
	n.display()
	n.insert(3)
	n.display()
	n.insert(15)
	n.display()
	n.insert(2)
	n.display()
	n.sort()
	n.display()
}
