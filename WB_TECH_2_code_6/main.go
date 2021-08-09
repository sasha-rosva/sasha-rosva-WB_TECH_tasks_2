package main
import (
	"fmt"
)
/*Выполним функцию modifySlice. Т.к. срез является модифицированным указателем на массив,
первая операция в функции изменит первое значение в массиве на который ссылается слайс
Функция append возвращает обновленный слайс, изменяя который мы не затрагиваем изначальный слайс */
func main() {
	var s = []string{"1", "2", "3"}
	modifySlice(s)
	fmt.Println(s)
}

func modifySlice(i []string) {
	i[0] = "3"
	i = append(i, "4")
	i[1] = "5"
	i = append(i, "6")
}
