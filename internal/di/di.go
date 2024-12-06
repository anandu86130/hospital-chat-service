package di

import (
	"log"

	"github.com/anandu86130/hospital-chat-service/config"
	"github.com/anandu86130/hospital-chat-service/internal/db"
	"github.com/anandu86130/hospital-chat-service/internal/handler"
	repository "github.com/anandu86130/hospital-chat-service/internal/repository"
	"github.com/anandu86130/hospital-chat-service/internal/server"
	"github.com/anandu86130/hospital-chat-service/internal/service"
)

func Init() {
	cnfg := config.LoadConfig()

	mongodb, err := db.ConnectToDatabase(cnfg)
	if err != nil {
		log.Fatalf("failed to connect to MongoDB: %v", err)
	}

	// mongoDB := mongoClient.Database(cnfg.DBname)

	chatRepo := repository.NewChatRepository(mongodb)
	chatUsecase := service.NewChatService(chatRepo)
	chatServer := handler.NewChatServiceServer(chatUsecase)
	err = server.NewGRPCServer(cnfg.ChatPort, chatServer)
	if err != nil {
		log.Fatalf("failed to start gRPC server %v", err.Error())
	}
}
