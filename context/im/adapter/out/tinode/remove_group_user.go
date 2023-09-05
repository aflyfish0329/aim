package tinode

import (
	"test/context/im/application/port/out/chat"

	"github.com/pkg/errors"
	"github.com/tinode/chat/pbx"
)

func (s *Service) RemoveGroupUser(req chat.RemoveGroupUserReq) (chat.RemoveGroupUserResp, error) {
	client, err := s.newMessageClient()
	if err != nil {
		return chat.RemoveGroupUserResp{}, err
	}

	messageID := NewMessageID()

	err = s.hi(client, messageID)
	if err != nil {
		return chat.RemoveGroupUserResp{}, err
	}

	_, err = s.loginByToken(client, messageID.Next(), req.Token)
	if err != nil {
		return chat.RemoveGroupUserResp{}, err
	}

	err = s.subTopic(client, req.Group, "data sub desc", 24)
	if err != nil {
		return chat.RemoveGroupUserResp{}, err
	}

	clientMessage := &pbx.ClientMsg{
		Message: &pbx.ClientMsg_Del{
			Del: &pbx.ClientDel{
				Id:     messageID.Next().String(),
				Topic:  req.Group,
				UserId: req.User,
				What:   pbx.ClientDel_SUB,
			},
		},
	}

	err = client.Send(clientMessage)
	if err != nil {
		return chat.RemoveGroupUserResp{}, err
	}

	serverMsg, err := client.Recv()
	if err != nil {
		return chat.RemoveGroupUserResp{}, err
	}
	if serverMsg.GetCtrl().GetCode() > 300 {
		return chat.RemoveGroupUserResp{}, errors.New(serverMsg.GetCtrl().GetText())
	}

	resp := chat.RemoveGroupUserResp{}

	return resp, nil
}
