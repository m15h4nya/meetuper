package session

import (
	"github.com/m15h4nya/meetupper/config"
	"github.com/m15h4nya/meetupper/meetup"
	"github.com/thethanos/go-containers/containers"

	"github.com/bwmarrin/discordgo"
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
	queue := containers.NewHeap[meetup.Meetup](
		func(a, b meetup.Meetup) bool {
			return b.Time.After(a.Time)
		})
	meetupQueuer := meetup.NewMeetupQueuer(ch, &queue)

	var err error
	session, err := discordgo.New("Bot " + cfg.Token)
	if err != nil {
		log.Fatal(err)
	}

	session.StateEnabled = false

	return Bot{
		Session: session,
		queuer:  &meetupQueuer,
		log:     log,
	}

	// handler := &hndlrs.Handler{Cfg: cfg, OptState: ""}
	// handlers := []interface{}{
	// 	handler.MessageCreate,
	// 	handler.MessageEdit,
	// 	handler.MessageDelete,
	// 	handler.MessageDeleteBulk,
	// 	handler.Ready,
	// }

	// hndlrs.AddHandlers(bot.Session, handlers)
}

func (b *Bot) StartSession() {
	err := b.Open()
	if err != nil {
		b.log.Errorf("StartSession(): %s", err)
	}

	b.Session.GuildIntegrationCreate()
}

func (b *Bot) StopSession() {
	err := b.Session.Close()
	if err != nil {
		b.log.Errorf("StopSession(): %s", err)
	}
}
