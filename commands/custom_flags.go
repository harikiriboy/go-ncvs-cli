package commands

import (
	"encoding/json"

	"github.com/harikiriboy/go-ncvs-cli/ncvs"
)

type (
	rulePackageNames []string
	scanTargets      []ncvs.ScanTarget
)

func (i *rulePackageNames) String() string {
	return "the list of rulePackageName"
}

func (i *rulePackageNames) Set(value string) error {
	*i = append(*i, value)
	return nil
}

func (i *scanTargets) String() string {
	return "the list of scanTarget"
}

func (i *scanTargets) Set(value string) error {
	var scanTarget ncvs.ScanTarget
	json.Unmarshal([]byte(value), &scanTarget)
	*i = append(*i, scanTarget)
	return nil
}
