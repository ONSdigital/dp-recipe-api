package api

import (
	"bytes"
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
	mu         sync.Mutex
	recipeTest = `{
					"id" : "123", 
					"alias" : "test", 
					"format" : "v4", 
					"files" : [
						{
							"description" : "test files"
						}
					], 
					"output_instances" : [
						{
							"dataset_id" : "1234",
							"editions" : [ 
                				"edition-test"
            				],
							"title" : "test",
							"code_lists" : [ 
                				{
                    				"id" : "12345",
                    				"href" : "http://localhost:22400/code-lists/12345",
                    				"name" : "codelist-test-name",
                    				"is_hierarchy" : false
								}
							]
						}
					]
				}`
)

// GetAPIWithMocks also used in other tests, so exported
func GetAPIWithMocks(mockedDataStore store.Storer) *RecipeAPI {
	mu.Lock()
	defer mu.Unlock()
	ctx := context.Background()
	cfg, err := config.Get()
	So(err, ShouldBeNil)
	cfg.MongoConfig.EnableMongoData = true
	cfg.MongoConfig.EnableMongoImport = true
	cfg.MongoConfig.EnableAuthImport = true

	return NewRecipeAPI(ctx, *cfg, mux.NewRouter(), store.DataStore{Backend: mockedDataStore})
}

func TestGetRecipesReturnsOK(t *testing.T) {
	t.Parallel()
	Convey("A successful request to get recipe returns 200 OK response", t, func() {
		r := httptest.NewRequest("GET", "http://localhost:22300/recipes", nil)
		w := httptest.NewRecorder()
		mockedDataStore := &storetest.StorerMock{
			GetRecipesFunc: func(ctx context.Context) ([]recipe.Response, error) {
				return []recipe.Response{}, nil
			},
		}

		api := GetAPIWithMocks(mockedDataStore)
		api.Router.ServeHTTP(w, r)

		So(w.Code, ShouldEqual, http.StatusOK)
		So(len(mockedDataStore.GetRecipesCalls()), ShouldEqual, 1)

	})

	Convey("When the api cannot find any recipes return 200 OK response", t, func() {
		r := httptest.NewRequest("GET", "http://localhost:22300/recipes", nil)
		w := httptest.NewRecorder()
		mockedDataStore := &storetest.StorerMock{
			GetRecipesFunc: func(ctx context.Context) ([]recipe.Response, error) {
				return nil, errs.ErrRecipesNotFound
			},
		}

		api := GetAPIWithMocks(mockedDataStore)
		api.Router.ServeHTTP(w, r)

		So(w.Code, ShouldEqual, http.StatusOK)
		So(len(mockedDataStore.GetRecipesCalls()), ShouldEqual, 1)
	})
}

func TestGetRecipesReturnsError(t *testing.T) {
	t.Parallel()
	Convey("When the api cannot connect to datastore return an internal server error", t, func() {
		r := httptest.NewRequest("GET", "http://localhost:22300/recipes", nil)
		w := httptest.NewRecorder()
		mockedDataStore := &storetest.StorerMock{
			GetRecipesFunc: func(ctx context.Context) ([]recipe.Response, error) {
				return nil, errs.ErrInternalServer
			},
		}

		api := GetAPIWithMocks(mockedDataStore)
		api.Router.ServeHTTP(w, r)

		So(w.Code, ShouldEqual, http.StatusInternalServerError)
		So(len(mockedDataStore.GetRecipesCalls()), ShouldEqual, 1)
	})
}

func TestGetRecipeReturnsOK(t *testing.T) {
	t.Parallel()
	Convey("A successful request to get specific recipe by id returns 200 OK response", t, func() {
		r := httptest.NewRequest("GET", "http://localhost:22300/recipes/123", nil)
		w := httptest.NewRecorder()
		mockedDataStore := &storetest.StorerMock{
			GetRecipeFunc: func(ID string) (*recipe.Response, error) {
				return &recipe.Response{ID: "123"}, nil
			},
		}

		api := GetAPIWithMocks(mockedDataStore)
		api.Router.ServeHTTP(w, r)

		So(w.Code, ShouldEqual, http.StatusOK)
		So(len(mockedDataStore.GetRecipeCalls()), ShouldEqual, 1)

	})
}

