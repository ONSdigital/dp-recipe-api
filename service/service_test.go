package service_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"testing"

	"github.com/ONSdigital/dp-healthcheck/healthcheck"
	"github.com/ONSdigital/dp-recipe-api/config"
	"github.com/ONSdigital/dp-recipe-api/service"
	"github.com/ONSdigital/dp-recipe-api/service/mock"
	serviceMock "github.com/ONSdigital/dp-recipe-api/service/mock"
	"github.com/ONSdigital/dp-recipe-api/store"
	storeMock "github.com/ONSdigital/dp-recipe-api/store/datastoretest"
	. "github.com/smartystreets/goconvey/convey"
)

var (
	ctx           = context.Background()
	testBuildTime = "BuildTime"
	testGitCommit = "GitCommit"
	testVersion   = "Version"
)

var (
	errMongo       = errors.New("MongoDB error")
	errServer      = errors.New("HTTP Server error")
	errHealthcheck = errors.New("healthCheck error")
)

var funcDoGetHealthcheckErr = func(cfg *config.Configuration, buildTime string, gitCommit string, version string) (service.HealthChecker, error) {
	return nil, errHealthcheck
}

var funcDoGetMongoDBErr = func(ctx context.Context, cfg config.MongoConfig) (store.MongoDB, error) {
	return nil, errMongo
}

