package service

import (
	"test/context/im/application/port/out/chat"
	"test/context/im/application/port/out/db/repository"
)

func NewChatService(
	userRepository repository.UserRepository,
	chatService chat.ChatService,
) ChatService {
	return ChatService{
		userRepository: userRepository,
		chatService:    chatService,
	}
}

type ChatService struct {
	userRepository repository.UserRepository
	chatService    chat.ChatService
}
