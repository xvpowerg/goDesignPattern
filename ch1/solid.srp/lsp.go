package main

import "fmt"

type Sized interface {
	GetWidth() int
	SetWidth(width int)
	GetHeight() int
	SetHeight(height int)
}
type Rectangle struct {
	width, height int
}

func (r *Rectangle) GetWidth() int {
	return r.width
}
func (r *Rectangle) GetHeight() int {
	return r.height
}

func (r *Rectangle) SetWidth(width int) {
	r.width = width
}
func (r *Rectangle) SetHeight(height int) {
	r.height = height
}

//使用Embedding
type Square struct {
	Rectangle
}

func (s *Square) SetWidth(width int) {
	s.width = width
	s.height = width
}
func (s *Square) SetHeight(height int) {
	s.width = height
	s.height = height
}

func Uselt(sized Sized) {
	width := sized.GetWidth()
	height := sized.GetHeight()
	fmt.Printf("%d * %d = %d\n", width, height, width*height)
}

func main() {
	rc := &Rectangle{2, 3}
	sq := &Square{}
	sq.SetHeight(10)
	Uselt(rc)
	Uselt(sq)
}
