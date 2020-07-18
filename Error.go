package go_dj

type errorStruct struct {
	Message string
}

func (e errorStruct) Error() string {
	return e.Message
}

func newError(message string) *errorStruct {
	return &errorStruct{Message: message}
}
