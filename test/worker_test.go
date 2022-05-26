package test

import (
	"example/internal/workers/myworker"
	"example/pkg/services/myservice"
	"testing"
)

func TestWorker(t *testing.T) {
	testWorker := myworker.New("Test", myservice.Mock())
	switch {
	case testWorker.Host() != "Test":
		t.Fatal("unexpected host")
	case testWorker.ServiceHost() != "MockServiceHost":
		t.Fatal("unexpected worker host")
	}
}
