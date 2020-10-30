package main

import (
	"fmt"
)

// Tree 菜单
type Tree struct {
	ID       int
	Pid      int
	Title    string
	Children []Tree
}

// GetTree 获取菜单
func GetTree(menuList []Tree, pid int) []Tree {
	var treeList []Tree
	for _, v := range menuList {
		if v.Pid == pid {
			child := GetTree(menuList, v.ID)
			node := Tree{ID: v.ID, Title: v.Title, Pid: v.Pid}
			node.Children = child
			treeList = append(treeList, node)
		}
	}
	return treeList
}

func printAsTree(tree []Tree, pre string) {
	if len(tree) == 0 {
		return
	}
	for _, v := range tree {
		fmt.Printf("%sID:%v\n%sPid:%v\n%sTitle:%v\n", pre, v.ID, pre, v.Pid, pre, v.Title)
		printAsTree(v.Children, fmt.Sprintf("%s|--", pre))
	}
}

func main() {
	myTree := []Tree{
		{ID: 1, Pid: 0, Title: "父节点1"},
		{ID: 2, Pid: 0, Title: "父节点2"},
		{ID: 3, Pid: 0, Title: "父节点3"},
		{ID: 4, Pid: 2, Title: "子节点2.1"},
		{ID: 5, Pid: 2, Title: "子节点2.2"},
		{ID: 6, Pid: 4, Title: "孙子节点2.1.1"},
		{ID: 7, Pid: 6, Title: "重孙子节点2.1.1.1"},
	}
	printAsTree(GetTree(myTree, 0), "|--")
}
