package main

import "fmt"

type typer interface {
	getType() string
	accept(visitor)
}
type visitor interface {
	visitMap(*Map)
	visitArray(*Array)
}
type Map struct{ myType string }

func (m *Map) getType() string  { return m.myType }
func (m *Map) accept(v visitor) { v.visitMap(m) }

type Array struct{ myType string }

func (a *Array) getType() string  { return a.myType }
func (a *Array) accept(v visitor) { v.visitArray(a) }

type Slice struct{ myType string }

func (s *Slice) visitMap(m *Map) {
	if s.myType != m.myType {
		fmt.Printf("Наши типы НЕ равны! У тебя %s, а у меня %s!\n", m.myType, s.myType)
	} else {
		fmt.Println("Наши типы равны!")
	}
}
func (s *Slice) visitArray(a *Array) {
	if s.myType != a.myType {
		fmt.Printf("Наши типы НЕ равны! У тебя %s, а у меня %s!\n", a.myType, s.myType)
	} else {
		fmt.Println("Наши типы равны!")
	}
}

func main() {
	map1 := &Map{"Map"}
	array1 := &Array{"Array"}
	slice := &Slice{"Slice"}

	map1.accept(slice)
	array1.accept(slice)
}
