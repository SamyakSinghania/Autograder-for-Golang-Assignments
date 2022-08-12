package sampleTask

import (
	_ "golang.org/x/net/html"
)

var NotCompleted = "not-completed"
var Completed = "completed"

// Task should return the value of Completed variable.
func Task() string {
	return Completed
}
