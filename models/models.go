package models

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/ONSdigital/log.go/v2/log"
)

const (
	v4                          = "v4"
	cantabularBlob              = "cantabular_blob"
	cantabularTable             = "cantabular_table"
	cantabularFlexibleTable     = "cantabular_flexible_table"
	cantabularMultivariateTable = "cantabular_multivariate_table"
)

// RecipeResults - struct for list of recipes
type RecipeResults struct {
	Count      int       `json:"count"`
	Offset     int       `json:"offset"`
	Limit      int       `json:"limit"`
	TotalCount int       `json:"total_count"`
	Items      []*Recipe `json:"items"`
}

// Recipe - struct for individual recipe
type Recipe struct {
	ID              string     `bson:"_id,omitempty" json:"id,omitempty"`
	Alias           string     `bson:"alias,omitempty" json:"alias,omitempty"`
	Format          string     `bson:"format,omitempty" json:"format,omitempty"`
	InputFiles      []file     `bson:"files,omitempty" json:"files,omitempty"`
	OutputInstances []Instance `bson:"output_instances,omitempty" json:"output_instances,omitempty"`
	CantabularBlob  string     `bson:"cantabular_blob,omitempty" json:"cantabular_blob,omitempty"`
}

// CodeList - Code lists for instance
type CodeList struct {
	ID                           string `bson:"id,omitempty" json:"id,omitempty"`
	HRef                         string `bson:"href,omitempty" json:"href,omitempty"`
	Name                         string `bson:"name,omitempty" json:"name,omitempty"`
	IsHierarchy                  *bool  `bson:"is_hierarchy,omitempty" json:"is_hierarchy,omitempty"`
	IsCantabularGeography        *bool  `bson:"is_cantabular_geography,omitempty" json:"is_cantabular_geography,omitempty"`
	IsCantabularDefaultGeography *bool  `bson:"is_cantabular_default_geography,omitempty" json:"is_cantabular_default_geography,omitempty"`
}

// Instance - struct for instance of recipe
type Instance struct {
	DatasetID       string     `bson:"dataset_id,omitempty" json:"dataset_id,omitempty"`
	Editions        []string   `bson:"editions,omitempty" json:"editions,omitempty"`
	Title           string     `bson:"title,omitempty" json:"title,omitempty"`
	CodeLists       []CodeList `bson:"code_lists,omitempty" json:"code_lists,omitempty"`
	LowestGeography string     `bson:"lowest_geography,omitempty" json:"lowest_geography,omitempty"`
}

type file struct {
	Description string `bson:"description,omitempty" json:"description,omitempty"`
}

// HRefURL - the href of all current code lists
const HRefURL = "http://localhost:22400/code-lists/"

var (
	validFormats = map[string]bool{
		v4:                          true,
		cantabularBlob:              true,
		cantabularTable:             true,
		cantabularFlexibleTable:     true,
		cantabularMultivariateTable: true,
	}
)

func (recipe *Recipe) IsCantabularFlexibleTable() bool {
	return recipe.Format == cantabularFlexibleTable || recipe.Format == cantabularMultivariateTable
}

func (recipe *Recipe) IsCantabularType() bool {
	return recipe.Format == cantabularBlob ||
		recipe.Format == cantabularTable ||
		recipe.Format == cantabularFlexibleTable ||
		recipe.Format == cantabularMultivariateTable
}

// validateInstance - checks if fields of OutputInstances are not empty for ValidateAddRecipe and ValidateAddInstance
func (instance *Instance) validateInstance(ctx context.Context, recipe *Recipe) (missingFields []string, invalidFields []string) {
	if instance.DatasetID == "" {
		missingFields = append(missingFields, "dataset_id")
	}

	if instance.Editions != nil && len(instance.Editions) > 0 {
		for j, edition := range instance.Editions {
			if edition == "" {
				missingFields = append(missingFields, "editions["+strconv.Itoa(j)+"]")
			}
		}
	} else {
		missingFields = append(missingFields, "editions")
	}

	if instance.Title == "" {
		missingFields = append(missingFields, "title")
	}

	if instance.CodeLists != nil && len(instance.CodeLists) > 0 {
		for i, codelist := range instance.CodeLists {
			missing, invalid := codelist.validateCodelist(ctx, recipe)

			if len(missing) > 0 {
				for j, field := range missing {
					missing[j] = "code-lists[" + strconv.Itoa(i) + "]." + field
				}
			}
			if len(invalid) > 0 {
				for j, field := range invalid {
					invalid[j] = "code-lists[" + strconv.Itoa(i) + "]." + field
				}
			}

			missingFields = append(missingFields, missing...)
			invalidFields = append(invalidFields, invalid...)
		}
	} else {
		missingFields = append(missingFields, "code-lists")
	}

	if recipe.IsCantabularType() {
		if len(instance.LowestGeography) == 0 {
			missingFields = append(missingFields, "lowest_geography")
		}
	}
	return
}

