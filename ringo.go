package ringo

import "sync"

// RingBuffer ...
type RingBuffer struct {
	data  []interface{}
	count int
	ptr   int
	cap   int
	mu    sync.Mutex
}

// NewRingBuffer with given capacity
func NewRingBuffer(capacity int) RingBuffer {
	return RingBuffer{
		count: 0,
		ptr:   -1,
		data:  make([]interface{}, capacity),
		cap:   capacity,
	}
}

// IsFull returns true if the buffer is full
func (rb *RingBuffer) IsFull() bool { return rb.count >= rb.cap }

// IsEmpty returns true if the buffer is empty
func (rb *RingBuffer) IsEmpty() bool { return rb.count == 0 }

// Push an item to the buffer
func (rb *RingBuffer) Push(val interface{}) {
	rb.mu.Lock()
	defer rb.mu.Unlock()
	rb.ptr = rb.count % rb.cap
	rb.data[rb.ptr] = val
	rb.count++
}

// GetNewest item added to the buffer. Nil if empty.
func (rb *RingBuffer) GetNewest() interface{} {
	if rb.IsEmpty() {
		return nil
	}
	return rb.data[rb.ptr]
}

// GetOldest item added to the buffer. Nil if empty.
func (rb *RingBuffer) GetOldest() interface{} {
	if rb.IsEmpty() {
		return nil
	}
	if rb.IsFull() {
		return rb.data[(rb.ptr+1)%rb.cap]
	}
	return rb.data[0]
}

// Snapshot returns a slice of items from least to most recent. Nil if empty.
func (rb *RingBuffer) Snapshot() (snap []interface{}) {
	rb.mu.Lock()
	defer rb.mu.Unlock()
	if rb.IsEmpty() {
		return nil
	}
	if !rb.IsFull() {
		snap = make([]interface{}, rb.count)
		_ = copy(snap, rb.data[:rb.count])
		return snap
	}
	snap = append(snap, rb.data[(rb.ptr+1)%rb.cap:]...)
	if (rb.ptr+1)%rb.cap == 0 {
		return snap
	}
	return append(snap, rb.data[:rb.ptr+1]...)
}
