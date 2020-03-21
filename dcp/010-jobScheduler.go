package dcp

import "time"

// Implement a job scheduler which takes in a function f and an integer n, and calls f after n milliseconds.

func schedule(f func(), n int) {
	executeAfter := func(function func(), delay time.Duration) {
		time.Sleep(delay)
		f()
	}

	var delay time.Duration = time.Duration(n) * time.Millisecond
	go executeAfter(f, delay)
}
