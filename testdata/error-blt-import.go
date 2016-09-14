// Test that importing log will error.

// Package pkg ...
package pkg

import "log" // MATCH /Do not import log directly, use common/log instead. If necessary name import std_log/

var _ log.Log
