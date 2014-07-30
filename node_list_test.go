package chash_test

import (
	"chash"
	"strconv"
	"testing"
)

func BenchmarkSequentialLarge(b *testing.B) {
	nodes := getNodes(10000)
	nodeList := chash.NewNodeList()
	chash.NewSharder(nodes, nodeList).FillNodeList()
	ring := chash.NewRing2(nodeList)
	for n := 0; n < b.N; n++ {
		ring.GetNode([]byte(strconv.Itoa(n)))
	}
}

func BenchmarkSequentialSmall(b *testing.B) {
	nodes := getNodes(10)
	nodeList := chash.NewNodeList()
	chash.NewSharder(nodes, nodeList).FillNodeList()
	ring := chash.NewRing2(nodeList)
	for n := 0; n < b.N; n++ {
		ring.GetNode([]byte(strconv.Itoa(n)))
	}
}

func BenchmarkBinarySearchLarge(b *testing.B) {
	nodes := getNodes(10000)
	nodeList := chash.NewBinSortNodeList()
	chash.NewSharder(nodes, nodeList).FillNodeList()
	ring := chash.NewRing2(nodeList)
	for n := 0; n < b.N; n++ {
		ring.GetNode([]byte(strconv.Itoa(n)))
	}
}

func BenchmarkBinarySearchSmall(b *testing.B) {
	nodes := getNodes(10)
	nodeList := chash.NewBinSortNodeList()
	chash.NewSharder(nodes, nodeList).FillNodeList()
	ring := chash.NewRing2(nodeList)
	for n := 0; n < b.N; n++ {
		ring.GetNode([]byte(strconv.Itoa(n)))
	}
}

func BenchmarkHashSpeed(b *testing.B) {
	hasher := chash.NewMurmur3Hasher()
	for n := 0; n < b.N; n++ {
		hasher.Hash([]byte(strconv.Itoa(n)))
	}
}
