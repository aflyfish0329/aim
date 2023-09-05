package tinode

import (
	"test/context/im/application/port/out/chat"

	"github.com/pkg/errors"
	"github.com/tinode/chat/pbx"
)

func (s *Service) DeleteUser(req chat.DeleteUserReq) (chat.DeleteUserResp, error) {
	client, err := s.newMessageClient()
	if err != nil {
		return chat.DeleteUserResp{}, err
	}

	messageID := NewMessageID()

	err = s.hi(client, messageID)
	if err != nil {
		return chat.DeleteUserResp{}, err
	}

	_, err = s.loginByToken(client, messageID.Next(), req.Token)
	if err != nil {
		return chat.DeleteUserResp{}, err
	}

	clientMessage := &pbx.ClientMsg{
		Message: &pbx.ClientMsg_Del{
			Del: &pbx.ClientDel{
				Id:     messageID.Next().String(),
				What:   4,
				UserId: req.User,
				Hard:   req.IsHardDelete,
			},
		},
	}

	err = client.Send(clientMessage)
	if err != nil {
		return chat.DeleteUserResp{}, err
	}

	serverMsg, err := client.Recv()
	if err != nil {
		return chat.DeleteUserResp{}, err
	}
	if serverMsg.GetCtrl().GetCode() > 300 {
		return chat.DeleteUserResp{}, errors.New(serverMsg.GetCtrl().GetText())
	}

	resp := chat.DeleteUserResp{}

	return resp, nil
}
