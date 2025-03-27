package main

import (
	"encoding/json"
	"example/proto"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"

	"github.com/google/uuid"
)

var subjects []string = []string{"The dog", "A cat", "The bird", "My friend", "The teacher"}
var verbs []string = []string{"eats", "plays with", "jumps over", "runs towards", "sings to"}
var objects []string = []string{"a ball", "the tree", "a friend", "the fence", "a toy"}

func main() {

	for {
		if rand.Intn(2) == 0 {
			getRequests()
		} else {
			postMessage()
		}
	}
}

type getRequestsResponse struct {
	Requests int `json:"requests"`
}

func getRequests() {
	resp, err := http.Get("http://localhost:8080/requests")

	if err != nil {
		time.Sleep(time.Second * 2)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		time.Sleep(time.Second * 2)
		return
	}

	getRequestsResponse := getRequestsResponse{}

	json.Unmarshal(body, &getRequestsResponse)

	fmt.Println("Requests:", getRequestsResponse.Requests)
}

func postMessage() {
	subject := subjects[rand.Intn(len(subjects))]
	verb := verbs[rand.Intn(len(verbs))]
	object := objects[rand.Intn(len(objects))]

	sentence := fmt.Sprintf("%s %s %s.", subject, verb, object)

	chat := proto.NewChat(uuid.NewString(), sentence)
	resp, err := http.Post("http://localhost:8080/message", "application/octet-stream", proto.Encode(chat))

	if err != nil {
		time.Sleep(time.Second * 2)
		return
	}

	if resp.StatusCode == http.StatusBadRequest {
		fmt.Println("Response:", resp.Status)
	}

	defer resp.Body.Close()

}
