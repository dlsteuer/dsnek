package main

import "fmt"

type GameStartRequest struct {
	GameId string `json:"game_id"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}

type GameStartResponse struct {
	Color    string  `json:"color"`
	HeadUrl  *string `json:"head_url,omitempty"`
	Name     string  `json:"name"`
	Taunt    *string `json:"taunt,omitempty"`
	HeadType *string `json:"head_type,omitempty"`
	TailType *string `json:"tail_type,omitempty"`
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

// interface Snake {
// 	id: string;
// 	object: 'snake';
// 	body: List<Point>;
// 	health: number;
// 	taunt: string;
// 	name: string;
//   }

//   interface List<T> {
// 	object: 'list';
// 	data: T[];
//   }

//   interface Point {
// 	object: 'point';
// 	x: number;
// 	y: number;
//   }

//   interface World {
// 	object: 'world';
// 	id: number;
// 	you: Snake;
// 	snakes: List<Snake>;
// 	height: number;
// 	width: number;
// 	turn: number;
// 	food: List<Point>;
//   }

type MoveRequest struct {
	ID     int `json:"id"`
	You    Snake
	Snakes []Snake
	Height int
	Width  int
	Turn   int
	Food   []Point
}

type Snake struct {
	Body   []Point
	Health int
	ID     string `json:"id"`
	Name   string `json:"name"`
	Taunt  string `json:"taunt"`
}

type MoveResponse struct {
	Move  string  `json:"move"`
	Taunt *string `json:"taunt,omitempty"`
}

type Vector Point
type Point struct {
	X int
	Y int
}
