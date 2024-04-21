package status

import (
	"github.com/bwmarrin/discordgo"
	"math/rand"
	"time"
)

var listeningStates = []string{
	"Photoshop",
	"Multi Millionär",
	"Babawagen",
	"Everyday Saturday",
	"Yalla Habibi",
	"Numero Uno",
	"Range Rover Mansory",
	"Billo",
}

var watchingStates = []string{
	"How to fake being rich?",
	"How to Steuern (für Anfänger)",
	"Finanzamt austricksen",
	"Muss ich Gewerbesteuern zahlen?",
}

var gameStates = []string{
	"Game of life",
	"Fortnite",
	"Counter-Strike 2",
}

var customStates = []string{
	"ist im Insi-Modus",
	"kocht Burgis",
	"ist Multi Millionär",
	"hat zu viel Bargeld",
	"ist der Maincharacter",
}

// Idle states of status updates
var idle0 = 0
var idle1 = 1

// Updates the status of main character bot
func UpdateStatusWorker(c chan *discordgo.Activity) {
	var activity *discordgo.Activity = nil
	for {
		stateType := rand.Intn(4)
		switch stateType {
		case 0:
			activity = &discordgo.Activity{Type: discordgo.ActivityTypeListening, Name: listeningStates[rand.Intn(len(listeningStates))]}
			break
		case 1:
			activity = &discordgo.Activity{Type: discordgo.ActivityTypeWatching, Name: watchingStates[rand.Intn(len(watchingStates))]}
			break
		case 2:
			activity = &discordgo.Activity{Type: discordgo.ActivityTypeGame, Name: gameStates[rand.Intn(len(gameStates))]}
			break
		case 3:
			activity = &discordgo.Activity{Type: discordgo.ActivityTypeCustom, State: customStates[rand.Intn(len(customStates))], Name: "custom status"}
			break

		}
		c <- activity
		i := 0

		// Sends every second empty activity
		// After 10 seconds the activity is reset
		for i < 10 {
			i++
			c <- activity
			time.Sleep(1 * time.Second)
		}
	}
}

func UpdateGuildDoNotDisrupt(s *discordgo.Session, c chan *discordgo.Activity) {
	var activity *discordgo.Activity = nil
	for {
		activity = <-c

		// Sets the idle states in dependence of
		// the activity type of the activity from channel
		idle := &idle1
		if activity.Type == discordgo.ActivityTypeCustom {
			idle = nil
		} else if activity.Type == discordgo.ActivityTypeListening {
			idle = &idle0
		}

		voiceConns := s.VoiceConnections
		if len(voiceConns) > 0 {
			s.UpdateStatusComplex(discordgo.UpdateStatusData{
				AFK:        false,
				Activities: []*discordgo.Activity{activity},
				Status:     string(discordgo.StatusDoNotDisturb),
				IdleSince:  idle,
			})
		} else {
			s.UpdateStatusComplex(discordgo.UpdateStatusData{
				AFK:        false,
				Activities: []*discordgo.Activity{activity},
				Status:     string(discordgo.StatusOnline),
				IdleSince:  idle,
			})
		}
	}
}
