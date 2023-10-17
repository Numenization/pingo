package ui

type StateError struct {
	str string
}

func (err *StateError) Error() string {
	return err.str
}
