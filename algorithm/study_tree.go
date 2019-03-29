package main

import "fmt"

func main() {
	var t = new(tree2)
	t.add2(31)
	t.add2(5)
	t.add2(7)
	t.add2(55)
	t.add2(3)
	t.add2(12)
	t.add2(9)
	t.add2(88)
	t.add2(42)
	t.add2(25)
	t.add2(90)
	t.add2(11)
	t.add2(10)
	t.add2(77)
	t.add2(4)
	t.add2(2)

	preSort2(t)
	fmt.Println("\n")
	midSort2(t)
	fmt.Println("\n")
	afterSort2(t)

	t.del2(5)
	fmt.Println("\n")
	afterSort2(t)
}

type tree2 struct {
	data        int
	left, right *tree2
}

// 前序遍历
func preSort2(t *tree2) {
	if t != nil {
		fmt.Printf("%v ", t.data)
		preSort2(t.left)
		preSort2(t.right)
	}
}

// 中序遍历
func midSort2(t *tree2) {
	if t != nil {
		midSort2(t.left)
		fmt.Printf("%v ", t.data)
		midSort2(t.right)
	}
}

// 后序遍历
func afterSort2(t *tree2) {
	if t != nil {
		afterSort2(t.left)
		afterSort2(t.right)
		fmt.Printf("%v ", t.data)
	}
}

func (t *tree2) del2(val int) bool {
	var parent = new(tree2)
	var current = t
	var isLeft bool

	// 查找到要删除的节点的位置, 找不到就返回false
	for current.data != val {
		parent = current
		if current.data > val { // 在左节点
			isLeft = true
			current = parent.left
		} else { // 在右节点
			isLeft = false
			current = parent.right
		}
		if current == nil { // 没找到
			return false
		}
	}

	// 如果该节点没有子节点, 直接删除
	if current.left == nil && current.right == nil {
		if current == t { // 只有一个元素, 刚好就是要删除的元素, 清空树
			t = nil
		} else if isLeft { // 直接删除该节点
			parent.left = nil
		} else { // 直接删除该节点
			parent.right = nil
		}
		return true

	} else if current.left != nil && current.right != nil { // 该节点没有2个子节点, 则使用中序后继节点
		successor := getSuccessor(current)
		if (current == t) {
			t = successor
		} else if (isLeft) {
			parent.left = successor
		} else {
			parent.right = successor
		}
		successor.left = current.left
	} else { // 如果该节点只有一个子节点, 直接用子节点替换
		if current.left == nil { // 只有右子节点
			if current == t {
				t = current.right
			} else if isLeft {
				parent.left = current.right
			} else {
				parent.right = current.right
			}
		} else { // 只有左子节点
			if current == t {
				t = current.left
			} else if isLeft {
				parent.left = current.left
			} else {
				parent.right = current.left
			}
		}
		return true
	}

	return false
}

func getSuccessor(delNode *tree2) *tree2 {
	successorParent := delNode
	successor := delNode
	current := delNode.right
	for current != nil {
		successorParent = successor
		successor = current
		current = current.left
	}
	//后继节点不是删除节点的右子节点，将后继节点替换删除节点
	if (successor != delNode.right) {
		successorParent.left = successor.right
		successor.right = delNode.right
	}

	return successor
}

// 创建二叉树
func (t *tree2) add2(val int) bool {
	var newTree = &tree2{
		data: val,
	}
	if t == nil {
		t = newTree
		return true
	}
	var current = t
	for current != nil {
		if current.data > val { //当前值比插入值大，搜索左子节点
			if current.left == nil {
				current.left = newTree
				return true
			}
		} else {
			if current.right == nil {
				current.right = newTree
				return true
			}
		}
	}

	return false
}
