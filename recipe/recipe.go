package recipe

import (
	"fmt"
	"strconv"
)

//List - struct for list of recipes
type List struct {
	Count      int        `bson:"count" json:"count"`
	Offset     int        `bson:"offset_index" json:"offset_index"`
	Limit      int        `bson:"limit" json:"limit"`
	Items      []Response `bson:"items" json:"items"`
	TotalCount int        `bson:"total_count" json:"total_count"`
}

//Response - struct for individual recipe
type Response struct {
	ID              string     `bson:"_id" json:"id"`
	Alias           string     `bson:"alias" json:"alias"`
	Format          string     `bson:"format" json:"format"`
	InputFiles      []file     `bson:"files" json:"files"`
	OutputInstances []instance `bson:"output_instances" json:"output_instances"`
}

//CodeList - Code lists for instance
type CodeList struct {
	ID          string `bson:"id" json:"id"`
	HRef        string `bson:"href" json:"href"`
	Name        string `bson:"name" json:"name"`
	IsHierarchy *bool  `bson:"is_hierarchy" json:"is_hierarchy"`
}

type instance struct {
	DatasetID string     `bson:"dataset_id" json:"dataset_id"`
	Editions  []string   `bson:"editions" json:"editions"`
	Title     string     `bson:"title" json:"title"`
	CodeLists []CodeList `bson:"code_lists" json:"code_lists"`
}

type file struct {
	Description string `bson:"description" json:"description"`
}

//Validate - checks if all the fields are non-empty
func (recipe *Response) Validate() error {
	var missingFields []string
	if recipe.ID == "" {
		missingFields = append(missingFields, "id")
	}
	if recipe.Alias == "" {
		missingFields = append(missingFields, "alias")
	}
	if recipe.Format == "" {
		missingFields = append(missingFields, "format")
	}

	if recipe.InputFiles != nil {
		for i, file := range recipe.InputFiles {
			if file.Description == "" {
				missingFields = append(missingFields, "input-files["+strconv.Itoa(i)+"].description")
			}
		}
	} else {
		missingFields = append(missingFields, "input-files")
	}

	if recipe.OutputInstances != nil {
		for i, instance := range recipe.OutputInstances {
			if instance.DatasetID == "" {
				missingFields = append(missingFields, "output-instances["+strconv.Itoa(i)+"].dataset-id")
			}

			if instance.Editions != nil {
				for j, edition := range instance.Editions {
					if edition == "" {
						missingFields = append(missingFields, "output-instances["+strconv.Itoa(i)+"].editions["+strconv.Itoa(j)+"]")
					}
				}
			} else {
				missingFields = append(missingFields, "output-instances["+strconv.Itoa(i)+"].editions")
			}

			if instance.Title == "" {
				missingFields = append(missingFields, "output-instances["+strconv.Itoa(i)+"].title")
			}

			if instance.CodeLists != nil {
				for j, codelist := range instance.CodeLists {
					if codelist.ID == "" {
						missingFields = append(missingFields, "output-instances["+strconv.Itoa(i)+"].codelists["+strconv.Itoa(j)+"].id")
					}

					if codelist.HRef == "" {
						missingFields = append(missingFields, "output-instances["+strconv.Itoa(i)+"].codelists["+strconv.Itoa(j)+"].href")
					}

					if codelist.Name == "" {
						missingFields = append(missingFields, "output-instances["+strconv.Itoa(i)+"].codelists["+strconv.Itoa(j)+"].name")
					}

					if codelist.IsHierarchy == nil {
						missingFields = append(missingFields, "output-instances["+strconv.Itoa(i)+"].codelists["+strconv.Itoa(j)+"].isHierarchy")
					}
				}
			} else {
				missingFields = append(missingFields, "output-instances["+strconv.Itoa(i)+"].codelists")
			}

		}
	} else {
		missingFields = append(missingFields, "output-instances")
	}

	if missingFields != nil {
		return fmt.Errorf("missing mandatory fields: %v", missingFields)
	}

	return nil

}
