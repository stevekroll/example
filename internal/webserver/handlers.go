package webserver

import (
	"crypto/rand"
	"math/big"
	"net/http"
	"time"
)

const sleepDuration = 250

func sleep() {
	num, _ := rand.Int(rand.Reader, big.NewInt(sleepDuration))
	time.Sleep(time.Millisecond * time.Duration(num.Int64()))
}

var (
	CreateBundle = http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		sleep()
		w.WriteHeader(http.StatusAccepted)
	})
	GetBundle = http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		sleep()
		w.WriteHeader(http.StatusOK)
	})
	UpdateBundle = http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		sleep()
		w.WriteHeader(http.StatusNoContent)
	})
	DeleteBundle = http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		sleep()
		w.WriteHeader(http.StatusNoContent)
	})
)
