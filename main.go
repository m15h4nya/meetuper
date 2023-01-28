package main

import (
	"github.com/m15h4nya/meetupper/config"
	"github.com/m15h4nya/meetupper/logger"
	"github.com/m15h4nya/meetupper/session"
)

func main() {
	log := logger.NewLogger()
	cfg := config.ParseConfig(log)
	b := session.CreateBot(cfg, log)
	b.StartSession()
	select {}
}
