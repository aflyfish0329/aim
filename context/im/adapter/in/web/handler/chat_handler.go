package handler

import "test/context/im/application/port/in/usecase"

func NewChatHandler(uc usecase.ChatUseCase) ChatHandler {
	return ChatHandler{
		ChatUseCase: uc,
	}
}

type ChatHandler struct {
	ChatUseCase usecase.ChatUseCase
}
