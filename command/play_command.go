package command

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

// GetPlayCommand gets the internal play command for the main character
func GetPlayCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Hey there! Congratulations, you just executed your first slash command",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
}
