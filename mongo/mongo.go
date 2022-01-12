package mongo

import (
	"context"
	"strconv"

	"github.com/ONSdigital/dp-healthcheck/healthcheck"
	mongohealth "github.com/ONSdigital/dp-mongodb/v3/health"
	mongodriver "github.com/ONSdigital/dp-mongodb/v3/mongodb"
	errs "github.com/ONSdigital/dp-recipe-api/apierrors"
	"github.com/ONSdigital/dp-recipe-api/models"
	"github.com/ONSdigital/log.go/v2/log"
	"go.mongodb.org/mongo-driver/bson"
)

type Mongo struct {
	mongodriver.MongoConnectionConfig

	connection   *mongodriver.MongoConnection
	healthClient *mongohealth.CheckMongoClient
}

// NewDatastore creates a new mongodb.MongoConnection with the given configuration
func NewDatastore(_ context.Context, cfg mongodriver.MongoConnectionConfig) (m *Mongo, err error) {
	m = &Mongo{MongoConnectionConfig: cfg}

	m.connection, err = mongodriver.Open(&m.MongoConnectionConfig)
	if err != nil {
		return nil, err
	}

	databaseCollectionBuilder := make(map[mongohealth.Database][]mongohealth.Collection)
	databaseCollectionBuilder[(mongohealth.Database)(m.Database)] = []mongohealth.Collection{(mongohealth.Collection)(m.Collection)}

	m.healthClient = mongohealth.NewClientWithCollections(m.connection, databaseCollectionBuilder)

	return m, nil
}

// GetRecipes retrieves all recipe documents from Mongo
func (m *Mongo) GetRecipes(ctx context.Context, offset int, limit int) (*models.RecipeResults, error) {

	var recipeItems []*models.Recipe
	query := m.connection.GetConfiguredCollection().Find(bson.D{})
	totalCount, err := query.Count(ctx)
	if err != nil {
		log.Error(ctx, "error counting items", err)
		return nil, err
	}

	// query the items corresponding to the provided offset and limit (only if necessary)
	// guaranteeing at least one document will be found
	if totalCount > 0 && limit > 0 && offset < totalCount {
		if err = query.Sort(bson.M{"_id": 1}).Skip(offset).Limit(limit).IterAll(ctx, &recipeItems); err != nil {
			return nil, err
		}
	}

	return &models.RecipeResults{
		Items:      recipeItems,
		Count:      len(recipeItems),
		TotalCount: totalCount,
		Offset:     offset,
		Limit:      limit,
	}, nil
}

// GetRecipe retrieves a recipe document
func (m *Mongo) GetRecipe(ctx context.Context, id string) (*models.Recipe, error) {
	var recipe models.Recipe
	err := m.connection.GetConfiguredCollection().FindOne(ctx, bson.M{"_id": id}, &recipe)
	if err != nil {
		if mongodriver.IsErrNoDocumentFound(err) {
			return nil, errs.ErrRecipeNotFound
		}
		return nil, err
	}

	return &recipe, nil
}

// AddRecipe adds a recipe document
func (m *Mongo) AddRecipe(ctx context.Context, item models.Recipe) error {
	_, err := m.connection.GetConfiguredCollection().UpsertById(ctx, item.ID, bson.M{"$set": item})

	return err
}

// UpdateAllRecipe updates an existing recipe document
func (m *Mongo) UpdateAllRecipe(ctx context.Context, id string, update bson.M) (err error) {
	_, err = m.connection.GetConfiguredCollection().Must().UpdateById(ctx, id, update)
	if err != nil {
		if mongodriver.IsErrNoDocumentFound(err) {
			return errs.ErrRecipeNotFound
		}
	}

	return err
}

// UpdateRecipe prepares updates in bson.M and then updates existing recipe document
func (m *Mongo) UpdateRecipe(ctx context.Context, recipeID string, updates models.Recipe) (err error) {
	return m.UpdateAllRecipe(ctx, recipeID, bson.M{"$set": updates})
}

// UpdateInstance updates specific instance to existing recipe document
func (m *Mongo) UpdateInstance(ctx context.Context, recipeID string, instanceIndex int, updates models.Instance) (err error) {
	return m.UpdateAllRecipe(ctx, recipeID, bson.M{"$set": bson.M{"output_instances." + strconv.Itoa(instanceIndex): updates}})
}

// AddCodelist adds codelist to a specific instance in existing recipe document
func (m *Mongo) AddCodelist(ctx context.Context, recipeID string, instanceIndex int, currentRecipe *models.Recipe) (err error) {
	return m.UpdateAllRecipe(ctx, recipeID, bson.M{"$set": bson.M{"output_instances." + strconv.Itoa(instanceIndex): currentRecipe}})
}

// UpdateCodelist updates specific codelist of a specific instance in existing recipe document
func (m *Mongo) UpdateCodelist(ctx context.Context, recipeID string, instanceIndex int, codelistIndex int, updates models.CodeList) (err error) {
	return m.UpdateAllRecipe(ctx, recipeID, bson.M{"$set": bson.M{"output_instances." + strconv.Itoa(instanceIndex) + ".code_lists." + strconv.Itoa(codelistIndex): updates}})
}

// Close closes the mongo session and returns any error
func (m *Mongo) Close(ctx context.Context) error {
	return m.connection.Close(ctx)
}

// Checker is called by the healthcheck library to check the health state of this mongoDB instance
func (m *Mongo) Checker(ctx context.Context, state *healthcheck.CheckState) error {
	return m.healthClient.Checker(ctx, state)
}
