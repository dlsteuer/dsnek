package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"

	"github.com/icrowley/fake"
	"github.com/lucasb-eyer/go-colorful"
)

var heads = []string{"bendr", "dead", "fang", "pixel", "regular", "safe", "sand-worm", "shades", "smile", "tongue"}
var tails = []string{"small-rattle", "skinny-tail", "round-bum", "regular", "pixel", "freckled", "fat-rattle", "curled", "block-bum"}

func str(str string) *string {
	return &str
}

func start(w http.ResponseWriter, r *http.Request) {
	var requestData GameStartRequest
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return
	}
	err = json.Unmarshal(body, &requestData)
	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("Game starting - %v\n", string(body))
	responseData := GameStartResponse{
		Color:    getColor(),
		Name:     fake.Word(),
		HeadUrl:  str("https://s3.amazonaws.com/john-box-o-mysteries/pacman+ghosts/inky.png"),
		HeadType: str("fang"),
		TailType: str("pixel"),
	}
	b, err := json.Marshal(responseData)
	if err != nil {
		log.Println("%v", err)
		return
	}
	w.Write(b)
}

func getColor() string {
	funcs := []func() colorful.Color{
		colorful.FastWarmColor,
		colorful.FastHappyColor,
	}

	return funcs[rand.Intn(len(funcs))]().Hex()
}

func pp(val []byte) {

	var prettyJSON bytes.Buffer
	json.Indent(&prettyJSON, val, "", "\t")
	fmt.Println(prettyJSON.String())
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

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("UP!"))
}
