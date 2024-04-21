package internal

import "github.com/bwmarrin/discordgo"

// Gets a voice channel by user Id
func GetVoiceChannelID(guild *discordgo.Guild, userID string) string {
	for j := range guild.VoiceStates {
		if guild.VoiceStates[j].UserID == userID {
			return guild.VoiceStates[j].ChannelID
		}
	}
	return ""
}

func AllVoiceStatesExcept(guild *discordgo.Guild, userID string) []*discordgo.VoiceState {
	var states []*discordgo.VoiceState
	for _, v := range guild.VoiceStates {
		if v.UserID != userID {
			states = append(states, v)
		}
	}
	return states
}
