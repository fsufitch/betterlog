package betterlog

import "fmt"

type LoggingError string

func (err LoggingError) Error() string {
	return fmt.Sprintf("logging error: %s", string(err))
}
