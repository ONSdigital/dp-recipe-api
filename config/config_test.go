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
				So(cfg.ZebedeeURL, ShouldEqual, "http://localhost:8082")
				So(cfg.GracefulShutdownTimeout, ShouldEqual, 5*time.Second)
				So(cfg.HealthCheckInterval, ShouldEqual, 30*time.Second)
				So(cfg.HealthCheckCriticalTimeout, ShouldEqual, 90*time.Second)
				So(cfg.MongoConfig.ClusterEndpoint, ShouldEqual, "localhost:27017")
				So(cfg.MongoConfig.Collection, ShouldEqual, "recipes")
				So(cfg.MongoConfig.Database, ShouldEqual, "recipes")
				So(cfg.DefaultLimit, ShouldEqual, 20)
				So(cfg.DefaultMaxLimit, ShouldEqual, 1000)
				So(cfg.DefaultOffset, ShouldEqual, 0)
				So(cfg.MongoConfig.Username, ShouldEqual, "")
				So(cfg.MongoConfig.Password, ShouldEqual, "")
				So(cfg.MongoConfig.IsSSL, ShouldEqual, false)
			})
		})
	})
}
