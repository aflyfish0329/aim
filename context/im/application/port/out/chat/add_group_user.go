package chat

type AddGroupUserReq struct {
	Token string
	Group string
	User  string
}

type AddGroupUserResp struct{}
