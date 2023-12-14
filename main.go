package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func v1_routes(r *gin.RouterGroup) {
	r.POST("/counties", postCountyHandler)
	r.POST("/patients", postPatientHandler)
	r.POST("/patients/:patient_id/appointments", postPatientAppointmentsHandler)
}

func main() {
	router := gin.Default()
	v1_router := router.Group("/v1")
	v1_routes(v1_router)
	server := &http.Server{
		Addr:    os.Getenv("SERVER_ADDRESS"),
		Handler: router,
	}
	go func() {
		//service connections
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdwon Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done() timeout 3 seconds
	select {
	case <-ctx.Done():
		log.Println("timeout of 2 seconds")
	}
	log.Println("Server exiting")
}
