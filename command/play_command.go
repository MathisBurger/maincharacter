package command

import (
	"encoding/binary"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io"
	"log"
	"os"
	"time"
)

var buffer = make([][]byte, 0)

// GetPlayCommand gets the internal play command for the main character
func GetPlayCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	guild, err := s.State.Guild(i.GuildID)
	if err != nil {
		log.Println("Error getting guild data", err)
	}
	for j := range guild.VoiceStates {
		if guild.VoiceStates[j].Member.User.ID == i.Member.User.ID {

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

func loadSound() error {

	file, err := os.Open("maincharacter.dca")
	if err != nil {
		fmt.Println("Error opening dca file :", err)
		return err
	}

	var opuslen int16

	for {
		// Read opus frame length from dca file.
		err = binary.Read(file, binary.LittleEndian, &opuslen)

		// If this is the end of the file, just return.
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			err := file.Close()
			if err != nil {
				return err
			}
			return nil
		}

		if err != nil {
			fmt.Println("Error reading from dca file :", err)
			return err
		}

		// Read encoded pcm from dca file.
		InBuf := make([]byte, opuslen)
		err = binary.Read(file, binary.LittleEndian, &InBuf)

		// Should not be any end of file errors
		if err != nil {
			fmt.Println("Error reading from dca file :", err)
			return err
		}

		// Append encoded pcm data to the buffer.
		buffer = append(buffer, InBuf)
	}
}

func playSound(s *discordgo.Session, guildID, channelID string) (err error) {
	vc, err := s.ChannelVoiceJoin(guildID, channelID, false, true)
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
