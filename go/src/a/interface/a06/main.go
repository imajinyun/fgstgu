package main

import "fmt"

// Flyer interface.
type Flyer interface {
	Fly()
}

// Walker interface.
type Walker interface {
	Walk()
}

type bird struct{}

// Fly func.
func (b *bird) Fly() {
	fmt.Println("Bird: call fly func")
}

// Walk func.
func (b *bird) Walk() {
	fmt.Println("Bird: call walk func")
}

type pig struct{}

func (p *pig) Walk() {
	fmt.Println("Pig: call walk func")
}

func main() {
	animals := map[string]interface{}{
		"bird": new(bird),
		"pig":  new(pig),
	}
	for name, ins := range animals {
		f, isFlyer := ins.(Flyer)
		w, isWalker := ins.(Walker)

		fmt.Printf("Name: %s, isFlyer: %v, isWalker: %v\n", name, isFlyer, isWalker)

		if isFlyer {
			f.Fly()
		}
		if isWalker {
			w.Walk()
		}
	}

	p1 := new(pig)
	var a Walker = p1
	// a.(*bird) -> panic: interface conversion: main.Walker is *main.pig, not *main.bird
	p2 := a.(*pig)
	fmt.Printf("p1=%p p2=%p\n", p1, p2)
}
