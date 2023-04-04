package main

import "fmt"

type node struct {
	data int
	next *node
}
type linkedlist struct {
	head   *node
	length int
}

func (l *linkedlist) insert(val int) {
	if l.length == 0 {
		var n node
		n.data = val
		l.head = &n
		l.length++
		n.next = nil
		return
	} else {
		ptr := l.head
		for i := 1; i <= l.length; i++ {
			if ptr.data == val {
				return
			}
			if ptr.next != nil {
				ptr = ptr.next
			}
		}
		var n node
		n.data = val
		ptr.next = &n
		n.next = nil
		l.length++
		return
	}
}
func (l linkedlist) display() {
	ptr := l.head
	for i := 0; i < l.length; i++ {
		fmt.Println(ptr.data)
		ptr = ptr.next
	}
}
func main() {
	var arr = []int{1, 1, 2, 3, 2, 4, 1}
	var n linkedlist
	for _, v := range arr {
		n.insert(v)
	}
	n.display()
}