func TestRun(t *testing.T) {

	Convey("Having a set of mocked dependencies", t, func() {

		cfg, err := config.Get()
		So(err, ShouldBeNil)

		hcMock := &mock.HealthCheckerMock{
			AddCheckFunc: func(name string, checker healthcheck.Checker) error { return nil },
			StartFunc:    func(ctx context.Context) {},
		}

		serverWg := &sync.WaitGroup{}
		serverMock := &mock.HTTPServerMock{
			ListenAndServeFunc: func() error {
				serverWg.Done()
				return nil
			},
		}

		failingServerMock := &serviceMock.HTTPServerMock{
			ListenAndServeFunc: func() error {
				serverWg.Done()
				return errServer
			},
		}

		funcDoGetHealthcheckOk := func(cfg *config.Configuration, buildTime string, gitCommit string, version string) (service.HealthChecker, error) {
			return hcMock, nil
		}

		funcDoGetHTTPServer := func(bindAddr string, router http.Handler) service.HTTPServer {
			return serverMock
		}

		funcDoGetFailingHTTPSerer := func(bindAddr string, router http.Handler) service.HTTPServer {
			return failingServerMock
		}

		funcDoGetMongoDBOk := func(ctx context.Context, cfg config.MongoConfig) (store.MongoDB, error) {
			return &storeMock.MongoDBMock{}, nil
		}

		Convey("Given that initialising MongoDB returns an error", func() {
			initMock := &serviceMock.InitialiserMock{
				DoGetMongoDBFunc: funcDoGetMongoDBErr,
			}
			svcErrors := make(chan error, 1)
			svcList := service.NewServiceList(initMock)
			svc := service.New(cfg, svcList)
			err := svc.Run(ctx, testBuildTime, testGitCommit, testVersion, svcErrors)

			Convey("Then service Run fails with the same error and the flag is not set. No further initialisations are attempted", func() {
				So(err, ShouldResemble, errMongo)
				So(svcList.MongoDB, ShouldBeFalse)
				So(svcList.HealthCheck, ShouldBeFalse)
			})
		})

		Convey("Given that initialising Helthcheck returns an error", func() {
			initMock := &serviceMock.InitialiserMock{
				DoGetMongoDBFunc:     funcDoGetMongoDBOk,
				DoGetHealthCheckFunc: funcDoGetHealthcheckErr,
			}
			svcErrors := make(chan error, 1)
			svcList := service.NewServiceList(initMock)
			svc := service.New(cfg, svcList)
			err := svc.Run(ctx, testBuildTime, testGitCommit, testVersion, svcErrors)

			Convey("Then service Run fails with the same error and the flag is not set. No further initialisations are attempted", func() {
				So(err, ShouldResemble, errHealthcheck)
				So(svcList.MongoDB, ShouldBeTrue)
				So(svcList.HealthCheck, ShouldBeFalse)
			})
		})

		Convey("Given that Checkers cannot be registered", func() {

			errAddheckFail := errors.New("Error(s) registering checkers for healthcheck")
			hcMockAddFail := &serviceMock.HealthCheckerMock{
				AddCheckFunc: func(name string, checker healthcheck.Checker) error { return errAddheckFail },
				StartFunc:    func(ctx context.Context) {},
			}

			initMock := &serviceMock.InitialiserMock{
				DoGetMongoDBFunc: funcDoGetMongoDBOk,
				DoGetHealthCheckFunc: func(cfg *config.Configuration, buildTime string, gitCommit string, version string) (service.HealthChecker, error) {
					return hcMockAddFail, nil
				},
			}
			svcErrors := make(chan error, 1)
			svcList := service.NewServiceList(initMock)
			svc := service.New(cfg, svcList)
			err := svc.Run(ctx, testBuildTime, testGitCommit, testVersion, svcErrors)

			Convey("Then service Run fails, but all checks try to register", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldResemble, fmt.Sprintf("unable to register checkers: %s", errAddheckFail.Error()))
				So(svcList.MongoDB, ShouldBeTrue)
				So(svcList.HealthCheck, ShouldBeTrue)
				So(len(hcMockAddFail.AddCheckCalls()), ShouldEqual, 2)
				So(hcMockAddFail.AddCheckCalls()[0].Name, ShouldResemble, "Zebedee")
				So(hcMockAddFail.AddCheckCalls()[1].Name, ShouldResemble, "Mongo DB")
			})
		})

		Convey("Given that all dependencies are successfully initialised", func() {
			initMock := &serviceMock.InitialiserMock{
				DoGetMongoDBFunc:     funcDoGetMongoDBOk,
				DoGetHealthCheckFunc: funcDoGetHealthcheckOk,
				DoGetHTTPServerFunc:  funcDoGetHTTPServer,
			}
			svcErrors := make(chan error, 1)
			svcList := service.NewServiceList(initMock)
			svc := service.New(cfg, svcList)
			serverWg.Add(1)
			err := svc.Run(ctx, testBuildTime, testGitCommit, testVersion, svcErrors)

			Convey("Then service Run succeeds and all the flags are set", func() {
				So(err, ShouldBeNil)
				So(svcList.MongoDB, ShouldBeTrue)
				So(svcList.HealthCheck, ShouldBeTrue)
			})

			Convey("The checkers are registered and the healthcheck and http server started", func() {
				So(len(hcMock.AddCheckCalls()), ShouldEqual, 2)
				So(hcMock.AddCheckCalls()[0].Name, ShouldResemble, "Zebedee")
				So(hcMock.AddCheckCalls()[1].Name, ShouldResemble, "Mongo DB")
				So(len(initMock.DoGetHTTPServerCalls()), ShouldEqual, 1)
				So(initMock.DoGetHTTPServerCalls()[0].BindAddr, ShouldEqual, ":22300")
				So(len(hcMock.StartCalls()), ShouldEqual, 1)
				serverWg.Wait() // Wait for HTTP server go-routine to finish
				So(len(serverMock.ListenAndServeCalls()), ShouldEqual, 1)
			})
		})

		Convey("Given that all dependencies are successfully initialised but the http server fails", func() {

			initMock := &serviceMock.InitialiserMock{
				DoGetMongoDBFunc:     funcDoGetMongoDBOk,
				DoGetHealthCheckFunc: funcDoGetHealthcheckOk,
				DoGetHTTPServerFunc:  funcDoGetFailingHTTPSerer,
			}
			svcErrors := make(chan error, 1)
			svcList := service.NewServiceList(initMock)
			svc := service.New(cfg, svcList)
			serverWg.Add(1)
			err := svc.Run(ctx, testBuildTime, testGitCommit, testVersion, svcErrors)
			So(err, ShouldBeNil)

			Convey("Then the error is returned in the error channel", func() {
				sErr := <-svcErrors
				So(sErr.Error(), ShouldResemble, fmt.Sprintf("failure in http listen and serve: %s", errServer.Error()))
				So(len(failingServerMock.ListenAndServeCalls()), ShouldEqual, 1)
			})
		})
	})
}