// validateCodelists - checks if fields of CodeList are not empty for ValidateAddRecipe, ValidateAddInstance, ValidateAddCodelist
func (c *CodeList) validateCodelist(ctx context.Context, recipe *Recipe) (missingFields []string, invalidFields []string) {

	if c.ID == "" {
		missingFields = append(missingFields, "id")
	}

	if c.HRef == "" {
		missingFields = append(missingFields, "href")
	} else {
		if !c.validateCodelistHRef(ctx) {
			invalidFields = append(invalidFields, "href should be in format (URL/id)")
		}
	}

	if c.Name == "" {
		missingFields = append(missingFields, "name")
	}

	if recipe.IsCantabularType() && c.Name != c.ID {
		invalidFields = append(invalidFields, "name and id should be matching values")
	}

	if c.IsHierarchy == nil {
		missingFields = append(missingFields, "isHierarchy")
	}

	if recipe.IsCantabularType() && c.IsCantabularGeography == nil {
		missingFields = append(missingFields, "isCantabularGeography")
	}

	if recipe.IsCantabularType() && c.IsCantabularDefaultGeography == nil {
		missingFields = append(missingFields, "isCantabularDefaultGeography")
	}

	return
}

// ValidateCodelistHRef - checks if the format of the codelist.href is correct
func (c *CodeList) validateCodelistHRef(ctx context.Context) bool {
	href, err := url.Parse(c.HRef)
	if err != nil {
		log.Error(ctx, "error parsing codelist.href", err)
		return false
	}
	validPath := strings.Contains(href.Path, "/code-lists") && strings.Contains(href.Path, c.ID)
	if href.Scheme != "" && href.Host != "" && validPath {
		return true
	}
	return false
}

// ValidateAddRecipe - checks if all the fields of the recipe are non-empty
func (r *Recipe) ValidateAddRecipe(ctx context.Context) error {
	var missingFields, invalidFields []string

	// recipe.ID generated by API if ID not given so never missing (generates a V4 UUID)

	if r.Alias == "" {
		missingFields = append(missingFields, "alias")
	}
	if r.Format == "" {
		missingFields = append(missingFields, "format")
	} else {
		if !validFormats[r.Format] {
			invalidFields = append(invalidFields, "format is not valid")
		}
	}

	if r.IsCantabularType() && r.CantabularBlob == "" {
		missingFields = append(missingFields, r.Format)
	}

	if r.Format == v4 {
		if r.InputFiles != nil && len(r.InputFiles) > 0 {
			for i, file := range r.InputFiles {
				if file.Description == "" {
					missingFields = append(missingFields, "input-files["+strconv.Itoa(i)+"].description")
				}
			}
		} else {
			missingFields = append(missingFields, "input-files")
		}
	}

	if r.OutputInstances != nil && len(r.OutputInstances) > 0 {
		for i, instance := range r.OutputInstances {
			missing, invalid := instance.validateInstance(ctx, r)
			if len(missing) > 0 {
				for j, field := range missing {
					missing[j] = "output-instances[" + strconv.Itoa(i) + "]." + field
				}
			}
			if len(invalid) > 0 {
				for j, field := range invalid {
					invalid[j] = "output-instances[" + strconv.Itoa(i) + "]." + field
				}
			}
			missingFields = append(missingFields, missing...)
			invalidFields = append(invalidFields, invalid...)
		}
	} else {
		missingFields = append(missingFields, "output-instances")
	}

	if missingFields != nil {
		return fmt.Errorf("missing mandatory fields: %v", missingFields)
	}

	if invalidFields != nil {
		return fmt.Errorf("invalid fields: %v", invalidFields)
	}

	return nil
}

// ValidateAddInstance - checks if fields of OutputInstances are not empty
func (instance *Instance) ValidateAddInstance(ctx context.Context, recipe *Recipe) error {
	var missingFields, invalidFields []string

	missing, invalid := instance.validateInstance(ctx, recipe)
	missingFields = append(missingFields, missing...)
	invalidFields = append(invalidFields, invalid...)

	if missingFields != nil {
		return fmt.Errorf("missing mandatory fields: %v", missingFields)
	}

	if invalidFields != nil {
		return fmt.Errorf("invalid fields: %v", invalidFields)
	}

	return nil
}

// ValidateAddCodelist - checks if fields of Codelist are not empty
func (c *CodeList) ValidateAddCodelist(ctx context.Context, recipe *Recipe) error {
	var missingFields, invalidFields []string

	missing, invalid := c.validateCodelist(ctx, recipe)
	missingFields = append(missingFields, missing...)
	invalidFields = append(invalidFields, invalid...)

	if missingFields != nil {
		return fmt.Errorf("missing mandatory fields: %v", missingFields)
	}

	if invalidFields != nil {
		return fmt.Errorf("invalid fields: %v", invalidFields)
	}

	return nil
}

