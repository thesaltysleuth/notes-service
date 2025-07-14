package errorwraps

import "fmt"

// WrapIf wraps err with msg only when err != nil
// so callers can keep doing: return errorwraps.WrapIf(err, "reading cfg")

func WrapIf(err error, msg string) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%w: %s", err, msg)
}

