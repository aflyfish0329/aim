package tinode

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/pkg/errors"
	"github.com/tinode/chat/pbx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (s *Service) newMessageClient() (pbx.Node_MessageLoopClient, error) {
	conn, err := grpc.Dial(s.serverUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	c := pbx.NewNodeClient(conn)
	client, err := c.MessageLoop(context.Background())
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (s *Service) hi(client pbx.Node_MessageLoopClient, messageID *MessageID) error {
	clientMessage := &pbx.ClientMsg{
		Message: &pbx.ClientMsg_Hi{
			Hi: &pbx.ClientHi{
				Id:        messageID.String(),
				UserAgent: "Golang_Spider_Bot/3.0",
				Ver:       "0.22.9",
				Lang:      "EN",
			},
		},
	}

	err := client.Send(clientMessage)
	if err != nil {
		return err
	}

	serverMsg, err := client.Recv()
	if err != nil {
		return err
	}

	result := serverMsgToHiMsg(serverMsg)
	if result.Code > 300 {
		return errors.New(result.Text)
	}

	return nil
}

func (s *Service) loginByUser(client pbx.Node_MessageLoopClient, messageID *MessageID, username, password string) (LoginMsg, error) {
	secret := fmt.Sprintf("%s:%s", username, password)
	clientMessage := &pbx.ClientMsg{
		Message: &pbx.ClientMsg_Login{
			Login: &pbx.ClientLogin{
				Id:     messageID.String(),
				Scheme: "basic",
				Secret: []byte(secret),
			},
		},
	}

	err := client.Send(clientMessage)
	if err != nil {
		return LoginMsg{}, err
	}

	serverMsg, err := client.Recv()
	if err != nil {
		return LoginMsg{}, err
	}

	result := serverMsgToLoginMsg(serverMsg)
	if result.Code > 300 {
		return LoginMsg{}, errors.New(result.Text)
	}

	return result, nil
}

func (s *Service) loginByToken(client pbx.Node_MessageLoopClient, messageID *MessageID, token string) (LoginMsg, error) {
	tb, err := s.decodeToken(token)
	if err != nil {
		return LoginMsg{}, err
	}

	clientMessage := &pbx.ClientMsg{
		Message: &pbx.ClientMsg_Login{
			Login: &pbx.ClientLogin{
				Id:     messageID.String(),
				Scheme: "token",
				Secret: []byte(tb),
			},
		},
	}

	err = client.Send(clientMessage)
	if err != nil {
		return LoginMsg{}, err
	}
	serverMsg, err := client.Recv()
	if err != nil {
		return LoginMsg{}, err
	}

	result := serverMsgToLoginMsg(serverMsg)
	if result.Code > 300 {
		return LoginMsg{}, errors.New(result.Text)
	}

	return result, nil
}

func (s *Service) subTopic(client pbx.Node_MessageLoopClient, topic, what string, limit int32) error {
	clientMessage := &pbx.ClientMsg{
		Message: &pbx.ClientMsg_Sub{
			Sub: &pbx.ClientSub{
				Topic: topic,
				GetQuery: &pbx.GetQuery{
					What: what,
					Data: &pbx.GetOpts{
						Limit: limit,
					},
				},
			},
		},
	}

	err := client.Send(clientMessage)
	if err != nil {
		return err
	}

	serverMsg, err := client.Recv()
	if err != nil {
		return err
	}

	result := serverMsgToHiMsg(serverMsg)
	if result.Code > 300 {
		return errors.New(result.Text)
	}

	return nil
}

func (s *Service) decodeToken(token string) ([]byte, error) {
	decodedSecret := make([]byte, base64.StdEncoding.DecodedLen(len(token)))
	n, err := base64.StdEncoding.Decode(decodedSecret, []byte(token))
	if err != nil {
		return nil, err
	}

	return decodedSecret[:n], nil
}
