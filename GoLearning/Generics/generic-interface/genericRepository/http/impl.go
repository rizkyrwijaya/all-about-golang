package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	genericrepository "github.com/rizkyrwijaya/all-about-golang/GoLearning/Generics/generic-interface/genericRepository"
)

type client[T any] struct {
	url     string
	mockRes []byte
}

func New[T any](url string, mockRes []byte) genericrepository.Requester[T] {
	return &client[T]{
		url:     url,
		mockRes: mockRes,
	}
}

func (r *client[T]) GetResponse() (res T) {
	fmt.Println(r.url)
	return
}

func (r *client[T]) MockResponse() (res T) {
	if r.mockRes == nil && len(r.mockRes) == 0 {
		return
	}

	bbuf := bytes.NewBuffer(r.mockRes)
	dec := json.NewDecoder(bbuf)
	err := dec.Decode(&res)
	if err != nil {
		log.Print("Failed to Parse due to: ", err)
	}
	return
}
