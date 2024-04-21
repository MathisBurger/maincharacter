package command

import (
	"github.com/bwmarrin/discordgo"
)

// GetCommands Gets all commands that are registered as slash commands
func GetCommands() []*discordgo.ApplicationCommand {
	return []*discordgo.ApplicationCommand{
		{
			Name:        "play",
			Description: "Plays the song of the maincharacter",
		},
	}
}

func GetCommandHandlers() map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){

		// The functionality of the play command
		"play": GetPlayCommand,
	}
}
