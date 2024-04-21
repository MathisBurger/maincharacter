package message

import (
	"github.com/bwmarrin/discordgo"
	"math/rand"
	"strings"
)

var keywords = []string{
	"burgi",
	"red",
	"insi",
	"hasen",
	"hasis",
	"broke",
	"broki",
	"fick",
	"cs",
	"rush",
	"mid",
	"minecraft",
	"fortnite",
}

var responses = []string{
	"Bist du broki moki oder was?",
	"Ayo, mach dir doch mal APO RED Burgis",
	"Willst du mich ficken?",
	"Steuern meuern",
	"Broki moki",
	"Bin mies am Steuern sparen",
	"Bin im Insi-Modus dikka",
	"Wer bist du, dass du so mit APO RED dem Allerechten redest?",
	"Hallo meine kleinen Hasis",
	"RED is back",
	"100k in meinem Prada bag",
	"Fahre range",
	"Meine Karre ist ein Batmobil digga",
	"Heute neuer Prank auf meinem Main Channel, weil ich der Maincharacter bin dikka",
}

// Respondes to a message
func MessageResponder(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	if strings.ToLower(m.Content) == "rush a" {
		s.ChannelMessageSend(m.ChannelID, "Alter, Rush A? Digga, das ist voll mein Ding! Einfach rein in die Action, volle Power voraus! Keine Zeit für Smalltalk, wir ballern uns durch, als ob es kein Morgen gäbe. Krass, wie das Adrenalin durch meine Adern pumpt, wenn ich mit Vollgas vorpresche! Echt jetzt, das ist der ultimative Kick! Wer kommt mit? Let's go, meine Gönner!")
	}
	if strings.ToLower(m.Content) == "rush b" {
		s.ChannelMessageSend(m.ChannelID, "Yo, Alter! Rush B, Digga? Das ist ja mal echt eine krasse Ansage! B wie Beastmode, Baby! Keine Zeit zu verlieren, wir stürmen den B-Bombenplatz wie eine wild gewordene Horde. Keine Angst, keine Gnade - wir machen kurzen Prozess! Hase, halt dich bereit, denn wir rocken den B-Spot und lassen nichts und niemanden stehen! B ist für Badass, und das werden wir allen zeigen! Also, lasst uns B zum Beben bringen, Habibis!")
	}
	for _, keyword := range keywords {
		if strings.Contains(m.Content, keyword) {
			s.ChannelMessageSend(m.ChannelID, responses[rand.Intn(len(responses))])
		}
	}
}
