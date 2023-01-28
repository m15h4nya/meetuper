package config

import (
	"os"

	toml "github.com/pelletier/go-toml"
	"go.uber.org/zap"
)

type Config struct {
	Session Session
	IDs     IDs
}

type Session struct {
	Token string `toml:"token"`
}

type IDs struct {
	AnnouncementChannelID string `toml:"annoincement_channel_id"`
	GuildID               string `toml:"guild_id"`
}

func ParseConfig(log *zap.SugaredLogger) *Config {
	cfg := &Config{}
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
