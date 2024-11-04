package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type CircularQueue struct {
	values     []int
	frontIndex int
	backIndex  int
	count      int
}

func NewCircularQueue(size int) CircularQueue {
	if size < 0 {
		panic("size can not be negative")
	}

	return CircularQueue{
		values:     make([]int, size),
		frontIndex: -1,
		backIndex:  -1,
	}
}

func (q *CircularQueue) Push(value int) bool {
	if q.Full() {
		return false
	}

	if q.frontIndex == -1 {
		q.frontIndex = 0
	}

	q.backIndex = (q.backIndex + 1) % len(q.values)
	q.values[q.backIndex] = value
	q.count += 1

	return true
}

func (q *CircularQueue) Pop() bool {
	if q.Empty() {
		return false
	}

	if q.frontIndex == q.backIndex {
		q.backIndex = 0
	}

	q.frontIndex = (q.frontIndex + 1) % len(q.values)
	q.count -= 1

	return true
}

func (q *CircularQueue) Front() int {
	if q.Empty() {
		return -1
	}

	return q.values[q.frontIndex]
}

func (q *CircularQueue) Back() int {
	if q.Empty() {
		return -1
	}

	return q.values[q.backIndex]
}

func (q *CircularQueue) Empty() bool {
	return q.count == 0
}

func (q *CircularQueue) Full() bool {
	return q.count == len(q.values)
}

func TestCircularQueue(t *testing.T) {
	const queueSize = 3
	queue := NewCircularQueue(queueSize)

	assert.True(t, queue.Empty())
	assert.False(t, queue.Full())

	assert.Equal(t, -1, queue.Front())
	assert.Equal(t, -1, queue.Back())
	assert.False(t, queue.Pop())

	assert.True(t, queue.Push(1))
	assert.True(t, queue.Push(2))
	assert.True(t, queue.Push(3))
	assert.False(t, queue.Push(4))

	assert.True(t, reflect.DeepEqual([]int{1, 2, 3}, queue.values))

	assert.False(t, queue.Empty())
	assert.True(t, queue.Full())

	assert.Equal(t, 1, queue.Front())
	assert.Equal(t, 3, queue.Back())

	assert.True(t, queue.Pop())
	assert.False(t, queue.Empty())
	assert.False(t, queue.Full())
	assert.True(t, queue.Push(4))

	assert.True(t, reflect.DeepEqual([]int{4, 2, 3}, queue.values))

	assert.Equal(t, 2, queue.Front())
	assert.Equal(t, 4, queue.Back())

	assert.True(t, queue.Pop())
	assert.True(t, queue.Pop())
	assert.True(t, queue.Pop())
	assert.False(t, queue.Pop())

	assert.True(t, queue.Empty())
	assert.False(t, queue.Full())
}
