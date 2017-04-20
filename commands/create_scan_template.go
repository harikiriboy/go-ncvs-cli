package commands

import (
	"context"
	"flag"

	"github.com/google/subcommands"
	"github.com/harikiriboy/go-ncvs-cli/ncvs"
	"github.com/harikiriboy/go-ncvs-cli/writer"
)

// CreateScanTemplate is command of CreateScanTemplate
type CreateScanTemplate struct {
	Client ncvs.Client
	Writer writer.Writer

	scanTemplateName      string
	sshPort               string
	description           string
	useCustomRulePackages bool
	rulePackageNames      rulePackageNames
	scanTargets           scanTargets
}

// Name return command name
func (*CreateScanTemplate) Name() string {
	return "create-scan-template"
}

// Synopsis return synopsis
func (*CreateScanTemplate) Synopsis() string {
	return "Create ScanTemplates"
}

// Usage return usage
func (*CreateScanTemplate) Usage() string {
	return `create-scan-template:
	create-scan-template
	  [-scan-template-name]
	  [-ssh-port]
	  [-description]
	  [-use-custom-rule-packages]
	  [--rule-package-names]
	  [--scan-targets <{'Region':'region','InstanceUniqueId': 'instanceUniqueId'}>]
`
}

// SetFlags set flag
func (p *CreateScanTemplate) SetFlags(f *flag.FlagSet) {
	f.StringVar(&p.scanTemplateName, "scan-template-name", "", "scan template name(default: empty)")
	f.StringVar(&p.sshPort, "ssh-port", "", "ssh port(default: empty)")
	f.StringVar(&p.description, "description", "", "description(default: empty)")
	f.BoolVar(&p.useCustomRulePackages, "use-custom-rule-packages", false, "use custom rulePackages(default: false)")
	f.Var(&p.rulePackageNames, "-rule-package-names", "the list of rulePackageName(default: [])")
	f.Var(&p.scanTargets, "-scan-targets", "the list of scanTarget(default: [])")
}

// Execute execute
func (p *CreateScanTemplate) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	res, err := p.Client.CreateScanTemplate(ncvs.CreateScanTemplateParams{
		ScanTemplateName:      p.scanTemplateName,
		SSHPort:               p.sshPort,
		Description:           p.description,
		UseCustomRulePackages: p.useCustomRulePackages,
		RulePackageNames:      p.rulePackageNames,
		ScanTargets:           p.scanTargets,
	})
	if err != nil {
		p.Writer.Print("API request Error " + err.Error())
		return subcommands.ExitFailure
	}

	p.Writer.Print(res)
	return subcommands.ExitSuccess
}
