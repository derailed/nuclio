package main

import (
	"encoding/json"
	"fmt"
	"log"

	nuclio "github.com/nuclio/nuclio-sdk-go"
)

// Handler the yo handler
func Handler(context *nuclio.Context, event nuclio.Event) (interface{}, error) {
	log.Printf("Go called with %s", string(event.GetBody()))
	log.Println(event.GetID(), event.GetFields(), event.GetPath(), event.GetMethod())

	input := struct {
		Body string
	}{}
	err := json.Unmarshal(event.GetBody(), &input)
	if err != nil {
		return "", fmt.Errorf("Doh! something went wrong %s", err)
	}

	body := struct {
		Message string
	}{
		Message: "Hello from F1 " + input.Body,
	}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return "", fmt.Errorf("Doh! something went wrong %s", err)
	}
	log.Println("JSON Resp", string(jsonBody))

	return nuclio.Response{
		StatusCode:  200,
		ContentType: "application/json",
		Body:        jsonBody,
	}, nil
}

func main() {}
