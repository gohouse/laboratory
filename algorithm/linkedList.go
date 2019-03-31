package main

import "fmt"

type ILinkedList interface {
	Insert(index int, data interface{})
	Get(index int) *LinkedList
	Delete(index int) bool
	IndexOf(value interface{}) int
	Reverse()
}

func main() {
	var lkl = NewLinkedList()
	lkl.Insert(1)
	lkl.Insert(2)
	lkl.Insert(3)
	lkl.Insert(4)
	fmt.Println("head: ", lkl)
	fmt.Println("长度：", lkl.Length())
	fmt.Println("search: ", lkl.Search(3))
}

type LinkNode struct {
	Data interface{}
	Prev,Next *LinkedList
}

type LinkedList struct {
	Head *LinkNode
	Tail *LinkNode
	Size int
}

func (l *LinkedList) Insert(index int, data interface{}) {
	if l== nil {
		//l.Data = data
		//l.Next = &LinkedList{}
		l = &LinkedList{data, &LinkedList{}}
		return true
	} else {
		return l.Next.Insert(data)
	}
}

func (l *LinkedList) Search(data interface{}) *LinkedList {
	if l.Data == nil {
		return nil
	} else if l.Data == data {
		return l
	}
	return l.Next.Search(data)
}

func (l *LinkedList) Delete(data interface{}) bool {
	findNode := l.Search(data)
	if findNode == nil {
		return false
	}
	l.replaceWithNode(findNode, findNode.Next)
	return true
}

func (l *LinkedList) Length() int {
	if l.Data==nil {
		return 0
	}

	return 1+l.Length()
}










func (l *LinkedList) replaceWithNode(dst, source *LinkedList) {
	dst.Data, dst.Next = source.Data, source.Next
	source.Data = nil
	source.Next = nil
}
