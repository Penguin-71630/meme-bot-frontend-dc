package bot

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

// Message listener for "ciallo"
func (b *Bot) onMessageCreate(
	s *discordgo.Session,
	m *discordgo.MessageCreate,
) {
	// Ignore messages from the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Check if message is "ciallo" (case insensitive)
	if strings.ToLower(strings.TrimSpace(m.Content)) == "ciallo" {
		s.ChannelMessageSend(m.ChannelID, "Ciallo!")
	}
}
