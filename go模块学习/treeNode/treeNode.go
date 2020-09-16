package main

import (
	"fmt"
)

//定义二叉树的节点
type treeNode struct {
	value int
	left,right *treeNode
}

// 定义接口
type Tree interface {
	Add(val int )
	Print()
}
//创建二叉树
func (tree *treeNode) Add(val int ){
	tree.value = val
}

// 打印树
func (tree *treeNode) Print() {
	if tree != nil {
		fmt.Printf("%d ",tree.value)
	}
}

func main() {
	t := &treeNode{}


	t.Add(10)
	t.left.Add(11)
	t.right.Add(12)

	t.Print()
	//t.left.Print()

}