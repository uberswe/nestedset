package main

import (
	"encoding/json"
	"fmt"
	"github.com/uberswe/nestedset"
	"strings"
)

type MySomeType struct {
	NodeId    int64  `json:"-"`
	NodeName  string `json:"-"`
	NodeLevel int64  `json:"-"`
	NodeLeft  int64  `json:"-"`
	NodeRight int64  `json:"-"`
	NodeKey   string `json:"key"`
	NodeValue string `json:"value"`
	ID        int64  `json:"-"`
}

func (t *MySomeType) Key() string {
	return t.NodeKey
}

func (t *MySomeType) Value() string {
	return t.NodeValue
}

func (t *MySomeType) SetKey(s string) {
	t.NodeKey = s
}

func (t *MySomeType) SetValue(s string) {
	t.NodeValue = s
}

// Init it in instance creation
func NewMySomeType() *MySomeType {
	return &MySomeType{}
}

// You can redefine NodeInterface functions

// Return your type
func (t *MySomeType) Type() string {
	return "generic"
}

// Return your inner id
func (t *MySomeType) Id() int64 {
	return t.ID
}

// Set your inner id or generate it
func (t *MySomeType) SetId(id int64) {
	t.ID = id // or t.MyId = getSomeNewId()
}

func (t MySomeType) Level() int64 {

	return t.NodeLevel
}

func (t MySomeType) Left() int64 {
	return t.NodeLeft
}

func (t MySomeType) Right() int64 {
	return t.NodeRight
}

func (t *MySomeType) SetLevel(level int64) {
	t.NodeLevel = level
}

func (t *MySomeType) SetLeft(left int64) {
	t.NodeLeft = left
}

func (t *MySomeType) SetRight(right int64) {
	t.NodeRight = right
}

func main() {

	ns := nestedset.NewNestedSet()

	// create 3 new nodes
	node1 := NewMySomeType()
	node1.NodeKey = "node_1"
	node2 := NewMySomeType()
	node2.NodeKey = "node_2"
	node3 := NewMySomeType()
	node3.NodeKey = "node_3"
	node3.NodeValue = "Node 3 value"
	node4 := NewMySomeType()
	node4.NodeKey = "node_4"
	node4.NodeValue = "Node 4 value"

	err := ns.Add(node1, nil)
	if err != nil {
		return
	} // add node to root
	err = ns.Add(node2, nil)
	if err != nil {
		panic(err)
	} // add node to root
	err = ns.Add(node3, node1)
	if err != nil {
		panic(err)
	} // add node to node1

	err = ns.Move(node3, node2)
	if err != nil {
		panic(err)
	} // move node3 from node1 to node2

	err = ns.Add(node4, node1)
	if err != nil {
		panic(err)
	}

	branch := ns.Branch(nil) // get full tree

	// print tree with indents
	for _, n := range branch {
		fmt.Print(strings.Repeat("..", int(n.Level())))
		fmt.Printf("%s lvl:%d, left:%d, right:%d\n", n.Key(), n.Level(), n.Left(), n.Right())
	}

	j, err := json.Marshal(ns)
	if err != nil {
		panic(err)
	}

	// Print the json object
	fmt.Println(string(j))
}
