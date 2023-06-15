package handlers

import (
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/m15h4nya/meetupper/config"
	"github.com/m15h4nya/meetupper/meetup"
	"go.uber.org/zap"
)

type Handler struct {
	Cfg         *config.Config
	Log         *zap.SugaredLogger
	Queuer      *meetup.MeetupQueuer
	MainMessage []string
}

func (h Handler) handlerError(err error, id string, s *discordgo.Session, i *discordgo.InteractionCreate) {
	h.Queuer.DeleteDraft(id)
	h.Log.Error(err)
	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
		Data: InvalidInput,
	})
	if err != nil {
		h.Log.Error(err)
	}
	time.Sleep(time.Second * 5)

	baseMessageEditContent := "Create meetup"
	BaseMessageEdit := &discordgo.MessageEdit{
		Channel: h.MainMessage[0],
		ID:      h.MainMessage[1],
		Content: &baseMessageEditContent,
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
	_, err = s.ChannelMessageEditComplex(BaseMessageEdit)
	if err != nil {
		panic(err)
	}
}
