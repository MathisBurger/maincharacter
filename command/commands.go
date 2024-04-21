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
		{
			Name:        "create-event",
			Description: "Creates a new event",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "name",
					Description: "Der Name des Events",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "start-date",
					Description: "Startdatum im Format 1.1.1990",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "start-hours",
					Description: "Startzeit stunden",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "start-min",
					Description: "Startzeit minuten",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "duration",
					Description: "Dauer in Stunden",
					Required:    true,
				},
			},
		},
	}
}

// Gets all command handlers that are registered
func GetCommandHandlers() map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){

		// The functionality of the play command
		"play":         GetPlayCommand,
		"stop":         StopCommand,
		"move":         MoveCommand,
		"was-rushen":   RushCommand,
		"create-event": CreateEventCommand,
	}
}
