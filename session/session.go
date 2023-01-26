package session

import (
	hndlrs "discord_logger/botSession/handlers"

	"github.com/m15h4nya/meetupper/config"

	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

type Bot struct {
	*discordgo.Session
	log *zap.SugaredLogger
}

func (b *Bot) CreateSession(cfg *config.Config, log *zap.SugaredLogger) {
	handler := &hndlrs.Handler{Cfg: cfg, OptState: ""}
	handlers := []interface{}{
		handler.MessageCreate,
		handler.MessageEdit,
		handler.MessageDelete,
		handler.MessageDeleteBulk,
		handler.Ready,
	}

	var err error
	b.Session, err = discordgo.New("Bot " + handler.Cfg.Token)
	if err != nil {
		log.Fatal(err)
	}

	b.Session.StateEnabled = false
	hndlrs.AddHandlers(b.Session, handlers)
}

func (b *Bot) StartSession() {
	err := b.Open()
	if err != nil {
		b.log.Errorf("StartSession(): %s", err)
	}
}

func (b *Bot) StopSession() {
	err := b.Session.Close()
	if err != nil {
		b.log.Errorf("StopSession(): %s", err)
	}
}
