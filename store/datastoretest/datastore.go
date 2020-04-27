package storetest

import (
	"context"
	"sync"

	"github.com/ONSdigital/dp-recipe-api/recipe"
)

var (
	lockStorerMockGetRecipe    sync.RWMutex
	lockStorerMockGetRecipes   sync.RWMutex
	lockStorerMockAddRecipe    sync.RWMutex
	lockStorerMockUpdateRecipe sync.RWMutex
)

//StorerMock contains initialises methods in Storer interface for mock
type StorerMock struct {
	// GetRecipeFunc mocks the GetRecipe method.
	GetRecipeFunc func(ID string) (*recipe.Response, error)

	// GetRecipesFunc mocks the GetRecipes method.
	GetRecipesFunc func(ctx context.Context) ([]recipe.Response, error)

	// AddRecipeFunc mocks the AddRecipe method.
	AddRecipeFunc func(item recipe.Response) error

	// UpdateRecipeFunc mocks the UpdateRecipe method.
	UpdateRecipeFunc func(ID string, recipeUpdate recipe.Response) error

	// calls tracks calls to the methods.
	calls struct {
		// GetRecipe holds details about calls to the GetRecipe method.
		GetRecipe []struct {
			// ID is the ID argument value.
			ID string
		}
		// GetRecipes holds details about calls to the GetRecipes method.
		GetRecipes []struct {
			// ctx is context
			ctx context.Context
		}
		//AddRecipe holds details about calls to the AddRecipe method
		AddRecipe []struct {
			// item represents each recipe
			item recipe.Response
		}
		//UpdateRecipe holds details about calls to the UpdateRecipe method
		UpdateRecipe []struct {
			// ID is the ID argument value.
			ID string
			// recipeUpdate represents each recipe to be updated
			recipeUpdate recipe.Response
		}
	}
}

// GetRecipe calls GetRecipeFunc.
func (mock *StorerMock) GetRecipe(ID string) (*recipe.Response, error) {
	if mock.GetRecipeFunc == nil {
		panic("StorerMock.GetRecipeFunc: method is nil but Storer.GetRecipe was just called")
	}
	callInfo := struct {
		ID string
	}{
		ID: ID,
	}
	lockStorerMockGetRecipe.Lock()
	mock.calls.GetRecipe = append(mock.calls.GetRecipe, callInfo)
	lockStorerMockGetRecipe.Unlock()
	return mock.GetRecipeFunc(ID)
}

// GetRecipeCalls gets all the calls that were made to GetRecipe.
// Check the length with:
//     len(mockedStorer.GetRecipeCalls())
func (mock *StorerMock) GetRecipeCalls() []struct {
	ID string
} {
	var calls []struct {
		ID string
	}
	lockStorerMockGetRecipe.RLock()
	calls = mock.calls.GetRecipe
	lockStorerMockGetRecipe.RUnlock()
	return calls
}

//GetRecipes calls GetRecipesFunc.
func (mock *StorerMock) GetRecipes(ctx context.Context) ([]recipe.Response, error) {
	if mock.GetRecipesFunc == nil {
		panic("StorerMock.GetRecipesFunc: method is nil but Storer.GetRecipes was just called")
	}
	callInfo := struct {
		ctx context.Context
	}{
		ctx: ctx,
	}
	lockStorerMockGetRecipes.Lock()
	mock.calls.GetRecipes = append(mock.calls.GetRecipes, callInfo)
	lockStorerMockGetRecipes.Unlock()
	return mock.GetRecipesFunc(ctx)
}

//GetRecipesCalls gets all the calls that were made to GetRecipes.
// Check the length with:
//     len(mockedStorer.GetRecipesCalls())
func (mock *StorerMock) GetRecipesCalls() []struct {
	ctx context.Context
} {
	var calls []struct {
		ctx context.Context
	}
	lockStorerMockGetRecipes.RLock()
	calls = mock.calls.GetRecipes
	lockStorerMockGetRecipes.RUnlock()
	return calls
}

//AddRecipe calls AddRecipeFunc.
func (mock *StorerMock) AddRecipe(item recipe.Response) error {
	if mock.AddRecipeFunc == nil {
		panic("StorerMock.AddRecipeFunc: method is nil but Storer.AddRecipe was just called")
	}
	callInfo := struct {
		item recipe.Response
	}{
		item: item,
	}
	lockStorerMockAddRecipe.Lock()
	mock.calls.AddRecipe = append(mock.calls.AddRecipe, callInfo)
	lockStorerMockAddRecipe.Unlock()
	return mock.AddRecipeFunc(item)
}

// AddRecipeCalls gets all the calls that were made to AddRecipe.
// Check the length with:
//     len(mockedStorer.AddRecipeCalls())
func (mock *StorerMock) AddRecipeCalls() []struct {
	item recipe.Response
} {
	var calls []struct {
		item recipe.Response
	}
	lockStorerMockAddRecipe.RLock()
	calls = mock.calls.AddRecipe
	lockStorerMockAddRecipe.RUnlock()
	return calls
}

// UpdateRecipe calls UpdateRecipeFunc.
func (mock *StorerMock) UpdateRecipe(ID string, recipeUpdate recipe.Response) error {
	if mock.UpdateRecipeFunc == nil {
		panic("StorerMock.UpdateRecipeFunc: method is nil but Storer.UpdateRecipe was just called")
	}
	callInfo := struct {
		ID           string
		recipeUpdate recipe.Response
	}{
		ID:           ID,
		recipeUpdate: recipeUpdate,
	}
	lockStorerMockUpdateRecipe.Lock()
	mock.calls.UpdateRecipe = append(mock.calls.UpdateRecipe, callInfo)
	lockStorerMockUpdateRecipe.Unlock()
	return mock.UpdateRecipeFunc(ID, recipeUpdate)
}

// UpdateRecipeCalls gets all the calls that were made to UpdateRecipe.
// Check the length with:
//     len(mockedStorer.UpdateRecipeCalls())
func (mock *StorerMock) UpdateRecipeCalls() []struct {
	ID           string
	recipeUpdate recipe.Response
} {
	var calls []struct {
		ID           string
		recipeUpdate recipe.Response
	}
	lockStorerMockUpdateRecipe.RLock()
	calls = mock.calls.UpdateRecipe
	lockStorerMockUpdateRecipe.RUnlock()
	return calls
}
