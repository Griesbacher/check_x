package check_x

import (
	"time"
	"fmt"
)

//StartTimeout starts an timeout, which will end the program with an unknown code
func StartTimeout(duration time.Duration) {
	go func() {
		time.Sleep(duration)
		ExitWithoutPerfdata(Unknown, fmt.Sprintf("Timeout reached after %fs", duration.Seconds()))
	}()
}
