package chash_test

import (
	"chash"
	"testing"
)

func TestSharderSplitsNodesEvenly(t *testing.T) {
	testData := []struct {
		numPoints int
		numNodes  int
		index     int
		expected  int
	}{
		{110, 11, 0, 10},
		{110, 11, 1, 20},
		{110, 11, 10, 110},
		{111, 11, 10, 111},
	}

	for _, test := range testData {
		nodes := getNodes(test.numNodes)
		nodeList := chash.NewNodeList()
		sharder := chash.NewSharder2(test.numPoints, nodes, nodeList)
		maxPoint := sharder.GetMaxPoint(test.index)
		if maxPoint != uint32(test.expected) {
			t.Error("Expected", test.expected, "Got", maxPoint)
		}
	}
}
