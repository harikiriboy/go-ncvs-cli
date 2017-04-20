package commands

import (
	"bytes"
	"context"
	"flag"
	"testing"

	"github.com/harikiriboy/go-ncvs-cli/writer"
	. "github.com/smartystreets/goconvey/convey"
)

func TestDescribeRulePackageAttributes(t *testing.T) {
	Convey("Tests DescribeRulePackageAttributes", t, func() {
		client := newMockClient()
		out := new(bytes.Buffer)

		cmd := DescribeRulePackageAttributes{Writer: writer.New(out), Client: client}
		ctx := context.Background()
		flg := flag.NewFlagSet("", flag.ContinueOnError)

		Convey("should success name", func() {
			So(cmd.Name(), ShouldEqual, "describe-rule-package-attributes")
		})
		Convey("should success synopsis", func() {
			So(cmd.Synopsis(), ShouldEqual, "Describe RulePackageAttributes")
		})
		Convey("should success usage", func() {
			So(cmd.Usage(), ShouldEqual, "describe-rule-package-attributes:\n	describe-rule-package-attributes\n	  [-rule-package-name]\n	  [-marker]\n	  [-max-results]\n")
		})
		Convey("should success set flags", func() {
			cmd.SetFlags(flg)

			So(flg, ShouldNotBeNil)
		})
		Convey("should success execute", func() {
			cmd.Execute(ctx, flg)
			So(out.String(), ShouldEqual, "this is dummy response.")
		})
		Convey("should success return error message", func() {
			errClient := newMockErrorClient()
			errCmd := DescribeRulePackageAttributes{Writer: writer.New(out), Client: errClient}
			errCmd.Execute(ctx, flg)
			So(out.String(), ShouldEqual, "API request Error this is dummy error.")
		})
	})
}
