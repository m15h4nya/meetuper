package handlers

import (
	"github.com/bwmarrin/discordgo"
)

func (h *Handler) Ready(s *discordgo.Session, m *discordgo.Ready) {
	msg := &discordgo.MessageSend{
		Content: "Test message",
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					// discordgo.SelectMenu{
					// 	MenuType:  discordgo.UserSelectMenu,
					// 	CustomID:  "UsersToMention",
					// 	MaxValues: 25,
					// },
					discordgo.Button{
						Label:    "Create meetup",
						Style:    discordgo.PrimaryButton,
						Disabled: false,
						CustomID: "create_meetup_button",
					},
				},
			},
		},
	}

	s.ChannelMessageSendComplex(h.Cfg.IDs.AnnouncementChannelID, msg)
}
