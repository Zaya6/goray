package game

import "fmt"

// Node is a type of object that represents logic, interactivity, or visuals in a game
type Node interface {
	Init()
	Update(delta float32)
	Draw()
	AddNode(nNode *Node)
	RemoveChild(rNode *Node)
	AddParent(pNode *Node)
	RemoveParent()
	Destroy() // TODO add ability to delay destruction
	GetChildren() []*Node
	Print()
	AddUpdate(f func())
	AddDraw(f func())
	AddInit(f func())
}

// node is the underlying structure of every other node
type node struct {
	Pos        Vec
	parent     *Node
	children   []*Node
	updateList []func()
	drawList   []func()
	initList   []func()
}

// test prints stuff
func (n *node) Print() { // TODO : print the type
	fmt.Println("testers")
}

// Init an update function to update the node
func (n *node) Init() {
	for _, child := range n.children {
		(*child).Init()
	}
	for _, f := range n.initList {
		f()
	}
}

// Update an update function to update the node
func (n *node) Update(delta float32) {
	for _, child := range n.children {
		(*child).Update(delta)
	}
	for _, f := range n.updateList {
		f()
	}
}

// Draw is the function that draws something to screen
func (n *node) Draw() {
	for _, child := range n.children {
		(*child).Draw()
	}
	for _, f := range n.drawList {
		f()
	}
}

// AddNode ads a node to the child list
func (n *node) AddNode(nNode *Node) {
	(*nNode).AddParent(n.N())
	n.children = append(n.children, nNode)
}

// RemoveChild removes a child from the child list
func (n *node) RemoveChild(rNode *Node) {
	for key, child := range n.children {
		if child == rNode {
			//swap child with the last and set the slice to everything but the last
			n.children[len(n.children)-1], n.children[key] = n.children[key], n.children[len(n.children)-1] //that's cool
			n.children = n.children[:len(n.children)-1]
			(*rNode).RemoveParent()
		}
	}
}

func (n *node) AddParent(pNode *Node) {
	n.parent = pNode
}

func (n *node) RemoveParent() {
	n.parent = nil
}

// Destroy allows a node to destroy itself and remove it from its parent
func (n *node) Destroy() { //need to test
	(*n.parent).RemoveChild(n.N())
	n = nil
}

// GetChildren allows a node to destroy itself and remove it from its parent
func (n *node) GetChildren() []*Node {
	return n.children
}

// N returns the node interface{}
func (n *node) N() *Node {
	var nd Node = n
	return &nd
}

// AddUpdate adds an anonymous function to the update loop
func (n *node) AddUpdate(f func()) {
	n.updateList = append(n.updateList, f)
}

// AddDraw adds an anonymous function to the update loop
func (n *node) AddDraw(f func()) {
	n.drawList = append(n.drawList, f)
}

// AddDraw adds an anonymous function to the update loop
func (n *node) AddInit(f func()) {
	n.initList = append(n.initList, f)
}
