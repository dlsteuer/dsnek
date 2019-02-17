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

var heads = []string{"beluga", "bendr", "dead", "evil", "fang", "pixel", "regular", "safe", "sand-worm", "shades", "silly", "smile", "tongue"}
var tails = []string{"block-bum", "bolt", "curled", "fat-rattle", "freckled", "hook", "pixel", "regular", "round-bum", "sharp", "skinny", "small-rattle"}
var taunts = []string{
	"A token of gratitude is nonsensical, much like me.",
	"Lucky number slevin has its world rocked by trees (or rocks).",
	"The body of mind slips on a banana peel.",
	"Sixty-four jumps both ways.",
	"Camouflage paint is not yet ready to die.",
	"Organizational culture brings both pleasure and pain.",
}

func str(str string) *string {
	return &str
}

func end(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("Game end - %v\n", string(body))
}

func start(w http.ResponseWriter, r *http.Request) {
	var requestData GameStartRequest
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("Game starting - %v\n", string(body))
	err = json.Unmarshal(body, &requestData)
	if err != nil {
		log.Println(err)
		return
	}

	responseData := GameStartResponse{
		Color:    getColor(),
		Name:     fake.Word(),
		HeadUrl:  str("https://picsum.photos/50/50"),
		HeadType: str(heads[rand.Intn(len(heads))]),
		TailType: str(tails[rand.Intn(len(tails))]),
		Taunt:    str(taunts[rand.Intn(len(taunts))]),
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
	var requestData SnakeRequest
	val, err := ioutil.ReadAll(r.Body)
	json.Unmarshal(val, &requestData)
	responseData := MoveResponse{
		Move:  requestData.GenerateMove(),
		Taunt: str(taunts[rand.Intn(len(taunts))]),
	}
	log.Printf("Move request - direction:%v - taunt: %v\n", responseData.Move, *responseData.Taunt)
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
