package chash

type Node interface {
	Id() int
}

type RingNode interface {
	Node
	MaxPoint() uint32
	Node() Node
}

type ringNode struct {
	node     Node
	maxPoint uint32
}

func NewRingNode(node Node, maxPoint uint32) RingNode {
	return &ringNode{node, maxPoint}
}

func (me *ringNode) Id() int {

	return me.node.Id()
}

func (me *ringNode) MaxPoint() uint32 {
	return me.maxPoint
}

func (me *ringNode) Node() Node {
	return me.node
}
