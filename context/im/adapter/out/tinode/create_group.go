package tinode

import (
	"encoding/json"
	"fmt"
	"test/context/im/application/port/out/chat"

	"github.com/pkg/errors"
	"github.com/tinode/chat/pbx"
)

func (s *Service) CreateGroup(req chat.CreateGroupReq) (chat.CreateGroupResp, error) {
	client, err := s.newMessageClient()
	if err != nil {
		return chat.CreateGroupResp{}, err
	}

	messageID := NewMessageID()

	err = s.hi(client, messageID)
	if err != nil {
		return chat.CreateGroupResp{}, err
	}

	_, err = s.loginByToken(client, messageID.Next(), req.Token)
	if err != nil {
		return chat.CreateGroupResp{}, err
	}

	type descPublic struct {
		Fn string `json:"fn"`
	}

	dp := descPublic{
		Fn: req.GroupName,
	}

	dpb, err := json.Marshal(dp)
	if err != nil {
		return chat.CreateGroupResp{}, err
	}

	clientMessage := &pbx.ClientMsg{
		Message: &pbx.ClientMsg_Sub{
			Sub: &pbx.ClientSub{
				Id:    messageID.Next().String(),
				Topic: fmt.Sprintf("new%s", messageID.String()),
				SetQuery: &pbx.SetQuery{
					Desc: &pbx.SetDesc{
						Public: dpb,
					},
				},
			},
		},
	}

	err = client.Send(clientMessage)
	if err != nil {
		return chat.CreateGroupResp{}, err
	}

	serverMsg, err := client.Recv()
	if err != nil {
		return chat.CreateGroupResp{}, err
	}

	if serverMsg.GetCtrl().GetCode() > 300 {
		return chat.CreateGroupResp{}, errors.New(serverMsg.GetCtrl().GetText())
	}

	resp := chat.CreateGroupResp{
		Group: serverMsg.GetCtrl().Topic,
	}

	return resp, nil
}
