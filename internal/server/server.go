package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ory/viper"
)

type Configs struct {
	Host string
	Port int
}

func loadConfigs() (*Configs, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("../configs/server/")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	var c Configs
	if err := viper.Unmarshal(&c); err != nil {
		return nil, err
	}
	return &c, nil
}

func Run() {
	configs, err := loadConfigs()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", configs.Host, configs.Port),
		Handler: GetRouter(),
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen:%s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	log.Println("Server exiting")
}
