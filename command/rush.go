package command

import (
	"github.com/bwmarrin/discordgo"
	"math/rand"
)

var rushMessages = []string{
	"Alter, du musst A rushen! Keine Zeit zu verlieren, Hase! Zeig diesen Gegnern, was du drauf hast! Stürm den A-Bombenplatz mit vollem Einsatz und lass dich von nichts und niemandem aufhalten! Du bist der Boss, der die Richtung vorgibt, und alle anderen sollen dir folgen! Zeig ihnen, dass du der wahre Champion bist! Also los, mein Gönner, A steht für Action! Mach es krass und rush A!",
	"Alter, du musst B rushen! B steht für bombastisch geile Gewinne und Dinge, wie mein Range Rover Mansory. Also rein in B meine Hasen",
	"Alter, du musst Mid rushen dikka",
}

// Rush command that tells us where to rush in CS
func RushCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: rushMessages[rand.Intn(len(rushMessages))],
		},
	})
}
