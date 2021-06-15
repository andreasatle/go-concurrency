package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

func (c circle) String() string {
	return fmt.Sprintf("Circle{radius: %f}", c.radius)
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type triangle struct {
	a, b, c float64
}

func (t triangle) String() string {
	return fmt.Sprintf("Triangle{sides: %f, %f, %f}", t.a, t.b, t.c)
}

func (t triangle) area() float64 {
	s := 0.5 * (t.a + t.b + t.c)
	return math.Sqrt(s * (s - t.a) * (s - t.b) * (s - t.c))
}

func (t triangle) angles() []float64 {
	return []float64{angle(t.b, t.c, t.a), angle(t.a, t.c, t.b), angle(t.a, t.b, t.c)}
}

func angle(a, b, c float64) float64 {
	return math.Acos((a*a+b*b-c*c)/(2.0*a*b)) * 180.0 / math.Pi
}

type rectangle struct {
	h, w float64
}

func (r rectangle) String() string {
	return fmt.Sprintf("Rectangle{sides: %f, %f}", r.h, r.w)
}

func (r rectangle) area() float64 {
	return r.h * r.w
}

type shape interface {
	area() float64
}

func main() {
	shapes := []shape{
		circle{radius: 2.0},
		rectangle{h: 3.0, w: 4.0},
		triangle{a: 2.0, b: 3.4, c: 4.0},
	}

	for _, s := range shapes {
		switch s.(type) {
		case circle:
			fmt.Println(s, "I go around in circles")
		case triangle:
			fmt.Println(s, "Looks edgy to me")
		case rectangle:
			fmt.Println(s, "Almost fair and square")
		}
	}

	types := []interface{}{int(4), float64(math.Pi), "foo, bar"}
	for _, t := range types {
		switch t.(type) {
		case int:
			fmt.Println(t, "Integer")
		case float64:
			fmt.Println(t, "Float64")
		case string:
			fmt.Println(t, "String")
		}
	}
}
