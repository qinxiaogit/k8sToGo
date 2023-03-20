package main

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	v1 "github.com/qinxiaogit/k8sToGo/mic/api/protobuf/blog/v1"
	"github.com/qinxiaogit/k8sToGo/mic/internal/blog"
	"github.com/qinxiaogit/k8sToGo/mic/internal/pkg/interceptor"
	"github.com/qinxiaogit/k8sToGo/mic/internal/pkg/jwt"
	"github.com/qinxiaogit/k8sToGo/mic/internal/pkg/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"context"

	//"github.com/qinxiaogit/k8sToGo/mic/internal/blog"
	//"github.com/qinxiaogit/k8sToGo/mic/internal/pkg/config"
	//"github.com/qinxiaogit/k8sToGo/mic/internal/pkg/log"
	"github.com/qinxiaogit/k8sToGo/mic/internal/pkg/config"
	flag "github.com/spf13/pflag"
)

var flagConfig = flag.String("config", "./configs/config.yaml", "path to config file")

func main() {
	logger := log.New()
	conf, err := config.Load(*flagConfig)
	if err != nil {
		logger.Fatal("load config failed", err)
	}
	blogService, err := InitServer(logger, conf)

	healthServer := health.NewServer()
	jwtManager := jwt.NewManage(logger, conf)

	authInterceptor := interceptor.NewInterceptor(logger, jwtManager, blog.AuthMethod)

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		grpc_recovery.UnaryServerInterceptor(),
		grpc_prometheus.UnaryServerInterceptor,
		grpc_validator.UnaryServerInterceptor(),
		authInterceptor.Unary(),
	)))

	v1.RegisterBlogServiceServer(grpcServer, blogService)
	grpc_health_v1.RegisterHealthServer(grpcServer, healthServer)


	listen, err := net.Listen("tcp", conf.Blog.Server.GRPC.Port)
	if err != nil {
		logger.Fatal("listen failed", err)
	}

	ch := make(chan os.Signal, 1)

	// Start gRPC server
	logger.Infof("gRPC Listening on port %s", conf.Blog.Server.GRPC.Port)
	go func() {
		if err = grpcServer.Serve(listen); err != nil {
			logger.Fatal("grpc serve failed", err)
		}
	}()

	// Start HTTP server
	logger.Infof("HTTP Listening on port %s", conf.Blog.Server.HTTP.Port)
	httpMux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	if err = v1.RegisterBlogServiceHandlerFromEndpoint(context.Background(), httpMux, conf.Blog.Server.GRPC.Port, opts); err != nil {
		logger.Fatal(err)
	}
	httpServer := &http.Server{
		Addr:    conf.Blog.Server.HTTP.Port,
		Handler: httpMux,
	}
	go func() {
		if err = httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal(err)
		}
	}()

	// Start Metrics server
	logger.Infof("Metrics Listening on port %s", conf.Blog.Server.Metrics.Port)
	metricsMux := http.NewServeMux()
	metricsMux.Handle("/metrics", promhttp.Handler())
	metricsServer := &http.Server{
		Addr:    conf.Blog.Server.Metrics.Port,
		Handler: metricsMux,
	}

	go func() {
		if err = metricsServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal(err)
		}
	}()

	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	grpcServer.GracefulStop()
	if err = httpServer.Shutdown(ctx); err != nil {
		logger.Fatal(err)
	}
	if err = metricsServer.Shutdown(ctx); err != nil {
		logger.Fatal(err)
	}
	<-ctx.Done()
	close(ch)
	logger.Info("Graceful Shutdown end")
}
