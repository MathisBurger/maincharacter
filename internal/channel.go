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

func GetActiveVoiceSession(guild *discordgo.Guild, s *discordgo.Session) *discordgo.VoiceConnection {
	for _, conn := range s.VoiceConnections {
		if conn.GuildID == guild.ID {
			return conn
		}
	}
	return nil
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
