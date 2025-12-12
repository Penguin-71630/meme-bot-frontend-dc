package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func (b *Bot) handleSlashGreet(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
) {
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
