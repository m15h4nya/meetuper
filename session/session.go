package session

import (
	"github.com/bwmarrin/discordgo"
	"github.com/m15h4nya/meetupper/config"
	"github.com/m15h4nya/meetupper/meetup"
	hndlrs "github.com/m15h4nya/meetupper/session/handlers"
	"github.com/thethanos/go-containers/containers"
	"go.uber.org/zap"
)

type Bot struct {
	*discordgo.Session
	queuer *meetup.MeetupQueuer
	log    *zap.SugaredLogger
	cfg    *config.Config
}

func CreateBot(cfg *config.Config, log *zap.SugaredLogger) Bot {
	ch := make(chan meetup.Meetup, 10)
	queue := containers.NewHeap(
		func(a, b meetup.Meetup) bool {
			return b.Time.After(a.Time)
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
	}

	handler := &hndlrs.Handler{Cfg: cfg, Log: log}
	handlers := []interface{}{
		// handler.MessageCreate,
		// handler.MessageEdit,
		// handler.MessageDelete,
		// handler.MessageDeleteBulk,
		handler.Ready,
		handler.InteractionCreate,
	}

	AddHandlers(bot.Session, handlers)
	return bot
}

func (b *Bot) StartSession() {
	err := b.Open()
	if err != nil {
		b.log.Errorf("StartSession(): %s", err)
	}

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
