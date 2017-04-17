package commands

import (
	"context"
	"flag"

	"github.com/google/subcommands"
	"github.com/harikiriboy/go-ncvs-cli/writer"
)

var version string

// Version is command of version
type Version struct {
	Version string
	Writer  writer.Writer
}

// Name return command name
func (*Version) Name() string {
	return "version"
}

// Synopsis return synopsis
func (*Version) Synopsis() string {
	return "Show version"
}

// Usage return usage
func (*Version) Usage() string {
	return `version:
	version
`
}

// SetFlags set flag
func (p *Version) SetFlags(f *flag.FlagSet) {
	return
}

// Execute execute
func (p *Version) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	p.Writer.Printf("ncvs-cli %s\n", p.Version)
	return subcommands.ExitSuccess
}
