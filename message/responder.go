package message

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

var keywords = []string{
	"burgi",
	"red",
	"insi",
}

func MessageResponder(s *discordgo.Session, m *discordgo.MessageCreate) {
	for _, keyword := range keywords {
		if strings.Contains(m.Content, keyword) {
			s.ChannelMessageSend(m.ChannelID, "Bist du broki moki oder was")
		}
	}
}
