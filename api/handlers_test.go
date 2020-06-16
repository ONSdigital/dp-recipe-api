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
	mu sync.Mutex

	// Variables to point to bool for isHierarchy in CodeLists
	falseVal    = false
	falseValPtr = &falseVal
	trueVal     = true
	trueValPtr  = &trueVal

	// Test data for request body
	recipeTest   = `{"id":"123","alias":"test","format":"v4","files":[{"description":"test files"}],"output_instances":[` + instanceTest + `]}`
	instanceTest = `{"dataset_id":"1234","editions":["edition-test"],"title":"test","code_lists" :[` + codelistTest + `]}`
	codelistTest = `{"id":"12345", "href":"http://localhost:22400/code-lists/12345", "name":"codelist-test-name", "is_hierarchy":false}`

	// Test data of recipe retrieved from GetRecipe()
	initialCodelist = []recipe.CodeList{
		{
			ID:          "789",
			Name:        "codelist-test",
			HRef:        "http://localhost:22400/code-lists/789",
			IsHierarchy: falseValPtr,
		},
	}
	initialInstance = []recipe.Instance{
		{
			DatasetID: "456",
			Editions:  []string{"editions"},
			Title:     "test",
			CodeLists: initialCodelist,
		},
	}
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
		r := httptest.NewRequest("PUT", "http://localhost:22300/recipes/123", bytes.NewBufferString(`{"alias":"Test"}`))
		w := httptest.NewRecorder()
		mockedDataStore := &storetest.StorerMock{
			UpdateRecipeFunc: func(ID string, updates recipe.Response) error {
				return nil
			},
		}
		api := GetAPIWithMocks(mockedDataStore)
		api.Router.ServeHTTP(w, r)

		So(w.Code, ShouldEqual, http.StatusOK)
		So(len(mockedDataStore.UpdateRecipeCalls()), ShouldEqual, 1)
	})
}

func TestUpdateRecipeReturnsBadRequest(t *testing.T) {
	t.Parallel()
	Convey("A request to process an incomplete update recipe to mongo returns 400 Bad Request response", t, func() {
		r := httptest.NewRequest("PUT", "http://localhost:22300/recipes/123", bytes.NewBufferString(`{"format":"v5"}`))
		w := httptest.NewRecorder()
		mockedDataStore := &storetest.StorerMock{
			UpdateRecipeFunc: func(ID string, updates recipe.Response) error {
				return errs.ErrAddUpdateRecipeBadRequest
			},
		}

		api := GetAPIWithMocks(mockedDataStore)
		api.Router.ServeHTTP(w, r)

		So(w.Code, ShouldEqual, http.StatusBadRequest)
		So(len(mockedDataStore.UpdateRecipeCalls()), ShouldEqual, 0)
	})
}

func TestUpdateRecipeReturnsError(t *testing.T) {
	t.Parallel()
	Convey("When the api cannot update recipe in mongo return an internal server error", t, func() {
		r := httptest.NewRequest("PUT", "http://localhost:22300/recipes/123", bytes.NewBufferString(`{`))
		w := httptest.NewRecorder()
		mockedDataStore := &storetest.StorerMock{
			UpdateRecipeFunc: func(ID string, updates recipe.Response) error {
				return errs.ErrAddUpdateRecipeBadRequest
			},
		}

		api := GetAPIWithMocks(mockedDataStore)
		api.Router.ServeHTTP(w, r)

		So(w.Code, ShouldEqual, http.StatusInternalServerError)
		So(len(mockedDataStore.UpdateRecipeCalls()), ShouldEqual, 0)
	})
}

