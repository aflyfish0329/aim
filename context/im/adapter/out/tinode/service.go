package tinode

func NewService(serverUrl string) Service {
	return Service{
		serverUrl: serverUrl,
	}
}

type Service struct {
	serverUrl string
}
