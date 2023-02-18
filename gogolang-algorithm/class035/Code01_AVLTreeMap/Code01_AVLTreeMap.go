package main

type AVLNode struct {
	k int // Comparable
	v int
	l *AVLNode
	r *AVLNode
	h int
}

func NewAVLNode(key, value int) *AVLNode {
	return &AVLNode{
		k: key,
		v: value,
		h: 1,
	}
}

// 右旋
func (avl *AVLNode) rightRotate(cur *AVLNode) *AVLNode {
	left := cur.l
	cur.l = left.r
	left.r = cur

	if cur.l != nil {
		cur.h = cur.l.h
	} else {
		cur.h = 0
	}
	if cur.r != nil {
		cur.h = Max(cur.h, cur.r.h)
	} else {
		cur.h = 0
	}
	cur.h += 1
	//left.h = Max((left.l != nil ? left.l.h : 0), (left.r != nil ? left.r.h : 0)) + 1

	return left
}

// 左旋
func (avl *AVLNode) leftRotate(cur *AVLNode) *AVLNode {
	right := cur.r
	cur.r = right.l
	right.l = cur
	//cur.h = Max((cur.l != nil ? cur.l.h : 0), (cur.r != nil ? cur.r.h : 0)) + 1
	//right.h = Max((right.l != nil ? right.l.h : 0), (right.r != nil ? right.r.h : 0)) + 1
	return right
}

// 四种方法查询，调整平衡
func (avl *AVLNode) maintain(cur *AVLNode) *AVLNode {
	if cur == nil {
		return nil
	}
	//int leftHeight = cur.l != null ? cur.l.h : 0;
	//int rightHeight = cur.r != null ? cur.r.h : 0;
	//if (Math.abs(leftHeight - rightHeight) > 1) {
	//	if (leftHeight > rightHeight) {
	//		int leftLeftHeight = cur.l != null && cur.l.l != null ? cur.l.l.h : 0;
	//		int leftRightHeight = cur.l != null && cur.l.r != null ? cur.l.r.h : 0;
	//		if (leftLeftHeight >= leftRightHeight) {
	//			cur = rightRotate(cur);
	//		} else {
	//			cur.l = leftRotate(cur.l);
	//			cur = rightRotate(cur);
	//		}
	//	} else {
	//		int rightLeftHeight = cur.r != null && cur.r.l != null ? cur.r.l.h : 0;
	//		int rightRightHeight = cur.r != null && cur.r.r != null ? cur.r.r.h : 0;
	//		if (rightRightHeight >= rightLeftHeight) {
	//			cur = leftRotate(cur);
	//		} else {
	//			cur.r = rightRotate(cur.r);
	//			cur = leftRotate(cur);
	//		}
	//	}
	//}
	return cur
}

//func (avl *AVLNode) findLastIndex() {
//
//}
//
//func (avl *AVLNode) findLastNoSmallIndex() {
//
//}
//
//func (avl *AVLNode) findLastNoBigIndex() {
//
//}

func (avl *AVLNode) add(cur *AVLNode, key, value int) *AVLNode {
	if cur == nil {
		return NewAVLNode(key, value)
	} else {
		if cur.k < key {
			cur.l = avl.add(cur.l, key, value)
		} else {
			cur.r = avl.add(cur.r, key, value)
		}
		//cur.h = Math.max(cur.l != null ? cur.l.h : 0, cur.r != null ? cur.r.h : 0) + 1;
		return avl.maintain(cur)
	}
}

// 在cur这棵树上，删掉key所代表的节点
// 返回cur这棵树的新头部
func (avl *AVLNode) delete(cur *AVLNode, key int) *AVLNode{
	if cur.k > key {
		cur.r = avl.delete(cur.r, key);
	} else if  cur.k < key {
		cur.l = avl.delete(cur.l, key);
	} else {
		if (cur.l == nil && cur.r == nil) {
			cur = nil;
		} else if (cur.l == nil && cur.r != nil) {
			cur = cur.r;
		} else if (cur.l != nil && cur.r == nil) {
			cur = cur.l;
		} else {
			 des := cur.r;
			for (des.l != nil) {
				des = des.l;
			}
			cur.r =  avl.delete(cur.r, des.k);
			des.l = cur.l;
			des.r = cur.r;
			cur = des;
		}
	}
	if (cur != nil) {
		cur.h = Max(cur.l != nil ? cur.l.h : 0, cur.r != nil ? cur.r.h : 0) + 1;
	}
	return avl.maintain(cur);
}

//func (avl *AVLNode) size() {
//
//}
//
//func (avl *AVLNode) containsKey() {
//
//}
//
//func (avl *AVLNode) put() {
//
//}
//
//func (avl *AVLNode) remove() {
//
//}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
