package tinode

import "github.com/tinode/chat/pbx"

type LoginMsg struct {
	Code  int32
	Text  string
	Token string
	User  string
}

func serverMsgToLoginMsg(msg *pbx.ServerMsg) LoginMsg {
	return LoginMsg{
		Code:  msg.GetCtrl().GetCode(),
		Text:  msg.GetCtrl().GetText(),
		Token: string(msg.GetCtrl().Params["token"]),
		User:  string(msg.GetCtrl().Params["user"]),
	}
}

type HiMsg struct {
	Code int32
	Text string
}

func serverMsgToHiMsg(msg *pbx.ServerMsg) HiMsg {
	return HiMsg{
		Code: msg.GetCtrl().GetCode(),
		Text: msg.GetCtrl().GetText(),
	}
}
