package router

import (
	"test/context/common/pkg/gin/middleware"
	"test/context/im/adapter/in/web/handler"
	"test/context/im/adapter/out/db/repository"
	"test/context/im/adapter/out/tinode"
	"test/context/im/application/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(c *gin.Engine, db *gorm.DB) {
	permissionValidator := middleware.NewPermissionMiddleware()
	userRepository := repository.NewUserRepository(db)
	tinodeService := tinode.NewService("localhost:16060")
	chatService := service.NewChatService(userRepository, &tinodeService)

	// chatHandler
	chatHandler := handler.NewChatHandler(chatService)
	c.POST("/user_login", permissionValidator, chatHandler.UserLogin)
}
