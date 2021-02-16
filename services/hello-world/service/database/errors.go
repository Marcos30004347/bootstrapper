package database

import "fmt"

type DatabaseError struct {
	code    int
	message string
}

func (e DatabaseError) Error() string {
	return fmt.Sprintf("%d - %s", e.code, e.message)
}
