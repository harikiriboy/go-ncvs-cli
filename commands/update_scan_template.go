package commands

import (
	"context"
	"flag"

	"github.com/google/subcommands"
	"github.com/harikiriboy/go-ncvs-cli/ncvs"
	"github.com/harikiriboy/go-ncvs-cli/writer"
)

// UpdateScanTemplate is command of UpdateScanTemplate
type UpdateScanTemplate struct {
	Client ncvs.Client
	Writer writer.Writer

	scanTemplateName       string
	updateScanTemplateName string
	sshPort                string
	description            string
	rulePackageNames       rulePackageNames
	scanTargets            scanTargets
}

// Name return command name
func (*UpdateScanTemplate) Name() string {
	return "update-scan-template"
}

// Synopsis return synopsis
func (*UpdateScanTemplate) Synopsis() string {
	return "Update ScanTemplates"
}

// Usage return usage
func (*UpdateScanTemplate) Usage() string {
	return `update-scan-template:
	update-scan-template
	  [-scan-template-name]
	  [-update-scan-template-name]
	  [-ssh-port]
	  [-description]
	  [--rule-package-names]
	  [--scan-targets <{'Region':'region','InstanceUniqueId': 'instanceUniqueId'}>]
`
}

// SetFlags set flag
func (p *UpdateScanTemplate) SetFlags(f *flag.FlagSet) {
	f.StringVar(&p.scanTemplateName, "scan-template-name", "", "scan template name(default: empty)")
	f.StringVar(&p.updateScanTemplateName, "update-scan-template-name", "", "scan template name(default: empty)")
	f.StringVar(&p.sshPort, "ssh-port", "", "ssh port(default: empty)")
	f.StringVar(&p.description, "description", "", "description(default: empty)")
	f.Var(&p.rulePackageNames, "rule-package-names", "the list of rulePackageName(default: [])")
	f.Var(&p.scanTargets, "scan-targets", "the list of scanTarget(default: [])")
}

// Execute execute
func (p *UpdateScanTemplate) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	res, err := p.Client.UpdateScanTemplate(ncvs.UpdateScanTemplateParams{
		ScanTemplateName:       p.scanTemplateName,
		UpdateScanTemplateName: p.updateScanTemplateName,
		SSHPort:                p.sshPort,
		Description:            p.description,
		RulePackageNames:       p.rulePackageNames,
		ScanTargets:            p.scanTargets,
	})
	if err != nil {
		p.Writer.Print("API request Error " + err.Error())
		return subcommands.ExitFailure
	}

	p.Writer.Print(res)
	return subcommands.ExitSuccess
}
