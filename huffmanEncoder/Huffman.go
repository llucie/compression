package main

import (
	"container/heap"
	"fmt"
)

type TreeElement interface {
	frequency() int
}

type Tree []TreeElement

func (tree *Tree) Push(ele interface{}) { *tree = append(*tree, ele.(TreeElement)) }

func (tree *Tree) Pop() (popped interface{}) {
	popped = (*tree)[len(*tree)-1]
	*tree = (*tree)[:len(*tree)-1]
	return
}

func (tree Tree) Len() int           { return len(tree) }
func (tree Tree) Less(i, j int) bool { return tree[i].frequency() < tree[j].frequency() }
func (tree Tree) Swap(i, j int)      { tree[i], tree[j] = tree[j], tree[i] }

type Leaf struct {
	freq   int
	Number uint32
}

type Node struct {
	left  TreeElement
	right TreeElement
	freq  int
}

func (leaf Leaf) frequency() int { return leaf.freq }
func (node Node) frequency() int { return node.freq }

func frequencyMap(image []uint32) map[uint32]int {
	freqMap := make(map[uint32]int)
	for _, elem := range image {
		freqMap[elem]++
	}
	return freqMap
}

func huffmanCompression(image []uint32) []byte {
	freqMap := frequencyMap(image)
	var tree Tree
	for nb, freq := range freqMap {
		tree = append(tree, Leaf{freq: freq, Number: nb})
	}

	heap.Init(&tree)
	for len(tree) > 1 {
		firstNode := heap.Pop(&tree).(TreeElement)
		secondNode := heap.Pop(&tree).(TreeElement)
		heap.Push(&tree, Node{left: firstNode, right: secondNode, freq: firstNode.frequency() + secondNode.frequency()})
	}
	//printCodes(heap.Pop(&tree).(TreeElement), []byte{})

	return encodeImage(image, heap.Pop(&tree).(TreeElement))
}

func printCodes(tree TreeElement, code []byte) {
	switch i := tree.(type) {
	case Leaf:
		// If this is a leaf node, then it contains one of the input
		// characters, print the character and its code from byte[]
		fmt.Printf("%s%d\t%s%d\t%s%s\n", "number: ", i.Number, " freq: ", i.freq, "code: ", string(code))
	case Node:
		// Assign 0 to left edge and recur
		code = append(code, '0')
		printCodes(i.left, code)
		code = code[:len(code)-1]

		// Assign 1 to right edge and recur
		code = append(code, '1')
		printCodes(i.right, code)
		code = code[:len(code)-1]
	}
}

func mapPixelToCode(tree TreeElement, codeMap map[uint32][]byte, code []byte) {
	switch i := tree.(type) {
	case Leaf:
		codeMap[i.Number] = code
	case Node:
		// Assign 0 to left edge and recur
		code = append(code, '0')
		mapPixelToCode(i.left, codeMap, code)
		code = code[:len(code)-1]

		// Assign 1 to right edge and recur
		code = append(code, '1')
		mapPixelToCode(i.right, codeMap, code)
		code = code[:len(code)-1]
	}
}

func encodeImage(image []uint32, tree TreeElement) []byte {
	codeMap := make(map[uint32][]byte)
	mapPixelToCode(tree, codeMap, []byte{})

	encodedImage := make([]byte, 0)
	for pixel := range image {
		code := codeMap[image[pixel]]
		for codeElement := range code {
			encodedImage = append(encodedImage, code[codeElement])
		}
	}

	return encodedImage
}
