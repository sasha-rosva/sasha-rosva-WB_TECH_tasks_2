package main

import "fmt"

type Writer interface {
	myWrite(s *Letter)
}
type Letter struct {
	str string
	strategy Writer
}
func (l *Letter) initAlgo(str string,w Writer) {
	l.str = str
	l.strategy = w
}
func (l *Letter) setAlgo(w Writer) {
	l.strategy = w
}
func (l *Letter) Write() {
	l.strategy.myWrite(l)
}
type oneStrategy struct{}
func (o *oneStrategy) myWrite(s *Letter){
	fmt.Println(s.str)
}

type doubleStrategy struct{}
func (d *doubleStrategy) myWrite(s *Letter){
	fmt.Println(s.str+s.str)
}

type tripleStrategy struct{}
func (t *tripleStrategy) myWrite(s *Letter){
	fmt.Println(s.str+s.str+s.str)
}
func main(){
	letter:=&Letter{}
	oneStrategy:=&oneStrategy{}
	doubleStrategy:=&doubleStrategy{}
	tripleStrategy:=&tripleStrategy{}
	letter.initAlgo("Hello!",oneStrategy)
	letter.Write()

	letter.setAlgo(doubleStrategy)
	letter.Write()

	letter.setAlgo(tripleStrategy)
	letter.Write()
}

