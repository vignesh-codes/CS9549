package main

import (
	"deployment-service/apps/repository/adapter"
	"deployment-service/apps/repository/instance"
	"deployment-service/apps/routes"
	"deployment-service/constants"

	"context"
	"deployment-service/logger"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	logger.InitConsoleLogger()
	logger.InitEventLogger()
	// configs := config.GetConfig()
	// aws := instance.GetAwsSession()
	// RedisDBConnection := instance.GetRedisConnection()
	// PSqlConnection := instance.GetPSqlConnection()
	MongoDBConnection := instance.GetMongoConnection()
	KubernetesConnection := instance.GetKubernetesConnection()
	repository := adapter.RepositoryAdapter(MongoDBConnection, KubernetesConnection)

	fmt.Println("Starting %s API server", "deployment-service")

	server := &http.Server{
		Addr:    constants.PORT,
		Handler: routes.NewRouter().SetRouters(repository),
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("Error::%v", err)
			fmt.Println("Failed to start %s service\n", "deployment-service")
		}
	}()

	fmt.Println("Listening on port %v ", server.Addr)

	// queue := svc.NewServiceRepo(repository).SQSService
	// go queue.InitSQS()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shutdown server
	fmt.Println("Shutting down server.")
	if err := server.Shutdown(ctx); err != nil {
		fmt.Println("Server forced to shutdown: %v\n", err)
	}
}
