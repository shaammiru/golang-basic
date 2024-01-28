package main

// Interface Implementation
type rectangle struct {
	width, height float32
}

func (r rectangle) area() float32 {
	return r.width * r.height
}

func (r rectangle) perimeter() float32 {
	return 2 * (r.width + r.height)
}

type square struct {
	side float32
}

func (s square) area() float32 {
	return s.side * s.side
}

func (s square) perimeter() float32 {
	return 4 * s.side
}
