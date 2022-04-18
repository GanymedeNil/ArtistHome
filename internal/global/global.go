package global

import (
	"github.com/GanymedeNil/GoFrameworkBase/config"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	VP     *viper.Viper
	RDB    *redis.Client
	CONFIG config.Server
	LOGGER *zap.Logger
)

const (
	AuthId   = "id"
	AuthName = "name"
)

const (
	DevMode     = "dev"
	ReleaseMode = "release"
	TestMode    = "test"
)