// ValidateUpdateRecipe - checks updates of recipe for PUT request
func (r *Recipe) ValidateUpdateRecipe(ctx context.Context) error {
	var missingFields, invalidFields []string

	// Validation to check at least one field of the recipe is updated
	stringFieldsEmpty := r.ID == "" && r.Format == "" && r.Alias == ""
	remainingFieldsNil := r.InputFiles == nil && r.OutputInstances == nil
	if stringFieldsEmpty && remainingFieldsNil {
		invalidFields = append(invalidFields, "no recipe fields updates given")
	}

	if r.ID != "" {
		invalidFields = append(invalidFields, "id cannot be changed")
	}
	if r.Format != "" {
		if !validFormats[r.Format] {
			invalidFields = append(invalidFields, "format is not valid")
		}
	}

	if r.InputFiles != nil && len(r.InputFiles) > 0 {
		for i, file := range r.InputFiles {
			if file.Description == "" {
				invalidFields = append(invalidFields, "empty input-files["+strconv.Itoa(i)+"].description given")
			}
		}
	}
	if r.InputFiles != nil && len(r.InputFiles) == 0 {
		invalidFields = append(invalidFields, "empty input-files update given")
	}

	// When doing the update, as recipe.OutputInstances is an array, it needs to make sure that all fields of the instance are complete
	// This functionality is already available in validateInstance
	if r.OutputInstances != nil && len(r.OutputInstances) > 0 {
		for i, instance := range r.OutputInstances {
			missing, invalid := instance.validateInstance(ctx, r)
			if len(missing) > 0 {
				for j, field := range missing {
					missing[j] = "output-instances[" + strconv.Itoa(i) + "]." + field
				}
			}
			if len(invalid) > 0 {
				for j, field := range invalid {
					invalid[j] = "output-instances[" + strconv.Itoa(i) + "]." + field
				}
			}
			missingFields = append(missingFields, missing...)
			invalidFields = append(invalidFields, invalid...)
		}
	}
	if r.OutputInstances != nil && len(r.OutputInstances) == 0 {
		invalidFields = append(invalidFields, "empty output-instances update given")
	}

	if missingFields != nil {
		return fmt.Errorf("missing mandatory fields: %v", missingFields)
	}

	if invalidFields != nil {
		return fmt.Errorf("invalid fields: %v", invalidFields)
	}

	return nil

}

// ValidateUpdateInstance - checks fields of instance before updating the instance of the recipe
func (instance *Instance) ValidateUpdateInstance(ctx context.Context, recipe *Recipe) error {
	var missingFields, invalidFields []string

	// Validation to check if at least one instance field is updated
	stringFieldsEmpty := instance.DatasetID == "" && instance.Title == ""
	remainingFieldsNil := instance.Editions == nil && instance.CodeLists == nil
	if stringFieldsEmpty && remainingFieldsNil {
		invalidFields = append(invalidFields, "no instance fields updates given")
	}

	if instance.Editions != nil && len(instance.Editions) > 0 {
		for j, edition := range instance.Editions {
			if edition == "" {
				missingFields = append(missingFields, "editions["+strconv.Itoa(j)+"]")
			}
		}
	}
	if instance.Editions != nil && len(instance.Editions) == 0 {
		missingFields = append(missingFields, "editions")
	}

	// When doing the update, as instance.Codelist is an array, it needs to make sure that all fields of the codelist are complete
	// This functionality is already available in validateCodelists
	if instance.CodeLists != nil && len(instance.CodeLists) > 0 {
		for i, codelist := range instance.CodeLists {
			missing, invalid := codelist.validateCodelist(ctx, recipe)
			if len(missing) > 0 {
				for j, field := range missing {
					missing[j] = "code-lists[" + strconv.Itoa(i) + "]." + field
				}
			}
			if len(invalid) > 0 {
				for j, field := range invalid {
					invalid[j] = "code-lists[" + strconv.Itoa(i) + "]." + field
				}
			}
			missingFields = append(missingFields, missing...)
			invalidFields = append(invalidFields, invalid...)
		}
	}

	// If any fields missing from the code lists of the instance
	if missingFields != nil {
		return fmt.Errorf("missing mandatory fields: %v", missingFields)
	}

	if invalidFields != nil {
		return fmt.Errorf("invalid fields: %v", invalidFields)
	}

	return nil

}

// ValidateUpdateCodeList - checks fields of codelist before updating the codelist in instance of the recipe
func (c *CodeList) ValidateUpdateCodeList(ctx context.Context) error {
	var invalidFields []string

	// Validation to check if at least one codelist field is updated
	stringFieldsEmpty := c.ID == "" && c.Name == "" && c.HRef == ""
	if stringFieldsEmpty && c.IsHierarchy == nil {
		invalidFields = append(invalidFields, "no codelist fields updates given")
	}

	if c.HRef != "" {
		if !c.validateCodelistHRef(ctx) {
			invalidFields = append(invalidFields, "href should be in format (URL/id)")
		}
	}

	if invalidFields != nil {
		return fmt.Errorf("invalid fields: %v", invalidFields)
	}

	return nil

}
