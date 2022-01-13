package config

import (
	"os"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBasicSpec(t *testing.T) {
	Convey("Given an environment with no environment variables set", t, func() {
		os.Clearenv()
		cfg = nil

		Convey("When the config values are retrieved", func() {
			cfg, err := Get()

			Convey("Then there should be no error returned", func() {
				So(err, ShouldBeNil)
			})

			Convey("The values should be set to the expected defaults", func() {
				So(cfg.BindAddr, ShouldEqual, ":22300")
				So(cfg.ZebedeeURL, ShouldEqual, "http://localhost:8082")
				So(cfg.GracefulShutdownTimeout, ShouldEqual, 5*time.Second)
				So(cfg.HealthCheckInterval, ShouldEqual, 30*time.Second)
				So(cfg.HealthCheckCriticalTimeout, ShouldEqual, 90*time.Second)
				So(cfg.DefaultLimit, ShouldEqual, 20)
				So(cfg.DefaultMaxLimit, ShouldEqual, 1000)
				So(cfg.DefaultOffset, ShouldEqual, 0)
				So(cfg.MongoConfig.ClusterEndpoint, ShouldEqual, "localhost:27017")
				So(cfg.MongoConfig.Username, ShouldEqual, "")
				So(cfg.MongoConfig.Password, ShouldEqual, "")
				So(cfg.MongoConfig.Database, ShouldEqual, "recipes")
				So(cfg.MongoConfig.Collections, ShouldResemble, map[string]string{RecipesCollection: "recipes"})
				So(cfg.MongoConfig.IsSSL, ShouldEqual, false)
			})
		})
	})
}

func TestEnvSpec(t *testing.T) {
	Convey("Given an environment with some config environment variables set", t, func() {
		os.Setenv("BIND_ADDR", "https://an-address:a-port")
		os.Setenv("MONGODB_DATABASE", "mongo-database")
		os.Setenv("MONGODB_COLLECTIONS", "MainCollection:bingo,SecondaryCollection:bongo")
		cfg = nil

		Convey("When the config values are retrieved", func() {
			cfg, err := Get()

			Convey("Then there should be no error returned", func() {
				So(err, ShouldBeNil)
			})

			Convey("The config variables that should receive their values from the environment are correctly set, and the remaining values should be set to the expected defaults", func() {
				So(cfg.BindAddr, ShouldEqual, "https://an-address:a-port")
				So(cfg.ZebedeeURL, ShouldEqual, "http://localhost:8082")
				So(cfg.GracefulShutdownTimeout, ShouldEqual, 5*time.Second)
				So(cfg.HealthCheckInterval, ShouldEqual, 30*time.Second)
				So(cfg.HealthCheckCriticalTimeout, ShouldEqual, 90*time.Second)
				So(cfg.DefaultLimit, ShouldEqual, 20)
				So(cfg.DefaultMaxLimit, ShouldEqual, 1000)
				So(cfg.DefaultOffset, ShouldEqual, 0)
				So(cfg.MongoConfig.ClusterEndpoint, ShouldEqual, "localhost:27017")
				So(cfg.MongoConfig.Username, ShouldEqual, "")
				So(cfg.MongoConfig.Password, ShouldEqual, "")
				So(cfg.MongoConfig.Database, ShouldEqual, "mongo-database")
				So(cfg.MongoConfig.Collections, ShouldResemble, map[string]string{"MainCollection": "bingo", "SecondaryCollection": "bongo"})
				So(cfg.MongoConfig.IsSSL, ShouldEqual, false)
			})
		})
	})
}
