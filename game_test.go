package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoveRequest_IsValidMove(t *testing.T) {
	snake := Snake{
		Body: []Point{
			{X: 1, Y: 1},
			{X: 1, Y: 2},
		},
		ID: "1",
	}
	m := SnakeRequest{
		Board: Board{
			Width:  20,
			Height: 20,
			Snakes: []Snake{
				snake,
			},
		},
		You: snake,
	}

	assert.True(t, m.IsValidMove(UP, false))
	assert.True(t, m.IsValidMove(LEFT, false))
	assert.True(t, m.IsValidMove(RIGHT, false))
	assert.False(t, m.IsValidMove(DOWN, false))

	m.Board.Snakes[0].Body[0].X = 0
	m.Board.Snakes[0].Body[0].Y = 0
	assert.False(t, m.IsValidMove(UP, false))
	assert.False(t, m.IsValidMove(LEFT, false))

	m.Board.Snakes[0].Body[0].X = m.Board.Width - 1
	m.Board.Snakes[0].Body[0].Y = m.Board.Height - 1
	assert.False(t, m.IsValidMove(DOWN, false))
	assert.False(t, m.IsValidMove(RIGHT, false))
}

func TestMoveRequest_IsLocationEmpty(t *testing.T) {
	m := SnakeRequest{
		Board: Board{
			Width:  20,
			Height: 20,
		},
	}
	snake := Snake{
		Body: []Point{
			{X: 1, Y: 1},
			{X: 1, Y: 2},
		},
	}
	m.Board.Snakes = append(m.Board.Snakes, snake)
	assert.False(t, m.IsLocationEmpty(Point{1, 1}))
	assert.True(t, m.IsLocationEmpty(Point{2, 1}))
}

func TestMoveRequest_FindMoveToNearestFood(t *testing.T) {
	snake := Snake{
		Body: []Point{
			{X: 1, Y: 1},
			{X: 1, Y: 2},
		},
		ID: "1",
	}
	m := SnakeRequest{
		Board: Board{
			Width:  20,
			Height: 20,
			Food: []Point{
				{X: 4, Y: 1},
			},
			Snakes: []Snake{
				snake,
			},
		},
		You: snake,
	}

	move := m.FindMoveToNearestFood()
	assert.Equal(t, RIGHT, move)

	m.Board.Food[0].X = 0
	move = m.FindMoveToNearestFood()
	assert.Equal(t, LEFT, move)

	m.Board.Food[0].X = 1
	m.Board.Food[0].Y = 0
	move = m.FindMoveToNearestFood()
	assert.Equal(t, UP, move)

	m.Board.Food[0].Y = 5
	move = m.FindMoveToNearestFood()
	assert.Equal(t, NOOP, move)
}

func TestMoveRequest_AddNodes(t *testing.T) {
	snake := Snake{
		Body: []Point{
			{X: 2, Y: 1},
			{X: 1, Y: 1},
			{X: 1, Y: 2},
			{X: 1, Y: 3},
			{X: 2, Y: 3},
			{X: 3, Y: 3},
			{X: 3, Y: 2},
		},
		ID: "1",
	}
	m := SnakeRequest{
		Board: Board{
			Width:  20,
			Height: 20,
			Food: []Point{
				{X: 4, Y: 1},
			},
			Snakes: []Snake{
				snake,
			},
		},
		You: snake,
	}
	assert.True(t, m.SearchForClosedArea(Point{2, 2}))
	assert.False(t, m.SearchForClosedArea(Point{2, 0}))
}

func TestMoveRequest_CheckForPossibleKills(t *testing.T) {
	snake1 := Snake{
		Body: []Point{
			{X: 1, Y: 1},
			{X: 1, Y: 2},
		},
		ID: "1",
	}
	snake2 := Snake{
		ID: "2",
		Body: []Point{
			{X: 2, Y: 0},
		},
	}
	m := SnakeRequest{
		Board: Board{
			Width:  20,
			Height: 20,
			Food: []Point{
				{X: 4, Y: 1},
			},
			Snakes: []Snake{
				snake1,
				snake2,
			},
		},
		You: snake1,
	}
	assert.Equal(t, UP, m.CheckForPossibleKills())

	m.Board.Snakes[1].Body = append(m.Board.Snakes[1].Body, Point{X: 2, Y: 1})
	assert.Equal(t, NOOP, m.CheckForPossibleKills())
}
