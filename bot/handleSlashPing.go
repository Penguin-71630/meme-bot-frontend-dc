package bot

import "github.com/bwmarrin/discordgo"

func (b *Bot) handleSlashPing(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "pong",
		},
	})
}
