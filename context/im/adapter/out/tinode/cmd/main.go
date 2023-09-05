package main

import (
	"context"
	"fmt"
	"time"

	"github.com/tinode/chat/pbx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:16060", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	c := pbx.NewNodeClient(conn)
	client, err := c.MessageLoop(context.Background())
	if err != nil {
		panic(err)
	}

	clientMessage := &pbx.ClientMsg{
		Message: &pbx.ClientMsg_Hi{
			Hi: &pbx.ClientHi{
				UserAgent: "Golang_Spider_Bot/3.0",
				Ver:       "0.22.9",
				Lang:      "EN",
			},
		},
	}

	err = client.Send(clientMessage)
	if err != nil {
		panic(err)
	}

	serverMsg, err := client.Recv()
	if err != nil {
		panic(err)
	}

	fmt.Println("# hi", serverMsg)

	secret := fmt.Sprintf("%s:%s", "carol", "carol123")
	clientMessage = &pbx.ClientMsg{
		Message: &pbx.ClientMsg_Login{
			Login: &pbx.ClientLogin{
				Scheme: "basic",
				Secret: []byte(secret),
			},
		},
	}

	err = client.SendMsg(clientMessage)
	if err != nil {
		panic(err)
	}

	serverMsg, err = client.Recv()
	if err != nil {
		panic(err)
	}

	fmt.Println("# login", serverMsg)

	clientMessage = &pbx.ClientMsg{
		Message: &pbx.ClientMsg_Sub{
			Sub: &pbx.ClientSub{
				Topic:    "grp5-nagvnQ1ek",
				GetQuery: nil,
			},
		},
	}

	err = client.Send(clientMessage)
	if err != nil {
		panic(err)
	}

	for {
		serverMsg, err = client.Recv()
		if err != nil {
			panic(err)
		}

		fmt.Println(serverMsg)
		time.Sleep(100 * time.Millisecond)
	}

}
