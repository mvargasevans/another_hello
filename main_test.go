package main

import "testing"

func TestGreet(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"no name", "", "Hello, World!"},
		{"with name", "Alice", "Hello, Alice!"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := greet(tt.input)
			if got != tt.expected {
				t.Errorf("greet(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}
