package ringo

import "sync"

// RingBuffer data structure implements a circular buffer of fixed size.
// It functions as an append/push-only buffer and once the buffer is full,
// further pushes overwrite the oldest written value.
type RingBuffer[T any] struct {
	data  []T
	count int
	ptr   int
	cap   int
	mu    sync.Mutex
}

// NewRingBuffer with given capacity
func NewRingBuffer[T any](capacity int) RingBuffer[T] {
	return RingBuffer[T]{
		count: 0,
		ptr:   -1,
		data:  make([]T, capacity),
		cap:   capacity,
	}
}

// IsFull returns true if the buffer is full
func (rb *RingBuffer[T]) IsFull() bool { return rb.count >= rb.cap }

// IsEmpty returns true if the buffer is empty
func (rb *RingBuffer[T]) IsEmpty() bool { return rb.count == 0 }

// Push an item to the buffer
func (rb *RingBuffer[T]) Push(val T) {
	rb.mu.Lock()
	defer rb.mu.Unlock()
	rb.ptr = rb.count % rb.cap
	rb.data[rb.ptr] = val
	rb.count++
}

// GetNewest item added to the buffer. Returns the zero-value of T if empty.
func (rb *RingBuffer[T]) GetNewest() (res T) {
	rb.mu.Lock()
	defer rb.mu.Unlock()
	if rb.IsEmpty() {
		return
	}
	return rb.data[rb.ptr]
}

// GetOldest item added to the buffer. Returns the zero-value of T if empty.
func (rb *RingBuffer[T]) GetOldest() (res T) {
	rb.mu.Lock()
	defer rb.mu.Unlock()
	if rb.IsEmpty() {
		return
	}
	if rb.IsFull() {
		return rb.data[(rb.ptr+1)%rb.cap]
	}
	return rb.data[0]
}

// Snapshot returns a slice of items from least to most recent. Nil if empty.
func (rb *RingBuffer[T]) Snapshot() (snap []T) {
	rb.mu.Lock()
	defer rb.mu.Unlock()
	if rb.IsEmpty() {
		return nil
	}
	if !rb.IsFull() {
		snap = make([]T, rb.count)
		_ = copy(snap, rb.data[:rb.count])
		return snap
	}
	snap = append(snap, rb.data[(rb.ptr+1)%rb.cap:]...)
	if (rb.ptr+1)%rb.cap == 0 {
		return snap
	}
	return append(snap, rb.data[:rb.ptr+1]...)
}
