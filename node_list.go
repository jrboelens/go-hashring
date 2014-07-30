package chash

import "sort"

type NodeList interface {
	AddNode(node RingNode)
	FindNode(point uint32) Node
}

type binSortNodeList struct {
	nodes []RingNode
}

func NewBinSortNodeList() NodeList {
	nodes := make([]RingNode, 0)
	return &binSortNodeList{nodes}
}

func (me *binSortNodeList) AddNode(node RingNode) {
	me.nodes = append(me.nodes, node)
}

// Assumes the nodes are sorted
func (me *binSortNodeList) FindNode(point uint32) Node {

	i := sort.Search(len(me.nodes), func(i int) bool {
		return point < me.nodes[i].MaxPoint()
	})

	if i > -1 {
		return me.nodes[i].Node()
	}
	return nil
}

// Don't use this, it's slow
// it was put here in order to benchmark the binary search performance
type nodeList struct {
	nodes []RingNode
}

func NewNodeList() NodeList {
	nodes := make([]RingNode, 0)
	return &nodeList{nodes}
}

func (me *nodeList) AddNode(node RingNode) {
	me.nodes = append(me.nodes, node)
}

// Assumes the nodes are sorted
func (me *nodeList) FindNode(point uint32) Node {
	for i := 0; i < len(me.nodes); i++ {
		if me.nodes[i] == nil {
			continue
		}

		if me.nodes[i].MaxPoint() > point {
			return me.nodes[i]
		}
	}
	return nil
}
