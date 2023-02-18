package main

/*
SB树（size-balance-tree）

1）让每一个叔叔节点为头的数，节点个数都不少于其任何一个侄子节点
2）也是从底层被影响节点开始向上做路径每个节点检查
3）与AVL树非常像，也是四种违规类型：LL、RR、LR、RL
4）与AVL树非常像，核心点是：
LL（做一次右旋）、RR（做一次左旋）
LR和RL（利用旋转让底层那个上到顶部）
5）与AVL树不同的是，每轮经过调整后，谁的孩子发生变化了，谁就再查
*/

/*
SB树在使用时候的改进

1）删除时候可以不用检查

2）就把平衡性的调整放在插入的时候

3）因为这种只要变就递归的特性，别的树没有

4）可以在节点上封装别的数据项，来增加功能
*/

type SBTNode struct {
	key   int // Comparable
	value int
	l     *SBTNode
	r     *SBTNode
	size  int // 不同的key的数量
}

func NewSBTNode(key, value int) *SBTNode {
	return &SBTNode{
		key:   key,
		value: value,
		size:  1,
	}
}

func (sbt *SBTNode) rightRotate(cur *SBTNode) *SBTNode {
	leftNode := cur.l
	cur.l = leftNode.r
	leftNode.r = cur
	leftNode.size = cur.size
	//cur.size = (cur.l != nil ? cur.l.size : 0) + (cur.r != nil ? cur.r.size : 0) + 1
	return leftNode
}

func (sbt *SBTNode) leftRotate(cur *SBTNode) *SBTNode {
	rightNode := cur.r
	cur.r = rightNode.l
	rightNode.l = cur
	rightNode.size = cur.size
	//cur.size = (cur.l != nil ? cur.l.size : 0) + (cur.r != nil ? cur.r.size : 0) + 1
	return rightNode
}

func (sbt *SBTNode) maintain(cur *SBTNode) *SBTNode {
	if cur == nil {
		return nil
	}
	leftSize := 0
	if cur.l != nil {
		leftSize = cur.l.size
	}
	leftLeftSize := 0
	if cur.l != nil && cur.l.l != nil {
		leftLeftSize = cur.l.l.size
	}
	leftRightSize := 0
	if cur.l != nil && cur.l.r != nil {
		leftRightSize = cur.l.r.size
	}
	rightSize := 0
	if cur.r != nil {
		rightSize = cur.r.size
	}
	rightLeftSize := 0
	if cur.r != nil && cur.r.l != nil {
		rightLeftSize = cur.r.l.size
	}
	rightRightSize := 0
	if cur.r != nil && cur.r.r != nil {
		rightRightSize = cur.r.r.size
	}
	if leftLeftSize > rightSize {
		cur = sbt.rightRotate(cur)
		cur.r = sbt.maintain(cur.r)
		cur = sbt.maintain(cur)
	} else if leftRightSize > rightSize {
		cur.l = sbt.leftRotate(cur.l)
		cur = sbt.rightRotate(cur)
		cur.l = sbt.maintain(cur.l)
		cur.r = sbt.maintain(cur.r)
		cur = sbt.maintain(cur)
	} else if rightRightSize > leftSize {
		cur = sbt.leftRotate(cur)
		cur.l = sbt.maintain(cur.l)
		cur = sbt.maintain(cur)
	} else if rightLeftSize > leftSize {
		cur.r = sbt.rightRotate(cur.r)
		cur = sbt.leftRotate(cur)
		cur.l = sbt.maintain(cur.l)
		cur.r = sbt.maintain(cur.r)
		cur = sbt.maintain(cur)
	}
	return cur
}

// 现在，以cur为头的树上，新增，加(key, value)这样的记录
// 加完之后，会对cur做检查，该调整调整
// 返回，调整完之后，整棵树的新头部
func (sbt *SBTNode) add(cur *SBTNode, key, value int) *SBTNode {
	if cur == nil {
		return NewSBTNode(key, value)
	} else {
		cur.size++
		if cur.key < key {
			cur.l = sbt.add(cur.l, key, value)
		} else {
			cur.r = sbt.add(cur.r, key, value)
		}
		return sbt.maintain(cur)
	}
}

func (sbt *SBTNode) delete(cur *SBTNode, key int) *SBTNode {
	cur.size--
	if cur.key > key {
		cur.r = sbt.delete(cur.r, key)
	} else if cur.key < key {
		cur.l = sbt.delete(cur.l, key)
	} else { // 当前要删掉cur
		if cur.l == nil && cur.r == nil {
			// free cur memory -> C++
			cur = nil
		} else if cur.l == nil && cur.r != nil {
			// free cur memory -> C++
			cur = cur.r
		} else if cur.l != nil && cur.r == nil {
			// free cur memory -> C++
			cur = cur.l
		} else { // 有左有右
			var pre *SBTNode = nil
			des := cur.r
			des.size--
			for des.l != nil {
				pre = des
				des = des.l
				des.size--
			}
			if pre != nil {
				pre.l = des.r
				des.r = cur.r
			}
			des.l = cur.l
			if des.r == nil {
				des.size = des.l.size + 1
			} else {
				des.size = des.l.size + des.r.size + 1
			}
			// free cur memory -> C++
			cur = des
		}
	}
	// cur = maintain(cur);
	return cur
}
