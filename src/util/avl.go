package util

// This code is an adoption of the original implementation at:
// https://www.golangprograms.com/golang-program-for-implementation-of-avl-trees.html

//type Key interface {
//	Less(Key) bool
//	Eq(Key) bool
//}

type AVLTree struct {
	root *AVLNode
}

type AVLNode struct {
	data    IComparable
	balance int
	link    [2]*AVLNode
}

// Put a node into the AVL tree.
func (t *AVLTree) Put(data IComparable) {
	if t.root == nil {
		t.root = &AVLNode{data: data}
		return
	} else {
		t.root, _ = t.root.putR(data)
	}
}

// Remove a single item from an AVL tree.
func (t *AVLTree) Remove(data IComparable) {
	t.root, _ = t.root.removeR(data)
}

func opp(dir int) int {
	return 1 - dir
}

// single rotation
func (root *AVLNode) single(dir int) *AVLNode {
	save := root.link[opp(dir)]
	root.link[opp(dir)] = save.link[dir]
	save.link[dir] = root
	return save
}

// double rotation
func (root *AVLNode) double(dir int) *AVLNode {
	save := root.link[opp(dir)].link[dir]

	root.link[opp(dir)].link[dir] = save.link[opp(dir)]
	save.link[opp(dir)] = root.link[opp(dir)]
	root.link[opp(dir)] = save

	save = root.link[opp(dir)]
	root.link[opp(dir)] = save.link[dir]
	save.link[dir] = root
	return save
}

// adjust valance factors after double rotation
func (root *AVLNode) adjustBalance(dir, bal int) {
	n := root.link[dir]
	nn := n.link[opp(dir)]
	switch nn.balance {
	case 0:
		root.balance = 0
		n.balance = 0
	case bal:
		root.balance = -bal
		n.balance = 0
	default:
		root.balance = 0
		n.balance = bal
	}
	nn.balance = 0
}

func (root *AVLNode) insertBalance(dir int) *AVLNode {
	n := root.link[dir]
	bal := 2*dir - 1
	if n.balance == bal {
		root.balance = 0
		n.balance = 0
		return root.single(opp(dir))
	}
	root.adjustBalance(dir, bal)
	return root.double(opp(dir))
}

func (root *AVLNode) putR(data IComparable) (*AVLNode, bool) {
	// if root == nil {
	// 	return &AVLNode{data: data}, false
	// }
	dir := 0
	if root.data.Equal(data) {
		root.data = data // if data is the same, replace the data and return
		return root, true
	} else if root.data.Compare(data) < 0 {
		dir = 1
	}
	var done bool
	root.link[dir], done = root.link[dir].putR(data)
	if done {
		return root, true
	}
	root.balance += 2*dir - 1
	switch root.balance {
	case 0:
		return root, true
	case 1, -1:
		return root, false
	}
	return root.insertBalance(dir), true
}

func (root *AVLNode) removeBalance(dir int) (*AVLNode, bool) {
	n := root.link[opp(dir)]
	bal := 2*dir - 1
	switch n.balance {
	case -bal:
		root.balance = 0
		n.balance = 0
		return root.single(dir), false
	case bal:
		root.adjustBalance(opp(dir), -bal)
		return root.double(dir), false
	}
	root.balance = -bal
	n.balance = bal
	return root.single(dir), true
}

func (root *AVLNode) removeR(data IComparable) (*AVLNode, bool) {
	if root == nil {
		return nil, false
	}
	if root.data.Equal(data) {
		switch {
		case root.link[0] == nil:
			return root.link[1], false
		case root.link[1] == nil:
			return root.link[0], false
		}
		heir := root.link[0]
		for heir.link[1] != nil {
			heir = heir.link[1]
		}
		root.data = heir.data
		data = heir.data
	}
	dir := 0
	if root.data.Compare(data) < 0 {
		dir = 1
	}
	var done bool
	root.link[dir], done = root.link[dir].removeR(data)
	if done {
		return root, true
	}
	root.balance += 1 - 2*dir
	switch root.balance {
	case 1, -1:
		return root, true
	case 0:
		return root, false
	}
	return root.removeBalance(dir)
}
