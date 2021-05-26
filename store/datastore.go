package store

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/ONSdigital/dp-recipe-api/models"
)

// DataStore provides a datastore.Storer interface used to store, retrieve, remove or update recipes
type DataStore struct {
	Backend Storer
}

//go:generate moq -out datastoretest/datastore.go -pkg storetest . Storer

// Storer represents basic data access via Get, Remove and Upsert methods (to be implemented in the future).
type Storer interface {
	GetRecipes(ctx context.Context, offset int, limit int) (*models.RecipeResults, error)
	GetRecipe(ctx context.Context, id string) (*models.Recipe, error)
	AddRecipe(ctx context.Context, item models.Recipe) error
	UpdateAllRecipe(ctx context.Context, id string, update bson.M) (err error)
	UpdateRecipe(ctx context.Context, recipeID string, updates models.Recipe) (err error)
	UpdateInstance(ctx context.Context, recipeID string, instanceIndex int, updates models.Instance) (err error)
	AddCodelist(ctx context.Context, recipeID string, instanceIndex int, currentRecipe *models.Recipe) (err error)
	UpdateCodelist(ctx context.Context, recipeID string, instanceIndex int, codelistIndex int, updates models.CodeList) (err error)
}
