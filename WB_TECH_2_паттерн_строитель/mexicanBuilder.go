package main

type mexicanBuilder struct {
	sideDish string
	soap   string
	mainDish      string
}

func newMexicanBuilder() *mexicanBuilder {
	return &mexicanBuilder{}
}

func (b *mexicanBuilder) cookSideDish() {
	b.sideDish = "Кесадилья"
}

func (b *mexicanBuilder) cookSoap() {
	b.soap = "Чили кон карне"
}

func (b *mexicanBuilder) cookMainDish() {
	b.mainDish = "Фахитос"
}

func (b *mexicanBuilder) getLaunch() launch {
	return launch{
		sideDish:   b.sideDish,
		soap: b.soap,
		mainDish:      b.mainDish,
	}
}

