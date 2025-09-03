package utils

import "fmt"

// for unique_voilation

type ModelExistsError struct {
	StatusCode int
	ModelName  string
	Cause      string
}

func (m *ModelExistsError) Error() string {
	return fmt.Sprintf(
		"%s already exists with this %s",
		m.ModelName,
		m.Cause,
	)
}

// for does not exists error

type ModelDoesNotExistsError struct {
	StatusCode int
	ModelName  string
}

func (m *ModelDoesNotExistsError) Error() string {
	return fmt.Sprintf(
		"%s does not exists",
		m.ModelName,
	)
}
