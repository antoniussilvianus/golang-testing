package main

import (
	"testing"
)

func TestPrintMessage(t *testing.T) {
	expected := "Hello, this is Anton with ID 10122197\n"
	if got := PrintMessage(); got != expected {
		t.Errorf("PrintMessage() = %q, want %q", got, expected)
	}
}
