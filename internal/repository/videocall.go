package repository

import (
	"context"
	"time"

	"github.com/anandu86130/hospital-chat-service/internal/model"
)

// LogVideoCall implements interfaces.ChatRepoInter.
func (c *ChatRepo) LogVideoCall(videoCall *model.VideoCall) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := c.Collection.InsertOne(ctx, videoCall)
	if err != nil {
		return err
	}
	return nil
}
