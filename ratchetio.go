package ratchetio

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"reflect"
)

const (
	endpoint       = "https://submit.ratchet.io/api/1/item/"
	post_data_type = "application/json"
)

var (
	Config = config{}
)

type config struct {
	APIKey string
}

type payload struct {
	AccessToken string       `json:"access_token"`
	Data        payload_data `json:"data"`
}

type payload_data struct {
	Env  string      `json:"environment"`
	Body interface{} `json:"body"`
}

type payload_message struct {
	Message struct {
		Body string `json:"body"`
	} `json:"message,omitempty"`
}

type parameters struct {
	Env string
}

func ReportMessage(m interface{}, p parameters) (err error) {
	var (
		body         payload_message
		payload_json []byte
		env, message string
		post_body    *bytes.Buffer
		item         payload
	)

	// Determine message type
	// only `string` or `error` are accepted
	switch reflect.TypeOf(m).String() {
	case "string":
		message = m.(string)
	case "*errors.errorString":
	case "runtime.errorString":
		message = m.(error).Error()
	default:
		return errors.New("Only String or Error")
	}

	// Get optional parameters and set default values
	if p.Env != "" {
		env = p.Env
	} else {
		env = "development" // Move to Config{}
	}

	// Set item Body
	body.Message.Body = message

	// Construct item
	item = payload{
		AccessToken: Config.APIKey,
		Data:        payload_data{env, body},
	}

	// Marshal item to JSON
	payload_json, err = json.Marshal(&item)

	// []byte to *Buffer
	post_body = bytes.NewBufferString(string(payload_json))

	// Send POST request
	resp, err := http.Post(endpoint, post_data_type, post_body)

	defer resp.Body.Close()
	// _, err = ioutil.ReadAll(resp.Body)

	return err
}

func CapturePanics() {
	if r := recover(); r != nil {
		log.Println(r)
		ReportMessage(r, parameters{})
	}
}
