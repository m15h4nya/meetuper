package config

import (
	"os"

	toml "github.com/pelletier/go-toml"
	"go.uber.org/zap"
)

type Config struct {
	Token               string `toml:"token"`
	AnnouncementChannel string `toml:"annoincement_channel_id"`
	GuildID             string `toml:"guild_id"`
}

func ParseConfig(log *zap.SugaredLogger) *Session {
	cfg := &Session{}
	file, err := os.ReadFile("./config/config.toml")

	if err != nil {
		log.Fatal(err)
	}

	err = toml.Unmarshal(file, cfg)

	if err != nil {
		log.Fatal(err)
	}

	return cfg
}

func Contains(str string, sl []string) bool {
	for _, v := range sl {
		if str == v {
			return true
		}
	}
	return false
}
