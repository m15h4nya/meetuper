package handlers

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
	_ "time/tzdata"

	"github.com/bwmarrin/discordgo"
	"github.com/m15h4nya/meetupper/tools"
)

func (h Handler) createMeetup(s *discordgo.Session, i *discordgo.InteractionCreate) {
	h.Queuer.NewDraft(i.Member.User.ID)

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseModal,
		Data: DescriptionModal,
	})
	if err != nil {
		h.handlerError(err, i.Member.User.ID, s, i)
		return
	}
}

func (h Handler) descriptionModal(s *discordgo.Session, i *discordgo.InteractionCreate) {
	infoRegExp := regexp.MustCompile(`->.*<-`)
	weekRegEpx := regexp.MustCompile(`\d*w`)
	dayRegEpx := regexp.MustCompile(`\d*d`)
	hourRegEpx := regexp.MustCompile(`\d*h`)

	raw := i.ModalSubmitData().Components[0].(*discordgo.ActionsRow).Components[0].(*discordgo.TextInput)
	data := infoRegExp.FindAllString(raw.Value, 4)
	for i, line := range data {
		asRune := []rune(line)
		data[i] = string(asRune[2 : len(asRune)-2])
	}

	loc, _ := time.LoadLocation("Local")
	start, err := time.ParseInLocation("15:04 2/1/2006", data[2], loc)
	if err != nil {
		h.handlerError(err, i.Member.User.ID, s, i)
		return
	}

	hours := hourRegEpx.FindString(data[3])
	interval, err := time.ParseDuration("0")
	if err != nil {
		h.handlerError(err, i.Member.User.ID, s, i)
		return
	}
	if hours != "" {
		hoursCount, err := time.ParseDuration(hours)
		if err != nil {
			h.handlerError(err, i.Member.User.ID, s, i)
			return
		}
		interval += hoursCount
	}
	weeks := weekRegEpx.FindString(data[3])
	weeksInDays := 0
	if weeks != "" {
		asRune := []rune(weeks)
		weeksCount, err := strconv.Atoi(string(asRune[:len(asRune)-1]))
		if err != nil {
			h.handlerError(err, i.Member.User.ID, s, i)
			return
		}
		weeksInDays = weeksCount * 7
	}
	days := dayRegEpx.FindString(data[3])
	daysCount := 0
	if days != "" {
		asRune := []rune(days)
		daysCount, err = strconv.Atoi(string(asRune[:len(asRune)-1]))
		if err != nil {
			h.handlerError(err, i.Member.User.ID, s, i)
			return
		}
	}
	daysCount += weeksInDays
	daysInhours, err := time.ParseDuration(fmt.Sprint(daysCount*24) + "h")
	if err != nil {
		h.handlerError(err, i.Member.User.ID, s, i)
		return
	}
	interval += daysInhours

	draft := h.Queuer.GetDraft(i.Member.User.ID)
	draft.Name = data[0]
	draft.Message = data[1]
	draft.Start = start       // data[2]
	draft.Interval = interval // data[3]
	h.Queuer.UpdateDraft(i.Member.User.ID, draft)
	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
		Data: MentionableSelect,
	})
	if err != nil {
		h.handlerError(err, i.Member.User.ID, s, i)
		return
	}
}

func (h Handler) mentionableSelect(s *discordgo.Session, i *discordgo.InteractionCreate) {
	raw := i.MessageComponentData().Resolved

	draft := h.Queuer.GetDraft(i.Member.User.ID)
	draft.Users = raw
	h.Queuer.UpdateDraft(i.Member.User.ID, draft)
	h.Queuer.PushDraft(i.Member.User.ID)
	isDisabled := true
	if len(h.Queuer.GetAllMeetups()) > 0 {
		isDisabled = false
	}
	baseMessage := &discordgo.InteractionResponseData{
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
						Disabled: isDisabled,
						CustomID: "delete_meetup_button",
					},
				},
			},
		},
	}
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
		Data: baseMessage,
	})
	if err != nil {
		h.handlerError(err, i.Member.User.ID, s, i)
		return
	}
}

func (h Handler) channelSelectButton(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
		Data: ChannelSelect,
	})
	if err != nil {
		h.handlerError(err, i.Member.User.ID, s, i)
		return
	}
}

func (h Handler) channelSelect(s *discordgo.Session, i *discordgo.InteractionCreate) {
	raw := i.MessageComponentData().Resolved
	channel := tools.GetKeys(raw.Channels)
	h.Cfg.IDs.AnnouncementChannelID = channel[0]
	isDisabled := true
	if len(h.Queuer.GetAllMeetups()) > 0 {
		isDisabled = false
	}
	baseMessage := &discordgo.InteractionResponseData{
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
						Disabled: isDisabled,
						CustomID: "delete_meetup_button",
					},
				},
			},
		},
	}
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
		Data: baseMessage,
	})
	if err != nil {
		h.handlerError(err, i.Member.User.ID, s, i)
		return
	}
}

func (h Handler) deleteMeetupButton(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := []discordgo.SelectMenuOption{}
	for _, val := range h.Queuer.GetAllMeetups() {
		options = append(options, discordgo.SelectMenuOption{
			Label:       val.Name,
			Value:       val.Name,
			Description: val.Message,
		})
	}

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
		Data: &discordgo.InteractionResponseData{
			CustomID: "MeetupDelete",
			Content:  "Choose meetups to delete",
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.SelectMenu{
							CustomID:    "meetup_delete",
							MenuType:    discordgo.StringSelectMenu,
							Options:     options,
							MinValues:   &minValuesChannel,
							MaxValues:   len(options),
							Placeholder: "Click here",
						},
					},
				},
			},
		},
	})
	if err != nil {
		h.handlerError(err, i.Member.User.ID, s, i)
		return
	}
}

func (h Handler) deleteMeetup(s *discordgo.Session, i *discordgo.InteractionCreate) {
	raw := i.MessageComponentData().Values
	for _, val := range raw {
		h.Queuer.DeleteFromQueue(val)
	}
	isDisabled := true
	if len(h.Queuer.GetAllMeetups()) > 0 {
		isDisabled = false
	}
	baseMessage := &discordgo.InteractionResponseData{
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
						Disabled: isDisabled,
						CustomID: "delete_meetup_button",
					},
				},
			},
		},
	}
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
		Data: baseMessage,
	})
	if err != nil {
		h.handlerError(err, i.Member.User.ID, s, i)
		return
	}
}

var (
	componentsHandlers = map[string]func(h Handler, s *discordgo.Session, i *discordgo.InteractionCreate){
		"create_meetup_button":        Handler.createMeetup,
		"description_modal":           Handler.descriptionModal,
		"mentionable_select":          Handler.mentionableSelect,
		"channel_select":              Handler.channelSelect,
		"delete_meetup_button":        Handler.deleteMeetupButton,
		"meetup_delete":               Handler.deleteMeetup,
		"change_notification_channel": Handler.channelSelectButton,
	}
)
