package main

import "fmt"

type GameStartRequest struct {
	GameId int `json:"id"`
	Height int `json:"height"`
	Width  int `json:"width"`
}

type GameStartResponse struct {
	Color    string  `json:"color"`
	HeadUrl  *string `json:"head_url,omitempty"`
	Name     string  `json:"name"`
	Taunt    *string `json:"taunt,omitempty"`
	HeadType *string
	TailType *string
}

func DereferenceStringSafely(str *string) string {
	if str == nil {
		return ""
	}
	return *str
}

func (gsr GameStartResponse) String() string {
	return fmt.Sprintf("Color: %v\nHeadUrl: %v\nName: %v\nTaunt: %v\nHeadType: %v\nTailType: %v\n",
		gsr.Color,
		DereferenceStringSafely(gsr.HeadUrl),
		gsr.Name,
		DereferenceStringSafely(gsr.Taunt),
		DereferenceStringSafely(gsr.HeadType),
		DereferenceStringSafely(gsr.TailType))
}

type SnakeRequest struct {
	Game  Game  `json:"game"`
	Turn  int64 `json:"turn"`
	Board Board `json:"board"`
	You   Snake `json:"you"`
}

// Game represents the current game state
type Game struct {
	ID string `json:"id"`
}

// Board provides information about the game board
type Board struct {
	Height int64   `json:"height"`
	Width  int64   `json:"width"`
	Food   []Point `json:"food"`
	Snakes []Snake `json:"snakes"`
}

// Snake represents information about a snake in the game
type Snake struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Health int64   `json:"health"`
	Body   []Point `json:"body"`
}

type MoveResponse struct {
	Move  string  `json:"move"`
	Taunt *string `json:"taunt,omitempty"`
}

type Vector Point
type Point struct {
	X int64
	Y int64
}
