package command

import (
	"github.com/bwmarrin/discordgo"
	"strconv"
	"strings"
	"time"
)

func CreateEventCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options
	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}
	startDateString := strings.Split(optionMap["start-date"].StringValue(), ".")
	if len(startDateString) != 3 {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Falsches Datum dikka",
			},
		})
		return
	}
	startDate := time.Date(
		safeParseInt(startDateString[2], time.Now().Year()+2),
		time.Month(safeParseInt(startDateString[1], 12)),
		safeParseInt(startDateString[0], 31),
		int(optionMap["start-hours"].IntValue()),
		int(optionMap["start-min"].IntValue()),
		0,
		0,
		time.UTC,
	)
	duration := optionMap["duration"].IntValue()
	endTime := startDate.Add(time.Duration(duration) * time.Hour)
	s.GuildScheduledEventCreate(i.GuildID, &discordgo.GuildScheduledEventParams{
		Name:               optionMap["name"].StringValue(),
		Description:        "Hammer geiles ApoRED event meine Hasen",
		ScheduledStartTime: &startDate,
		ScheduledEndTime:   &endTime,
		EntityType:         discordgo.GuildScheduledEventEntityTypeStageInstance,
		ChannelID:          i.ChannelID,
		PrivacyLevel:       discordgo.GuildScheduledEventPrivacyLevelGuildOnly,
	})
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Ich hab ein RED tastisches Event erstellt dikka",
		},
	})
}

func safeParseInt(s string, maxValue int) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 1
	}
	if i > maxValue {
		return maxValue
	}
	return i
}
