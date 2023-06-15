package session

import (
	"github.com/bwmarrin/discordgo"
	"github.com/m15h4nya/meetupper/config"
	containers "github.com/m15h4nya/meetupper/heap"
	"github.com/m15h4nya/meetupper/meetup"
	hndlrs "github.com/m15h4nya/meetupper/session/handlers"
	"github.com/m15h4nya/meetupper/tools"
	"go.uber.org/zap"
)

type Bot struct {
	*discordgo.Session
	queuer *meetup.MeetupQueuer
	log    *zap.SugaredLogger
	meetup <-chan meetup.Meetup
	cfg    *config.Config
}

func CreateBot(cfg *config.Config, log *zap.SugaredLogger) Bot {
	ch := make(chan meetup.Meetup, 10)
	queue := containers.NewHeap(
		func(a, b meetup.Meetup) bool {
			return b.Start.After(a.Start)
		})
	meetupQueuer := meetup.NewMeetupQueuer(ch, &queue)

	var err error
	session, err := discordgo.New("Bot " + cfg.Session.Token)
	if err != nil {
		log.Fatal(err)
	}

	session.StateEnabled = true

	bot := Bot{
		Session: session,
		queuer:  &meetupQueuer,
		log:     log,
		meetup:  ch,
		cfg:     cfg,
	}

	handler := &hndlrs.Handler{Cfg: cfg, Log: log, Queuer: &meetupQueuer, MainMessage: make([]string, 2, 2)}
	handlers := []interface{}{
		handler.Ready,
		handler.InteractionHandler,
	}

	AddHandlers(bot.Session, handlers)
	return bot
}

func (b *Bot) StartSession() {
	err := b.Open()
	if err != nil {
		b.log.Errorf("StartSession(): %s", err)
	}
	go func() {
		for {
			meetup := <-b.meetup
			msgText := ""
			for _, id := range tools.GetKeys(meetup.Users.Members) {
				msgText += "<@" + id + "> "
			}
			for _, id := range tools.GetKeys(meetup.Users.Roles) {
				msgText += "<@&" + id + "> "
			}
			message := discordgo.MessageSend{
				Content: msgText,
				Embeds: []*discordgo.MessageEmbed{
					{
						Type:        discordgo.EmbedTypeArticle,
						Title:       meetup.Name,
						Description: meetup.Message,
					},
				},
			}

			if meetup.Interval != 0 {
				b.queuer.PushMeetup(meetup.AddDuration(meetup.Interval))
			}

			_, err := b.ChannelMessageSendComplex(b.cfg.IDs.AnnouncementChannelID, &message)
			if err != nil {
				b.log.Error(err)
			}
		}
	}()
	go b.queuer.RunMeetupQueue()
	b.log.Info("Bot is up")
}

func (b *Bot) StopSession() {
	err := b.Session.Close()
	if err != nil {
		b.log.Errorf("StopSession(): %s", err)
	}
}

func AddHandlers(s *discordgo.Session, handlers []interface{}) {
	for _, handler := range handlers {
		s.AddHandler(handler)
	}
}
