package nagios

import (
	"fmt"
	"os"
)

const (
	OK       = 0
	Warning  = 1
	Critical = 2
	Unknown  = 3
)

var nagiosStatusText = []string{
	"OK",
	"WARNING",
	"CRITICAL",
	"UNKNOWN",
}

func Exit(rc int, output string) {
	fmt.Printf("%s - %s\n", nagiosStatusText[rc], output)
	os.Exit(rc)
}

func ExitError(err error) {
	Exit(Unknown, err.Error())
}
