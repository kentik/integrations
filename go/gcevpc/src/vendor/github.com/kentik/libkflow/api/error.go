package api

import "fmt"

type Error struct {
	StatusCode int
	Message    string
}

func (e *Error) Error() string {
	switch {
	case e.StatusCode == 403:
		return fmt.Sprintf("api: unauthorized (403)")
	case e.Message != "":
		return fmt.Sprintf("api: %s (%d)", e.Message, e.StatusCode)
	default:
		return fmt.Sprintf("api: HTTP status code %d", e.StatusCode)
	}
}

func IsErrorWithStatusCode(err error, code int) bool {
	if err, ok := err.(*Error); ok {
		return err.StatusCode == code
	}
	return false
}
