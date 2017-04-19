package commands

import (
	"context"
	"flag"

	"github.com/google/subcommands"
	"github.com/harikiriboy/go-ncvs-cli/ncvs"
	"github.com/harikiriboy/go-ncvs-cli/writer"
)

// DeleteScanTemplate is command of DeleteScanTemplate
type DeleteScanTemplate struct {
	Client ncvs.Client
	Writer writer.Writer

	scanTemplateName string
}

// Name return command name
func (*DeleteScanTemplate) Name() string {
	return "delete-scan-template"
}

// Synopsis return synopsis
func (*DeleteScanTemplate) Synopsis() string {
	return "Delete ScanTemplates"
}

// Usage return usage
func (*DeleteScanTemplate) Usage() string {
	return `delete-scan-template:
	delete-scan-template
	  [-scan-template-name]
`
}

// SetFlags set flag
func (p *DeleteScanTemplate) SetFlags(f *flag.FlagSet) {
	f.StringVar(&p.scanTemplateName, "scan-template-name", "", "scan template name(default: empty)")
}

// Execute execute
func (p *DeleteScanTemplate) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	res, err := p.Client.DeleteScanTemplate(ncvs.DeleteScanTemplateParams{ScanTemplateName: p.scanTemplateName})
	if err != nil {
		p.Writer.Print("API request Error " + err.Error())
		return subcommands.ExitFailure
	}

	p.Writer.Print(res)
	return subcommands.ExitSuccess
}
