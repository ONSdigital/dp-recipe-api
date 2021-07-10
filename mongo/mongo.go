package mongo

import (
	"context"
	"errors"
	dpMongoDriver "github.com/ONSdigital/dp-mongodb/v2/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"strconv"

	errs "github.com/ONSdigital/dp-recipe-api/apierrors"
	"github.com/ONSdigital/dp-recipe-api/models"
	"github.com/ONSdigital/log.go/log"
)

const (
	connectTimeoutInSeconds = 5
	queryTimeoutInSeconds   = 15
)

// Mongo represents a simplistic MongoDB configuration.
type Mongo struct {
	Collection string
	Database   string
	Connection *dpMongoDriver.MongoConnection
	Username   string
	Password   string
	URI        string
	IsSSL      bool
}

func (m *Mongo) getConnectionConfig(shouldEnableReadConcern, shouldEnableWriteConcern bool) *dpMongoDriver.MongoConnectionConfig {
	return &dpMongoDriver.MongoConnectionConfig{
		IsSSL:                   m.IsSSL,
		ConnectTimeoutInSeconds: connectTimeoutInSeconds,
		QueryTimeoutInSeconds:   queryTimeoutInSeconds,

		Username:                      m.Username,
		Password:                      m.Password,
		ClusterEndpoint:               m.URI,
		Database:                      m.Database,
		Collection:                    m.Collection,
		IsWriteConcernMajorityEnabled: shouldEnableWriteConcern,
		IsStrongReadConcernEnabled:    shouldEnableReadConcern,
	}
}

// Init creates a new mgo.Connection with a strong consistency and a write mode of "majority".
func (m *Mongo) Init(ctx context.Context, shouldEnableReadConcern, shouldEnableWriteConcern bool) (err error) {
	if m.Connection != nil {
		return errors.New("Datastore Connection already exists")
	}
	mongoConnection, err := dpMongoDriver.Open(m.getConnectionConfig(shouldEnableReadConcern, shouldEnableWriteConcern))
	if err != nil {
		return err
	}
	m.Connection = mongoConnection
	return nil
}

// GetRecipes retrieves all recipe documents from Mongo
func (m *Mongo) GetRecipes(ctx context.Context, offset int, limit int) (*models.RecipeResults, error) {

	query := m.Connection.GetConfiguredCollection().Find(bson.D{})
	totalCount, err := query.Count(ctx)
	if err != nil {
		if dpMongoDriver.IsErrNoDocumentFound(err) {
			return emptyRecipeResults(offset, limit), nil
		}
		log.Event(ctx, "error counting items", log.ERROR, log.Error(err))
		return nil, err
	}

	var recipeItems []*models.Recipe
	if limit > 0 {
		err := query.
			Sort(nil).
			Skip(offset).
			Limit(limit).
			IterAll(ctx, &recipeItems)
		if err != nil {
			if dpMongoDriver.IsErrNoDocumentFound(err) {
				return emptyRecipeResults(offset, limit), nil
			}
			return nil, err
		}
	}

	return &models.RecipeResults{
		Items:      recipeItems,
		Count:      len(recipeItems),
		TotalCount: int(totalCount),
		Offset:     offset,
		Limit:      limit,
	}, nil
}

func emptyRecipeResults(offset int, limit int) *models.RecipeResults {
	return &models.RecipeResults{
		Items:      []*models.Recipe{},
		Count:      0,
		TotalCount: 0,
		Offset:     offset,
		Limit:      limit,
	}
}

// GetRecipe retrieves a recipe document
func (m *Mongo) GetRecipe(ctx context.Context, id string) (*models.Recipe, error) {
	var recipe models.Recipe
	err := m.Connection.GetConfiguredCollection().FindOne(ctx, bson.M{"_id": id}, &recipe)
	if err != nil {
		if dpMongoDriver.IsErrNoDocumentFound(err) {
			return nil, errs.ErrRecipeNotFound
		}
		return nil, err
	}

	return &recipe, nil
}

//AddRecipe adds a recipe document
func (m *Mongo) AddRecipe(ctx context.Context, item models.Recipe) error {
	_, err := m.Connection.GetConfiguredCollection().UpsertId(ctx, item.ID, item)
	return err
}

//UpdateAllRecipe updates an existing recipe document
func (m *Mongo) UpdateAllRecipe(ctx context.Context, id string, update bson.M) (err error) {
	_, err = m.Connection.GetConfiguredCollection().UpdateId(ctx, id, update)
	if err != nil {
		if dpMongoDriver.IsErrNoDocumentFound(err) {
			return errs.ErrRecipeNotFound
		}
	}
	return err
}

//UpdateRecipe prepares updates in bson.M and then updates existing recipe document
func (m *Mongo) UpdateRecipe(ctx context.Context, recipeID string, updates models.Recipe) (err error) {
	update := bson.M{"$set": updates}
	return m.UpdateAllRecipe(ctx, recipeID, update)
}

//UpdateInstance updates specific instance to existing recipe document
func (m *Mongo) UpdateInstance(ctx context.Context, recipeID string, instanceIndex int, updates models.Instance) (err error) {
	update := bson.M{"$set": bson.M{"output_instances." + strconv.Itoa(instanceIndex): updates}}
	return m.UpdateAllRecipe(ctx, recipeID, update)
}

//AddCodelist adds codelist to a specific instance in existing recipe document
func (m *Mongo) AddCodelist(ctx context.Context, recipeID string, instanceIndex int, currentRecipe *models.Recipe) (err error) {
	update := bson.M{"$set": bson.M{"output_instances." + strconv.Itoa(instanceIndex): currentRecipe}}
	return m.UpdateAllRecipe(ctx, recipeID, update)
}

//UpdateCodelist updates specific codelist of a specific instance in existing recipe document
func (m *Mongo) UpdateCodelist(ctx context.Context, recipeID string, instanceIndex int, codelistIndex int, updates models.CodeList) (err error) {
	update := bson.M{"$set": bson.M{"output_instances." + strconv.Itoa(instanceIndex) + ".code_lists." + strconv.Itoa(codelistIndex): updates}}
	return m.UpdateAllRecipe(ctx, recipeID, update)
}
