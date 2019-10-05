package ignite

import "github.com/kingdonb/footloose/pkg/exec"

// Remove removes an Ignite VM
func Remove(name string) error {
	return exec.CommandWithLogging(execName, "rm", "-f", name)
}
