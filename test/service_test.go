package test

import (
	"example/internal/services/myservice"
	"example/pkg/workers/myworker"
	"testing"
)

func TestService(t *testing.T) {
	testService := myservice.New("Test", myworker.Mock())
	switch {
	case testService.Host() != "Test":
		t.Fatal("unexpected host")
	case testService.WorkerHost() != "MockWorkerHost":
		t.Fatal("unexpected worker host")
	}
}
