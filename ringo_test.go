package ringo

import (
	"testing"
)

func TestRingBuffer(t *testing.T) {
	buf := NewRingBuffer[string](2)
	if !buf.IsEmpty() {
		t.Errorf("IsEmpty() = %t, want %t", false, true)
	}
	if old := buf.GetOldest(); old != "" {
		t.Errorf("GetOldest() = %q, want %q", old, "")
	}
	if n := buf.GetNewest(); n != "" {
		t.Errorf("GetNewest() = %q, want %q", n, "")
	}
	if snap := buf.Snapshot(); snap != nil {
		t.Errorf("Snapshot() = %q, want 'nil'", snap)
	}
	buf.Push("a")
	if buf.IsEmpty() {
		t.Errorf("IsEmpty() = %t, want %t", true, false)
	}
	if buf.IsFull() {
		t.Errorf("IsFull() = %t, want %t", true, false)
	}
	if old := buf.GetOldest(); old != "a" {
		t.Errorf("GetOldest() = %q, want %q", old, "a")
	}
	if n := buf.GetNewest(); buf.GetOldest() != n {
		t.Errorf("GetNewest() = %q, want %q", n, buf.GetOldest())
	}
	if snap := buf.Snapshot(); snap == nil {
		t.Error("Snapshot() = nil, want non-nil")
	}
	if snap := buf.Snapshot(); len(snap) != 1 {
		t.Errorf("len(Snapshot()) = %q, want '1'", len(snap))
	}
	if snap := buf.Snapshot(); snap[0] != "a" {
		t.Errorf("Snapshot()[0] = %q, want 'a'", snap[0])
	}
	buf.Push("b")
	if !buf.IsFull() {
		t.Errorf("IsFull() = %t, want %t", false, true)
	}
	buf.Push("c")
	if !buf.IsFull() {
		t.Errorf("IsFull() = %t, want %t", false, true)
	}
	if old := buf.GetOldest(); old != "b" {
		t.Errorf("GetOldest() = %q, want %q", old, "b")
	}
	if n := buf.GetNewest(); "c" != n {
		t.Errorf("GetNewest() = %q, want %q", n, "c")
	}
	if snap := buf.Snapshot(); snap == nil {
		t.Error("Snapshot() = nil, want non-nil")
	}
	if snap := buf.Snapshot(); len(snap) != 2 {
		t.Errorf("len(Snapshot()) = %q, want '2'", len(snap))
	}
	if snap := buf.Snapshot(); snap[0] != "b" {
		t.Errorf("Snapshot()[0] = %q, want 'b'", snap[0])
	}
}
