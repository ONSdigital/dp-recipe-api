package store

import (
	"context"

	"github.com/ONSdigital/dp-recipe-api/models"
	"github.com/globalsign/mgo/bson"
)

// DataStore provides a datastore.Storer interface used to store, retrieve, remove or update recipes
type DataStore struct {
	Backend Storer
}

//go:generate moq -out datastoretest/datastore.go -pkg storetest . Storer

// Storer represents basic data access via Get, Remove and Upsert methods (to be implemented in the future).
type Storer interface {
	GetRecipes(ctx context.Context, offset int, limit int) (*models.RecipeResults, error)
	GetRecipe(id string) (*models.Recipe, error)
	AddRecipe(item models.Recipe) error
	UpdateAllRecipe(id string, update bson.M) (err error)
	UpdateRecipe(recipeID string, updates models.Recipe) (err error)
	UpdateInstance(recipeID string, instanceIndex int, updates models.Instance) (err error)
	AddCodelist(recipeID string, instanceIndex int, currentRecipe *models.Recipe) (err error)
	UpdateCodelist(recipeID string, instanceIndex int, codelistIndex int, updates models.CodeList) (err error)
}
