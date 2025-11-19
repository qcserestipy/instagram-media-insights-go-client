package media

import "github.com/qcserestipy/instagram-api-go-client/pkg/client"

type Service struct {
	client *client.InstagramClient
}

func NewService(c *client.InstagramClient) *Service {
	return &Service{client: c}
}
