package tinode

import (
	"fmt"
	"test/context/im/application/port/out/chat"

	"github.com/pkg/errors"
	"github.com/tinode/chat/pbx"
)

func (s *Service) CreateUser(req chat.CreateUserReq) (chat.CreateUserResp, error) {
	client, err := s.newMessageClient()
	if err != nil {
		return chat.CreateUserResp{}, err
	}

	messageID := NewMessageID()

	err = s.hi(client, messageID)
	if err != nil {
		return chat.CreateUserResp{}, err
	}

	secret := fmt.Sprintf("%s:%s", req.Username, req.Password)
	clientMessage := &pbx.ClientMsg{
		Message: &pbx.ClientMsg_Acc{
			Acc: &pbx.ClientAcc{
				Id:     messageID.Next().String(),
				UserId: fmt.Sprintf("new%s", req.Username),
				Cred: []*pbx.ClientCred{
					{
						Method: "email",
						Value:  req.Email,
					},
				},
				Scheme: "basic",
				Secret: []byte(secret),
			},
		},
	}

	err = client.Send(clientMessage)
	if err != nil {
		return chat.CreateUserResp{}, err
	}

	serverMsg, err := client.Recv()
	if err != nil {
		return chat.CreateUserResp{}, err
	}
	if serverMsg.GetCtrl().GetCode() > 300 {
		return chat.CreateUserResp{}, errors.New(serverMsg.GetCtrl().GetText())
	}

	resp := chat.CreateUserResp{
		User: string(serverMsg.GetCtrl().Params["user"]),
	}

	return resp, nil
}
