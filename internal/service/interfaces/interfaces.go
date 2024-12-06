package interfaces

import pb "github.com/anandu86130/hospital-chat-service/internal/pbC"

type ChatServiceInter interface {
	CreateChatService(p *pb.Message) error
	FetchChatService(p *pb.ChatID) (*pb.ChatHistory, error)
	StartVideoCallService(p *pb.VideoCallRequest) (*pb.VideoCallResponse, error)
}
