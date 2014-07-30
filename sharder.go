package chash

type Sharder struct {
	keyspaceSize int
	nodes        []Node
	nodeList     NodeList
}

func NewSharder(node []Node, nodeList NodeList) *Sharder {
	keyspaceSize := 4294967295 // max uint32
	return NewSharder2(keyspaceSize, node, nodeList)
}

func NewSharder2(keyspaceSize int, node []Node, nodeList NodeList) *Sharder {
	return &Sharder{keyspaceSize, node, nodeList}
}

func (me *Sharder) FillNodeList() {
	for i := 0; i < len(me.nodes); i++ {
		me.nodeList.AddNode(NewRingNode(me.nodes[i], me.GetMaxPoint(i)))
	}
}

func (me *Sharder) GetMaxPoint(nodeIndex int) uint32 {
	if nodeIndex+1 == len(me.nodes) {
		return uint32(me.keyspaceSize)
	}

	pointsPerNode := me.keyspaceSize / len(me.nodes)
	return uint32(pointsPerNode * (nodeIndex + 1))
}
