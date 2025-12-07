package bot

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// Message listener for "ciallo"
func (b *Bot) onMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore messages from the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Check if message is "ciallo" (case insensitive)
	if strings.ToLower(strings.TrimSpace(m.Content)) == "ciallo" {
		s.ChannelMessageSend(m.ChannelID, "Ciallo!")
	}
}

// Slash command handlers
func (b *Bot) handleSlashPing(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "pong",
		},
	})
}

func (b *Bot) handleSlashGreet(s *discordgo.Session, i *discordgo.InteractionCreate) {
	var username string
	if i.Member != nil {
		username = i.Member.User.Username
	} else if i.User != nil {
		username = i.User.Username
	} else {
		username = "Unknown"
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("Ciallo, %s!", username),
		},
	})
}

func (b *Bot) handleSlashEcho(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options
	if len(options) == 0 {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "No message provided!",
			},
		})
		return
	}

	message := options[0].StringValue()

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: message,
		},
	})
}
