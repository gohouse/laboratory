package queue

import "testing"

func TestQueue(t *testing.T) {
	var q = New()
	q.EnQueue(1)
	q.EnQueue(8)
	q.EnQueue(4)
	t.Log(q.Peek())

	q.Show()

	t.Log(q.DeQueue())
	t.Log(q.Peek())
	t.Log(q.Length())

	q.Show()
}
