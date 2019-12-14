package composite

import "testing"

func TestNewTree(t *testing.T) {
	root := NewTree(RLSubtree, "root")
	s1 := NewTree(RLSubtree, "s1")
	s2 := NewTree(RLSubtree, "s2")
	s3 := NewTree(RLSubtree, "s3")
	l1 := NewTree(LeafNode, "l1")
	l2 := NewTree(LeafNode, "l2")
	root.AddChild(s1)
	root.AddChild(s2)
	s1.AddChild(s3)
	s1.SetParent(root)
	s2.SetParent(root)
	s3.SetParent(s1)
	s1.AddChild(l1)
	s2.AddChild(l2)
	root.Print("")

}
