//go:build profile

package main

import _ "net/http/pprof"

func profileEnabled() bool {
	return true
}
