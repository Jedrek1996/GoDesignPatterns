package main

import "fmt"

//LSP - the behaviour

/*
of implementors of a particular type, like in this case, the size the interface should not break the
core fundamental behaviors that you rely on.
So you should be able to continue taking sized objects instead of somehow figuring out in here, for
example, by doing type checks whether you have a rectangle or a square, it should still work in the
generalised case

--objects of a superclass should be replaceable with objects of its subclasses without breaking the application.
*/
type Sized interface {
	GetWidth() int
	SetWidth(width int)
	GetHeight() int
	SetHeight(height int)
}

type Rectangle struct {
	width, height int
}

func (r Rectangle) GetWidth() int {
	return r.width
}

func (r *Rectangle) SetWidth(width int) {
	r.width = width
}

func (r *Rectangle) GetHeight() int {
	return r.height
}

func (r *Rectangle) SetHeight(height int) {
	r.height = height
}

type Square struct {
	Rectangle
}

func NewSquare(size int) *Square {
	sq := Square{}
	sq.width = size
	sq.height = size
	return &sq
}

func UseIt(sized Sized) {
	width := sized.GetHeight()
	sized.SetHeight(10)
	expectedArea := 10 * width
	actualArea := sized.GetWidth() * sized.GetHeight()

	fmt.Printf("Expected area: %v, Actual area:%v\n", expectedArea, actualArea)

}

func main() {
	rc := &Rectangle{2, 3}
	UseIt(rc)

}
