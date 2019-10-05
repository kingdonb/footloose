package ignite

import (
	"github.com/kingdonb/footloose/pkg/exec"
)

// Start starts an Ignite VM
func Start(name string) error {
	return exec.CommandWithLogging(execName, "start", name)
}
