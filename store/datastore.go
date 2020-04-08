package store

import (
	"github.com/ONSdigital/dp-recipe-api/recipe"
)

// DataStore provides a datastore.Storer interface used to store, retrieve, remove or update recipes
type DataStore struct {
	Backend Storer
}

//go:generate moq -out datastoretest/datastore.go -pkg storetest . Storer

// Storer represents basic data access via Get, Remove and Upsert methods (to be implemented in the future).
type Storer interface {
	GetRecipes() ([]recipe.Response, error)
	GetRecipe(id string) (*recipe.Response, error)
	AddRecipe(item recipe.Response) error
	UpdateRecipe(id string, recipeUpdate recipe.Response) (err error)
}
