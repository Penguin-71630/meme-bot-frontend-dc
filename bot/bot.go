package bot

import (
	"context"
	"log"

	"github.com/Penguin-71630/meme-bot-frontend-dc/api"
	"github.com/Penguin-71630/meme-bot-frontend-dc/tracing"
	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var commands = []*discordgo.ApplicationCommand{
	{
		Name:        "ping",
		Description: "Check if bot is responsive",
	},
	{
		Name:        "greet",
		Description: "Get a friendly greeting",
	},
	{
		Name:        "echo",
		Description: "Bot repeats what you say",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "message",
				Description: "Message to echo",
				Required:    true,
			},
		},
	},
	{
		Name:        "web",
		Description: "Get a login link to the web interface",
	},
}

type Bot struct {
	session   *discordgo.Session
	apiClient *api.Client
}

func New() (*Bot, error) {
	// Create Discord session
	session, err := discordgo.New("Bot " + viper.GetString("discord-bot-token"))
	if err != nil {
		return nil, err
	}

	bot := &Bot{
		session:   session,
		apiClient: api.NewClient(),
	}

	// Register handlers
	bot.registerHandlers()

	// Set intents - only need guild messages for the ciallo listener
	session.Identify.Intents = discordgo.IntentsGuildMessages |
		discordgo.IntentsDirectMessages |
		discordgo.IntentsMessageContent

	return bot, nil
}

func (b *Bot) registerSlashCommands(ctx context.Context) error {
	for _, cmd := range commands {
		_, err := b.session.ApplicationCommandCreate(
			b.session.State.User.ID, "", cmd)
		if err != nil {
			tracing.Logger.Ctx(ctx).
				Error("failed to create command",
					zap.String("command", cmd.Name),
					zap.Error(err))
			return err
		}
	}

	return nil
}

func (b *Bot) clearSlashCommands(guildID string) error {
	commands, err := b.session.ApplicationCommands(b.session.State.User.ID, guildID)
	if err != nil {
		return err
	}

	for _, cmd := range commands {
		err := b.session.ApplicationCommandDelete(b.session.State.User.ID, guildID, cmd.ID)
		if err != nil {
			log.Printf("Failed to delete command %s: %v", cmd.Name, err)
		}
	}
	return nil
}

func (b *Bot) registerHandlers() {
	b.session.AddHandler(b.onReady)
	b.session.AddHandler(b.onMessageCreate)
	b.session.AddHandler(b.onInteractionCreate)
}

func (b *Bot) onReady(
	s *discordgo.Session,
	event *discordgo.Ready,
) {
	ctx := context.Background()
	tracing.Logger.Ctx(ctx).
		Info("logged in",
			zap.String("username", s.State.User.Username),
			zap.String("discriminator", s.State.User.Discriminator))

	// For development: set your guild ID here for instant updates
	// For production: use "" for global commands
	guildID := "1377176828833169468" // Replace with your Discord server ID for faster testing

	// clear slash commands
	if err := b.clearSlashCommands(guildID); err != nil {
		log.Printf("Error clearing slash commands: %v", err)
	}

	// Register slash commands
	if err := b.registerSlashCommands(ctx); err != nil {
		tracing.Logger.Ctx(ctx).
			Error("failed to register slash commands",
				zap.Error(err))
		return
	}

	// Set bot status
	err := s.UpdateGameStatus(0, "/ping to check status")
	if err != nil {
		tracing.Logger.Ctx(ctx).
			Error("failed to set status",
				zap.Error(err))
		return
	}
}

func (b *Bot) onInteractionCreate(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
) {
	if i.Type != discordgo.InteractionApplicationCommand {
		return
	}

	switch i.ApplicationCommandData().Name {
	case "ping":
		b.handleSlashPing(s, i)
	case "greet":
		b.handleSlashGreet(s, i)
	case "echo":
		b.handleSlashEcho(s, i)
	case "web":
		b.handleSlashWeb(s, i)
	}
}

func (b *Bot) Start() error {
	return b.session.Open()
}

func (b *Bot) Stop() {
	if b.session != nil {
		b.session.Close()
	}
}
