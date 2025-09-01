package main

import (
	"testing"
)

func TestHello1(t *testing.T) {
	name := "World"
	expected := "Hello, World !"
	actual := Hello(name)

	if actual != expected {
		t.Errorf("Actual %q not equal to expected %q", actual, expected)
	}
}

func TestHello2(t *testing.T) {
	name := "<3"
	expected := "Hello, <3!"
	actual := Hello(name)

	if actual != expected {
		t.Errorf("Actual %q not equal to expected %q", actual, expected)
	}
}
