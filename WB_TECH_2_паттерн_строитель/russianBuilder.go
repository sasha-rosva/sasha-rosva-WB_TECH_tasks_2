package main

type russianBuilder struct {
	sideDish string
	soap   string
	mainDish      string
}

func newRussianBuilder() *russianBuilder {
	return &russianBuilder{}
}

func (b *russianBuilder) cookSideDish() {
	b.sideDish = "Заливное из судака"
}

func (b *russianBuilder) cookSoap() {
	b.soap = "Борщ"
}

func (b *russianBuilder) cookMainDish() {
	b.mainDish = "Бефстроганов из говядины"
}

func (b *russianBuilder) getLaunch() launch {
	return launch{
		sideDish:   b.sideDish,
		soap: b.soap,
		mainDish:      b.mainDish,
	}
}
