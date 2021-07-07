package main

import (
	"context"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	"github.com/sirupsen/logrus"
	pb "news/proto"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	consulReg := consul.NewRegistry(
		registry.Addrs("192.168.1.103"),
	)

	microService := micro.NewService(
		micro.Name("go-news"),
		micro.Registry(consulReg),
	)

	//servers, err := consulReg.GetService("go-news-user-grpc")
	//if err != nil {
	//	panic(err)
	//}
	//
	//if len(servers) == 0 {
	//	panic("Services find empty")
	//}
	//
	//next := selector.Random(servers)
	//node, err := next()
	//fmt.Println(node.Id)
	//fmt.Println(node.Address)
	//fmt.Println(node.Metadata)
	//conn, err := grpc.Dial(node.Address, grpc.WithInsecure())
	//if err != nil {
	//	grpclog.Fatal(err)
	//	logrus.Errorf("Start GRPC Error: %s", err.Error())
	//}

	//httpServer := http.Init()
	//
	//err := httpServer.Run(":8899")
	//if err != nil {
	//	panic("micro server register error!")
	//}

	microService.Init()
	microService.Run()
	userService := pb.NewUserService("go-news-user-grpc", microService.Client())

	req := &pb.TokenCheckRequest{
		JwtToken: "213213122",
	}
	ctx := context.Background()
	res, err := userService.TokenCheck(ctx, req)
	if err != nil {
		logrus.Errorf(err.Error())
	}
	logrus.Errorf("grpc res: %+v", res)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		switch <-c {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			return
		case syscall.SIGHUP:
		}
	}
}
