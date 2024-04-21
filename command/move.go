package command

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"maincharacter/internal"
)

func MoveCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	guild, err := s.State.Guild(i.GuildID)
	if err != nil {
		log.Fatal(err)
		return
	}
	userId := i.Member.User.ID

	channelID := internal.GetVoiceChannelID(guild, userId)
	if channelID == "" {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "JO du bist in keinem channel",
			},
		})
		return
	}
	states := internal.AllVoiceStatesExcept(guild, userId)
	for _, state := range states {
		err := s.GuildMemberMove(state.GuildID, state.UserID, &channelID)
		if err != nil {
			return
		}
	}
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Die APO RED crew ist wieder vereint",
		},
	})
}
