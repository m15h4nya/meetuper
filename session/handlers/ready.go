package handlers

import (
	"github.com/bwmarrin/discordgo"
)

func (h *Handler) Ready(s *discordgo.Session, m *discordgo.Ready) {
	msg := &discordgo.MessageSend{
		Content: "Create meetup",
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
				},
			},
		},
	}
	st, err := s.ChannelMessageSendComplex(h.Cfg.IDs.AnnouncementChannelID, msg)
	if err != nil {
		panic(err)
	}
	h.MainMessage[0] = st.ChannelID
	h.MainMessage[1] = st.ID
}
