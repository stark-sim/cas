package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// 当有多个可执行文件时，应该将 main.go 放在 /cmd/myapp/ 目录下
// 而对于微服务形式来说，只需要放在根目录就够直观了

func main() {
	// 监听 TCP 端口
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8080))
	if err != nil {
		logrus.Fatalf("failed to listen: %v", err)
	}
	// gRPC 服务初始化
	// 要将业务注册进该服务中
	grpcServer := grpc.NewServer()
	// 同步信道监听结束信号
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGQUIT, syscall.SIGTERM)
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
	logrus.Printf("Staring server")
	err = grpcServer.Serve(lis)
	if err != nil {
		logrus.Fatalf("failed at start server: %v", err)
	}
	// 确保 Stop 协程执行完，主线程才能结束
	wg.Wait()
}
