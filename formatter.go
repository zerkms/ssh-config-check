package main

import (
	"fmt"
)

var formatStatus = map[bool]string{
	true:  "Ok",
	false: "Error",
}

func simpleStringFormatter(r result) string {
	hostPort := fmt.Sprintf("%s:%d", r.host.hostname, r.host.port)

	if r.host.id != r.host.hostname {
		hostPort = fmt.Sprintf("%s (%s)", r.host.id, hostPort)
	}

	return fmt.Sprintf("%s - %s", formatStatus[r.accessible], hostPort)
}
