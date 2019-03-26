package imageProcess

type unknownImageTypeError struct {
}

func (e unknownImageTypeError) Error() string {
	return "UnknownImageType."
}
