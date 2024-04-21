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
		{
			Name:        "stop",
			Description: "Stops the song",
		},
		{
			Name:        "move",
			Description: "Moves all users to the pinging user",
		},
		{
			Name:        "was-rushen",
			Description: "Apored sagt dir was alles gerusht werden soll",
		},
	}
}

// Gets all command handlers that are registered
func GetCommandHandlers() map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){

		// The functionality of the play command
		"play":       GetPlayCommand,
		"stop":       StopCommand,
		"move":       MoveCommand,
		"was-rushen": RushCommand,
	}
}
