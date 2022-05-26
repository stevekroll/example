package myworker

import (
	"example/pkg/services/myservice"
	"example/pkg/workers/myworker"
)

func New(host string, service myservice.Client) myworker.Client {
	return &Client{
		host:    host,
		service: service,
	}
}

type Client struct {
	service myservice.Client
	host    string
}

func (c Client) Host() string {
	return c.host
}

func (c Client) ServiceHost() string {
	return c.service.Host()
}
