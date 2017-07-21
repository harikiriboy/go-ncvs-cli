package main

import (
	"context"
	"flag"
	"os"
	"strconv"

	"github.com/google/subcommands"
	"github.com/harikiriboy/go-ncvs-cli/commands"
	"github.com/harikiriboy/go-ncvs-cli/ncvs"
	"github.com/harikiriboy/go-ncvs-cli/writer"
)

func main() {
	writer := writer.New(os.Stdout)
	client := ncvs.NewClient(getAPIEndpoint(), isIgnoreSSLCertsErrors())

	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")

	subcommands.Register(&commands.Version{Writer: writer}, "version")
	subcommands.Register(&commands.CreateScanTemplate{Client: client, Writer: writer}, "create-scan-template")
	subcommands.Register(&commands.DeleteScanTemplate{Client: client, Writer: writer}, "delete-scan-template")
	subcommands.Register(&commands.DescribeRulePackages{Client: client, Writer: writer}, "describe-rule-packages")
	subcommands.Register(&commands.DescribeRulePackageAttributes{Client: client, Writer: writer}, "describe-rule-package-attributes")
	subcommands.Register(&commands.DescribeScanHistories{Client: client, Writer: writer}, "describe-scan-histories")
	subcommands.Register(&commands.DescribeScanResults{Client: client, Writer: writer}, "describe-scan-results")
	subcommands.Register(&commands.DescribeScanTemplates{Client: client, Writer: writer}, "describe-scan-templates")
	subcommands.Register(&commands.DownloadScanResults{Client: client, Writer: writer}, "download-scan-results")
	subcommands.Register(&commands.ExecuteScan{Client: client, Writer: writer}, "execute-scan")
	subcommands.Register(&commands.UpdateScanTemplate{Client: client, Writer: writer}, "update-scan-template")

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

// IsIgnoreSSLCertsErrors return ignoreSSLCertsErrors flag
func isIgnoreSSLCertsErrors() bool {
	skipVerify := os.Getenv("NIFTY_CLOUD_VSS_SKIP_VERIFY")
	result, _ := strconv.ParseBool(skipVerify)
	return result
}
