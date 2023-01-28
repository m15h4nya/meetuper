package session

import "github.com/bwmarrin/discordgo"

var (
	commands = []*discordgo.ApplicationCommand{
		{
			Name:        "Create",
			Description: "Create new meetup",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "Название митапа",
					Description: "Название митапа для ясности и удобства его редактирования",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "Время митапа в формате ...", //TODO добавить формат
					Description: "String option",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "string-option",
					Description: "String option",
					Required:    true,
				},
			},
		},
		{
			Name:        "Delete",
			Description: "Delete existing meetup",
		},
		{
			Name:        "Change",
			Description: "Change existring meetup",
		},
	}
)
