package test

import (
	"testing"

	"github.com/go-test/deep"
)

func CheckEqual(t *testing.T, actual, expected interface{}) {
	t.Helper()
	diff := deep.Equal(actual, expected)
	if diff == nil {
		return
	} else if len(diff) > 0 {
		for _, d := range diff {
			t.Log("\t -- \t", d)
		}
		t.Error("checkEqual failed.")
	}
}
