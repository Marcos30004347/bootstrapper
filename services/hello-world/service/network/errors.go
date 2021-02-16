package network

import "fmt"

type networkError struct {
	code    int
	message string
}

func (e *networkError) Error() string {
	return fmt.Sprintf("%d - %s", e.code, e.message)
}
