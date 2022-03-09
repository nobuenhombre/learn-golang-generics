package main

import "fmt"

type Cloner[C any] interface {
	Clone() C
}

type Cloneable int

func (c Cloneable) Clone() Cloneable {
	return c
}

func CloneAll[T Cloner[T]](originals []T) (clones []T) {
	for _, c := range originals {
		clones = append(clones, c.Clone())
	}

	return clones
}

func main() {
	origs := []Cloneable{1, 2, 3, 4}
	fmt.Println(CloneAll(origs))
}
