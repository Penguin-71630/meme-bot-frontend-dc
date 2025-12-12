package bot

import "github.com/bwmarrin/discordgo"

func (b *Bot) handleSlashEcho(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
) {
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
