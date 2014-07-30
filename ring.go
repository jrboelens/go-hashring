package chash

type Ring struct {
	nodeList NodeList
	hasher   Hasher
}

// Convenience function for creating a default Ring from a slice of Nodes
func NewRing(nodes []Node) *Ring {
	nodeList := NewBinSortNodeList()
	NewSharder(nodes, nodeList).FillNodeList()
	return NewRing2(nodeList)
}

// Convenience function for creating a default Ring from NodeList
func NewRing2(nodeList NodeList) *Ring {
	hasher := NewMurmur3Hasher()
	return &Ring{nodeList, hasher}
}

func NewRing3(nodeList NodeList, hasher Hasher) *Ring {
	return &Ring{nodeList, hasher}
}

func (me *Ring) GetNode(data []byte) Node {
	point := me.hasher.Hash(data)
	return me.nodeList.FindNode(point)
}
