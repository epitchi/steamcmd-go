//go:build windows

package console

import (
	"context"
	"fmt"
	"github.com/UserExistsError/conpty"
	"io"
	"os/exec"
	"strings"
)

type Console struct {
	stdout      io.Writer // stdout is the writer where the output will be written to.
	commandLine string    // commandLine is the command that will be executed.
	conPTY      *conpty.ConPty
	args        []string
	exitCode    uint32
	Parser      *Parser
}

// Run executes the command in steamcmd and returns the exit code. Exit code does not need to be 0 to return no errors (error is for executing the pseudoconsole)
func (c *Console) Run() (uint32, error) {
	var err error

	if c.commandLine == "" {
		_, err := exec.LookPath(c.commandLine)
		if err != nil {
			return 1, fmt.Errorf("Steamcmd not found: %v\n", err)
		}
	}

	var a []string
	a = append(a, c.commandLine)
	a = append(a, c.args...)

	c.conPTY, err = conpty.Start(c.commandLine + " " + strings.Join(a, " "))
	if err != nil {
		return 1, err
	}
	defer c.conPTY.Close()

	d := &duplicateWriter{
		writer1: c.Parser,
		writer2: c.stdout,
	}

	go io.Copy(d, c.conPTY)

	c.exitCode, err = c.conPTY.Wait(context.Background())
	if err != nil {
		return 1, err
	}

	return c.exitCode, nil
}
