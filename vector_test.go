package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPoint_DistanceTo(t *testing.T) {
	p1 := Point{0, 0}
	p2 := Point{3, 0}
	v := p1.DistanceTo(p2)
	assert.Equal(t, float64(3), v.Magnitude())
}
