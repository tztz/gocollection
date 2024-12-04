package main

import "testing"

func TestMainFunc(t *testing.T) {
	main()
}

func TestExampleFunc(t *testing.T) {
	results := example()

	amount := 16
	if len(results) != amount {
		t.Errorf("Expected %v results, got %v", amount, len(results))
	}
}
