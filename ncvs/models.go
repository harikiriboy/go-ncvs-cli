package ncvs

// ScanTemplate struct
type ScanTemplate struct {
	ScanTemplateName string
	RulePackageNames []string
	ScanTargets      []ScanTarget
	Description      string
	SSHPort          string
	CreatedTime      string
}

// ScanTarget struct
type ScanTarget struct {
	Type             string
	IPAddress        string
	Region           string
	InstanceUniqueID string `json:"InstanceUniqueId"`
}

// ScanHistory struct
type ScanHistory struct {
	ScanHistoryUUID      string
	StartTime            string
	EndTime              string
	Status               string
	ScannedScanTargets   []ScanTarget
	UnScannedScanTargets []ScanTarget
	ScanTemplateName     string
}

// ScanResult struct
type ScanResult struct {
	Rule              Rule
	Severity          int
	Count             int
	ScanResultTargets []ScanResultTarget
	ScanHistoryUUID   string
}

// Rule struct
type Rule struct {
	RuleName      string
	Synopsis      string
	Description   string
	Solution      string
	RiskFactor    string
	CVSSBaseScore string
}

// ScanResultTarget struct
type ScanResultTarget struct {
	IPAddress string
	Port      string
}

// RulePackage struct
type RulePackage struct {
	RulePackageName string
	RuleCount       int
}
