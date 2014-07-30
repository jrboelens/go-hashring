package chash_test

import (
	"chash"
	"strconv"
	"testing"
)

func TestGetNodeReturnsTheSameValueTwice(t *testing.T) {
	ring := getRing(100)
	node := ring.GetNode([]byte("the itsy bitsy spider went up the water spout "))
	node1 := ring.GetNode([]byte("the itsy bitsy spider went up the water spout "))
	if node != node1 {
		t.Error("Node mismatch", node, node, 1)
	}
}

func TestDistributionByBruteForce(t *testing.T) {
	ring := getRing(100)
	counter := make(map[int]int)
	for i := 0; i < 100000; i++ {
		node := ring.GetNode([]byte(strconv.Itoa(i)))
		counter[node.Id()]++
	}

	vari := 100
	expected := 1000
	for key, val := range counter {
		if val < expected-vari || val > expected+vari {
			t.Error("Uneven distribution", key, val, expected)
		}
	}

}

type testNode struct {
	id int
}

func (me *testNode) Id() int {
	return me.id
}

func getRing(numNodes int) *chash.Ring {
	nodes := getNodes(numNodes)
	return chash.NewRing(nodes)
}

func getNodes(num int) []chash.Node {
	nodes := make([]chash.Node, num)
	for i := 1; i <= num; i++ {
		nodes[i-1] = &testNode{i}
	}
	return nodes
}