func TestGetRecipeReturnsError(t *testing.T) {
	t.Parallel()
	Convey("When the api cannot connect to datastore return an internal server error", t, func() {
		r := httptest.NewRequest("GET", "http://localhost:22300/recipes/123", nil)
		w := httptest.NewRecorder()
		mockedDataStore := &storetest.StorerMock{
			GetRecipeFunc: func(ID string) (*recipe.Response, error) {
				return nil, errs.ErrInternalServer
			},
		}

		api := GetAPIWithMocks(mockedDataStore)
		api.Router.ServeHTTP(w, r)

		So(w.Code, ShouldEqual, http.StatusInternalServerError)
		So(len(mockedDataStore.GetRecipeCalls()), ShouldEqual, 1)
	})

	Convey("When the api cannot find the recipe return status not found, 404", t, func() {
		r := httptest.NewRequest("GET", "http://localhost:22300/recipes/123", nil)
		w := httptest.NewRecorder()
		mockedDataStore := &storetest.StorerMock{
			GetRecipeFunc: func(ID string) (*recipe.Response, error) {
				return nil, errs.ErrRecipeNotFound
			},
		}

		api := GetAPIWithMocks(mockedDataStore)
		api.Router.ServeHTTP(w, r)

		So(w.Code, ShouldEqual, http.StatusNotFound)
		So(len(mockedDataStore.GetRecipeCalls()), ShouldEqual, 1)
	})
}

func TestAddAllRecipesReturnsOK(t *testing.T) {
	t.Parallel()
	Convey("A successful request to add all recipes to mongo returns 200 OK response", t, func() {
		r := httptest.NewRequest("POST", "http://localhost:22300/allrecipes", nil)
		w := httptest.NewRecorder()
		mockedDataStore := &storetest.StorerMock{
			AddRecipeFunc: func(item recipe.Response) error {
				return nil
			},
		}

		api := GetAPIWithMocks(mockedDataStore)
		api.Router.ServeHTTP(w, r)

		So(w.Code, ShouldEqual, http.StatusOK)
		So(len(mockedDataStore.AddRecipeCalls()), ShouldEqual, len(recipe.FullList.Items))

	})
}

func TestAddAllRecipesReturnsError(t *testing.T) {
	t.Parallel()
	Convey("When the api cannot add all recipes to mongo return an internal server error", t, func() {
		r := httptest.NewRequest("POST", "http://localhost:22300/allrecipes", nil)
		w := httptest.NewRecorder()
		mockedDataStore := &storetest.StorerMock{
			AddRecipeFunc: func(item recipe.Response) error {
				return errs.ErrInternalServer
			},
		}

		api := GetAPIWithMocks(mockedDataStore)
		api.Router.ServeHTTP(w, r)

		So(w.Code, ShouldEqual, http.StatusInternalServerError)
		So(len(mockedDataStore.AddRecipeCalls()), ShouldEqual, 1)
	})
}

func TestAddRecipeReturnsOK(t *testing.T) {
	t.Parallel()
	Convey("A successful request to add recipe to mongo returns 200 OK response", t, func() {
		r := httptest.NewRequest("POST", "http://localhost:22300/recipes", bytes.NewBufferString(recipeTest))
		w := httptest.NewRecorder()
		mockedDataStore := &storetest.StorerMock{
			AddRecipeFunc: func(item recipe.Response) error {
				return nil
			},
		}
		api := GetAPIWithMocks(mockedDataStore)
		api.Router.ServeHTTP(w, r)

		So(w.Code, ShouldEqual, http.StatusOK)
		So(len(mockedDataStore.AddRecipeCalls()), ShouldEqual, 1)

	})
}

func TestAddRecipeReturnsBadRequest(t *testing.T) {
	t.Parallel()
	Convey("A request to add an incomplete recipe to mongo returns 400 Bad Request response", t, func() {
		r := httptest.NewRequest("POST", "http://localhost:22300/recipes", bytes.NewBufferString(`{"alias":"test"}`))
		w := httptest.NewRecorder()
		mockedDataStore := &storetest.StorerMock{
			AddRecipeFunc: func(item recipe.Response) error {
				return nil
			},
		}

		api := GetAPIWithMocks(mockedDataStore)
		api.Router.ServeHTTP(w, r)

		So(w.Code, ShouldEqual, http.StatusBadRequest)
		So(len(mockedDataStore.AddRecipeCalls()), ShouldEqual, 0)

	})
}

