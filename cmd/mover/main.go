package main

import (
	"fmt"
	"reflect"
)

type Substractable interface {
	~int | ~int32 | ~int64 | ~float32 | ~float64 | ~uint | ~uint32 | ~uint64
}

func Subtract[V Substractable](a, b V) V {
	return a - b
}

// Moveable is a interface that is used to handle many objects that are moveable
type Moveable[S Substractable] interface {
	Move(S)
}

// Person is a person, implements Moveable
type Person[S Substractable] struct {
	Name string
}

func (p Person[S]) Move(meters S) {
	fmt.Printf("%s moved %d meters\n", p.Name, meters)
}

// Car is a test struct for cars, implements Moveable
type Car[S Substractable] struct {
	Name string
}

func (c Car[S]) Move(meters S) {
	fmt.Printf("%s moved %d meters\n", c.Name, meters)
}

// Move is a generic function that takes in a Moveable and moves it
func Move[V Moveable[S], S Substractable](v V, distance S, meters S) S {
	v.Move(meters)
	return Subtract(distance, meters)
}

func main() {
	// John is travelling to his Job
	// His car travel is counted in int
	// And his walking in Float32
	p := Person[float64]{Name: "John"}
	c := Car[int]{Name: "Ferrari"}

	// John has 100 miles to his job
	milesToDestination := 100
	// John moves with the Car
	distanceLeft := Move(c, milesToDestination, 95)
	// John has 5 miles left to walk after parking (phew)
	fmt.Println("DistanceLeft: ", distanceLeft)
	fmt.Println("DistanceType: ", reflect.TypeOf(distanceLeft))

	// Jumps out of Car and Walks to Building
	// Again we need to define the data type to use for the Move function, or else it will default to int
	// So here we have to tell Move to initialize with a Person type, with a float64 value,
	// And that the Subtract data type is also float64
	// [Move[float64], float64]
	// distanceLeft is also a INT, since Move defaulted to int in previous call, so we need to convert it
	newDistanceLeft := Move[Person[float64], float64](p, float64(distanceLeft), 5)
	fmt.Println("newDistanceLeft: ", newDistanceLeft)
	fmt.Println("newDistanceType: ", reflect.TypeOf(newDistanceLeft))
}
