package restclient

import (
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
)

const (
	// userAPIURL is a context url to config API
	messageAPIURL = "/messages/%d"
)

type MessageAPI struct {
	restAPI
}

type Message struct {
	Text string `json:"text"`
}

func NewMessageAPI() *MessageAPI {
	api := new(MessageAPI)

	// Create a Resty Client
	client := resty.New()
	// You probably need to configure timeouts
	// and retries strategies for message API

	api.readClient = client

	//api.logger = &apiLogger{}
	return api
}

func (api *MessageAPI) readURL(id string) string {
	return fmt.Sprintf(messageAPIURL, id)
}

func (api *MessageAPI) Get(id string) (*Message, error) {
	url := api.readURL(id)

	msg := new(Message)
	res, err := api.get(url, http.Header{}, msg)

	// type assertion
	msg, _ = res.(*Message)

	return msg, err
}