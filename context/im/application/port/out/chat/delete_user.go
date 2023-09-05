package chat

type DeleteUserReq struct {
	Token        string
	User         string
	IsHardDelete bool
}

type DeleteUserResp struct{}
