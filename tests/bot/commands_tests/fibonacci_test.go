package commands_tests_test

import (
	"myTest/src/bot/commands"
	"testing"
)

func TestFibonacci(t *testing.T) {
	got := commands.Fibonacci(44)
	if got != 701408733 {
		t.Errorf("Fibonacci(44) = %d; want 701408733", got)
	}
}
