package admin

import (
	"fmt"

	"insulation/server/admin/internal/web"
	"insulation/server/base/pkg/args"
	"insulation/server/base/pkg/config"
	"insulation/server/base/pkg/logger"
	redisutil "insulation/server/base/pkg/redis_util"
)

func Start() {
	if args.ConfigPath == nil || args.ConfigName == nil {
		config.DefaultInitialize()
	} else {
		config.Initialize(*args.ConfigPath, *args.ConfigName)
	}
	err := redisutil.InitRedis()
	if err != nil {
		panic(err)
	}
	defer logger.CloseAllLog()
	// fmt.Printf("%v\n", string(jsonp_pretty.Pretty(config.Global())))
	web.Start(fmt.Sprintf(`:%d`, config.Global().Web.Port))
}
