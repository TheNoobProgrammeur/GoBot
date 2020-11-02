package commands_tests

import (
	"github.com/bwmarrin/discordgo"
	c "myTest/src/bot/commands"
	"testing"
)


func TestPing(t *testing.T) {
	user := discordgo.User{}
	got := c.Ping(user)
	if got != "Pong" {
		t.Errorf("Ping = %s; want Pong", got)
	}
}
