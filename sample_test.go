package chash_test

import (
	"chash"
	"fmt"
	"testing"
)

func TestSample(t *testing.T) {

	// First we create our own  implementation of node capable of carrying around the information we care about.
	// this might be a handle to a connection pool, an open connection or just a pile of connection information
	nodes := make([]chash.Node, 10)

	// In this sample we're reusing the SqlConn object.  In real-life each node would likely have it's own connection
	var conn SqlConn
	for i := 0; i < len(nodes); i++ {
		nodes[i] = NewSqlNode(i, conn)
	}

	// Create a new ring out of the provided nodes.
	// If you're previously sharded the ring, it's possible to create a ring directly out of RingNode objects using NewRing2.
	ring := chash.NewRing(nodes)

	// Hash a value, get a node.
	// The returned value is of type Node, but can always be type asserted into the node type used when creating the ring
	node := ring.GetNode([]byte("foo")).(*SqlNode)

	// Yay! It worked
	fmt.Printf("Found Node of type %T with values %#v\n", node, node)
}

type SqlConn interface {
	Query(query string) (interface{}, error)
}

type SqlNode struct {
	id   int
	conn SqlConn
}

func NewSqlNode(id int, conn SqlConn) *SqlNode {
	return &SqlNode{id, conn}
}

func (me *SqlNode) Id() int {
	return me.id
}
