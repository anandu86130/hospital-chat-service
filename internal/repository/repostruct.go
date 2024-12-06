package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	inter "github.com/anandu86130/hospital-chat-service/internal/repository/interfaces"
)

type ChatRepo struct {
	Collection          *mongo.Collection
	VideoCallCollection *mongo.Collection
	//ChatCollection *mongo.Collection
}

func NewChatRepository(mongo *mongo.Database) inter.ChatRepoInter {
	return &ChatRepo{
		Collection:          mongo.Collection("myNotification"),
		VideoCallCollection: mongo.Collection("video-call"),
		//ChatCollection: mongo.Collection("myNotification"),
	}
}
