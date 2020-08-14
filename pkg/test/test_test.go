package test

import (
	"testing"
)

func TestCheckEqual(t *testing.T) {
	// Yes, this is a pretty dumb test.

	data := map[string]int{"Hello": 42}

	expected := make(map[string]int)

	for k, v := range data {
		expected[k] = v
	}

	CheckEqual(t, data, expected)
}
