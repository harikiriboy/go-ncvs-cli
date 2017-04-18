package commands

import (
	"context"
	"flag"

	"github.com/google/subcommands"
	"github.com/harikiriboy/go-ncvs-cli/ncvs"
	"github.com/harikiriboy/go-ncvs-cli/writer"
)

// DescribeScanResults is command of DescribeScanResults
type DescribeScanResults struct {
	Client ncvs.Client
	Writer writer.Writer

	scanHistoryUUID string
}

// Name return command name
func (*DescribeScanResults) Name() string {
	return "describe-scan-results"
}

// Synopsis return synopsis
func (*DescribeScanResults) Synopsis() string {
	return "Describe ScanResults"
}

// Usage return usage
func (*DescribeScanResults) Usage() string {
	return `describe-scan-results:
	describe-scan-results
	  [-scan-history-uuid]
`
}

// SetFlags set flag
func (p *DescribeScanResults) SetFlags(f *flag.FlagSet) {
	f.StringVar(&p.scanHistoryUUID, "scan-history-uuid", "", "scan history uuid(default: empty)")
}

// Execute execute
func (p *DescribeScanResults) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	res, err := p.Client.DescribeScanResults(ncvs.DescribeScanResultsParams{ScanHistoryUUID: p.scanHistoryUUID})
	if err != nil {
		p.Writer.Print("API request Error " + err.Error())
		return subcommands.ExitFailure
	}

	p.Writer.Print(res)
	return subcommands.ExitSuccess
}
