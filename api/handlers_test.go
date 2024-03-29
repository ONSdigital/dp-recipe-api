package api

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	errs "github.com/ONSdigital/dp-recipe-api/apierrors"
	"github.com/ONSdigital/dp-recipe-api/models"
	storetest "github.com/ONSdigital/dp-recipe-api/store/datastoretest"
	. "github.com/smartystreets/goconvey/convey"
)

var (

	// Variables to point to bool for isHierarchy in CodeLists
	falseVal    = false
	falseValPtr = &falseVal
	trueVal     = true
	trueValPtr  = &trueVal

	// Test data for request body
	recipeTest   = `{"id":"123","alias":"test","format":"v4","files":[{"description":"test files"}],"output_instances":[` + instanceTest + `]}`
	instanceTest = `{"dataset_id":"1234","editions":["edition-test"],"title":"test","code_lists" :[` + codelistTest + `]}`
	codelistTest = `{"id":"12345", "href":"http://localhost:22400/code-lists/12345", "name":"codelist-test-name", "is_hierarchy":false, "is_cantabular_geography":true, "is_cantabular_default_geography":true}`

	// Test data of recipe retrieved from GetRecipe()
	initialCodelist = []models.CodeList{
		{
			ID:                           "789",
			Name:                         "codelist-test",
			HRef:                         "http://localhost:22400/code-lists/789",
			IsHierarchy:                  falseValPtr,
			IsCantabularGeography:        trueValPtr,
			IsCantabularDefaultGeography: trueValPtr,
		},
	}
	initialInstance = []models.Instance{
		{
			DatasetID: "456",
			Editions:  []string{"editions"},
			Title:     "test",
			CodeLists: initialCodelist,
		},
	}
)

