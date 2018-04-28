package main

import (
	"fmt"
)

// 二叉树定义
type tree struct {
	left  *tree
	right *tree
	data  int
}

// 各个方法的值
type Treer struct {
	NodeNum int
}

type Binary interface {
	// 创建二叉树
	Create(index int, value []int)
	//1. 求二叉树中的节点个数
	GetNodeNum(pRoot *tree)
	//2. 求二叉树的深度
	//3. 前序遍历，中序遍历，后序遍历
	//4.分层遍历二叉树（按层次从上往下，从左往右）
	//5. 将二叉查找树变为有序的双向链表
	//6. 求二叉树第K层的节点个数
	//7. 求二叉树中叶子节点的个数
	//8. 判断两棵二叉树是否结构相同
	//9. 判断二叉树是不是平衡二叉树
	//10. 求二叉树的镜像
	//11. 求二叉树中两个节点的最低公共祖先节点
	//12. 求二叉树中节点的最大距离
	//13. 由前序遍历序列和中序遍历序列重建二叉树
	//14.判断二叉树是不是完全二叉树
}

// 创建二叉树
func Create(index int, value []int) (T *tree) {
	T = &tree{}
	T.data = value[index-1]
	//fmt.Printf("index %v value %v \n", index, T.data)
	if index < len(value)-1 &&		// 限制数据超限
		2*index+1 <= len(value) {	// 判断是否到了最底层
		T.left = Create(2*index, value)
		T.right = Create(2*index+1, value)
	}
	return T
}
// 前序
func PreSort(treeNode *tree) {
	if treeNode != nil {
		fmt.Printf("%v ", treeNode.data)
		//ch<-treeNode.data
		//list = append(list, treeNode.data)
		if treeNode.left != nil {
			PreSort(treeNode.left)
		}
		if treeNode.right != nil {
			PreSort(treeNode.right)
		}
	} else {
		return
	}
}
// 中序
func MiddleSort(treeNode *tree) {
	if treeNode != nil {
		if treeNode.left != nil {
			MiddleSort(treeNode.left)
		}
		fmt.Printf("%v ", treeNode.data)
		if treeNode.right != nil {
			MiddleSort(treeNode.right)
		}
	} else {
		return
	}
}
// 后序
func AfterSort(treeNode *tree) {
	if treeNode != nil {
		if treeNode.left != nil {
			AfterSort(treeNode.left)
		}
		if treeNode.right != nil {
			AfterSort(treeNode.right)
		}
		fmt.Printf("%v ", treeNode.data)
	} else {
		return
	}
}

//1. 求二叉树中的节点个数
func (t *Treer) GetNodeNum(treeNode *tree) *Treer {
	//fmt.Println(tree.data)
	if treeNode.data == 0 {
		return t
	}
	t.NodeNum += 1
	if treeNode.left != nil {
		t.GetNodeNum(treeNode.left)
	}
	if treeNode.right != nil {
		t.GetNodeNum(treeNode.right)
	}
	return t
}
//2. 求二叉树的深度
//3. 前序遍历，中序遍历，后序遍历
//4.分层遍历二叉树（按层次从上往下，从左往右）
//5. 将二叉查找树变为有序的双向链表
//6. 求二叉树第K层的节点个数
//7. 求二叉树中叶子节点的个数
//8. 判断两棵二叉树是否结构相同
//9. 判断二叉树是不是平衡二叉树
//10. 求二叉树的镜像
//11. 求二叉树中两个节点的最低公共祖先节点
//12. 求二叉树中节点的最大距离
//13. 由前序遍历序列和中序遍历序列重建二叉树
//14.判断二叉树是不是完全二叉树

func main() {
	var t *Treer
	value := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	TreeRoot := Create(1, value)
	//var list []int
	PreSort(TreeRoot)

	//os.Exit(1)
	//AfterSort(TreeRoot)
	num := t.GetNodeNum(TreeRoot).NodeNum
	fmt.Println("\n", num)
	//1. 求二叉树中的节点个数
	//2. 求二叉树的深度
	//3. 前序遍历，中序遍历，后序遍历
	//4.分层遍历二叉树（按层次从上往下，从左往右）
	//5. 将二叉查找树变为有序的双向链表
	//6. 求二叉树第K层的节点个数
	//7. 求二叉树中叶子节点的个数
	//8. 判断两棵二叉树是否结构相同
	//9. 判断二叉树是不是平衡二叉树
	//10. 求二叉树的镜像
	//11. 求二叉树中两个节点的最低公共祖先节点
	//12. 求二叉树中节点的最大距离
	//13. 由前序遍历序列和中序遍历序列重建二叉树
	//14.判断二叉树是不是完全二叉树
}
