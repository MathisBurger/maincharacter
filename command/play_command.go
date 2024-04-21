package command

import (
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io"
	"log"
	"maincharacter/internal"
	"os"
	"time"
)

var buffer [][]byte = make([][]byte, 0)

// GetPlayCommand gets the internal play command for the main character
func GetPlayCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	guild, err := s.State.Guild(i.GuildID)
	if conn := internal.GetActiveVoiceSession(guild, s); conn != nil {
		conn.Close()
		conn.Disconnect()
	}
	if err != nil {
		log.Println("Error getting guild data", err)
	}
	for j := range guild.VoiceStates {
		if guild.VoiceStates[j].Member.User.ID == i.Member.User.ID {
			buffer, err := loadSound()
			if err != nil {
				log.Println("Error loading sound", err)
			}
			err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Jo APO RED der maker master erzählt dir jetzt von seinem Leben",
				},
			})
			if err != nil {
				log.Println("Error interaction respond", err)
			}
			err = playSound(s, guild.ID, guild.VoiceStates[j].ChannelID, buffer)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Ayo du bist in keinem Channel dikka",
		},
	})
	if err != nil {
		log.Println("Error interaction respond", err)
	}
}

// Loads the sound
func loadSound() ([][]byte, error) {

	file, err := os.Open("audiofiles/maincharacter.dca")
	if err != nil {
		fmt.Println("Error opening dca file :", err)
		return nil, err
	}

	var opuslen int16

	for {
		// Read opus frame length from dca file.
		err = binary.Read(file, binary.LittleEndian, &opuslen)

		// If this is the end of the file, just return.
		if err == io.EOF || errors.Is(err, io.ErrUnexpectedEOF) {
			err := file.Close()
			if err != nil {
				return nil, err
			}
			return buffer, nil
		}

		if err != nil {
			fmt.Println("Error reading from dca file :", err)
			return nil, err
		}
		// Read encoded pcm from dca file.
		InBuf := make([]byte, opuslen)
		err = binary.Read(file, binary.LittleEndian, &InBuf)

		// Should not be any end of file errors
		if err != nil {
			fmt.Println("Error reading from dca file :", err)
			return nil, err
		}

		// Append encoded pcm data to the buffer.
		buffer = append(buffer, InBuf)
	}
}

// Plays a sound
func playSound(s *discordgo.Session, guildID, channelID string, buffer [][]byte) (err error) {
	vc, err := s.ChannelVoiceJoin(guildID, channelID, false, false)
	if err != nil {
		return err
	}

	// Sleep for a specified amount of time before playing the sound
	time.Sleep(250 * time.Millisecond)

	// Start speaking.
	err = vc.Speaking(true)
	if err != nil {
		return err
	}

	// Send the buffer data.
	for _, buff := range buffer {
		vc.OpusSend <- buff
	}

	// Stop speaking
	err = vc.Speaking(false)
	if err != nil {
		return err
	}

	// Sleep for a specificed amount of time before ending.
	time.Sleep(250 * time.Millisecond)

	// Disconnect from the provided voice channel.
	err = vc.Disconnect()
	if err != nil {
		return err
	}

	return nil
}
