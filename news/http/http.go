package http

//func Init() *gin.Engine {
//	consulReg := consul.NewRegistry(
//		registry.Addrs("192.168.1.103"),
//	)
//
//	microService := micro.NewService(
//		micro.Name("go-news"),
//	)
//	microService.Init()
//
//	consulReg.GetService("go-news-user-grpc")
//
//	gin.SetMode(gin.DebugMode)
//
//	r := gin.New()
//
//	userGroup := r.Group("/news")
//
//	userGroup.GET("/list", func(c *gin.Context) {
//		tokenCheckRequest := new(pb.TokenCheckRequest)
//
//		c.AbortWithStatusJSON(
//			http.StatusOK,
//			gin.H{
//				"user_id": "",
//			})
//		return
//	})
//
//	return r
//}
