package handlers

import "github.com/bwmarrin/discordgo"

func (h *Handler) createMeetup(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseModal,
		Data: &discordgo.InteractionResponseData{
			CustomID: "Modal",
			Title:    "Text for meetup notification",
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.TextInput{
							CustomID: "TextInput",
							Label:    "Label",
							Style:    discordgo.TextInputParagraph,
						},
					},
				},
			},
		},
	})
	if err != nil {
		panic(err)
	}
}
