package mongo

import (
	"context"
	"errors"
	"strconv"

	"github.com/globalsign/mgo"

	"github.com/globalsign/mgo/bson"

	errs "github.com/ONSdigital/dp-recipe-api/apierrors"
	"github.com/ONSdigital/dp-recipe-api/models"
	"github.com/ONSdigital/log.go/log"
)

// Mongo represents a simplistic MongoDB configuration.
type Mongo struct {
	Collection string
	Database   string
	Session    *mgo.Session
	URI        string
}

// Init creates a new mgo.Session with a strong consistency and a write mode of "majority".
func (m *Mongo) Init() (session *mgo.Session, err error) {
	if session != nil {
		return nil, errors.New("session already exists")
	}

	if session, err = mgo.Dial(m.URI); err != nil {
		return nil, err
	}

	session.EnsureSafe(&mgo.Safe{WMode: "majority"})
	session.SetMode(mgo.Strong, true)
	return session, nil
}

// GetRecipes retrieves all recipe documents from Mongo
func (m *Mongo) GetRecipes(ctx context.Context, offset int, limit int) (*models.RecipeResults, error) {
	s := m.Session.Copy()
	defer s.Close()

	query := s.DB(m.Database).C(m.Collection).Find(nil)
	totalCount, err := query.Count()
	if err != nil {
		log.Event(ctx, "error counting items", log.ERROR, log.Error(err))
		if err == mgo.ErrNotFound {
			return &models.RecipeResults{
				Items:      []*models.Recipe{},
				Count:      0,
				TotalCount: 0,
				Offset:     offset,
				Limit:      limit,
			}, nil
		}
		return nil, err
	}

	var recipeItems []*models.Recipe
	if limit > 0 {
		iter := query.Sort().Skip(offset).Limit(limit).Iter()
		defer func() {
			err := iter.Close()
			if err != nil {
				log.Event(ctx, "error closing job iterator", log.ERROR, log.Error(err), log.Data{})
			}
		}()

		if err := iter.All(&recipeItems); err != nil {
			if err == mgo.ErrNotFound {
				return &models.RecipeResults{
					Items:      []*models.Recipe{},
					Count:      0,
					TotalCount: totalCount,
					Offset:     offset,
					Limit:      limit,
				}, nil
			}
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
func (m *Mongo) GetRecipe(id string) (*models.Recipe, error) {
	s := m.Session.Copy()
	defer s.Close()
	var recipe models.Recipe
	err := s.DB(m.Database).C(m.Collection).Find(bson.M{"_id": id}).One(&recipe)
	if err != nil {
		if err == mgo.ErrNotFound {
			return nil, errs.ErrRecipeNotFound
		}
		return nil, err
	}

	return &recipe, nil
}

//AddRecipe adds a recipe document
func (m *Mongo) AddRecipe(item models.Recipe) error {
	s := m.Session.Copy()
	defer s.Close()
	_, err := s.DB(m.Database).C(m.Collection).UpsertId(item.ID, item)
	return err
}

//UpdateAllRecipe updates an existing recipe document
func (m *Mongo) UpdateAllRecipe(id string, update bson.M) (err error) {
	s := m.Session.Copy()
	defer s.Close()

	err = s.DB(m.Database).C("recipes").UpdateId(id, update)
	if err != nil {
		if err == mgo.ErrNotFound {
			return errs.ErrRecipeNotFound
		}
	}
	return err
}

//UpdateRecipe prepares updates in bson.M and then updates existing recipe document
func (m *Mongo) UpdateRecipe(recipeID string, updates models.Recipe) (err error) {
	update := bson.M{"$set": updates}
	return m.UpdateAllRecipe(recipeID, update)
}

//UpdateInstance updates specific instance to existing recipe document
func (m *Mongo) UpdateInstance(recipeID string, instanceIndex int, updates models.Instance) (err error) {
	update := bson.M{"$set": bson.M{"output_instances." + strconv.Itoa(instanceIndex): updates}}
	return m.UpdateAllRecipe(recipeID, update)
}

//AddCodelist adds codelist to a specific instance in existing recipe document
func (m *Mongo) AddCodelist(recipeID string, instanceIndex int, currentRecipe *models.Recipe) (err error) {
	update := bson.M{"$set": bson.M{"output_instances." + strconv.Itoa(instanceIndex): currentRecipe}}
	return m.UpdateAllRecipe(recipeID, update)
}

//UpdateCodelist updates specific codelist of a specific instance in existing recipe document
func (m *Mongo) UpdateCodelist(recipeID string, instanceIndex int, codelistIndex int, updates models.CodeList) (err error) {
	update := bson.M{"$set": bson.M{"output_instances." + strconv.Itoa(instanceIndex) + ".code_lists." + strconv.Itoa(codelistIndex): updates}}
	return m.UpdateAllRecipe(recipeID, update)
}
