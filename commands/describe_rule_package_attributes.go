package commands

import (
	"context"
	"flag"

	"github.com/google/subcommands"
	"github.com/harikiriboy/go-ncvs-cli/ncvs"
	"github.com/harikiriboy/go-ncvs-cli/writer"
)

// DescribeRulePackageAttributes is command of DescribeRulePackageAttributes
type DescribeRulePackageAttributes struct {
	Client ncvs.Client
	Writer writer.Writer

	rulePackageName string
	marker          string
	maxResults      int
}

// Name return command name
func (*DescribeRulePackageAttributes) Name() string {
	return "describe-rule-package-attributes"
}

// Synopsis return synopsis
func (*DescribeRulePackageAttributes) Synopsis() string {
	return "Describe RulePackageAttributes"
}

// Usage return usage
func (*DescribeRulePackageAttributes) Usage() string {
	return `describe-rule-package-attributes:
	describe-rule-package-attributes
	  [-rule-package-name]
	  [-marker]
	  [-max-results]
`
}

// SetFlags set flag
func (p *DescribeRulePackageAttributes) SetFlags(f *flag.FlagSet) {
	f.StringVar(&p.rulePackageName, "rule-package-name", "", "rule package name(default: empty)")
	f.StringVar(&p.marker, "marker", "", "acquisition position for page nation(default: empty)")
	f.IntVar(&p.maxResults, "max-results", 100, "max result count(default: 100)")
}

// Execute execute
func (p *DescribeRulePackageAttributes) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	res, err := p.Client.DescribeRulePackageAttributes(ncvs.DescribeRulePackageAttributesParams{RulePackageName: p.rulePackageName, Marker: p.marker, MaxResults: p.maxResults})
	if err != nil {
		p.Writer.Print("API request Error " + err.Error())
		return subcommands.ExitFailure
	}

	p.Writer.Print(res)
	return subcommands.ExitSuccess
}
