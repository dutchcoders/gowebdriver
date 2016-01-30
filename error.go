package webdriver

import "fmt"

type Error struct {
	State  string
	Status int64
}

func (de *Error) Error() string {
	return fmt.Sprintf("%s (%d)", de.State, de.Status)
}
