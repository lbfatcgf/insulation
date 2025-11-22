package admin

import (
	"fmt"
	"time"

	"insulation/server/admin/internal/controller"
	"insulation/server/admin/internal/service"
	"insulation/server/base/pkg/args"
	"insulation/server/base/pkg/config"
	"insulation/server/base/pkg/logger"
	redisutil "insulation/server/base/pkg/redis_util"
)

func Start() {
	time.LoadLocation("Asia/Shanghai")
	if args.ConfigPath == nil || args.ConfigName == nil {
		config.DefaultInitialize()
	} else {
		config.Initialize(*args.ConfigPath, *args.ConfigName)
	}
	defer logger.CloseAllLog()
	if args.Initialize != nil && *args.Initialize {
		service.InitSys()
		return
	}
	err := redisutil.InitRedis()
	if err != nil {
		panic(err)
	}
	// fmt.Printf("%v\n", string(jsonp_pretty.Pretty(config.Global())))
	controller.Start(fmt.Sprintf(`:%d`, config.Global().Web.Port))
}
