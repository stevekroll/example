package myservice

type Client interface {
	Host() string
	WorkerHost() string
}
