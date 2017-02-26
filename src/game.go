package main

import (
	"math/rand"
	"sort"
)

const UP = "up"
const DOWN = "down"
const LEFT = "left"
const RIGHT = "right"
const NOOP = "no-op"

var directions = []string{UP, DOWN, LEFT, RIGHT}

func (m MoveRequest) GenerateMove() string {
	snake := m.MySnake()
	foodVectors := m.GetFoodVectors()
	if snake.HealthPoints < 35 || foodVectors[0].Magnitude() < 5 {
		dir := m.FindMoveToNearestFood()
		if dir != NOOP {
			return dir
		}
	}

	// fill space
	for _, i := range rand.Perm(4) {
		dir := directions[i]
		if m.IsValidMove(dir, true) {
			return dir
		}
	}
	// once more without space check, maybe we can prolong death
	for _, i := range rand.Perm(4) {
		dir := directions[i]
		if m.IsValidMove(dir, false) {
			return dir
		}
	}
	// gonna die, just go up
	return UP
}

func (m MoveRequest) GetFoodVectors() Vectors {
	head := m.MySnake().Head()
	vectors := Vectors{}
	// Move to closest food
	for _, food := range m.GetFood() {
		vectors = append(vectors, head.DistanceTo(food))
	}

	sort.Sort(vectors)
	return vectors
}

func (m MoveRequest) FindMoveToNearestFood() string {
	vectors := m.GetFoodVectors()
	for _, closestFood := range vectors {
		dir := NOOP
		if closestFood.X < 0 {
			dir = LEFT
		} else if closestFood.X > 0 {
			dir = RIGHT
		} else if closestFood.Y > 0 {
			dir = DOWN
		} else if closestFood.Y < 0 {
			dir = UP
		}
		if m.IsValidMove(dir, true) {
			return dir
		}
	}
	return NOOP
}

func (m MoveRequest) IsValidMove(dir string, spaceCheck bool) bool {
	snake := m.MySnake()
	head := snake.Head()
	newLocation := head.Add(dir)
	empty := m.IsLocationEmpty(newLocation)
	if !empty {
		return false
	}

	potentialDeath := m.CheckForPotentialDeath(newLocation)
	if potentialDeath {
		return false
	}

	if spaceCheck {
		blocked := m.SearchForClosedArea(newLocation)
		return !blocked
	}
	return empty
}

func (m MoveRequest) CheckForPotentialDeath(p Point) bool {
	me := m.MySnake()
	for _, dir := range directions {
		check := p.Add(dir)
		for _, snake := range m.Snakes {
			head := snake.Head()
			if head.Equals(check) && len(snake.Coords) > len(me.Coords) && !head.Equals(me.Head()) {
				return true
			}
		}
	}
	return false
}

func (m MoveRequest) SearchForClosedArea(p Point) bool {
	availableNodes := Points{p}
	toSearch := Stack{}
	toSearch = toSearch.Push(p)
	var current Point

	for {
		if len(toSearch) == 0 || len(availableNodes) > len(m.MySnake().Coords) {
			break
		}

		toSearch, current = toSearch.Pop()
		newNodes := m.AddNodes(current)
		for _, node := range newNodes {
			if !availableNodes.Contains(node) {
				availableNodes = append(availableNodes, node)
				toSearch = toSearch.Push(node)
			}
		}
	}

	return len(availableNodes) < len(m.MySnake().Coords)
}

func (m MoveRequest) AddNodes(p Point) []Point {
	availableNeighbours := []Point{}
	for _, dir := range directions {
		newPoint := p.Add(dir)
		if m.IsLocationEmpty(newPoint) {
			availableNeighbours = append(availableNeighbours, newPoint)
		}
	}
	return availableNeighbours
}

func (m MoveRequest) IsLocationEmpty(p Point) bool {
	if p.X < 0 || p.Y < 0 {
		return false
	}

	if p.X >= m.Width || p.Y >= m.Height {
		return false
	}

	for _, snake := range m.Snakes {
		for _, part := range snake.GetBody() {
			if p.Equals(part) {
				return false
			}
		}
	}

	return true
}

func (m MoveRequest) GetFood() []Point {
	points := []Point{}
	for _, p := range m.Food {
		points = append(points, Point{p[0], p[1]})
	}
	return points
}

func (m MoveRequest) MySnake() *Snake {
	for _, snake := range m.Snakes {
		if snake.Id == m.You {
			return &snake
		}
	}
	return nil
}
