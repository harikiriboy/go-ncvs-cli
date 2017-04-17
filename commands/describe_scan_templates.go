package commands

import (
	"context"
	"flag"

	"github.com/google/subcommands"
	"github.com/harikiriboy/go-ncvs-cli/ncvs"
	"github.com/harikiriboy/go-ncvs-cli/writer"
)

// DescribeScanTemplates is command of DescribeScanTemplates
type DescribeScanTemplates struct {
	Client ncvs.Client
	Writer writer.Writer

	scanTemplateName string
}

// Name return command name
func (*DescribeScanTemplates) Name() string {
	return "describe-scan-templates"
}

// Synopsis return synopsis
func (*DescribeScanTemplates) Synopsis() string {
	return "Describe ScanTemplates"
}

// Usage return usage
func (*DescribeScanTemplates) Usage() string {
	return `describe-scan-templates:
	describe-scan-templates
	  [-scan-template-name]
`
}

// SetFlags set flag
func (p *DescribeScanTemplates) SetFlags(f *flag.FlagSet) {
	f.StringVar(&p.scanTemplateName, "scan-template-name", "", "scan template name(default: empty)")
}

// Execute execute
func (p *DescribeScanTemplates) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	res, err := p.Client.DescribeScanTemplates(ncvs.DescribeScanTemplatesParams{ScanTemplateName: p.scanTemplateName})
	if err != nil {
		p.Writer.Print("API request Error " + err.Error())
		return subcommands.ExitFailure
	}

	p.Writer.Print(res)
	return subcommands.ExitSuccess
}
