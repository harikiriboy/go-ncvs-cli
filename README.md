# go-ncvs-cli

[![Build Status](https://travis-ci.org/harikiriboy/go-ncvs-cli.svg?branch=master)](https://travis-ci.org/harikiriboy/go-ncvs-cli)
[![Go Doc](https://godoc.org/github.com/harikiriboy/go-ncvs-cli?status.svg)](http://godoc.org/github.com/harikiriboy/go-ncvs-cli)
[![Go Report](https://goreportcard.com/badge/github.com/harikiriboy/go-ncvs-cli)](https://goreportcard.com/report/github.com/harikiriboy/go-ncvs-cli)
[![Coverage Status](https://coveralls.io/repos/github/harikiriboy/go-ncvs-cli/badge.svg?branch=master)](https://coveralls.io/github/harikiriboy/go-ncvs-cli?branch=master)

## Install

```
$ curl -L -o /usr/local/bin/ncvs-cli https://github.com/harikiriboy/go-ncvs-cli/releases/download/v1.0.0/ncvs-cli
$ chmod +x /usr/local/bin/ncvs-cli
$ ncvs-cli version
```

## Usage

### Using your Credentials

Set NIFTY Cloud credentials to environment variables with your credentials. 

```
$ export AWS_ACCESS_KEY_ID=XXXXXXXXX
$ export AWS_SECRET_ACCESS_KEY=XXXXXXXX
```

### Subcommands

```
Usage: ncvs-cli <flags> <subcommand> <subcommand args>

Subcommands:
        commands         list all command names
        flags            describe all known top-level flags
        help             describe subcommands and their syntax

Subcommands for create-scan-template:
        create-scan-template  Create ScanTemplates

Subcommands for delete-scan-template:
        delete-scan-template  Delete ScanTemplates

Subcommands for describe-rule-package-attributes:
        describe-rule-package-attributes  Describe RulePackageAttributes

Subcommands for describe-rule-packages:
        describe-rule-packages  Describe RulePackages

Subcommands for describe-scan-histories:
        describe-scan-histories  Describe ScanHistories

Subcommands for describe-scan-results:
        describe-scan-results  Describe ScanResults

Subcommands for describe-scan-templates:
        describe-scan-templates  Describe ScanTemplates

Subcommands for download-scan-results:
        download-scan-results  Download ScanResults

Subcommands for execute-scan:
        execute-scan     Execute Scan

Subcommands for update-scan-template:
        update-scan-template  Update ScanTemplates

Subcommands for version:
        version          Show version
```

### Examples

#### create-scan-template

```
$ ncvs-cli create-scan-template \
  -scan-template-name test \
  -ssh-port 22 \
  -description memo \
  -use-custom-rule-packages \
  --rule-package-names Backdoors \
  --rule-package-names DNS \
  --scan-targets '{"Region": "east-1", "InstanceUniqueId": "i-0ew602z8"}' \
  --scan-targets '{"Region": "west-1", "InstanceUniqueId": "i-06xh9pth"}'
```

#### delete-scan-template

```
$ ncvs-cli delete-scan-template -scan-template-name test
```

#### describe-rule-package-attributes

```
$ ncvs-cli describe-rule-package-attributes \
  -rule-package-name DNS \
  -max-results 20 \
  -marker MDAwMDAyNTEzNA==
```

#### describe-rule-packages

```
$ ncvs-cli describe-rule-packages
```

#### describe-scan-histories

```
$ ncvs-cli describe-scan-histories -scan-template-name test
```

#### describe-scan-results

```
$ ncvs-cli describe-scan-results -scan-history-uuid 501289ad-89bf-4fee-9c89-9555f20c96ca
```

#### describe-scan-templates

```
$ ncvs-cli describe-scan-templates -scan-template-name test
```

#### download-scan-results

```
$ ncvs-cli download-scan-results -scan-history-uuid 501289ad-89bf-4fee-9c89-9555f20c96ca
```

#### execute-scan

```
$ ncvs-cli execute-scan -scan-template-name test
```

#### update-scan-template

```
$ ncvs-cli update-scan-template \
  -scan-template-name test \
  -update-scan-template-name test2 \
  -ssh-port 22 \
  -description memo \
  -use-custom-rule-packages \
  --rule-package-names Backdoors \
  --rule-package-names DNS \
  --scan-targets '{"Region": "east-1", "InstanceUniqueId": "i-0ew602z8"}' \
  --scan-targets '{"Region": "west-1", "InstanceUniqueId": "i-06xh9pth"}'
```
