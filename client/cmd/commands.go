package cmd

import (
	cli "github.com/spf13/cobra"
)

// New constructs the 'root' command. It houses all other sub commands
// default output of logging should go to stdout
// interactive output like progress bars should go to stderr
// unless the stdout/err is a tty, colors/progressbar should be disabled
func New() *cli.Command { _ = "STUB: not implemented"; return nil }

// Client related commands

// Will decide later, to add it server side or not
