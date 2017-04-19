package commands

import (
	"context"
	"flag"

	"github.com/google/subcommands"
	"github.com/harikiriboy/go-ncvs-cli/ncvs"
	"github.com/harikiriboy/go-ncvs-cli/writer"
)

// DownloadScanResults is command of DownloadScanResults
type DownloadScanResults struct {
	Client ncvs.Client
	Writer writer.Writer

	scanHistoryUUID string
}

// Name return command name
func (*DownloadScanResults) Name() string {
	return "download-scan-results"
}

// Synopsis return synopsis
func (*DownloadScanResults) Synopsis() string {
	return "Download ScanResults"
}

// Usage return usage
func (*DownloadScanResults) Usage() string {
	return `download-scan-results:
	download-scan-results
	  [-scan-history-uuid]
`
}

// SetFlags set flag
func (p *DownloadScanResults) SetFlags(f *flag.FlagSet) {
	f.StringVar(&p.scanHistoryUUID, "scan-history-uuid", "", "scan history uuid(default: empty)")
}

// Execute execute
func (p *DownloadScanResults) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	res, err := p.Client.DownloadScanResults(ncvs.DownloadScanResultsParams{ScanHistoryUUID: p.scanHistoryUUID})
	if err != nil {
		p.Writer.Print("API request Error " + err.Error())
		return subcommands.ExitFailure
	}

	p.Writer.Print(res)
	return subcommands.ExitSuccess
}
