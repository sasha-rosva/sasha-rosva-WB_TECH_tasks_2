package main

import "fmt"

func main() {
	russianBuilder := getBuilder("russian")
	mexicanBuilder := getBuilder("mexican")

	director := newDirector(russianBuilder)
	russianLaunch := director.cookLaunch()

	fmt.Printf("Русский ланч (Закуска): %s\n", russianLaunch.sideDish)
	fmt.Printf("Русский ланч (Суп): %s\n", russianLaunch.soap)
	fmt.Printf("Русский ланч (Основное блюдо): %s\n", russianLaunch.mainDish)

	director.setBuilder(mexicanBuilder)
	mexicanLaunch := director.cookLaunch()

	fmt.Printf("\nМексиканский ланч (Закуска): %s\n", mexicanLaunch.sideDish)
	fmt.Printf("Мексиканский ланч (Суп): %s\n", mexicanLaunch.soap)
	fmt.Printf("Мексиканский ланч (Основное блюдо): %s\n", mexicanLaunch.mainDish)

}
