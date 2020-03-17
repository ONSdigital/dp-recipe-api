package storetest

import (
	"sync"

	"github.com/ONSdigital/dp-recipe-api/recipe"
)

var (
	lockStorerMockGetRecipes sync.RWMutex
)

//StorerMock contains initialises methods in Storer interface for mock
type StorerMock struct {
	// GetRecipesFunc mocks the GetRecipes method.
	GetRecipesFunc func() ([]recipe.Response, error)

	// calls tracks calls to the methods.
	calls struct {
		// GetRecipes holds details about calls to the GetRecipes method.
		GetRecipes []struct {
		}
	}
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
