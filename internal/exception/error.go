package exception

import "fmt"

type Error struct {
	Key  string
	Data any
}

func (error *Error) Error() string {
	return fmt.Sprintf("KEY: %s\nDATA: %s", error.Key, error.Data)
}
