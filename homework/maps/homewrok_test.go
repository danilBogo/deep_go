package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type node struct {
	key       int
	value     int
	isDeleted bool

	left  *node
	right *node
}

type OrderedMap struct {
	size int
	root *node
}

func NewOrderedMap() OrderedMap {
	return OrderedMap{}
}

func (m *OrderedMap) Insert(key, value int) {
	defer func() {
		m.size += 1
	}()

	if m.size == 0 {
		m.root = &node{
			key:   key,
			value: value,
		}

		return
	}

	neededNode, sameKey := findNode(key, m.root)
	if sameKey {
		neededNode.isDeleted = false
		return
	}

	newNode := &node{key: key, value: value}
	if key < neededNode.key {
		neededNode.left = newNode
	} else {
		neededNode.right = newNode
	}
}

func (m *OrderedMap) Erase(key int) {
	defer func() {
		m.size -= 1
	}()

	if m.size == 1 {
		m.root = nil
		return
	}

	neededNode, sameKey := findNode(key, m.root)
	if sameKey && !neededNode.isDeleted {
		neededNode.isDeleted = true
	}
}

func (m *OrderedMap) Contains(key int) bool {
	neededNode, sameKey := findNode(key, m.root)
	return sameKey && !neededNode.isDeleted
}

func (m *OrderedMap) Size() int {
	return m.size
}

func (m *OrderedMap) ForEach(action func(int, int)) {
	traverse(m.root, action)
}

func findNode(key int, currentNode *node) (*node, bool) {
	if key < currentNode.key && currentNode.left != nil {
		return findNode(key, currentNode.left)
	}

	if key > currentNode.key && currentNode.right != nil {
		return findNode(key, currentNode.right)
	}

	return currentNode, currentNode.key == key
}

func traverse(currentNode *node, action func(int, int)) {
	if currentNode.left != nil {
		traverse(currentNode.left, action)
	}

	if !currentNode.isDeleted {
		action(currentNode.key, currentNode.value)
	}

	if currentNode.right != nil {
		traverse(currentNode.right, action)
	}
}

func TestCircularQueue(t *testing.T) {
	data := NewOrderedMap()
	assert.Zero(t, data.Size())

	data.Insert(10, 10)
	data.Insert(5, 5)
	data.Insert(15, 15)
	data.Insert(2, 2)
	data.Insert(4, 4)
	data.Insert(12, 12)
	data.Insert(14, 14)

	assert.Equal(t, 7, data.Size())
	assert.True(t, data.Contains(4))
	assert.True(t, data.Contains(12))
	assert.False(t, data.Contains(3))
	assert.False(t, data.Contains(13))

	var keys []int
	expectedKeys := []int{2, 4, 5, 10, 12, 14, 15}
	data.ForEach(func(key, _ int) {
		keys = append(keys, key)
	})

	assert.True(t, reflect.DeepEqual(expectedKeys, keys))

	data.Erase(15)
	data.Erase(14)
	data.Erase(2)

	assert.Equal(t, 4, data.Size())
	assert.True(t, data.Contains(4))
	assert.True(t, data.Contains(12))
	assert.False(t, data.Contains(2))
	assert.False(t, data.Contains(14))

	keys = nil
	expectedKeys = []int{4, 5, 10, 12}
	data.ForEach(func(key, _ int) {
		keys = append(keys, key)
	})

	assert.True(t, reflect.DeepEqual(expectedKeys, keys))
}
