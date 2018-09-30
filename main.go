package main

import (
	"context"
	"fmt"
	"golang-tdd-rest-api/database"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"google.golang.org/grpc"

	todoController "golang-tdd-rest-api/todo/controller"
	todoRepo "golang-tdd-rest-api/todo/repository"
	todoGRPCTransport "golang-tdd-rest-api/todo/transport/gRPC"
	todoHTTPTransport "golang-tdd-rest-api/todo/transport/http"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("unable to read .env file")
	}

	gormTodo, err := database.GetGormTodo()
	if err != nil {
		log.Fatal(err)
	}
	defer database.CloseGormTodo()

	todoRepo := todoRepo.NewGormRepo(gormTodo)
	todoController := todoController.NewController(todoRepo)

	// HTTP
	e := echo.New()
	todoHTTPTransport.NewHandler(e, todoController)
	webPort := os.Getenv("WEB_PORT")
	if webPort == "" {
		webPort = "8080"
	}
	go func() {
		if err := e.Start(fmt.Sprintf(":%s", webPort)); err != nil {
			e.Logger.Info("shutting down the server")
		}
	}()

	// gRPC
	gRPCPort := os.Getenv("GRPC_PORT")
	if gRPCPort == "" {
		gRPCPort = "9000"
	}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", gRPCPort))
	grpcServer := grpc.NewServer()
	todoGRPCTransport.RegisterTodoServer(grpcServer, todoGRPCTransport.NewHandler(todoController))
	go func() {
		grpcServer.Serve(lis)
	}()

	// Gracful Shutdown
	gracfulShutdownChan := make(chan os.Signal)
	signal.Notify(gracfulShutdownChan, syscall.SIGTERM)
	signal.Notify(gracfulShutdownChan, syscall.SIGINT)
	<-gracfulShutdownChan
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal("Something wrong")
	}
	grpcServer.GracefulStop()
}
