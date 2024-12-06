package service

import (
	"fmt"
	"net/url"
	"time"

	"github.com/anandu86130/hospital-chat-service/internal/model"
	pb "github.com/anandu86130/hospital-chat-service/internal/pbC"
	"github.com/anandu86130/hospital-chat-service/internal/repository/interfaces"
	inter "github.com/anandu86130/hospital-chat-service/internal/service/interfaces"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type chatService struct {
	repo interfaces.ChatRepoInter
}

func generateJitsiRoomURL(userID, receiverID uint) string {
	// Combine userID and receiverID to form a unique room name
	roomName := fmt.Sprintf("call-%d-%d", userID, receiverID)

	// Create the base Jitsi Meet URL
	baseURL := "https://meet.jit.si/" + roomName

	// Add query parameters for user and receiver IDs
	params := url.Values{}
	params.Add("user_id", fmt.Sprintf("%d", userID))
	params.Add("receiver_id", fmt.Sprintf("%d", receiverID))

	// Return the full Jitsi room URL with query parameters
	return baseURL + "?" + params.Encode()
}

// StartVideoCall generates a new video call URL using Jitsi and stores it
func (c *chatService) StartVideoCallService(p *pb.VideoCallRequest) (*pb.VideoCallResponse, error) {
	roomURL := generateJitsiRoomURL(uint(p.User_ID), uint(p.Receiver_ID))
	videoCall := &model.VideoCall{
		UserID:     uint(p.User_ID),
		ReceiverID: uint(p.Receiver_ID),
		RoomURL:    roomURL,
		Timestamp:  primitive.NewDateTimeFromTime(time.Now()),
	}
	//store video call information to db
	if err := c.repo.LogVideoCall(videoCall); err != nil {
		return nil, err
	}

	return &pb.VideoCallResponse{RoomUrl: roomURL}, nil
}

func NewChatService(repo interfaces.ChatRepoInter) inter.ChatServiceInter {
	return &chatService{
		repo: repo,
	}
}
