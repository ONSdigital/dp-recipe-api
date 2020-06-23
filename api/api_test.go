package api

import (
	"context"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/ONSdigital/dp-authorisation/auth"
	"github.com/ONSdigital/dp-recipe-api/config"
	"github.com/ONSdigital/dp-recipe-api/store"
	storetest "github.com/ONSdigital/dp-recipe-api/store/datastoretest"
	"github.com/gorilla/mux"
	. "github.com/smartystreets/goconvey/convey"
)

var (
	mu sync.Mutex
)

func getAuthorisationHandlerMock() *AuthHandlerMock {
	return &AuthHandlerMock{
		RequireFunc: func(required auth.Permissions, handler http.HandlerFunc) http.HandlerFunc {
			return func(w http.ResponseWriter, r *http.Request) {
				handler.ServeHTTP(w, r)
			}
		},
	}
}

// GetAPIWithMocks also used in other tests, so exported
func GetAPIWithMocks(mockedDataStore store.Storer) *RecipeAPI {
	mu.Lock()
	defer mu.Unlock()
	ctx := context.Background()
	cfg, err := config.Get()
	So(err, ShouldBeNil)
	recipePermissionsMock := getAuthorisationHandlerMock()
	permissionsMock := getAuthorisationHandlerMock()

	return NewRecipeAPI(ctx, *cfg, mux.NewRouter(), store.DataStore{Backend: mockedDataStore}, recipePermissionsMock, permissionsMock)
}

func hasRoute(r *mux.Router, path, method string) bool {
	req := httptest.NewRequest(method, path, nil)
	match := &mux.RouteMatch{}
	return r.Match(req, match)
}

func TestNewRecipeAPI(t *testing.T) {
	Convey("Given a public API instance", t, func() {
		mockedDataStore := &storetest.StorerMock{}
		api := GetAPIWithMocks(mockedDataStore)

		Convey("When created the following routes should have been added", func() {
			So(hasRoute(api.Router, "/health", "GET"), ShouldBeTrue)
			So(hasRoute(api.Router, "/recipes", "GET"), ShouldBeTrue)
			So(hasRoute(api.Router, "/recipes/{id}", "GET"), ShouldBeTrue)
		})
	})

	Convey("Given a private API instance", t, func() {
		mockedDataStore := &storetest.StorerMock{}
		api := GetAPIWithMocks(mockedDataStore)

		Convey("When created the following routes should have been added", func() {
			So(hasRoute(api.Router, "/health", "GET"), ShouldBeTrue)
			So(hasRoute(api.Router, "/recipes", "GET"), ShouldBeTrue)
			So(hasRoute(api.Router, "/recipes/{id}", "GET"), ShouldBeTrue)

			So(hasRoute(api.Router, "/recipes", "POST"), ShouldBeTrue)
			So(hasRoute(api.Router, "/recipes/{id}/instances", "POST"), ShouldBeTrue)
			So(hasRoute(api.Router, "/recipes/{id}/instances/{instance_id}/codelists", "POST"), ShouldBeTrue)

			So(hasRoute(api.Router, "/recipes/{id}", "PUT"), ShouldBeTrue)
			So(hasRoute(api.Router, "/recipes/{id}/instances/{instance_id}", "PUT"), ShouldBeTrue)
			So(hasRoute(api.Router, "/recipes/{id}/instances/{instance_id}/codelists/{codelist_id}", "PUT"), ShouldBeTrue)
		})
	})
}
