package myworker

func Mock() Client {
	return new(mock)
}

type mock struct{}

func (m mock) Host() string {
	return "MockWorkerHost"
}

func (m mock) ServiceHost() string {
	return "MockServiceHost"
}
