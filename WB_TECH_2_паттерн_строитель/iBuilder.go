package main

type iBuilder interface {
	cookSideDish()
	cookSoap()
	cookMainDish()
	getLaunch() launch
}

func getBuilder(builderType string) iBuilder {
	if builderType == "russian" {
		return &russianBuilder{}
	}

	if builderType == "mexican" {
		return &mexicanBuilder{}
	}
	return nil
}
