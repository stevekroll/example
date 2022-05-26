package myservice

func Mock() Client {
	return new(mock)
}

type mock struct{}

func (m mock) Host() string {
	return "MockServiceHost"
}

func (m mock) WorkerHost() string {
	return "MockWorkerHost"
}
