package version

import (
	"fmt"
	"runtime"
)

const Name = "playground"

var (
	Version   = "UNKNOWN"
	Branch    = "UNKNOWN"
	Commit    = "UNKNOWN"
	BuildUser = "UNKNOWN"
	BuildDate = "UNKNOWN"
)

func Message() string {
	const format = `%s:%s (Branch: %s, Revision: %s)
  build user: %s
  build date: %s
  go version: %s
`
	return fmt.Sprintf(format, Name, Version, Branch, Commit, BuildUser, BuildDate, runtime.Version())
}
