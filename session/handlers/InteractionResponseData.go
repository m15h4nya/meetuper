package handlers

import "github.com/bwmarrin/discordgo"

var ( // all the components that are shown to user
	BaseMessage = &discordgo.InteractionResponseData{
		CustomID: "base_message",
		Content:  "Create meetup",
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.Button{
						Label:    "Create meetup",
						Style:    discordgo.PrimaryButton,
						Disabled: false,
						CustomID: "create_meetup_button",
					},
					discordgo.Button{
						Label:    "Change notification channel",
						Style:    discordgo.PrimaryButton,
						Disabled: false,
						CustomID: "change_notification_channel",
					},
					discordgo.Button{
						Label:    "Delete notification",
						Style:    discordgo.PrimaryButton,
						Disabled: false,
						CustomID: "delete_meetup_button",
					},
				},
			},
		},
	}
	raw = `Meetup name: -><-
Description: -><-
Start date: ->hh:mm dd/mm/yyyy<-
Interval: ->1w1d1h<- (if needed)`
	DescriptionModal = &discordgo.InteractionResponseData{
		CustomID: "description_modal",
		Title:    "Text for meetup notification",
		Content:  "Enter the description",
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.TextInput{
						CustomID: "DescriptionModal",
						Label:    "Remind message",
						Value:    raw,
						Style:    discordgo.TextInputParagraph,
					},
				},
			},
		},
	}
	minValues         = 0
	maxValues         = 25
	MentionableSelect = &discordgo.InteractionResponseData{
		CustomID: "MentionableSelect",
		Content:  "Choose who to mention",
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.SelectMenu{
						CustomID:    "mentionable_select",
						MenuType:    discordgo.MentionableSelectMenu,
						MinValues:   &minValues,
						MaxValues:   maxValues,
						Placeholder: "Click here",
					},
				},
			},
		},
	}
	minValuesChannel = 1
	maxValuesChannel = 1
	ChannelSelect    = &discordgo.InteractionResponseData{
		CustomID: "ChannelSelect",
		Content:  "Choose who to mention",
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.SelectMenu{
						CustomID:    "channel_select",
						MenuType:    discordgo.ChannelSelectMenu,
						MinValues:   &minValuesChannel,
						MaxValues:   maxValuesChannel,
						Placeholder: "Click here",
					},
				},
			},
		},
	}
	InvalidInput = &discordgo.InteractionResponseData{
		CustomID: "InvalidInput",
		Content:  "**ERROR: InvalidInput**",
	}
)
