package api

import (
	"context"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	errs "github.com/ONSdigital/dp-recipe-api/apierrors"
	"github.com/ONSdigital/dp-recipe-api/config"
	"github.com/ONSdigital/dp-recipe-api/recipe"
	"github.com/ONSdigital/dp-recipe-api/store"
	storetest "github.com/ONSdigital/dp-recipe-api/store/datastoretest"
	"github.com/gorilla/mux"
	. "github.com/smartystreets/goconvey/convey"
)

var (
	mu sync.Mutex
)

// GetAPIWithMocks also used in other tests, so exported
func GetAPIWithMocks(mockedDataStore store.Storer) *RecipeAPI {
	mu.Lock()
	defer mu.Unlock()
	ctx := context.Background()
	cfg, err := config.Get()
	So(err, ShouldBeNil)
	cfg.MongoConfig.EnableMongoData = true

	return NewRecipeAPI(ctx, *cfg, mux.NewRouter(), store.DataStore{Backend: mockedDataStore})
}

func TestGetRecipesReturnsOK(t *testing.T) {
	t.Parallel()
	Convey("A successful request to get recipe returns 200 OK response", t, func() {
		r := httptest.NewRequest("GET", "http://localhost:22300/recipes", nil)
		w := httptest.NewRecorder()
		mockedDataStore := &storetest.StorerMock{
			GetRecipesFunc: func() ([]recipe.Response, error) {
				return []recipe.Response{}, nil
			},
		}

		api := GetAPIWithMocks(mockedDataStore)
		api.Router.ServeHTTP(w, r)

		So(w.Code, ShouldEqual, http.StatusOK)
		So(len(mockedDataStore.GetRecipesCalls()), ShouldEqual, 1)

	})
}

func TestGetDatasetsReturnsError(t *testing.T) {
	t.Parallel()
	Convey("When the api cannot connect to datastore return an internal server error", t, func() {
		r := httptest.NewRequest("GET", "http://localhost:22300/recipes", nil)
		w := httptest.NewRecorder()
		mockedDataStore := &storetest.StorerMock{
			GetRecipesFunc: func() ([]recipe.Response, error) {
				return nil, errs.ErrInternalServer
			},
		}

		api := GetAPIWithMocks(mockedDataStore)
		api.Router.ServeHTTP(w, r)

		So(w.Code, ShouldEqual, http.StatusInternalServerError)
		So(len(mockedDataStore.GetRecipesCalls()), ShouldEqual, 1)
	})
}
