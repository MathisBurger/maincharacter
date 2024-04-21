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

	// ends current sessions on guild
	if conn := internal.GetActiveVoiceSession(guild, s); conn != nil {
		conn.Close()
		conn.Disconnect()
	}
	if err != nil {
		log.Println("Error getting guild data", err)
	}

	// Finds current user in voice channel
	for j := range guild.VoiceStates {

		// user is in voice channel
		if guild.VoiceStates[j].Member != nil && guild.VoiceStates[j].Member.User.ID == i.Member.User.ID {
			var buffer [][]byte = make([][]byte, 0)
			options := i.ApplicationCommandData().Options

			// sets the boost version dependent on data
			if len(options) == 1 {
				option := options[0].IntValue()
				switch option {
				case 1:
					buffer, err = loadSound("audiofiles/boost_1.dca")
					if err != nil {
						log.Println("Error loading sound", err)
					}
					break
				case 2:
					buffer, err = loadSound("audiofiles/boost_2.dca")
					if err != nil {
						log.Println("Error loading sound", err)
					}
					break
				default:
					buffer, err = loadSound("audiofiles/maincharacter.dca")
					if err != nil {
						log.Println("Error loading sound", err)
					}
					break
				}
			} else {
				buffer, err = loadSound("audiofiles/maincharacter.dca")
				if err != nil {
					log.Println("Error loading sound", err)
				}
			}

			// Sends now playing message
			err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Jo APO RED der maker master erzählt dir jetzt von seinem Leben",
				},
			})
			if err != nil {
				log.Println("Error interaction respond", err)
			}

			// Plays sound on voice channel
			err = playSound(s, guild.ID, guild.VoiceStates[j].ChannelID, buffer)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	// Send user not in voice channel message
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
func loadSound(version string) ([][]byte, error) {

	file, err := os.Open(version)
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
