package commands

import (
	"context"
	"flag"

	"github.com/google/subcommands"
	"github.com/harikiriboy/go-ncvs-cli/ncvs"
	"github.com/harikiriboy/go-ncvs-cli/writer"
)

// ExecuteScan is command of ExecuteScan
type ExecuteScan struct {
	Client ncvs.Client
	Writer writer.Writer

	scanTemplateName string
}

// Name return command name
func (*ExecuteScan) Name() string {
	return "execute-scan"
}

// Synopsis return synopsis
func (*ExecuteScan) Synopsis() string {
	return "Execute Scan"
}

// Usage return usage
func (*ExecuteScan) Usage() string {
	return `execute-scan:
	execute-scan
	  [-scan-template-name]
`
}

// SetFlags set flag
func (p *ExecuteScan) SetFlags(f *flag.FlagSet) {
	f.StringVar(&p.scanTemplateName, "scan-template-name", "", "scan template name(default: empty)")
}

// Execute execute
func (p *ExecuteScan) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	res, err := p.Client.ExecuteScan(ncvs.ExecuteScanParams{ScanTemplateName: p.scanTemplateName})
	if err != nil {
		p.Writer.Print("API request Error " + err.Error())
		return subcommands.ExitFailure
	}

	p.Writer.Print(res)
	return subcommands.ExitSuccess
}
