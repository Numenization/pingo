package ui

type StateError struct {
	str string
}

type InputError struct {
	str string
}

func (err *StateError) Error() string {
	return err.str
}

func (err *InputError) Error() string {
	return err.str
}
