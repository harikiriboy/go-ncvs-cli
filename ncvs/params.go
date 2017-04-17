package ncvs

import "time"

// DescribeScanTemplatesParams struct
type DescribeScanTemplatesParams struct {
	ScanTemplateName string
}

// CreateScanTemplateParams struct
type CreateScanTemplateParams struct {
	ScanTemplateName      string
	UseCustomRulePackages bool
	RulePackageNames      []string
	ScanTargets           []ScanTarget
	Description           string
	SSHPort               string
}

// DeleteScanTemplateParams struct
type DeleteScanTemplateParams struct {
	ScanTemplateName string
}

// UpdateScanTemplateParams struct
type UpdateScanTemplateParams struct {
	ScanTemplateName       string
	UpdateScanTemplateName string
	RulePackageNames       []string
	ScanTargets            []ScanTarget
	Description            string
	SSHPort                string
}

// ExecuteScanParams struct
type ExecuteScanParams struct {
	ScanTemplateName string
}

// DescribeScanHistoriesParams struct
type DescribeScanHistoriesParams struct {
	ScanTemplateName string
	StartTime        *time.Time
	EndTime          *time.Time
}

// DescribeScanResultsParams struct
type DescribeScanResultsParams struct {
	ScanHistoryUUID string
}

// DownloadScanResultsParams struct
type DownloadScanResultsParams struct {
	ScanHistoryUUID string
}

// DescribeRulePackagesParams struct
type DescribeRulePackagesParams struct{}

// DescribeRulePackageAttributesParams struct
type DescribeRulePackageAttributesParams struct {
	RulePackageName string
	Marker          string
	MaxResults      int
}
