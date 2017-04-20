package commands

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCustomFlags(t *testing.T) {
	Convey("Tests Custom Flags", t, func() {
		var rulePackageNames rulePackageNames
		var scanTargets scanTargets
		Convey("should success rulePackageNames flag", func() {
			rulePackageNames.Set("test")
			So(rulePackageNames[0], ShouldEqual, "test")
		})
		Convey("should success scanTargets flag", func() {
			scanTargets.Set(`{"Region": "test_region", "InstanceUniqueId": "test_unique_id"}`)
			So(scanTargets[0].Region, ShouldEqual, "test_region")
			So(scanTargets[0].InstanceUniqueID, ShouldEqual, "test_unique_id")
		})
	})
}
