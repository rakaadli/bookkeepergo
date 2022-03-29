package main

import (
	"fmt"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Raka")

	if result != "Hello Raka" {
		// error
		t.Error("Result must be 'Hello Raka'")
	}

	fmt.Println("TestHelloWorldRaka Done")
}
