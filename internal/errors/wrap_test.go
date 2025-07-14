package errorwraps

import (
	"errors"
	"testing"
)

var ErrSentinel = errors.New("sentinel")

func TestWrapIf(t *testing.T) {
	cases := []struct {
		name string
		inErr error
		wantIs bool //expect errors.Is(..., ErrSentinel)
	}{
		{"nil returns nil", nil, false},
		{"wrap sentinel", ErrSentinel, true},
	}

	for _,tc := range cases {
		got := WrapIf(tc.inErr, "ctx")
		if tc.inErr == nil && got != nil {
			t.Fatalf("%s: expected nil got %v", tc.name, got)
		}
		if tc.wantIs && !errors.Is(got, ErrSentinel) {
			t.Fatalf("%s: sentinel lost in chain", tc.name)
		}
	}
}
