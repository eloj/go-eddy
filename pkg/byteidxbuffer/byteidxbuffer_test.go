package byteidxbuffer

import (
	"testing"
	"reflect"
)

func TestNew(t *testing.T) {
	c := NewByteIdxBuffer()

	if c.Len() != 0 {
		t.Errorf("New buffer is not length zero")
	}

	c.AddString("1234")
	c.AddString("56")
	c.AddString("789")

	if c.Len() != 3 {
		t.Errorf("Invalid length of buffer")
	}

	if got := c.GetString(0); got != "1234" {
		t.Errorf("Expected [0] to be '1234', '%s'", got)
	}

	if got := c.GetSlice(2); !reflect.DeepEqual(got, []byte{55,56,57}) {
		t.Errorf("Expected [2] to be '[]{55,56,57}', got '%v'", got)
	}

	// Now some hackery to test that the panic() calls fire on OOB.

	expected_recovers := 4
	got_recovers := 0

	recoveryCounter := func() {
		if r := recover() ; r != nil {
			got_recovers++
		}
	}

	func() {
		defer recoveryCounter()
		t.Log("Testing underflow OOB")
		c.GetString(-1)
	}()

	func() {
		defer recoveryCounter()
		t.Log("Testing underflow OOB")
		c.GetSlice(-1)
	}()

	func() {
		defer recoveryCounter()
		t.Log("Testing OOB")
		c.GetString(3)
	}()

	func() {
		defer recoveryCounter()
		t.Log("Testing OOB")
		c.GetSlice(3)
	}()

	if expected_recovers != got_recovers {
		t.Errorf("Expect 2 panics, got %d", got_recovers)

	}
}
