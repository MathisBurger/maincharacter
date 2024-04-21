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

// Updates the status of main character bot
func UpdateStatusWorker(s *discordgo.Session) {

	for {
		stateType := rand.Intn(4)
		switch stateType {
		case 0:
			s.UpdateListeningStatus(listeningStates[rand.Intn(len(listeningStates))])
			break
		case 1:
			s.UpdateWatchStatus(1, watchingStates[rand.Intn(len(watchingStates))])
			break
		case 2:
			s.UpdateGameStatus(1, gameStates[rand.Intn(len(gameStates))])
			break
		case 3:
			s.UpdateCustomStatus(customStates[rand.Intn(len(customStates))])
			break

		}
		time.Sleep(5 * time.Second)
	}
}
