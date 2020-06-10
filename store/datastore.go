package store

import (
	"context"

	"github.com/ONSdigital/dp-recipe-api/recipe"
	"github.com/globalsign/mgo/bson"
)

// DataStore provides a datastore.Storer interface used to store, retrieve, remove or update recipes
type DataStore struct {
	Backend Storer
}

//go:generate moq -out datastoretest/datastore.go -pkg storetest . Storer

// Storer represents basic data access via Get, Remove and Upsert methods (to be implemented in the future).
type Storer interface {
	GetRecipes(ctx context.Context) ([]recipe.Response, error)
	GetRecipe(id string) (*recipe.Response, error)
	AddRecipe(item recipe.Response) error
	UpdateAllRecipe(id string, update bson.M) (err error)
	UpdateRecipe(recipeID string, updates recipe.Response) (err error)
	AddInstance(recipeID string, currentRecipe *recipe.Response) (err error)
	UpdateInstance(recipeID string, instanceIndex int, updates recipe.Instance) (err error)
	AddCodelist(recipeID string, instanceIndex int, currentRecipe *recipe.Response) (err error)
	UpdateCodelist(recipeID string, instanceIndex int, codelistIndex int, updates recipe.CodeList) (err error)
}