func TestAddInstanceReturnsOK(t *testing.T) {
	t.Parallel()
	Convey("A successful request to add instance in existing recipe to mongo returns 200 OK response", t, func() {
		r := httptest.NewRequest("POST", "http://localhost:22300/recipes/123/instances", bytes.NewBufferString(instanceTest))
		w := httptest.NewRecorder()
		mockedDataStore := &storetest.StorerMock{
			GetRecipeFunc: func(ID string) (*recipe.Response, error) {
				return &recipe.Response{ID: "123", Alias: "Original", Format: "v4"}, nil
			},
			UpdateRecipeFunc: func(recipeID string, currentRecipe recipe.Response) error {
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

func TestAddInstanceReturnsBadRequestError(t *testing.T) {
	t.Parallel()
	Convey("A request to process and add an incomplete instance to recipe to mongo returns 400 Bad Request response", t, func() {
		r := httptest.NewRequest("POST", "http://localhost:22300/recipes/123/instances", bytes.NewBufferString(`{"dataset_id":"1234"}`))
		w := httptest.NewRecorder()
		mockedDataStore := &storetest.StorerMock{
			GetRecipeFunc: func(ID string) (*recipe.Response, error) {
				return &recipe.Response{ID: "123", Alias: "Original", Format: "v4"}, nil
			},
			UpdateRecipeFunc: func(recipeID string, currentRecipe recipe.Response) error {
				return nil
			},
		}
		api := GetAPIWithMocks(mockedDataStore)
		api.Router.ServeHTTP(w, r)

		So(w.Code, ShouldEqual, http.StatusBadRequest)
		So(len(mockedDataStore.UpdateRecipeCalls()), ShouldEqual, 0)
		So(len(mockedDataStore.GetRecipeCalls()), ShouldEqual, 0)
	})
}

func TestAddInstanceReturnsError(t *testing.T) {
	t.Parallel()
	Convey("When the api cannot add instance in mongo return an internal server error", t, func() {
		r := httptest.NewRequest("POST", "http://localhost:22300/recipes/123/instances", bytes.NewBufferString(`{`))
		w := httptest.NewRecorder()
		mockedDataStore := &storetest.StorerMock{
			UpdateRecipeFunc: func(recipeID string, currentRecipe recipe.Response) error {
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

func TestUpdateInstanceReturnsOK(t *testing.T) {
	t.Parallel()
	Convey("A successful request to update instance of a recipe in mongo returns 200 OK response", t, func() {
		r := httptest.NewRequest("PUT", "http://localhost:22300/recipes/123/instances/456", bytes.NewBufferString(`{"title":"Test"}`))
		w := httptest.NewRecorder()
		mockedDataStore := &storetest.StorerMock{
			GetRecipeFunc: func(ID string) (*recipe.Response, error) {
				return &recipe.Response{ID: "123", Alias: "Original", Format: "v4", OutputInstances: initialInstance}, nil
			},
			UpdateInstanceFunc: func(recipeID string, instanceIndex int, updates recipe.Instance) error {
				return nil
			},
		}
		api := GetAPIWithMocks(mockedDataStore)
		api.Router.ServeHTTP(w, r)

		So(w.Code, ShouldEqual, http.StatusOK)
		So(len(mockedDataStore.UpdateInstanceCalls()), ShouldEqual, 1)
		So(len(mockedDataStore.GetRecipeCalls()), ShouldEqual, 1)
	})
}

func TestUpdateInstanceReturnsBadRequest(t *testing.T) {
	t.Parallel()
	Convey("A request to process an invalid update instance to mongo returns 400 Bad Request response", t, func() {
		r := httptest.NewRequest("PUT", "http://localhost:22300/recipes/123/instances/456", bytes.NewBufferString(`{"wrong-field":"Test"}`))
		w := httptest.NewRecorder()
		mockedDataStore := &storetest.StorerMock{
			GetRecipeFunc: func(ID string) (*recipe.Response, error) {
				return &recipe.Response{ID: "123", Alias: "Original", Format: "v4", OutputInstances: initialInstance}, nil
			},
			UpdateInstanceFunc: func(recipeID string, instanceIndex int, updates recipe.Instance) error {
				return nil
			},
		}

		api := GetAPIWithMocks(mockedDataStore)
		api.Router.ServeHTTP(w, r)

		So(w.Code, ShouldEqual, http.StatusBadRequest)
		So(len(mockedDataStore.UpdateInstanceCalls()), ShouldEqual, 0)
		So(len(mockedDataStore.GetRecipeCalls()), ShouldEqual, 0)
	})
}

func TestUpdateInstanceReturnsError(t *testing.T) {
	t.Parallel()
	Convey("When the api cannot update instance in mongo return an internal server error", t, func() {
		r := httptest.NewRequest("PUT", "http://localhost:22300/recipes/123/instances/456", bytes.NewBufferString(`{`))
		w := httptest.NewRecorder()
		mockedDataStore := &storetest.StorerMock{
			UpdateInstanceFunc: func(recipeID string, instanceIndex int, updates recipe.Instance) error {
				return errs.ErrAddUpdateRecipeBadRequest
			},
			GetRecipeFunc: func(ID string) (*recipe.Response, error) {
				return nil, errs.ErrInternalServer
			},
		}

		api := GetAPIWithMocks(mockedDataStore)
		api.Router.ServeHTTP(w, r)

		So(w.Code, ShouldEqual, http.StatusInternalServerError)
		So(len(mockedDataStore.UpdateInstanceCalls()), ShouldEqual, 0)
		So(len(mockedDataStore.GetRecipeCalls()), ShouldEqual, 0)
	})
}

func TestAddCodelistReturnsOK(t *testing.T) {
	t.Parallel()
	Convey("A successful request to add codelist in instance of existing recipe to mongo returns 200 OK response", t, func() {
		r := httptest.NewRequest("POST", "http://localhost:22300/recipes/123/instances/456/codelists", bytes.NewBufferString(codelistTest))
		w := httptest.NewRecorder()
		mockedDataStore := &storetest.StorerMock{
			GetRecipeFunc: func(ID string) (*recipe.Response, error) {
				return &recipe.Response{ID: "123", Alias: "Original", Format: "v4", OutputInstances: initialInstance}, nil
			},
			AddCodelistFunc: func(recipeID string, instanceIndex int, currentRecipe *recipe.Response) error {
				return nil
			},
		}
		api := GetAPIWithMocks(mockedDataStore)
		api.Router.ServeHTTP(w, r)

		So(w.Code, ShouldEqual, http.StatusOK)
		So(len(mockedDataStore.AddCodelistCalls()), ShouldEqual, 1)
		So(len(mockedDataStore.GetRecipeCalls()), ShouldEqual, 1)
	})
}

func TestAddCodelistReturnsBadRequestError(t *testing.T) {
	t.Parallel()
	Convey("A request to process and add an invalid codelist to recipe to mongo returns 400 Bad Request response", t, func() {
		r := httptest.NewRequest("POST", "http://localhost:22300/recipes/123/instances/456/codelists", bytes.NewBufferString(`{"name":"test"}`))
		w := httptest.NewRecorder()
		mockedDataStore := &storetest.StorerMock{
			GetRecipeFunc: func(ID string) (*recipe.Response, error) {
				return &recipe.Response{ID: "123", Alias: "Original", Format: "v4", OutputInstances: initialInstance}, nil
			},
			AddCodelistFunc: func(recipeID string, instanceIndex int, currentRecipe *recipe.Response) error {
				return nil
			},
		}
		api := GetAPIWithMocks(mockedDataStore)
		api.Router.ServeHTTP(w, r)

		So(w.Code, ShouldEqual, http.StatusBadRequest)
		So(len(mockedDataStore.AddCodelistCalls()), ShouldEqual, 0)
		So(len(mockedDataStore.GetRecipeCalls()), ShouldEqual, 0)
	})
}

func TestAddCodelistReturnsError(t *testing.T) {
	t.Parallel()
	Convey("When the api cannot add codelist in mongo return an internal server error", t, func() {
		r := httptest.NewRequest("POST", "http://localhost:22300/recipes/123/instances/456/codelists", bytes.NewBufferString(`{`))
		w := httptest.NewRecorder()
		mockedDataStore := &storetest.StorerMock{
			AddCodelistFunc: func(recipeID string, instanceIndex int, currentRecipe *recipe.Response) error {
				return errs.ErrAddUpdateRecipeBadRequest
			},
			GetRecipeFunc: func(ID string) (*recipe.Response, error) {
				return nil, errs.ErrInternalServer
			},
		}

		api := GetAPIWithMocks(mockedDataStore)
		api.Router.ServeHTTP(w, r)

		So(w.Code, ShouldEqual, http.StatusInternalServerError)
		So(len(mockedDataStore.AddCodelistCalls()), ShouldEqual, 0)
		So(len(mockedDataStore.GetRecipeCalls()), ShouldEqual, 0)
	})
}

func TestUpdateCodelistReturnsOK(t *testing.T) {
	t.Parallel()
	Convey("A successful request to update codelist of a recipe in mongo returns 200 OK response", t, func() {
		r := httptest.NewRequest("PUT", "http://localhost:22300/recipes/123/instances/456/codelists/789", bytes.NewBufferString(`{"name":"Test"}`))
		w := httptest.NewRecorder()
		mockedDataStore := &storetest.StorerMock{
			GetRecipeFunc: func(ID string) (*recipe.Response, error) {
				return &recipe.Response{ID: "123", Alias: "Original", Format: "v4", OutputInstances: initialInstance}, nil
			},
			UpdateCodelistFunc: func(recipeID string, instanceIndex int, codelistIndex int, updates recipe.CodeList) error {
				return nil
			},
		}
		api := GetAPIWithMocks(mockedDataStore)
		api.Router.ServeHTTP(w, r)

		So(w.Code, ShouldEqual, http.StatusOK)
		So(len(mockedDataStore.UpdateCodelistCalls()), ShouldEqual, 1)
		So(len(mockedDataStore.GetRecipeCalls()), ShouldEqual, 1)
	})
}

func TestUpdateCodelistReturnsBadRequest(t *testing.T) {
	t.Parallel()
	Convey("A request to process an invalid update codelist to mongo returns 400 Bad Request response", t, func() {
		r := httptest.NewRequest("PUT", "http://localhost:22300/recipes/123/instances/456/codelists/789", bytes.NewBufferString(`{"href":"Test"}`))
		w := httptest.NewRecorder()
		mockedDataStore := &storetest.StorerMock{
			GetRecipeFunc: func(ID string) (*recipe.Response, error) {
				return &recipe.Response{ID: "123", Alias: "Original", Format: "v4", OutputInstances: initialInstance}, nil
			},
			UpdateCodelistFunc: func(recipeID string, instanceIndex int, codelistIndex int, updates recipe.CodeList) error {
				return nil
			},
		}

		api := GetAPIWithMocks(mockedDataStore)
		api.Router.ServeHTTP(w, r)

		So(w.Code, ShouldEqual, http.StatusBadRequest)
		So(len(mockedDataStore.UpdateCodelistCalls()), ShouldEqual, 0)
		So(len(mockedDataStore.GetRecipeCalls()), ShouldEqual, 0)
	})
}

func TestUpdateCodelistReturnsError(t *testing.T) {
	t.Parallel()
	Convey("When the api cannot update codelist in mongo return an internal server error", t, func() {
		r := httptest.NewRequest("PUT", "http://localhost:22300/recipes/123/instances/456/codelists/789", bytes.NewBufferString(`{`))
		w := httptest.NewRecorder()
		mockedDataStore := &storetest.StorerMock{
			UpdateCodelistFunc: func(recipeID string, instanceIndex int, codelistIndex int, updates recipe.CodeList) error {
				return errs.ErrAddUpdateRecipeBadRequest
			},
			GetRecipeFunc: func(ID string) (*recipe.Response, error) {
				return nil, errs.ErrInternalServer
			},
		}

		api := GetAPIWithMocks(mockedDataStore)
		api.Router.ServeHTTP(w, r)

		So(w.Code, ShouldEqual, http.StatusInternalServerError)
		So(len(mockedDataStore.UpdateCodelistCalls()), ShouldEqual, 0)
		So(len(mockedDataStore.GetRecipeCalls()), ShouldEqual, 0)
	})
}
