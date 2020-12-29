package configs

import "os"

// type GetServerPort interface {
// 	GetPort(string) string
// }

type GetServerPort struct {
}

func (p *GetServerPort) GetPort(port string) string {
	serverPort := os.Getenv("PORT")
	if serverPort == "" {
		serverPort = port
	}

	return serverPort
}
