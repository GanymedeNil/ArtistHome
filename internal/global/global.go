package global

import (
	"ArtistHome/config"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	VP     *viper.Viper
	CONFIG config.Server
	LOGGER *zap.Logger
)

const (
	AuthId   = "id"
	AuthName = "name"
)
