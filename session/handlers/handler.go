package handlers

import (
	"time"

	"github.com/m15h4nya/meetupper/config"
	"go.uber.org/zap"
)

type Handler struct {
	Cfg    *config.Config
	Log    *zap.SugaredLogger
}


