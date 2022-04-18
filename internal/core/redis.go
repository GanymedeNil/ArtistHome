package core

import (
	"context"
	"fmt"

	"github.com/GanymedeNil/GoFrameworkBase/internal/global"

	"github.com/go-redis/redis/v8"
)

func Redis() {
	redis.SetLogger(&rdb{})
	opt, err := redis.ParseURL(fmt.Sprintf("redis://%s:%s@%s:%s/%s",
		global.CONFIG.Redis.Username,
		global.CONFIG.Redis.Password,
		global.CONFIG.Redis.Host,
		global.CONFIG.Redis.Port,
		global.CONFIG.Redis.DB))
	if err != nil {
		panic(err)
	}
	global.RDB = redis.NewClient(opt)
}

type rdb struct{}

func (r *rdb) Printf(ctx context.Context, format string, v ...interface{}) {
	global.LOGGER.Sugar().Infof(format, v)
}
