package chat

type CreateUserReq struct {
	Username string
	Password string
	Email    string
}

type CreateUserResp struct {
	User string
}