func TestAddRecipeReturnsError(t *testing.T) {
	t.Parallel()
	Convey("When the api cannot add recipe to mongo return an internal server error", t, func() {
		r := httptest.NewRequest("POST", "http://localhost:22300/recipes", bytes.NewBufferString(`{`))
		w := httptest.NewRecorder()
		mockedDataStore := &storetest.StorerMock{
			AddRecipeFunc: func(item recipe.Response) error {
				return errs.ErrAddUpdateRecipeBadRequest
			},
		}

		api := GetAPIWithMocks(mockedDataStore)
		api.Router.ServeHTTP(w, r)

		So(w.Code, ShouldEqual, http.StatusInternalServerError)
		So(len(mockedDataStore.AddRecipeCalls()), ShouldEqual, 0)
	})
}

func TestUpdateRecipeReturnsOK(t *testing.T) {
	t.Parallel()
	Convey("A successful request to update recipe in mongo returns 200 OK response", t, func() {
		r := httptest.NewRequest("PUT", "http://localhost:22300/recipes/123", bytes.NewBufferString(`{"output_instances":[{"title":"Test"}]}`))
		w := httptest.NewRecorder()
		mockedDataStore := &storetest.StorerMock{
			GetRecipeFunc: func(ID string) (*recipe.Response, error) {
				return &recipe.Response{ID: "123", Alias: "Original", Format: "v3"}, nil
			},
			UpdateRecipeFunc: func(ID string, recipeUpdate interface{}, instanceIndex int, codelistIndex int) error {
				return nil
			},
		}
		api := GetAPIWithMocks(mockedDataStore)
		api.Router.ServeHTTP(w, r)

		So(w.Code, ShouldEqual, http.StatusOK)
		So(len(mockedDataStore.UpdateRecipeCalls()), ShouldEqual, 1)
		So(len(mockedDataStore.GetRecipeCalls()), ShouldEqual, 1)
	})
}

func TestUpdateRecipeReturnsBadRequest(t *testing.T) {
	t.Parallel()
	Convey("A request to process an incomplete update recipe to mongo returns 400 Bad Request response", t, func() {
		r := httptest.NewRequest("PUT", "http://localhost:22300/recipes/123", bytes.NewBufferString(`{"format":"v5"}`))
		w := httptest.NewRecorder()
		mockedDataStore := &storetest.StorerMock{
			GetRecipeFunc: func(ID string) (*recipe.Response, error) {
				return &recipe.Response{ID: "123", Alias: "Original", Format: "v3"}, nil
			},
			UpdateRecipeFunc: func(ID string, recipeUpdate interface{}, instanceIndex int, codelistIndex int) error {
				return errs.ErrAddUpdateRecipeBadRequest
			},
		}

		api := GetAPIWithMocks(mockedDataStore)
		api.Router.ServeHTTP(w, r)

		So(w.Code, ShouldEqual, http.StatusBadRequest)
		So(len(mockedDataStore.UpdateRecipeCalls()), ShouldEqual, 0)
		So(len(mockedDataStore.GetRecipeCalls()), ShouldEqual, 0)
	})
}

func TestUpdateRecipeReturnsError(t *testing.T) {
	t.Parallel()
	Convey("When the api cannot update recipe in mongo return an internal server error", t, func() {
		r := httptest.NewRequest("PUT", "http://localhost:22300/recipes/123", bytes.NewBufferString(`{`))
		w := httptest.NewRecorder()
		mockedDataStore := &storetest.StorerMock{
			UpdateRecipeFunc: func(ID string, recipeUpdate interface{}, instanceIndex int, codelistIndex int) error {
				return errs.ErrAddUpdateRecipeBadRequest
			},
			GetRecipeFunc: func(ID string) (*recipe.Response, error) {
				return nil, errs.ErrInternalServer
			},
		}

		api := GetAPIWithMocks(mockedDataStore)
		api.Router.ServeHTTP(w, r)

		So(w.Code, ShouldEqual, http.StatusInternalServerError)
		So(len(mockedDataStore.UpdateRecipeCalls()), ShouldEqual, 0)
		So(len(mockedDataStore.GetRecipeCalls()), ShouldEqual, 0)
	})
}
