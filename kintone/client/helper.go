package client

import (
	"time"
)

func WaitUntil(maxAttempt int, fn func() (bool, error)) error {
	for i := 0; i < maxAttempt; i++ {

		ok, err := fn()
		if err != nil {
			return err
		}
		if ok {
			return nil
		}

		time.Sleep(5 * time.Second)
	}
	return nil
}
