package chat

type CreateGroupReq struct {
	Token     string
	GroupName string
}

type CreateGroupResp struct {
	Group string
}
