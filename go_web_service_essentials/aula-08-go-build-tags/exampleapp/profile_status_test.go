//go:build !profile

package main

import "testing"

func TestProfileDisabledByDefault(t *testing.T) {
	if profileEnabled() {
		t.Fatalf("sem a tag profile, profileEnabled() deveria retornar false")
	}
}
