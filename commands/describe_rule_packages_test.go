package commands

import (
	"bytes"
	"context"
	"flag"
	"testing"

	"github.com/harikiriboy/go-ncvs-cli/writer"
	. "github.com/smartystreets/goconvey/convey"
)

func TestDescribeRulePackages(t *testing.T) {
	Convey("Tests DescribeRulePackages", t, func() {
		client := newMockClient()
		out := new(bytes.Buffer)

		cmd := DescribeRulePackages{Writer: writer.New(out), Client: client}
		ctx := context.Background()
		flg := flag.NewFlagSet("", flag.ContinueOnError)

		Convey("should success name", func() {
			So(cmd.Name(), ShouldEqual, "describe-rule-packages")
		})
		Convey("should success synopsis", func() {
			So(cmd.Synopsis(), ShouldEqual, "Describe RulePackages")
		})
		Convey("should success usage", func() {
			So(cmd.Usage(), ShouldEqual, "describe-rule-packages:\n	describe-rule-packages\n")
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
			errCmd := DescribeRulePackages{Writer: writer.New(out), Client: errClient}
			errCmd.Execute(ctx, flg)
			So(out.String(), ShouldEqual, "API request Error this is dummy error.")
		})
	})
}
