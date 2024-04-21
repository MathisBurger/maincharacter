package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/kelseyhightower/envconfig"
	"log"
	"maincharacter/command"
	"maincharacter/internal"
	"maincharacter/message"
	"maincharacter/status"
	"os"
	"os/signal"
)

// Application entrypoint
func main() {

	// Load config
	var config internal.Config
	err := envconfig.Process("maincharacter", &config)
	if err != nil {
		log.Fatal(err)
	}

	// Create app
	discord, err := discordgo.New("Bot " + config.BotToken)
	if err != nil {
		log.Fatal(err)
	}

	discord.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})
	discord.AddHandler(message.MessageResponder)

	discord.Identify.Intents = discordgo.IntentsGuilds | discordgo.IntentsGuildMessages | discordgo.IntentsGuildVoiceStates | discordgo.IntentsMessageContent

	err = discord.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}

	go status.UpdateStatusWorker(discord)

	command.RegisterCommands(discord)

	defer discord.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop
}
