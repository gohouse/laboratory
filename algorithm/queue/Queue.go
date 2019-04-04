package queue

import "fmt"

type Node struct {
	data interface{}
	next *Node
}

type Queue struct {
	head *Node
	tail *Node
	size int
}

func New() *Queue {
	return &Queue{}
}
// EnQueue : 入列
func (q *Queue) EnQueue(data interface{}) bool {
	var tmp = &Node{data:data}
	if q.size==0 {
		q.head = tmp
		q.tail = tmp
	} else {
		q.tail.next = tmp
		q.tail = tmp
	}
	q.size++
	tmp = nil
	return true
}
// DeQueue : 出列
func (q *Queue) DeQueue() interface{} {
	tmp := q.head
	q.head = tmp.next
	q.size--
	tmp.next = nil
	return tmp.data
}
// Peek : 获取队列首部第一个数据
func (q *Queue) Peek() interface{} {
	if q.size==0 {
		return nil
	}
	return q.head.data
}
// IsEmpty
func (q *Queue) IsEmpty() bool {
	return q.size==0
}
// Show 打印队列
func (q *Queue) Show() {
	current := q.head
	fmt.Printf("%s: ","show")
	for current!=nil {
		fmt.Printf("%d ", current.data)
		current=current.next
	}
	fmt.Println("")
}
// Length 队列长度
func (q *Queue) Length() int {
	return q.size
}
