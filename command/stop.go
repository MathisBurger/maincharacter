package command

import "github.com/bwmarrin/discordgo"

// Stop command to stop current playback
func StopCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	for _, conn := range s.VoiceConnections {
		if conn.GuildID == i.GuildID {
			conn.Close()
			conn.Disconnect()
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Jo, ich hab meine geile Mukke ausgemacht meine Hasen. ApoRED der Allerechte out",
				},
			})
			return
		}
	}
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Du hörst meine geile Mukke doch grad gar nicht dikka",
		},
	})
}
