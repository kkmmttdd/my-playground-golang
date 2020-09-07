package error_sample

import (
	"errors"
	"fmt"
)

type QueryError struct {
	Query string
	Err   error
}

func (e *QueryError) Error() string { return "QueryError:" + e.Query + ":" + e.Err.Error() }

func Render() (string, error) {
	str, err := DB()
	return str, fmt.Errorf("FROM RENDER[%w]", err)
}

func DB() (string, error) {
	err := errors.New("something happened")
	return "string", fmt.Errorf("FROM DB[%w]", err)
}

func Render2() (string, error) {
	str, err := DB2()
	return str, fmt.Errorf("FROM RENDER[%w]", err)
}

func DB2() (string, error) {
	err := errors.New("something happened")
	return "string", &QueryError{Query: "select hogehoge", Err: err}
}

