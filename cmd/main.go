package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/Pet002/poc-grpc-server-otel/app/user"
	"github.com/Pet002/poc-grpc-server-otel/config"
	"github.com/Pet002/poc-grpc-server-otel/logger"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
)

func main() {

	ctx := context.Background()
	// init logger
	logger.InitLogger()

	tp, err := config.InitTrace()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	grpcEndpoint := fmt.Sprintf(":%s", port)

	lis, err := net.Listen("tcp", grpcEndpoint)
	if err != nil {
		log.Fatal(err)
	}

	// create a grpc server
	grpcServe := grpc.NewServer(
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
	)

	userService := user.NewService()
	userHandler := user.NewHandler(userService)

	// Register a server to handler
	user.RegisterHelloServiceServer(grpcServe, userHandler)

	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		<-ctx.Done()
		grpcServe.GracefulStop()
		defer wg.Done()
	}()

	log.Printf("starting grpc server at endpoint %s", grpcEndpoint)
	err = grpcServe.Serve(lis)
	if err != nil {
		log.Fatalf("could not serve: %v", err)
	}
	wg.Wait()
}
