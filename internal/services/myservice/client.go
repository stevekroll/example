package myservice

import (
	"example/pkg/services/myservice"
	"example/pkg/workers/myworker"
)

func New(host string, worker myworker.Client) myservice.Client {
	return &Client{
		host:   host,
		worker: worker,
	}
}

type Client struct {
	worker myworker.Client
	host   string
}

func (c Client) Host() string {
	return c.host
}

func (c Client) WorkerHost() string {
	return c.worker.Host()
}
