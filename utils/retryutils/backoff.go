package retryutls

import (
	"fmt"
	"time"
)

func RetryWithBackoff(operation func() error, retries int) error {
	delay := time.Second
	for i := 0; i < retries; i++ {
		err := operation()
		if err == nil {
			return nil
		}
		time.Sleep(delay)
		delay *= 2
	}
	return fmt.Errorf("operation failed after %d attempts", retries)
}
