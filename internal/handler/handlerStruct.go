package handler

import (
	"github.com/anandu86130/hospital-chat-service/internal/service/interfaces"
	pb "github.com/anandu86130/hospital-chat-service/internal/pbC"
)

type ChatServiceServer struct {
	pb.UnimplementedChatServiceServer
	svc interfaces.ChatServiceInter
}

func NewChatServiceServer(svc interfaces.ChatServiceInter) *ChatServiceServer {
	return &ChatServiceServer{
		svc: svc,
	}
}
