package chat

type UserLoginReq struct {
	Username string
	Password string
}

type UserLoginResp struct {
	User  string
	Token string
}
