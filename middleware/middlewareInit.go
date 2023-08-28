package middleware

import (
	"github.com/gin-gonic/gin"
	"simple_douyin/log"
	"simple_douyin/middleware/ratelimit"
)

//func init() {
//	log.Log.Info("Init Middleware")
//	Apirouter := gin.Default().Group("/douyin")
//	InitMiddleware(Apirouter)
//}

func InitMiddleware(apiRouter *gin.RouterGroup) {

	// 初始化 logrus
	log.LogrusInit()

	////初始化 xss
	//xssMiddleware := xss.XssMw{}
	//apiRouter.Use(xssMiddleware.RemoveXss())

	//初始化 Redis
	//redis.InitRedis()

	// 初始化限流器
	rateControlHandler, err := ratelimit.RateControl()
	if err != nil {
		log.Log.Fatalf("Error setting up rate control middleware: %v", err)
	}

	// 检查rateControlHandler是否正确初始化
	if rateControlHandler == nil {
		log.Log.Fatalf("Error initializing rate control handler")
	}

	// 将速率限制中间件添加到路由器
	apiRouter.Use(rateControlHandler)
}
