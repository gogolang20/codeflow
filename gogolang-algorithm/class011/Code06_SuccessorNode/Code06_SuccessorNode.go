package main

//给你二叉树中的某个节点，返回该节点的后继节点
//给定的结构
//后继节点：中序遍历的顺序，X节点的后面就是后继节点

type Node struct {
	value  int
	left   *Node
	right  *Node
	parent *Node
}

//1 X有右树：找右树上最左节点
//2 X无右树：找 Y 节点左树上的最右节点
//			一直往上找到 parent 是 Y 的左树，结果就是 Y；如果找不到就是没有后继

func getSuccessorNode(node *Node) *Node {
	if node == nil {
		return node
	}
	if node.right != nil {
		return getLeftMost(node.right) //找右树上最左节点
	} else { // 无右子树
		parent := node.parent
		for parent != nil && parent.right == node {
			node = parent
			parent = node.parent
		}
		return parent
	}
}

//找右树上最左节点
func getLeftMost(node *Node) *Node {
	if node == nil {
		return node
	}
	for node.left != nil {
		node = node.left
	}
	return node
}

func main() {

}
