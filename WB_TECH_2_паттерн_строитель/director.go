package main

type director struct {
	builder iBuilder
}

func newDirector(b iBuilder) *director {
	return &director{
		builder: b,
	}
}

func (d *director) setBuilder(b iBuilder) {
	d.builder = b
}

func (d *director) cookLaunch() launch {
	d.builder.cookSideDish()
	d.builder.cookSoap()
	d.builder.cookMainDish()
	return d.builder.getLaunch()
}
