package nagios

import (
	"fmt"
	"strings"
)

type Overall struct {
	OKs       int
	Warnings  int
	Criticals int
	Unknowns  int
	Summary   string
	Outputs   []string
}

func (o *Overall) Add(status int, output string) {
	switch status {
	case OK:
		o.OKs++
	case Warning:
		o.Warnings++
	case Critical:
		o.Criticals++
	default:
		o.Unknowns++
	}

	o.Outputs = append(o.Outputs, fmt.Sprintf("[%s] %s", nagiosStatusText[status], output))
}

func (o *Overall) GetStatus() int {
	if o.Criticals > 0 {
		return Critical
	} else if o.Unknowns > 0 {
		return Unknown
	} else if o.Warnings > 0 {
		return Warning
	} else if o.OKs > 0 {
		return OK
	} else {
		return Unknown
	}
}

func (o *Overall) GetSummary() string {
	if o.Summary != "" {
		return o.Summary
	}
	if o.Criticals > 0 {
		o.Summary += fmt.Sprintf("critical=%d ", o.Criticals)
	}
	if o.Unknowns > 0 {
		o.Summary += fmt.Sprintf("unknown=%d ", o.Unknowns)
	}
	if o.Warnings > 0 {
		o.Summary += fmt.Sprintf("warning=%d ", o.Warnings)
	}
	if o.OKs > 0 {
		o.Summary += fmt.Sprintf("ok=%d ", o.OKs)
	}
	if o.Summary == "" {
		return "No status information"
	}
	o.Summary = "states: " + strings.TrimSpace(o.Summary)
	return o.Summary
}

func (o *Overall) GetOutput() string {
	output := o.GetSummary() + "\n"

	for _, extra := range o.Outputs {
		output += extra + "\n"
	}

	return output
}
