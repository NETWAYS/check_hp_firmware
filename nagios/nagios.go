package nagios

import (
	"fmt"
	"github.com/mitchellh/go-ps"
	"os"
	"runtime/debug"
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

func Exit(rc int, output string, args ...interface{}) {
	fmt.Println(nagiosStatusText[rc], "-", fmt.Sprintf(output, args...))
	os.Exit(rc)
}

func ExitError(err error) {
	Exit(Unknown, err.Error())
}

func CatchPanic() {
	ppid := os.Getppid()
	if parent, err := ps.FindProcess(ppid); err == nil {
		if parent.Executable() == "dlv" {
			// seems to be a debugger, don't continue with recover
			return
		}
	}
	if r := recover(); r != nil {
		Exit(Unknown, "Golang encountered a panic: %s\n\n%s", r, debug.Stack())
	}
}
