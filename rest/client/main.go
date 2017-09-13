package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	resp, err := http.Get("http://localhost:8080/getScore/Rafael/vs/Serena")
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("couldnt read body response: %v", err)
	}

	var data ScoreModel
	json.Unmarshal(body, &data)

	log.Printf("The current score is: %s", data.Score)
}

type ScoreModel struct {
	Score string
}
