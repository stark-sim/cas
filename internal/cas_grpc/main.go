package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/stark-sim/cas/configs"
	"github.com/stark-sim/cas/internal/db"
	pb "github.com/stark-sim/cas/pkg/grpc/pb"
	"github.com/stark-sim/cas/pkg/grpc/servers"
	"github.com/stark-sim/cas/tools"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	var err error
	err = configs.InitLogger()
	err = configs.InitConfig()
	err = tools.Init()
	err = db.InitDB()
	if err != nil {
		return
	}
	// 监听 TCP 端口
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", configs.Conf.APIConfig.GrpcPort))
	if err != nil {
		logrus.Fatalf("failed to listen: %v", err)
	}
	// gRPC 服务初始化
	// 要将业务注册进该服务中
	grpcServer := grpc.NewServer()
	// Initialize the generated User service
	client := db.NewDBClient()
	svc := servers.UserServer{Client: client}
	// 注册 service 到 server 中
	pb.RegisterUserServiceServer(grpcServer, &svc)
	// 同步信道监听结束信号
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	// Sync 监听协程和主线程一致性
	wg := sync.WaitGroup{}
	wg.Add(1)
	// 启动监听协程
	go func() {
		s := <-sigCh
		logrus.Printf("got signal %v, attempint graceful shutdown", s)
		grpcServer.GracefulStop()
		wg.Done()
	}()
	// 启动服务，绑定在监听器上
	logrus.Printf("Staring server at :%d", configs.Conf.APIConfig.GrpcPort)
	err = grpcServer.Serve(lis)
	if err != nil {
		logrus.Fatalf("failed at start server: %v", err)
	}
	// 确保 Stop 协程执行完，主线程才能结束
	wg.Wait()
}
