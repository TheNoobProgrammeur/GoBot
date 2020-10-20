package commands_tests

import (
	c "myTest/src/bot/commands"
	"testing"
)

func TestPing(t *testing.T) {
	got := c.Ping()
	if got != "Pong" {
		t.Errorf("Ping = %s; want Pong", got)
	}
}