func TestClose(t *testing.T) {

	Convey("Having a correctly initialised service", t, func(c C) {

		cfg, err := config.Get()
		So(err, ShouldBeNil)

		hcStopped := false
		serverStopped := false

		// healthcheck Stop does not depend on any other service being closed/stopped
		hcMock := &serviceMock.HealthCheckerMock{
			AddCheckFunc: func(name string, checker healthcheck.Checker) error { return nil },
			StartFunc:    func(ctx context.Context) {},
			StopFunc:     func() { hcStopped = true },
		}

		// server Shutdown will fail if healthcheck is not stopped
		serverMock := &serviceMock.HTTPServerMock{
			ListenAndServeFunc: func() error { return nil },
			ShutdownFunc: func(ctx context.Context) error {
				if !hcStopped {
					return errors.New("Server was stopped before healthcheck")
				}
				serverStopped = true
				return nil
			},
		}

		funcClose := func(ctx context.Context) error {
			if !hcStopped {
				return errors.New("Dependency was closed before healthcheck")
			}
			if !serverStopped {
				return errors.New("Dependency was closed before http server")
			}
			return nil
		}

		// mongoDB will fail if healthcheck or http server are not stopped
		mongoMock := &storeMock.MongoDBMock{
			CloseFunc: funcClose,
		}

		Convey("Closing a service does not close uninitialised dependencies", func() {
			svcList := service.NewServiceList(nil)
			svcList.HealthCheck = true
			svc := service.New(cfg, svcList)
			svc.SetServer(serverMock)
			svc.SetHealthCheck(hcMock)
			err = svc.Close(context.Background())
			So(err, ShouldBeNil)
			So(len(hcMock.StopCalls()), ShouldEqual, 1)
			So(len(serverMock.ShutdownCalls()), ShouldEqual, 1)
		})

		fullSvcList := &service.ExternalServiceList{
			HealthCheck: true,
			MongoDB:     true,
			Init:        nil,
		}

		Convey("Closing the service results in all the initialised dependencies being closed in the expected order", func() {
			svc := service.New(cfg, fullSvcList)
			svc.SetServer(serverMock)
			svc.SetHealthCheck(hcMock)
			svc.SetMongoDB(mongoMock)
			err = svc.Close(context.Background())
			So(err, ShouldBeNil)
			So(len(hcMock.StopCalls()), ShouldEqual, 1)
			So(len(serverMock.ShutdownCalls()), ShouldEqual, 1)
			So(len(mongoMock.CloseCalls()), ShouldEqual, 1)
		})

		Convey("If services fail to stop, the Close operation tries to close all dependencies and returns an error", func() {
			failingserverMock := &serviceMock.HTTPServerMock{
				ListenAndServeFunc: func() error { return nil },
				ShutdownFunc: func(ctx context.Context) error {
					return errors.New("Failed to stop http server")
				},
			}

			svc := service.New(cfg, fullSvcList)
			svc.SetServer(failingserverMock)
			svc.SetHealthCheck(hcMock)
			svc.SetMongoDB(mongoMock)
			err = svc.Close(context.Background())
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldResemble, "failed to shutdown gracefully")
			So(len(hcMock.StopCalls()), ShouldEqual, 1)
			So(len(failingserverMock.ShutdownCalls()), ShouldEqual, 1)
			So(len(mongoMock.CloseCalls()), ShouldEqual, 1)
		})
	})
}
