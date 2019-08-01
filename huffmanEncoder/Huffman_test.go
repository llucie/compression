package main

import "testing"

func TestTreePush(t *testing.T) {
	var tree Tree
	leaf := Leaf{Number: 1, freq: 2}
	tree.Push(leaf)

	if len(tree) != 1 {
		t.Error("Tree size different from 1. Equals ", len(tree))
	}

	tree.Push(Node{freq: 42})

	if len(tree) != 2 {
		t.Error("Tree size different from 2. Equals ", len(tree))
	}
}

func TestTreePop(t *testing.T) {

}

func TestFrequencyMap(t *testing.T) {

}
