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

func (m SnakeRequest) GenerateMove() string {
	snake := m.You

	dir := m.CheckForPossibleKills()
	if dir != NOOP {
		return dir
	}

	//foodVectors := m.GetFoodVectors()
	if snake.Health < 35 {
		dir := m.FindMoveToNearestFood()
		if dir != NOOP {
			return dir
		}
	}

	// try and head towards the smallest snake
	smallestSnake := Snake{}
	for _, snake := range m.Board.Snakes {
		if len(smallestSnake.Body) == 0 {
			smallestSnake = snake
		} else if len(smallestSnake.Body) >= len(snake.Body) {
			smallestSnake = snake
		}
	}

	if !smallestSnake.Head().Equals(m.You.Head()) {
		directionVector := m.You.Head().DistanceTo(smallestSnake.Head())
		dir := directionVector.GetValidDirectionFrom(m, true)
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

func (m SnakeRequest) CheckForPossibleKills() string {
	head := m.You.Head()

	for _, dir := range directions {
		newLocation := head.Add(dir)
		if !m.IsLocationEmpty(newLocation) {
			continue
		}

		for _, dir2 := range directions {
			locationToCheck := newLocation.Add(dir2)
			for _, snake := range m.Board.Snakes {
				if snake.Head().Equals(head) {
					continue
				}
				if snake.Head().Equals(locationToCheck) && len(snake.Body) < len(m.You.Body) {
					if m.IsValidMove(dir, true) {
						return dir
					}
				}
			}
		}
	}

	return NOOP
}

func (m SnakeRequest) GetFoodVectors() Vectors {
	head := m.You.Head()
	vectors := Vectors{}
	// Move to closest food
	for _, food := range m.GetFood() {
		vectors = append(vectors, head.DistanceTo(food))
	}

	sort.Sort(vectors)
	return vectors
}

func (m SnakeRequest) FindMoveToNearestFood() string {
	vectors := m.GetFoodVectors()
	for _, closestFood := range vectors {
		dir := closestFood.GetValidDirectionFrom(m, true)
		if dir != NOOP {
			return dir
		}
	}
	return NOOP
}

func (m SnakeRequest) IsValidMove(dir string, spaceCheck bool) bool {
	snake := m.You
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

func (m SnakeRequest) CheckForPotentialDeath(p Point) bool {
	me := m.You
	for _, dir := range directions {
		check := p.Add(dir)
		for _, snake := range m.Board.Snakes {
			head := snake.Head()
			if head.Equals(check) && len(snake.Body) >= len(me.Body) && !head.Equals(me.Head()) {
				return true
			}
		}
	}
	return false
}

func (m SnakeRequest) SearchForClosedArea(p Point) bool {
	availableNodes := Points{p}
	toSearch := Stack{}
	toSearch = toSearch.Push(p)
	var current Point

	for {
		if len(toSearch) == 0 || len(availableNodes) > len(m.You.Body) {
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

	return len(availableNodes) < len(m.You.Body)
}

func (m SnakeRequest) AddNodes(p Point) []Point {
	availableNeighbours := []Point{}
	for _, dir := range directions {
		newPoint := p.Add(dir)
		if m.IsLocationEmpty(newPoint) {
			availableNeighbours = append(availableNeighbours, newPoint)
		}
	}
	return availableNeighbours
}

func (m SnakeRequest) IsLocationEmpty(p Point) bool {
	if p.X < 0 || p.Y < 0 {
		return false
	}

	if p.X >= m.Board.Width || p.Y >= m.Board.Height {
		return false
	}

	for _, snake := range m.Board.Snakes {
		for _, part := range snake.Body {
			if p.Equals(part) {
				return false
			}
		}
	}

	return true
}

func (m SnakeRequest) GetFood() []Point {
	points := []Point{}
	for _, p := range m.Board.Food {
		points = append(points, Point{p.X, p.Y})
	}
	return points
}
