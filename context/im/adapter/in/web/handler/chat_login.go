package handler

import (
	"net/http"
	"test/context/im/application/port/in/usecase"

	"github.com/gin-gonic/gin"
)

type UserLoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewUserLoginOutput(input usecase.UserLoginOutput) UserLoginOutput {
	return UserLoginOutput{
		Token: input.Token,
	}
}

type UserLoginOutput struct {
	Token string `json:"token"`
}

func (h ChatHandler) UserLogin(c *gin.Context) {
	var input UserLoginInput
	if err := c.ShouldBind(&input); err != nil {
		c.Error(err)

		return
	}

	command, err := usecase.NewUserLoginCommand(input.Username, input.Password)
	if err != nil {
		c.Error(err)

		return
	}

	output, err := h.ChatUseCase.UserLogin(c.Request.Context(), command)
	if err != nil {
		c.Error(err)

		return
	}

	c.JSON(http.StatusOK, NewUserLoginOutput(output))
}
