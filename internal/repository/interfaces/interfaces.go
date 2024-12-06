package interfaces

import "github.com/anandu86130/hospital-chat-service/internal/model"

type ChatRepoInter interface {
	Createchat(chat *model.History) error
	Findchat(userID, receiverID uint) (*[]model.History, error)
	LogVideoCall(videoCall *model.VideoCall) error
}
