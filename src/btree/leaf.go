package btree

import (
	"errors"
	"fmt"
	//"log"
	"sort"
)

func (leaf LeafNode) Search(key int) ([]int, error) {
	if len(leaf.keys) != len(leaf.values) {
		return nil, errors.New("LeafNode's keys and values should have similar number of items")
	}
	index := sort.SearchInts(leaf.keys, key)
	if index == len(leaf.keys) || leaf.keys[index] != key {
		return nil, errors.New(fmt.Sprintf("Key %d not found", key))
	}
	var result []int
	for ; leaf.keys[index] == key; index++ {
		result = append(result, leaf.values[index])
	}
	return result, nil
}

func (leaf *LeafNode) Insert(key int, value int, degree int) error {
	index := sort.SearchInts(leaf.keys, key)
	leaf.keys, leaf.values = insertInt(leaf.keys, index, key), insertInt(leaf.values, index, value)
	// number of keys is still lower than the maximum number
	if len(leaf.keys) < degree {
		return nil
	}
	numberOfKeys := degree / 2
	// otherwise, split the LeafNode and create new InternalNode
	rightSibling := newLeafNode(leaf.keys[numberOfKeys:], leaf.values[numberOfKeys:], leaf.parent, leaf.rightSibling)
	leaf.rightSibling, leaf.keys, leaf.values = rightSibling, leaf.keys[:numberOfKeys], leaf.values[:numberOfKeys]
	// if parent node is nil -> create new parent node
	if leaf.parent == nil {
		parent := newInternalNode([]int{rightSibling.keys[0]}, []interface{}{leaf, rightSibling})
		leaf.parent, rightSibling.parent = parent, parent
		return nil
	}
	// else, insert the split key into the parent
	return leaf.parent.Insert(rightSibling.keys[0], rightSibling, degree)
}
