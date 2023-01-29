package handlers

import "github.com/bwmarrin/discordgo"

var (
	CreateMeetupButtonResponseData = &discordgo.InteractionResponseData{
		CustomID: "DescriptionTextModalResponseData",
		Title:    "Text for meetup notification",
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.TextInput{
						CustomID: "DescriptionTextModal",
						Label:    "Text",
						Style:    discordgo.TextInputParagraph,
					},
				},
			},
		},
	}
)
