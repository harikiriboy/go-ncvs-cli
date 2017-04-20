package commands

import (
	"context"
	"flag"

	"github.com/google/subcommands"
	"github.com/harikiriboy/go-ncvs-cli/ncvs"
	"github.com/harikiriboy/go-ncvs-cli/writer"
)

// DescribeRulePackages is command of DescribeRulePackages
type DescribeRulePackages struct {
	Client ncvs.Client
	Writer writer.Writer

	scanTemplateName string
}

// Name return command name
func (*DescribeRulePackages) Name() string {
	return "describe-rule-packages"
}

// Synopsis return synopsis
func (*DescribeRulePackages) Synopsis() string {
	return "Describe RulePackages"
}

// Usage return usage
func (*DescribeRulePackages) Usage() string {
	return `describe-rule-packages:
	describe-rule-packages
`
}

// SetFlags set flag
func (p *DescribeRulePackages) SetFlags(f *flag.FlagSet) {
}

// Execute execute
func (p *DescribeRulePackages) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	res, err := p.Client.DescribeRulePackages(ncvs.DescribeRulePackagesParams{})
	if err != nil {
		p.Writer.Print("API request Error " + err.Error())
		return subcommands.ExitFailure
	}

	p.Writer.Print(res)
	return subcommands.ExitSuccess
}
