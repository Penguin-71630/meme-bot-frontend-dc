package bot

import (
	"fmt"
	"log"

	"github.com/Penguin-71630/meme-bot-frontend-dc/api"
	"github.com/Penguin-71630/meme-bot-frontend-dc/config"
	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	session   *discordgo.Session
	config    *Config
	apiClient *api.Client
}

type Config struct {
	Token     string
	APIClient *api.Client
	Prefix    string
}

func New(cfg *config.Config) (*Bot, error) {
	// Create Discord session
	session, err := discordgo.New("Bot " + cfg.DiscordToken)
	if err != nil {
		return nil, fmt.Errorf("error creating Discord session: %w", err)
	}

	// Create API client
	apiClient := api.NewClient(cfg.APIBaseURL)

	bot := &Bot{
		session:   session,
		apiClient: apiClient,
		config: &Config{
			Token:     cfg.DiscordToken,
			APIClient: apiClient,
			Prefix:    cfg.BotPrefix,
		},
	}

	// Register handlers
	bot.registerHandlers()

	// Set intents
	session.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsDirectMessages | discordgo.IntentsMessageContent

	return bot, nil
}

func (b *Bot) registerHandlers() {
	b.session.AddHandler(b.onReady)
	b.session.AddHandler(b.onMessageCreate)
}

func (b *Bot) onReady(s *discordgo.Session, event *discordgo.Ready) {
	log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)

	// Set bot status
	err := s.UpdateGameStatus(0, fmt.Sprintf("%shelp for commands", b.config.Prefix))
	if err != nil {
		log.Printf("Error setting status: %v", err)
	}
}

func (b *Bot) Start() error {
	if err := b.session.Open(); err != nil {
		return fmt.Errorf("error opening connection: %w", err)
	}
	return nil
}

func (b *Bot) Stop() {
	if b.session != nil {
		b.session.Close()
	}
}
