//go:build !darwin && !linux && !windows

package main

func founder() string {
	return "Unknown"
}
