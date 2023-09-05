package tinode

import (
	"test/context/im/application/port/out/chat"

	"github.com/pkg/errors"
	"github.com/tinode/chat/pbx"
)

func (s *Service) AddGroupUser(req chat.AddGroupUserReq) (chat.AddGroupUserResp, error) {
	client, err := s.newMessageClient()
	if err != nil {
		return chat.AddGroupUserResp{}, err
	}

	messageID := NewMessageID()

	err = s.hi(client, messageID)
	if err != nil {
		return chat.AddGroupUserResp{}, err
	}

	_, err = s.loginByToken(client, messageID.Next(), req.Token)
	if err != nil {
		return chat.AddGroupUserResp{}, err
	}

	err = s.subTopic(client, req.Group, "data sub desc", 24)
	if err != nil {
		return chat.AddGroupUserResp{}, err
	}

	clientMessage := &pbx.ClientMsg{
		Message: &pbx.ClientMsg_Set{
			Set: &pbx.ClientSet{
				Topic: req.Group,
				Query: &pbx.SetQuery{
					Sub: &pbx.SetSub{
						UserId: req.User,
					},
				},
			},
		},
	}

	err = client.Send(clientMessage)
	if err != nil {
		return chat.AddGroupUserResp{}, err
	}

	serverMsg, err := client.Recv()
	if err != nil {
		return chat.AddGroupUserResp{}, err
	}

	if serverMsg.GetCtrl().GetCode() > 300 {
		return chat.AddGroupUserResp{}, errors.New(serverMsg.GetCtrl().GetText())
	}

	resp := chat.AddGroupUserResp{}

	return resp, nil
}
