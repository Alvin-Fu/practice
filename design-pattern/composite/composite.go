package composite

import "fmt"
// 用于表示树
type Tree interface {
	Parent() Tree
	SetParent(Tree)
	Name()string
	SetName(string)
	AddChild(Tree)
	Print(string)
}

const   (
	LeafNode = iota
	RLSubtree
)
func NewTree(nodeType int, name string)Tree{
	var t Tree
	switch nodeType {
	case LeafNode:
		t = NewLeaf()
	case RLSubtree:
		t = NewSubtree()
	}
	t.SetName(name)
	return t
}

type tree struct{
	name   string
	parent Tree
}

func (t *tree) Parent()Tree{
	return t.parent
}

func (t *tree)SetParent(parent Tree){
	t.parent = parent
}

func (t *tree) Name()string{
	return t.name
}

func (t *tree) SetName(name string){
	t.name = name
}

func (t *tree)AddChild(Tree){}
func (t *tree) Print(string){}

type Leaf struct{
	tree
}
// 叶子节点
func NewLeaf()*Leaf{
	return &Leaf{}
}
func (l *Leaf) Print(str string){
	fmt.Printf("%s %s\n", str, l.name)
}
// 子树
type Subtree struct{
	tree
	childs []Tree
}

func NewSubtree()*Subtree{
	return &Subtree{
		childs: make([]Tree, 0),
	}
}
// 添加一个孩子节点
func (s *Subtree) AddChild(child Tree){
	child.SetParent(s)
	s.childs = append(s.childs, child)
}
func (s *Subtree) Print(str string){
	fmt.Printf("%s %s\n", str, s.Name())
	str += " "
	for _, c := range s.childs{
		c.Print(str)
	}
}
