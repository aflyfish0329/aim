package tinode

import (
	"test/context/im/application/port/out/chat"
)

func (s *Service) UserLogin(req chat.UserLoginReq) (chat.UserLoginResp, error) {
	client, err := s.newMessageClient()
	if err != nil {
		return chat.UserLoginResp{}, err
	}

	messageID := NewMessageID()

	err = s.hi(client, messageID)
	if err != nil {
		return chat.UserLoginResp{}, err
	}

	loginMsg, err := s.loginByUser(client, messageID.Next(), req.Username, req.Password)
	if err != nil {
		return chat.UserLoginResp{}, err
	}

	resp := chat.UserLoginResp{
		User:  loginMsg.User,
		Token: loginMsg.Token,
	}

	return resp, nil
}
