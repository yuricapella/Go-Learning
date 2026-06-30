//go:build profile

package main

import "testing"

func TestProfileEnabledWithBuildTag(t *testing.T) {
	if !profileEnabled() {
		t.Fatalf("com a tag profile, profileEnabled() deveria retornar true")
	}
}
