package commands

import (
	"bytes"
	"context"
	"flag"
	"testing"

	"github.com/harikiriboy/go-ncvs-cli/writer"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCreateScanTemplate(t *testing.T) {
	Convey("Tests CreateScanTemplate", t, func() {
		client := newMockClient()
		out := new(bytes.Buffer)

		cmd := CreateScanTemplate{Writer: writer.New(out), Client: client}
		ctx := context.Background()
		flg := flag.NewFlagSet("", flag.ContinueOnError)

		Convey("should success name", func() {
			So(cmd.Name(), ShouldEqual, "create-scan-template")
		})
		Convey("should success synopsis", func() {
			So(cmd.Synopsis(), ShouldEqual, "Create ScanTemplates")
		})
		Convey("should success usage", func() {
			So(cmd.Usage(), ShouldEqual, "create-scan-template:\n	create-scan-template\n	  [-scan-template-name]\n	  [-ssh-port]\n	  [-description]\n	  [-use-custom-rule-packages]\n	  [--rule-package-names]\n	  [--scan-targets <{'Region':'region','InstanceUniqueId': 'instanceUniqueId'}>]\n")
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
			errCmd := CreateScanTemplate{Writer: writer.New(out), Client: errClient}
			errCmd.Execute(ctx, flg)
			So(out.String(), ShouldEqual, "API request Error this is dummy error.")
		})
	})
}
