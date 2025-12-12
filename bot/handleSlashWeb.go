package bot

import (
	"context"

	"github.com/Penguin-71630/meme-bot-frontend-dc/tracing"
	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

func (b *Bot) handleSlashWeb(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
) {
	var userID string
	if i.Member != nil {
		userID = i.Member.User.ID
	} else if i.User != nil {
		userID = i.User.ID
	}

	// Call backend API
	loginURL, err := b.apiClient.PostGenLoginURL(userID)
	if err != nil {
		tracing.Logger.Ctx(context.Background()).
			Error("failed to generate login url",
				zap.Error(err))
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "❌ Failed to generate login URL",
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		})
		return
	}

	content := "🔗 **Click here to access the web page:**\n"+    
		loginURL + "\n\n"
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: content,
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
}
