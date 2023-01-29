package handlers

import "github.com/bwmarrin/discordgo"

var (
	componentsHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"create_meetup_button": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseModal,
				Data: CreateMeetupButtonResponseData,
			})
			if err != nil {
				panic(err)
			}
			
		},
	}
	commandsHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){}
	// MeetupNameTextModal = &discordgo.InteractionResponse{
	// 	Type: discordgo.InteractionResponseModal,
	// 	Data: &discordgo.InteractionResponseData{
	// 		CustomID: "MeetupNameTextModalResponseData",
	// 		Title:    "Name for meetup notification",
	// 		Components: []discordgo.MessageComponent{
	// 			discordgo.ActionsRow{
	// 				Components: []discordgo.MessageComponent{
	// 					discordgo.TextInput{
	// 						CustomID: "MeetupNameTextModal",
	// 						Label:    "Name",
	// 						Style:    discordgo.TextInputParagraph,
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
	// CreateMeetupButton = &discordgo.Button{}
)

func 