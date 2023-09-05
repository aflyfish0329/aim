package chat

type RemoveGroupUserReq struct {
	Token string
	Group string
	User  string
}

type RemoveGroupUserResp struct{}
