package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/derailed/nuclio/iconoflix"
	nuclio "github.com/nuclio/nuclio-sdk-go"
)

func main() {}

// Handler the yo handler
func Handler(context *nuclio.Context, event nuclio.Event) (interface{}, error) {
	log.Printf("Go called with %s", string(event.GetBody()))
	log.Println(event.GetID(), event.GetFields(), event.GetPath(), event.GetMethod())

	raw, err := json.Marshal(iconoflix.RandMovie())
	if err != nil {
		return "", fmt.Errorf("Doh! something went wrong %v", err)
	}
	log.Println("JSON Resp", string(raw))

	return nuclio.Response{
		StatusCode:  200,
		ContentType: "application/json",
		Body:        raw,
	}, nil
}
