package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/kelseyhightower/envconfig"
	"log"
	"maincharacter/command"
	"maincharacter/internal"
	"os"
	"os/signal"
)

// Application entrypoint
func main() {
	var config internal.Config
	err := envconfig.Process("maincharacter", &config)
	if err != nil {
		log.Fatal(err)
	}
	discord, err := discordgo.New("Bot " + config.BotToken)
	if err != nil {
		log.Fatal(err)
	}
	discord.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})

	discord.Identify.Intents = discordgo.IntentsGuilds | discordgo.IntentsGuildMessages | discordgo.IntentsGuildVoiceStates

	//err = command.LoadSound()
	if err != nil {
		log.Fatal(err)
	}

	err = discord.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}

	err = discord.UpdateCustomStatus("Ist im Insi-Modus")
	if err != nil {
		log.Fatal(err)
	}

	commands := command.GetCommands()
	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
	for i, v := range commands {
		cmd, err := discord.ApplicationCommandCreate(discord.State.User.ID, "", v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}

	commandHandlers := command.GetCommandHandlers()
	discord.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})

	defer discord.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop
}
