package config

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSpec(t *testing.T) {
	Convey("Given an environment with no environment variables set", t, func() {
		cfg, err := Get()

		Convey("When the config values are retrieved", func() {

			Convey("Then there should be no error returned", func() {
				So(err, ShouldBeNil)
			})

			Convey("The values should be set to the expected defaults", func() {
				So(cfg.BindAddr, ShouldEqual, ":22300")
				So(cfg.GracefulShutdownTimeout, ShouldEqual, 5*time.Second)
				So(cfg.HealthCheckInterval, ShouldEqual, 30*time.Second)
				So(cfg.HealthCheckCriticalTimeout, ShouldEqual, 90*time.Second)
				So(cfg.MongoConfig.BindAddr, ShouldEqual, "localhost:27017")
				So(cfg.MongoConfig.Collection, ShouldEqual, "recipes")
				So(cfg.MongoConfig.Database, ShouldEqual, "recipes")
				So(cfg.MongoConfig.EnableMongoData, ShouldEqual, false)
				So(cfg.MongoConfig.EnableMongoImport, ShouldEqual, false)
				So(cfg.MongoConfig.EnableAuthImport, ShouldEqual, false)
			})
		})
	})
}
