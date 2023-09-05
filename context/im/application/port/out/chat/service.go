package chat

type ChatService interface {
	UserLogin(UserLoginReq) (UserLoginResp, error)
}
