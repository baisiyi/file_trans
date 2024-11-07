package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/siyibai/file-transfer/config"
	fileServer "github.com/siyibai/file-transfer/gateway"
	pb "github.com/siyibai/file-transfer/pb"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
	"google.golang.org/grpc"
)

var (
	configPath = flag.String("config", "config/config.yaml", "path to config file")
)

func main() {

	flag.Parse()

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		panic(err)
	}

	// 1. 首先启动 gRPC 服务器
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.GRPCPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterFileServiceServer(grpcServer, fileServer.NewFileService())
	// 在后台运行 gRPC 服务器
	go func() {
		log.Printf("gRPC server starting on :50051")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// 创建 HTTP 网关
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err = pb.RegisterFileServiceHandlerFromEndpoint(
		ctx,
		mux,
		fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.GRPCPort),
		opts,
	)
	if err != nil {
		log.Fatal(err)
	}

	// 修改 CORS 配置
	handler := cors.New(cors.Options{
		AllowedOrigins: cfg.Cors.AllowedOrigins,
		AllowedMethods: cfg.Cors.AllowedMethods,
		AllowedHeaders: []string{
			"Content-Type",
			"Authorization",
			"X-Requested-With",
			"Accept",
			"Origin",
		},
		ExposedHeaders: []string{
			"Content-Length",
			"Content-Disposition",
		},
		AllowCredentials: true,
		Debug:            true, // 开启调试日志
	}).Handler(mux)

	// 添加日志中间件
	loggingHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})

	// 启动 HTTP 服务
	addr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.HTTPPort)
	log.Printf("HTTP GateWay starting on %s", addr)
	if err := http.ListenAndServe(addr, loggingHandler); err != nil {
		log.Fatal(err)
	}
}
