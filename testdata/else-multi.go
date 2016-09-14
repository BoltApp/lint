// Test of return+else warning; should not trigger on multi-branch if/else.
// OK

// Package pkg ...
package pkg

import std_log "log"

func f(x int) bool {
	if x == 0 {
		std_log.Print("x is zero")
	} else if x > 0 {
		return true
	} else {
		std_log.Printf("non-positive x: %d", x)
	}
	return false
}
