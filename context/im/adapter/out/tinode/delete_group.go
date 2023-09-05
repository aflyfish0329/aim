package tinode

import (
	"test/context/im/application/port/out/chat"

	"github.com/pkg/errors"
	"github.com/tinode/chat/pbx"
)

func (s *Service) DeleteGroup(req chat.DeleteGroupReq) (chat.DeleteGroupResp, error) {
	client, err := s.newMessageClient()
	if err != nil {
		return chat.DeleteGroupResp{}, err
	}

	messageID := NewMessageID()

	err = s.hi(client, messageID)
	if err != nil {
		return chat.DeleteGroupResp{}, err
	}

	_, err = s.loginByToken(client, messageID.Next(), req.Token)
	if err != nil {
		return chat.DeleteGroupResp{}, err
	}

	clientMessage := &pbx.ClientMsg{
		Message: &pbx.ClientMsg_Del{
			Del: &pbx.ClientDel{
				Id:    messageID.Next().String(),
				Hard:  true,
				Topic: req.Group,
				What:  pbx.ClientDel_TOPIC,
			},
		},
	}

	err = client.Send(clientMessage)
	if err != nil {
		return chat.DeleteGroupResp{}, err
	}

	serverMsg, err := client.Recv()
	if err != nil {
		return chat.DeleteGroupResp{}, err
	}

	if serverMsg.GetCtrl().GetCode() > 300 {
		return chat.DeleteGroupResp{}, errors.New(serverMsg.GetCtrl().GetText())
	}

	resp := chat.DeleteGroupResp{}

	return resp, nil
}
