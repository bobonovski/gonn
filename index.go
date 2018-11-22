package gonn

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

// Item represents the input data point, Id denotes
// the unique identifier of the data point and we
// assume the value is stored in float32 slice.
type Item struct {
	Id  string
	Val []float32
}

// Node represents binary search tree which is constructed
// by randomly split some dimension of the data point
const (
	INNODE int = iota
	LEAF
)

type Node struct {
	Boundary []float32
	Indices  []int
	NodeType int
	Left     *Node
	Right    *Node
}

// Index is the main data structure to store the nearest
// neighbors information. It contains user specified number
// of randomly splitted trees for fast item searching
type Index struct {
	Items   []*Item // input data item
	Nodes   []*Node // internal search tree
	Count   int     // the number of search trees to build
	MinNode int     // mininum number of items in a node
}

func NewIndex(c int, m int) *Index {
	return &Index{Count: c, MinNode: m}
}

// add data item to index, each item should have unique id
func (i *Index) AddItem(id string, val []float32) error {
	if id == "" {
		return fmt.Errorf("input id is empty")
	}
	if len(val) == 0 {
		return fmt.Errorf("input val is empty")
	}
	i.Items = append(i.Items, &Item{
		Id:  id,
		Val: val,
	})
	return nil
}

// build multiple binary search tree
func (i *Index) Build() error {
	if len(i.Items) == 0 {
		return fmt.Errorf("date items is empty")
	}
	var err error
	for c := 0; c < i.Count; c += 1 {
		n := Node{}
		for k := 0; k < len(i.Items); k += 1 {
			n.Indices = append(n.Indices, k)
		}
		err = i.Split(&n)
		if err != nil {
			return err
		}
		i.Nodes = append(i.Nodes, &n)
	}
	return nil
}

// recursively split node until the current node
// has fewer than MinNode items belong to it
func (i *Index) Split(n *Node) error {
	// choose two items randomly
	ridx := rand.Perm(len(n.Indices))
	p := i.Items[n.Indices[ridx[0]]]
	q := i.Items[n.Indices[ridx[1]]]
	// compute split boundary
	b, err := GetBoundary(p.Val, q.Val)
	if err != nil {
		return err
	}
	n.Boundary = b
	// split items
	left := Node{}
	right := Node{}
	for _, k := range n.Indices {
		v, err := EvalFormula(b, i.Items[k].Val)
		if err != nil {
			return err
		}
		if v <= 0.0 {
			left.Indices = append(left.Indices, k)
		} else {
			right.Indices = append(right.Indices, k)
		}
	}
	// split left child
	if len(left.Indices) > i.MinNode {
		left.NodeType = INNODE
		err = i.Split(&left)
		if err != nil {
			return err
		}
		n.Left = &left
	} else if len(left.Indices) > 0 {
		left.NodeType = LEAF
		n.Left = &left
	}
	// split right child
	if len(right.Indices) > i.MinNode {
		right.NodeType = INNODE
		err = i.Split(&right)
		if err != nil {
			return err
		}
		n.Right = &right
	} else if len(right.Indices) > 0 {
		right.NodeType = LEAF
		n.Right = &right
	}
	return nil
}
