package check_x

import (
	"os"
	"fmt"
)

//Exit returns with the given returncode and message and performancedata
func Exit(state State, msg string) {
	fmt.Printf("%s - %s|%s\n", state.name, msg, PrintPerformanceData())
	os.Exit(state.code)
}
//ExitWithoutPerfdata returns with the given returncode and message
func ExitWithoutPerfdata(state State, msg string) {
	fmt.Printf("%s - %s\n", state.name, msg)
	os.Exit(state.code)
}
