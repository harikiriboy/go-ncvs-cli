package commands

import (
	"context"
	"flag"
	"time"

	"github.com/google/subcommands"
	"github.com/harikiriboy/go-ncvs-cli/ncvs"
	"github.com/harikiriboy/go-ncvs-cli/writer"
)

// DescribeScanHistories is command of DescribeScanHistories
type DescribeScanHistories struct {
	Client ncvs.Client
	Writer writer.Writer

	scanTemplateName string
	startTime        string
	endTime          string
}

// Name return command name
func (*DescribeScanHistories) Name() string {
	return "describe-scan-histories"
}

// Synopsis return synopsis
func (*DescribeScanHistories) Synopsis() string {
	return "Describe ScanHistories"
}

// Usage return usage
func (*DescribeScanHistories) Usage() string {
	return `describe-scan-histories:
	describe-scan-histories
	  [-scan-template-name]
	  [-start-time]
	  [-end-time]
`
}

// SetFlags set flag
func (p *DescribeScanHistories) SetFlags(f *flag.FlagSet) {
	f.StringVar(&p.scanTemplateName, "scan-template-name", "", "scan template name(default: empty)")
	f.StringVar(&p.startTime, "start-time", "", "start time yyyy-mm-dd hh:mm:ss (jst)(default: empty)")
	f.StringVar(&p.endTime, "end-time", "", "end time yyyy-mm-dd hh:mm:ss (jst)(default: empty)")
}

// Execute execute
func (p *DescribeScanHistories) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	var st, et *time.Time
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)

	if p.startTime != "" {
		t, timeFormatErr := time.ParseInLocation("2006-01-02 15:04:05", p.startTime, jst)
		if timeFormatErr != nil {
			p.Writer.Print("start-time format error")
			return subcommands.ExitUsageError
		}
		st = &t
	}

	if p.endTime != "" {
		t, timeFormatErr := time.ParseInLocation("2006-01-02 15:04:05", p.endTime, jst)
		if timeFormatErr != nil {
			p.Writer.Print("end-time format error")
			return subcommands.ExitUsageError
		}
		et = &t
	}

	res, err := p.Client.DescribeScanHistories(ncvs.DescribeScanHistoriesParams{ScanTemplateName: p.scanTemplateName, StartTime: st, EndTime: et})
	if err != nil {
		p.Writer.Print("API request Error " + err.Error())
		return subcommands.ExitFailure
	}

	p.Writer.Print(res)
	return subcommands.ExitSuccess
}
