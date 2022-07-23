package global

import (
	"github.com/hackerchai/threatbook-ip-web/internal/app/config"
	"go.uber.org/zap"
)

var (
	CONFIG config.Config
	Logger *zap.Logger
)
