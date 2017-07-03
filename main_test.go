package main

import (
	"os"
	"testing"

	"github.com/harikiriboy/go-ncvs-cli/ncvs"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMain(t *testing.T) {
	Convey("Tests Main", t, func() {
		Convey("should success getAPIEndpoint for env", func() {
			os.Setenv("NIFTY_CLOUD_VSS_URL", "")
			So(getAPIEndpoint(), ShouldEqual, ncvs.DefaultAPIEndPoint)
		})
		Convey("should success getAPIEndpoint for default", func() {
			os.Setenv("NIFTY_CLOUD_VSS_URL", "test endpoint")
			So(getAPIEndpoint(), ShouldEqual, "test endpoint")
		})
		Convey("should success isIgnoreSSLCertsErrors true", func() {
			os.Setenv("NIFTY_CLOUD_VSS_SKIP_VERIFY", "true")
			So(isIgnoreSSLCertsErrors(), ShouldEqual, true)
		})
		Convey("should success isIgnoreSSLCertsErrors false", func() {
			os.Setenv("NIFTY_CLOUD_VSS_SKIP_VERIFY", "")
			So(isIgnoreSSLCertsErrors(), ShouldEqual, false)
		})
	})
}
