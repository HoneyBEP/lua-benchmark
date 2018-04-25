package implementation

import (
	"time"
	"fmt"
)

type result struct {
	name string
	time time.Duration
	result string
}

// Returns formatted resultstring
func (res *result) getString() string {
	return fmt.Sprintf("----- %s BENCHMARK -----\n" +
		"Time: %d\n" +
		"Returned: %s\n" +
		"Finished\n\n", res.name, res.time, res.result)
}

// Adds two result structs which have the same method
func (res *result) add(other *result) (*result, error) {
	if res.name != other.name {
		return nil, fmt.Errorf("methods should be the same, %s and %s are not equal", res, other)
	}
	if res.result != "" || other.result != "" {
		return nil, fmt.Errorf("one of the methods exited incorrectly")
	}

	return &result{res.name, res.time + other.time, ""}, nil
}
