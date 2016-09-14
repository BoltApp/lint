// Test of return+else warning.

// Package pkg ...
package pkg

import std_log "log"

func f(x int) bool {
	if x > 0 {
		return true
	} else { // MATCH /if.*return.*else.*outdent/
		std_log.Printf("non-positive x: %d", x)
	}
	return false
}

func g(f func() bool) string {
	if ok := f(); ok {
		return "it's okay"
	} else { // MATCH /if.*return.*else.*outdent.*short.*var.*declaration/
		return "it's NOT okay!"
	}
}
