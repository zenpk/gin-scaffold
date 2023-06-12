package main

import (
	"flag"
	"github.com/zenpk/gin-scaffold/internal/controller"
	"github.com/zenpk/gin-scaffold/pkg/viperpkg"
	"github.com/zenpk/gin-scaffold/pkg/zap"
	"log"
)

var (
	mode = flag.String("mode", "dev", "define program mode")
)

func main() {
	flag.Parse()
	// Viper
	if err := viperpkg.InitGlobalConfig(*mode); err != nil {
		log.Fatalf("failed to initialize Viper: %v", err)
	}
	// zap
	if err := zap.InitLogger(*mode); err != nil {
		log.Fatalf("failed to initialize zap: %v", err)
	}
	defer func() {
		if err := zap.Logger.Sync(); err != nil {
			log.Fatalf("failed to close zap: %v", err)
		}
	}()
	// GORM
	//if err := dal.InitDB(); err != nil {
	//	log.Fatalf("failed to initialize database: %v", err)
	//}
	// Gin
	if err := controller.InitGin(); err != nil {
		log.Fatalf("failed to initialize Gin: %v", err)
	}
}
