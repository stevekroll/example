package myworker

type Client interface {
	Host() string
	ServiceHost() string
}
