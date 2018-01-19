package main

func (s Snake) Head() Point {
	return s.Body.Data[0]
}
