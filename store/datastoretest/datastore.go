package storetest

import (
	"sync"

	"github.com/ONSdigital/dp-recipe-api/recipe"
)

var (
	lockStorerMockGetRecipe  sync.RWMutex
	lockStorerMockGetRecipes sync.RWMutex
)

//StorerMock contains initialises methods in Storer interface for mock
type StorerMock struct {
	// GetRecipeFunc mocks the GetRecipe method.
	GetRecipeFunc func(ID string) (*recipe.Response, error)

	// GetRecipesFunc mocks the GetRecipes method.
	GetRecipesFunc func() ([]recipe.Response, error)

	// calls tracks calls to the methods.
	calls struct {
		// GetRecipe holds details about calls to the GetRecipe method.
		GetRecipe []struct {
			// ID is the ID argument value.
			ID string
		}
		// GetRecipes holds details about calls to the GetRecipes method.
		GetRecipes []struct {
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
func (mock *StorerMock) GetRecipes() ([]recipe.Response, error) {
	if mock.GetRecipesFunc == nil {
		panic("StorerMock.GetDatasetsFunc: method is nil but Storer.GetDatasets was just called")
	}
	callInfo := struct {
	}{}
	lockStorerMockGetRecipes.Lock()
	mock.calls.GetRecipes = append(mock.calls.GetRecipes, callInfo)
	lockStorerMockGetRecipes.Unlock()
	return mock.GetRecipesFunc()
}

//GetRecipesCalls gets all the calls that were made to GetRecipes.
// Check the length with:
//     len(mockedStorer.GetRecipesCalls())
func (mock *StorerMock) GetRecipesCalls() []struct {
} {
	var calls []struct {
	}
	lockStorerMockGetRecipes.RLock()
	calls = mock.calls.GetRecipes
	lockStorerMockGetRecipes.RUnlock()
	return calls
}