func TestGetRecipesReturnsOK(t *testing.T) {
	t.Parallel()
	Convey("A successful request to get recipe returns 200 OK response", t, func() {
		r := httptest.NewRequest("GET", "http://localhost:22300/recipes", nil)
		w := httptest.NewRecorder()
		mockedDataStore := &storetest.StorerMock{
			GetRecipesFunc: func(ctx context.Context, offset int, limit int) (*models.RecipeResults, error) {
				return &models.RecipeResults{}, nil
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
			GetRecipesFunc: func(ctx context.Context, offset int, limit int) (*models.RecipeResults, error) {
				return &models.RecipeResults{}, nil
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
			GetRecipesFunc: func(ctx context.Context, offset int, limit int) (*models.RecipeResults, error) {
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
			GetRecipeFunc: func(ctx context.Context, ID string) (*models.Recipe, error) {
				return &models.Recipe{ID: "123"}, nil
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
			GetRecipeFunc: func(ctx context.Context, ID string) (*models.Recipe, error) {
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
			GetRecipeFunc: func(ctx context.Context, ID string) (*models.Recipe, error) {
				return nil, errs.ErrRecipeNotFound
			},
		}

		api := GetAPIWithMocks(mockedDataStore)
		api.Router.ServeHTTP(w, r)

		So(w.Code, ShouldEqual, http.StatusNotFound)
		So(len(mockedDataStore.GetRecipeCalls()), ShouldEqual, 1)
	})
}

func TestAddRecipeReturnsOK(t *testing.T) {
	t.Parallel()
	Convey("A successful request to add recipe to mongo returns 200 OK response", t, func() {
		r := httptest.NewRequest("POST", "http://localhost:22300/recipes", bytes.NewBufferString(recipeTest))
		w := httptest.NewRecorder()
		mockedDataStore := &storetest.StorerMock{
			AddRecipeFunc: func(ctx context.Context, item models.Recipe) error {
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
			AddRecipeFunc: func(ctx context.Context, item models.Recipe) error {
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
			AddRecipeFunc: func(ctx context.Context, item models.Recipe) error {
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
			UpdateRecipeFunc: func(ctx context.Context, ID string, updates models.Recipe) error {
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
			UpdateRecipeFunc: func(ctx context.Context, ID string, updates models.Recipe) error {
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
			UpdateRecipeFunc: func(ctx context.Context, ID string, updates models.Recipe) error {
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
			GetRecipeFunc: func(ctx context.Context, ID string) (*models.Recipe, error) {
				return &models.Recipe{ID: "123", Alias: "Original", Format: "v4"}, nil
			},
			UpdateRecipeFunc: func(ctx context.Context, recipeID string, currentRecipe models.Recipe) error {
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
			GetRecipeFunc: func(ctx context.Context, ID string) (*models.Recipe, error) {
				return &models.Recipe{ID: "123", Alias: "Original", Format: "v4"}, nil
			},
			UpdateRecipeFunc: func(ctx context.Context, recipeID string, currentRecipe models.Recipe) error {
				return nil
			},
		}
		api := GetAPIWithMocks(mockedDataStore)
		api.Router.ServeHTTP(w, r)

		So(w.Code, ShouldEqual, http.StatusBadRequest)
		So(len(mockedDataStore.UpdateRecipeCalls()), ShouldEqual, 0)
		So(len(mockedDataStore.GetRecipeCalls()), ShouldEqual, 1)
	})
}

func TestAddInstanceReturnsError(t *testing.T) {
	t.Parallel()
	Convey("When the api cannot add instance in mongo return an internal server error", t, func() {
		r := httptest.NewRequest("POST", "http://localhost:22300/recipes/123/instances", bytes.NewBufferString(`{`))
		w := httptest.NewRecorder()
		mockedDataStore := &storetest.StorerMock{
			UpdateRecipeFunc: func(ctx context.Context, recipeID string, currentRecipe models.Recipe) error {
				return errs.ErrAddUpdateRecipeBadRequest
			},
			GetRecipeFunc: func(ctx context.Context, ID string) (*models.Recipe, error) {
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
			GetRecipeFunc: func(ctx context.Context, ID string) (*models.Recipe, error) {
				return &models.Recipe{ID: "123", Alias: "Original", Format: "v4", OutputInstances: initialInstance}, nil
			},
			UpdateInstanceFunc: func(ctx context.Context, recipeID string, instanceIndex int, updates models.Instance) error {
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
			GetRecipeFunc: func(ctx context.Context, ID string) (*models.Recipe, error) {
				return &models.Recipe{ID: "123", Alias: "Original", Format: "v4", OutputInstances: initialInstance}, nil
			},
			UpdateInstanceFunc: func(ctx context.Context, recipeID string, instanceIndex int, updates models.Instance) error {
				return nil
			},
		}

		api := GetAPIWithMocks(mockedDataStore)
		api.Router.ServeHTTP(w, r)

		So(w.Code, ShouldEqual, http.StatusBadRequest)
		So(len(mockedDataStore.UpdateInstanceCalls()), ShouldEqual, 0)
		So(len(mockedDataStore.GetRecipeCalls()), ShouldEqual, 1)
	})
}

func TestUpdateInstanceReturnsError(t *testing.T) {
	t.Parallel()
	Convey("When the api cannot update instance in mongo return an internal server error", t, func() {
		r := httptest.NewRequest("PUT", "http://localhost:22300/recipes/123/instances/456", bytes.NewBufferString(`{`))
		w := httptest.NewRecorder()
		mockedDataStore := &storetest.StorerMock{
			UpdateInstanceFunc: func(ctx context.Context, recipeID string, instanceIndex int, updates models.Instance) error {
				return errs.ErrAddUpdateRecipeBadRequest
			},
			GetRecipeFunc: func(ctx context.Context, ID string) (*models.Recipe, error) {
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
		r := httptest.NewRequest("POST", "http://localhost:22300/recipes/123/instances/456/code-lists", bytes.NewBufferString(codelistTest))
		w := httptest.NewRecorder()
		mockedDataStore := &storetest.StorerMock{
			GetRecipeFunc: func(ctx context.Context, ID string) (*models.Recipe, error) {
				return &models.Recipe{ID: "123", Alias: "Original", Format: "v4", OutputInstances: initialInstance}, nil
			},
			AddCodelistFunc: func(ctx context.Context, recipeID string, instanceIndex int, currentRecipe *models.Recipe) error {
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
		r := httptest.NewRequest("POST", "http://localhost:22300/recipes/123/instances/456/code-lists", bytes.NewBufferString(`{"name":"test"}`))
		w := httptest.NewRecorder()
		mockedDataStore := &storetest.StorerMock{
			GetRecipeFunc: func(ctx context.Context, ID string) (*models.Recipe, error) {
				return &models.Recipe{ID: "123", Alias: "Original", Format: "v4", OutputInstances: initialInstance}, nil
			},
			AddCodelistFunc: func(ctx context.Context, recipeID string, instanceIndex int, currentRecipe *models.Recipe) error {
				return nil
			},
		}
		api := GetAPIWithMocks(mockedDataStore)
		api.Router.ServeHTTP(w, r)

		So(w.Code, ShouldEqual, http.StatusBadRequest)
		So(len(mockedDataStore.AddCodelistCalls()), ShouldEqual, 0)
		So(len(mockedDataStore.GetRecipeCalls()), ShouldEqual, 1)
	})
}

func TestAddCodelistReturnsError(t *testing.T) {
	t.Parallel()
	Convey("When the api cannot add codelist in mongo return an internal server error", t, func() {
		r := httptest.NewRequest("POST", "http://localhost:22300/recipes/123/instances/456/code-lists", bytes.NewBufferString(`{`))
		w := httptest.NewRecorder()
		mockedDataStore := &storetest.StorerMock{
			AddCodelistFunc: func(ctx context.Context, recipeID string, instanceIndex int, currentRecipe *models.Recipe) error {
				return errs.ErrAddUpdateRecipeBadRequest
			},
			GetRecipeFunc: func(ctx context.Context, ID string) (*models.Recipe, error) {
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
		r := httptest.NewRequest("PUT", "http://localhost:22300/recipes/123/instances/456/code-lists/789", bytes.NewBufferString(`{"name":"Test"}`))
		w := httptest.NewRecorder()
		mockedDataStore := &storetest.StorerMock{
			GetRecipeFunc: func(ctx context.Context, ID string) (*models.Recipe, error) {
				return &models.Recipe{ID: "123", Alias: "Original", Format: "v4", OutputInstances: initialInstance}, nil
			},
			UpdateCodelistFunc: func(ctx context.Context, recipeID string, instanceIndex int, codelistIndex int, updates models.CodeList) error {
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
		r := httptest.NewRequest("PUT", "http://localhost:22300/recipes/123/instances/456/code-lists/789", bytes.NewBufferString(`{"href":"Test"}`))
		w := httptest.NewRecorder()
		mockedDataStore := &storetest.StorerMock{
			GetRecipeFunc: func(ctx context.Context, ID string) (*models.Recipe, error) {
				return &models.Recipe{ID: "123", Alias: "Original", Format: "v4", OutputInstances: initialInstance}, nil
			},
			UpdateCodelistFunc: func(ctx context.Context, recipeID string, instanceIndex int, codelistIndex int, updates models.CodeList) error {
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
		r := httptest.NewRequest("PUT", "http://localhost:22300/recipes/123/instances/456/code-lists/789", bytes.NewBufferString(`{`))
		w := httptest.NewRecorder()
		mockedDataStore := &storetest.StorerMock{
			UpdateCodelistFunc: func(ctx context.Context, recipeID string, instanceIndex int, codelistIndex int, updates models.CodeList) error {
				return errs.ErrAddUpdateRecipeBadRequest
			},
			GetRecipeFunc: func(ctx context.Context, ID string) (*models.Recipe, error) {
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
