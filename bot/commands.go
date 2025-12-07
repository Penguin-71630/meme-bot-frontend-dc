package bot

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func (b *Bot) onMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore messages from the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Check if message starts with prefix
	if !strings.HasPrefix(m.Content, b.config.Prefix) {
		return
	}

	// Parse command
	content := strings.TrimPrefix(m.Content, b.config.Prefix)
	args := strings.Fields(content)
	if len(args) == 0 {
		return
	}

	command := strings.ToLower(args[0])
	commandArgs := args[1:]

	// Route commands
	switch command {
	case "help":
		b.handleHelp(s, m)
	case "meme":
		b.handleMeme(s, m, commandArgs)
	case "upload":
		b.handleUpload(s, m, commandArgs)
	case "search":
		b.handleSearch(s, m, commandArgs)
	case "aliases":
		b.handleAliases(s, m)
	case "ping":
		b.handlePing(s, m)
	default:
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Unknown command: `%s`. Use `%shelp` for available commands.", command, b.config.Prefix))
	}
}

func (b *Bot) handleHelp(s *discordgo.Session, m *discordgo.MessageCreate) {
	embed := &discordgo.MessageEmbed{
		Title:       "🤖 MemeBot Commands",
		Description: "Available commands for the MemeBot",
		Color:       0x00ff00,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   fmt.Sprintf("%shelp", b.config.Prefix),
				Value:  "Show this help message",
				Inline: false,
			},
			{
				Name:   fmt.Sprintf("%smeme <alias>", b.config.Prefix),
				Value:  "Get a meme by alias",
				Inline: false,
			},
			{
				Name:   fmt.Sprintf("%ssearch <keyword>", b.config.Prefix),
				Value:  "Search for memes by keyword",
				Inline: false,
			},
			{
				Name:   fmt.Sprintf("%supload <url> <aliases...>", b.config.Prefix),
				Value:  "Upload a meme with aliases",
				Inline: false,
			},
			{
				Name:   fmt.Sprintf("%saliases", b.config.Prefix),
				Value:  "List all available aliases",
				Inline: false,
			},
			{
				Name:   fmt.Sprintf("%sping", b.config.Prefix),
				Value:  "Check if bot is responsive",
				Inline: false,
			},
		},
	}

	s.ChannelMessageSendEmbed(m.ChannelID, embed)
}

func (b *Bot) handleMeme(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	if len(args) == 0 {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Usage: `%smeme <alias>`", b.config.Prefix))
		return
	}

	alias := strings.Join(args, " ")

	// TODO: Implement API call to get meme by alias
	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("🔍 Searching for meme with alias: `%s`\n(API integration pending)", alias))
}

func (b *Bot) handleSearch(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	if len(args) == 0 {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Usage: `%ssearch <keyword>`", b.config.Prefix))
		return
	}

	keyword := strings.Join(args, " ")

	// TODO: Implement API call to search memes
	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("🔍 Searching for: `%s`\n(API integration pending)", keyword))
}

func (b *Bot) handleUpload(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	// TODO: Implement meme upload functionality
	s.ChannelMessageSend(m.ChannelID, "📤 Upload functionality coming soon!\n(API integration pending)")
}

func (b *Bot) handleAliases(s *discordgo.Session, m *discordgo.MessageCreate) {
	// TODO: Implement API call to get all aliases
	s.ChannelMessageSend(m.ChannelID, "📋 Fetching all aliases...\n(API integration pending)")
}

func (b *Bot) handlePing(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "🏓 Pong! Bot is online.")
}
