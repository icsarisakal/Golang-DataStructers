package main

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func (t *Tree) addBranch(n int) {
	if t.Value == 0 {
		t.Value = n
		return
	}
	if n < t.Value && t.Left == nil {
		t.Left = new(Tree)
		t.Left.addBranch(n)
	} else if n < t.Value {
		t.Left.addBranch(n)
	}
	if n > t.Value && t.Right == nil {
		t.Right = new(Tree)
		t.Right.addBranch(n)
	} else if n > t.Value {
		t.Right.addBranch(n)
	}
	if t.Value == n {
		nT := new(Tree)
		nT.Left = t.Left
		t.Left = nT
		t.Left.addBranch(n)
	}
}

func main() {
	numArr := []int{1, 1, 2, 3, 5, 8, 6, 10, 11, 9, 9, 12, 13, -1}
	tree := new(Tree)
	for _, v := range numArr {
		tree.addBranch(v)
	}
}
