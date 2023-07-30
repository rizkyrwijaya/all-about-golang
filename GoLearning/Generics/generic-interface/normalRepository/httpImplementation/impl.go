package httpimplementation

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	"github.com/rizkyrwijaya/all-about-golang/GoLearning/Generics/generic-interface/normalRepository"
)

type request struct {
	url     string
	mockRes []byte
}

// New just example implementation of a day2day basis requester
func New(url string, mockRes []byte) normalRepository.Requester {
	return &request{
		url:     url,
		mockRes: mockRes,
	}
}

func (r *request) GetResponse() normalRepository.Response {
	fmt.Println(r.url)
	return normalRepository.Response{}
}

func (r *request) MockResponse() (res normalRepository.Response) {
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
