package main

import (
	"encoding/json"
	"fmt"
	"github.com/sendwithus/lib-go"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
)

var heads = []string{"bendr", "dead", "fang", "pixel", "regular", "safe", "sand-worm", "shades", "smile", "tongue"}
var tails = []string{"small-rattle", "skinny-tail", "round-bum", "pointed", "pixel", "freckled", "fat-rattle", "curled", "block-bum"}

func start(w http.ResponseWriter, r *http.Request) {
	var requestData GameStartRequest
	json.NewDecoder(r.Body).Decode(&requestData)

	log.Printf("Game starting - %v\n", requestData.GameId)
	responseData := GameStartResponse{
		Color:    "#00f8f8",
		Name:     "inky-snek",
		HeadUrl:  swu.String("https://s3.amazonaws.com/john-box-o-mysteries/pacman+ghosts/inky.png"),
		HeadType: swu.String(heads[rand.Intn(len(heads))]),
		TailType: swu.String(tails[rand.Intn(len(tails))]),
	}
	b, err := json.Marshal(responseData)
	if err != nil {
		log.Fatalf("%v", err)
		return
	}
	fmt.Println(string(b))
	w.Write(b)
}

func move(w http.ResponseWriter, r *http.Request) {
	var requestData MoveRequest
	val, err := ioutil.ReadAll(r.Body)
	json.Unmarshal(val, &requestData)
	responseData := MoveResponse{
		Move: requestData.GenerateMove(),
	}
	log.Printf("Move request - direction:%v\n", responseData.Move)
	if err != nil {
		fmt.Printf("ERR: %#v\n", err)
	}
	log.Printf("%v\n", string(val))
	b, err := json.Marshal(responseData)
	if err != nil {
		log.Fatalf("%v", err)
		return
	}
	w.Write(b)
}
