package btree

import (
	"errors"
	"fmt"
	"reflect"
)

func (tree *BPlusTree) search(key int) (*LeafNode, error) {
	var leaf *LeafNode
	switch root := tree.root.(type) {
	case *InternalNode:
		var err error
		leaf, err = root.searchPossibleLeafNode(key)
		if err != nil {
			return nil, err
		}
	case *LeafNode:
		leaf = root
	default:
		return nil, errors.New(fmt.Sprintf("Class of a node should be LeafNode or InternalNode insted of %s", reflect.TypeOf(root).String()))
	}
	_, err := leaf.getValue(key)
	if err != nil {
		return nil, err
	}
	return leaf, nil
}

func (tree *BPlusTree) insert(key int, value int) error {
	leaf, err := tree.search(key)
	if err != nil {
		return err
	}
	err = leaf.insertValue(key, value, tree.degree)
	if err != nil {
		return err
	}
	if leaf.parent == nil {
		tree.root = leaf
		return nil
	}
	newParent := leaf.parent
	for newParent.parent != nil {
		newParent = newParent.parent
	}
	tree.root = newParent
	return nil
}
