package main

import (
	"context"
	"flag"
	"os"

	"github.com/google/subcommands"
	"github.com/harikiriboy/go-ncvs-cli/commands"
	"github.com/harikiriboy/go-ncvs-cli/ncvs"
	"github.com/harikiriboy/go-ncvs-cli/writer"
)

var version string

func main() {

	writer := writer.New(os.Stdout)
	client := ncvs.NewClient(getAPIEndpoint())

	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")

	subcommands.Register(&commands.Version{Version: version, Writer: writer}, "version")
	subcommands.Register(&commands.DescribeScanHistories{Client: client, Writer: writer}, "describe-scan-histories")
	subcommands.Register(&commands.DescribeScanResults{Client: client, Writer: writer}, "describe-scan-results")
	subcommands.Register(&commands.DescribeScanTemplates{Client: client, Writer: writer}, "describe-scan-templates")
	subcommands.Register(&commands.DownloadScanResults{Client: client, Writer: writer}, "download-scan-results")
	subcommands.Register(&commands.ExecuteScan{Client: client, Writer: writer}, "execute-scan")

	flag.Parse()

	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}

// getAPIEndpoint return ncvs api endpoint
func getAPIEndpoint() string {
	endpoint := os.Getenv("NIFTY_CLOUD_VSS_URL")
	if endpoint == "" {
		endpoint = ncvs.DefaultAPIEndPoint
	}
	return endpoint
}
