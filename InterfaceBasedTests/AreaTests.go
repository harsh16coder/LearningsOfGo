package interfacebasedtests

type Shape interface {
	Area() float64
}
type Rectangle struct {
	width  float64
	height float64
}

type Square struct {
	length float64
}

type Triangle struct {
	base   float64
	height float64
}

func (r *Rectangle) Area() float64 {
	return r.width * r.height
}

func (r *Square) Area() float64 {
	return r.length * r.length
}

func (r *Triangle) Area() float64 {
	return (r.base * r.height) / 2
}
