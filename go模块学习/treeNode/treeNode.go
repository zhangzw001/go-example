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
func CreateTree(val int) *treeNode {
	return &treeNode{val,nil,nil}
}


// 打印树
func (tree *treeNode) Print() {
	if tree != nil {
		fmt.Printf("%d ",tree.value)
		if tree.left != nil && tree.right != nil {
			tree.left.Print()
			tree.right.Print()
		}
	}
}


func main() {
	var t1 *treeNode
	t1 = CreateTree(1)
	t1.left = CreateTree(2)
	t1.right = CreateTree(3)
	t1.Print()




	//(t.left).Add(11)
	//(t.right).Add(12)

	//t.Print()
	//t.left.Print()

}