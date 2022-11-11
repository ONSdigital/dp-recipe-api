package models

import (
	"context"
	"errors"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	ctx         = context.Background()
	falseVal    = false
	trueVal     = true
	falseValPtr = &falseVal
	trueValPtr  = &trueVal
)

func createCodeList() CodeList {
	return CodeList{
		ID:                           "789",
		Name:                         "789",
		HRef:                         "http://localhost:22400/code-lists/789",
		IsHierarchy:                  falseValPtr,
		IsCantabularGeography:        falseValPtr,
		IsCantabularDefaultGeography: trueValPtr,
	}
}

func createRecipeData() Recipe {
	return Recipe{
		ID:              "123",
		Alias:           "test",
		Format:          v4,
		InputFiles:      []file{{Description: "test files"}},
		OutputInstances: []Instance{createInstance()},
		CantabularBlob:  "blob1",
	}
}

func createInstance() Instance {
	return Instance{
		DatasetID:       "456",
		Editions:        []string{"editions"},
		Title:           "test",
		CodeLists:       []CodeList{createCodeList()},
		LowestGeography: "lowest_geo",
	}
}

func TestValidateInstance(t *testing.T) {
	t.Parallel()

	Convey("Non-empty missing field successfully returned", t, func() {

		Convey("when datasetID is missing", func() {
			recipe := createRecipeData()
			instance := createInstance()
			instance.DatasetID = ""
			missingFields, invalidFields := instance.validateInstance(ctx, &recipe)
			So(missingFields, ShouldNotBeNil)
			So(missingFields, ShouldResemble, []string{"dataset_id"})
			So(invalidFields, ShouldBeNil)
		})

		Convey("when editions is missing", func() {
			recipe := createRecipeData()
			instance := createInstance()
			instance.Editions = nil
			missingFields, invalidFields := instance.validateInstance(ctx, &recipe)
			So(missingFields, ShouldNotBeNil)
			So(missingFields, ShouldResemble, []string{"editions"})
			So(invalidFields, ShouldBeNil)
		})

		Convey("when an edition of editions is missing", func() {
			recipe := createRecipeData()
			instance := createInstance()
			instance.Editions = []string{""}
			missingFields, invalidFields := instance.validateInstance(ctx, &recipe)
			So(missingFields, ShouldNotBeNil)
			So(missingFields, ShouldResemble, []string{"editions[0]"})
			So(invalidFields, ShouldBeNil)

			instance.Editions = []string{"editions"} // Reset to original
		})

		Convey("when title is missing", func() {
			recipe := createRecipeData()
			instance := createInstance()
			instance.Title = ""
			missingFields, invalidFields := instance.validateInstance(ctx, &recipe)
			So(missingFields, ShouldNotBeNil)
			So(missingFields, ShouldResemble, []string{"title"})
			So(invalidFields, ShouldBeNil)
		})

		Convey("when code lists is missing", func() {
			recipe := createRecipeData()
			instance := createInstance()
			instance.CodeLists = nil
			missingFields, invalidFields := instance.validateInstance(ctx, &recipe)
			So(missingFields, ShouldNotBeNil)
			So(missingFields, ShouldResemble, []string{"code-lists"})
			So(invalidFields, ShouldBeNil)
		})

		Convey("when any field of code lists is missing", func() {
			recipe := createRecipeData()
			instance := createInstance()
			instance.CodeLists[0].Name = ""
			missingFields, invalidFields := instance.validateInstance(ctx, &recipe)
			So(missingFields, ShouldNotBeNil)
			So(missingFields, ShouldResemble, []string{"code-lists[0].name"})
			So(invalidFields, ShouldBeNil)

			instance.CodeLists[0].Name = "codelist-test" // Reset to original
		})

		Convey("when lowest_geography is missing", func() {
			recipe := createRecipeData()
			recipe.Format = "cantabular_flexible_table"
			instance := createInstance()
			instance.LowestGeography = ""
			missingFields, invalidFields := instance.validateInstance(ctx, &recipe)
			So(missingFields, ShouldNotBeNil)
			So(missingFields, ShouldResemble, []string{"lowest_geography"})
			So(invalidFields, ShouldBeNil)
		})

	})

	Convey("Empty missing field successfully returned", t, func() {

		Convey("when all fields are output-instances are given", func() {
			recipe := createRecipeData()
			instance := createInstance()
			missingFields, invalidFields := instance.validateInstance(ctx, &recipe)
			So(missingFields, ShouldBeNil)
			So(invalidFields, ShouldBeNil)
		})

	})

	Convey("Non-empty invalid field successfully returned", t, func() {

		Convey("when code-lists.href is incorrectly entered", func() {
			recipe := createRecipeData()
			instance := createInstance()
			instance.CodeLists[0].HRef = "incorrect-href"
			missingFields, invalidFields := instance.validateInstance(ctx, &recipe)
			So(missingFields, ShouldBeNil)
			So(invalidFields, ShouldNotBeNil)
			So(invalidFields, ShouldResemble, []string{"code-lists[0].href should be in format (URL/id)"})

			instance.CodeLists[0].HRef = "http://localhost:22400/code-lists/789" // Reset
		})

	})

	Convey("Empty invalid field successfully returned", t, func() {

		Convey("when code-lists.href is correctly entered", func() {
			recipe := createRecipeData()
			instance := createInstance()
			instance.CodeLists[0].HRef = "http://localhost:22400/code-lists/789"
			missingFields, invalidFields := instance.validateInstance(ctx, &recipe)
			So(missingFields, ShouldBeNil)
			So(invalidFields, ShouldBeNil)
		})

	})

}

func TestValidateCodelist(t *testing.T) {
	t.Parallel()

	Convey("Non-empty missing field successfully returned", t, func() {

		Convey("when id is missing", func() {
			recipe := createRecipeData()
			codelist := createCodeList()
			codelist.ID = ""
			// HRef Updated as the format of HRef follows the value from ID
			codelist.HRef = "http://localhost:22400/code-lists/"
			missingFields, invalidFields := codelist.validateCodelist(ctx, &recipe)
			So(missingFields, ShouldNotBeNil)
			So(missingFields, ShouldResemble, []string{"id"})
			So(invalidFields, ShouldBeNil)
		})

		Convey("when href is missing", func() {
			recipe := createRecipeData()
			codelist := createCodeList()
			codelist.HRef = ""
			missingFields, invalidFields := codelist.validateCodelist(ctx, &recipe)
			So(missingFields, ShouldNotBeNil)
			So(missingFields, ShouldResemble, []string{"href"})
			So(invalidFields, ShouldBeNil)
		})

		Convey("when name is missing", func() {
			recipe := createRecipeData()
			codelist := createCodeList()
			codelist.Name = ""
			missingFields, invalidFields := codelist.validateCodelist(ctx, &recipe)
			So(missingFields, ShouldNotBeNil)
			So(missingFields, ShouldResemble, []string{"name"})
			So(invalidFields, ShouldBeNil)
		})

		Convey("when ishierarchy is missing", func() {
			recipe := createRecipeData()
			codelist := createCodeList()
			codelist.IsHierarchy = nil
			missingFields, invalidFields := codelist.validateCodelist(ctx, &recipe)
			So(missingFields, ShouldNotBeNil)
			So(missingFields, ShouldResemble, []string{"isHierarchy"})
			So(invalidFields, ShouldBeNil)
		})

		Convey("when isCantabularGeography is missing", func() {
			recipe := createRecipeData()
			codelist := createCodeList()
			codelist.IsCantabularGeography = nil
			recipe.Format = "cantabular_flexible_table"
			missingFields, invalidFields := codelist.validateCodelist(ctx, &recipe)
			So(missingFields, ShouldNotBeNil)
			So(missingFields, ShouldResemble, []string{"isCantabularGeography"})
			So(invalidFields, ShouldBeNil)
		})

		Convey("when isCantabularDefaultGeography is missing", func() {
			recipe := createRecipeData()
			codelist := createCodeList()
			codelist.IsCantabularDefaultGeography = nil
			recipe.Format = "cantabular_flexible_table"
			missingFields, invalidFields := codelist.validateCodelist(ctx, &recipe)
			So(missingFields, ShouldNotBeNil)
			So(missingFields, ShouldResemble, []string{"isCantabularDefaultGeography"})
			So(invalidFields, ShouldBeNil)
		})

	})

	Convey("Empty missing field successfully returned", t, func() {

		Convey("when all fields are codelist are given", func() {
			recipe := createRecipeData()
			codelist := createCodeList()
			missingFields, invalidFields := codelist.validateCodelist(ctx, &recipe)
			So(missingFields, ShouldBeNil)
			So(invalidFields, ShouldBeNil)
		})

	})

	Convey("Non-empty invalid field successfully returned", t, func() {

		Convey("when href is incorrectly entered", func() {
			recipe := createRecipeData()
			codelist := createCodeList()
			codelist.HRef = "incorrect-href"
			missingFields, invalidFields := codelist.validateCodelist(ctx, &recipe)
			So(missingFields, ShouldBeNil)
			So(invalidFields, ShouldNotBeNil)
			So(invalidFields, ShouldResemble, []string{"href should be in format (URL/id)"})
		})

		Convey("when name and id fields do not match", func() {
			recipe := createRecipeData()
			recipe.Format = "cantabular_flexible_table"
			codelist := createCodeList()
			codelist.HRef = "http://localhost:22400/code-lists/789"
			codelist.Name = "123"
			codelist.ID = "789"
			missingFields, invalidFields := codelist.validateCodelist(ctx, &recipe)
			So(missingFields, ShouldBeNil)
			So(invalidFields, ShouldNotBeNil)
			So(invalidFields, ShouldResemble, []string{"name and id should be matching values"})
		})

	})

	Convey("Empty invalid field successfully returned", t, func() {

		Convey("when href is correctly entered", func() {
			recipe := createRecipeData()
			codelist := createCodeList()
			codelist.HRef = "http://localhost:22400/code-lists/789"
			missingFields, invalidFields := codelist.validateCodelist(ctx, &recipe)
			So(missingFields, ShouldBeNil)
			So(invalidFields, ShouldBeNil)
		})

		Convey("when name and id fields do match", func() {
			recipe := createRecipeData()
			recipe.Format = "cantabular_flexible_table"
			codelist := createCodeList()
			codelist.HRef = "http://localhost:22400/code-lists/789"
			codelist.Name = "789"
			codelist.ID = "789"
			missingFields, invalidFields := codelist.validateCodelist(ctx, &recipe)
			So(missingFields, ShouldBeNil)
			So(invalidFields, ShouldBeNil)
		})

	})
}

func TestValidateCodelistHRef(t *testing.T) {
	t.Parallel()

	Convey("Unsuccessful validation (false) returned", t, func() {

		Convey("when codelist.href is unable to be parsed into url format", func() {
			codelist := createCodeList()
			codelist.HRef = ":foo"
			valid := codelist.validateCodelistHRef(ctx)
			So(valid, ShouldBeFalse)
		})

		Convey("when codelist.href does not contain the scheme", func() {
			codelist := createCodeList()
			codelist.HRef = "localhost:22400/code-lists/"
			valid := codelist.validateCodelistHRef(ctx)
			So(valid, ShouldBeFalse)
		})

		Convey("when codelist.href does not contain the host information", func() {
			codelist := createCodeList()
			codelist.HRef = "/code-lists/123"
			valid := codelist.validateCodelistHRef(ctx)
			So(valid, ShouldBeFalse)
		})

		Convey("when codelist.href does not contain 'code-lists' path", func() {
			codelist := createCodeList()
			codelist.HRef = "/123"
			valid := codelist.validateCodelistHRef(ctx)
			So(valid, ShouldBeFalse)
		})

		Convey("when codelist.href does not contain codelist.id", func() {
			codelist := createCodeList()
			codelist.ID = "123"
			codelist.HRef = "http://localhost:22400/code-lists/"
			valid := codelist.validateCodelistHRef(ctx)
			So(valid, ShouldBeFalse)
		})
	})

	Convey("Successful validation (true) returned", t, func() {

		Convey("when codelist.href contains its scheme, host, path and codelist.id in appropriate url format ", func() {
			codelist := createCodeList()
			codelist.ID = "123"
			codelist.HRef = "http://localhost:22400/code-lists/123"
			valid := codelist.validateCodelistHRef(ctx)
			So(valid, ShouldBeTrue)
		})
	})
}

func TestValidateAddRecipe(t *testing.T) {
	t.Parallel()

	Convey("Error returned with missing field", t, func() {

		Convey("when alias is missing", func() {
			recipe := createRecipeData()
			recipe.Alias = ""
			err := recipe.ValidateAddRecipe(ctx)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldResemble, errors.New("missing mandatory fields: [alias]").Error())
		})

		Convey("when format is missing", func() {
			recipe := createRecipeData()
			recipe.Format = ""
			err := recipe.ValidateAddRecipe(ctx)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldResemble, errors.New("missing mandatory fields: [format]").Error())
		})

		Convey("when input-files is missing", func() {
			recipe := createRecipeData()
			recipe.InputFiles = nil
			err := recipe.ValidateAddRecipe(ctx)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldResemble, errors.New("missing mandatory fields: [input-files]").Error())
		})

		Convey("when format is cantabular_blob and cantabular_blob field is missing", func() {
			recipe := createRecipeData()
			recipe.Format = cantabularBlob
			recipe.CantabularBlob = ""
			err := recipe.ValidateAddRecipe(ctx)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldResemble, errors.New("missing mandatory fields: [cantabular_blob]").Error())
		})

		Convey("when format is cantabular_table and cantabular_blob field is missing", func() {
			recipe := createRecipeData()
			recipe.Format = cantabularTable
			recipe.CantabularBlob = ""
			err := recipe.ValidateAddRecipe(ctx)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldResemble, errors.New("missing mandatory fields: [cantabular_table]").Error())
		})

		Convey("when format is cantabular_flexible_table and cantabular_blob field is missing", func() {
			recipe := createRecipeData()
			recipe.Format = cantabularFlexibleTable
			recipe.CantabularBlob = ""
			err := recipe.ValidateAddRecipe(ctx)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldResemble, errors.New("missing mandatory fields: [cantabular_flexible_table]").Error())
		})

		Convey("when format is cantabular_multivariate_table and cantabular_blob field is missing", func() {
			recipe := createRecipeData()
			recipe.Format = cantabularMultivariateTable
			recipe.CantabularBlob = ""
			err := recipe.ValidateAddRecipe(ctx)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldResemble, errors.New("missing mandatory fields: [cantabular_multivariate_table]").Error())
		})

		Convey("when input-files.description is missing", func() {
			recipe := createRecipeData()
			recipe.InputFiles[0].Description = ""
			err := recipe.ValidateAddRecipe(ctx)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldResemble, errors.New("missing mandatory fields: [input-files[0].description]").Error())

			recipe.InputFiles[0].Description = "test files" // Reset to original
		})

		Convey("when output-instances is missing", func() {
			recipe := createRecipeData()
			recipe.OutputInstances = nil
			err := recipe.ValidateAddRecipe(ctx)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldResemble, errors.New("missing mandatory fields: [output-instances]").Error())
		})

		Convey("when any field of output-instances is missing", func() {
			recipe := createRecipeData()
			recipe.OutputInstances[0].Title = ""
			err := recipe.ValidateAddRecipe(ctx)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldResemble, errors.New("missing mandatory fields: [output-instances[0].title]").Error())

			recipe.OutputInstances[0].Title = "test" // Reset to original
		})

	})

	Convey("Successful with no missing fields (nil error returned)", t, func() {

		Convey("when all fields of recipe are given", func() {
			recipe := createRecipeData()
			err := recipe.ValidateAddRecipe(ctx)
			So(err, ShouldBeNil)
		})

		Convey("when format is cantabular_blob and input-files is missing", func() {
			recipe := createRecipeData()
			recipe.InputFiles = nil
			recipe.Format = cantabularBlob
			err := recipe.ValidateAddRecipe(ctx)
			So(err, ShouldBeNil)
		})

		Convey("when format is cantabular_table and input-files is missing", func() {
			recipe := createRecipeData()
			recipe.InputFiles = nil
			recipe.Format = cantabularTable
			err := recipe.ValidateAddRecipe(ctx)
			So(err, ShouldBeNil)
		})

		Convey("when format is cantabular_flexible_table and input-files is missing", func() {
			recipe := createRecipeData()
			recipe.InputFiles = nil
			recipe.Format = cantabularFlexibleTable
			err := recipe.ValidateAddRecipe(ctx)
			So(err, ShouldBeNil)
		})

		Convey("when format is cantabular_multivariate_table and input-files is missing", func() {
			recipe := createRecipeData()
			recipe.InputFiles = nil
			recipe.Format = cantabularMultivariateTable
			err := recipe.ValidateAddRecipe(ctx)
			So(err, ShouldBeNil)
		})

	})

	Convey("Error returned with invalid field", t, func() {

		Convey("when output-instances.code-lists.href is incorrectly entered", func() {
			recipe := createRecipeData()
			recipe.OutputInstances[0].CodeLists[0].HRef = "incorrect-href"
			err := recipe.ValidateAddRecipe(ctx)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldResemble, errors.New("invalid fields: [output-instances[0].code-lists[0].href should be in format (URL/id)]").Error())

			recipe.OutputInstances[0].CodeLists[0].HRef = "http://localhost:22400/code-lists/789" // Reset to original
		})

		Convey("when format is not valid", func() {
			recipe := createRecipeData()
			recipe.Format = "v1"
			err := recipe.ValidateAddRecipe(ctx)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldResemble, errors.New("invalid fields: [format is not valid]").Error())
		})

	})

	Convey("Successful with no invalid fields (nil error returned)", t, func() {

		Convey("when v4 format is valid", func() {
			recipe := createRecipeData()
			recipe.Format = v4
			err := recipe.ValidateAddRecipe(ctx)
			So(err, ShouldBeNil)
		})

		Convey("when cantabular_blob format is valid", func() {
			recipe := createRecipeData()
			recipe.Format = cantabularBlob
			err := recipe.ValidateAddRecipe(ctx)
			So(err, ShouldBeNil)
		})

		Convey("when cantabular_table format is valid", func() {
			recipe := createRecipeData()
			recipe.Format = cantabularTable
			err := recipe.ValidateAddRecipe(ctx)
			So(err, ShouldBeNil)
		})

		Convey("when cantabular_flexible_table format is valid", func() {
			recipe := createRecipeData()
			recipe.Format = cantabularFlexibleTable
			err := recipe.ValidateAddRecipe(ctx)
			So(err, ShouldBeNil)
		})

		Convey("when cantabular_multivariate_table format is valid", func() {
			recipe := createRecipeData()
			recipe.Format = cantabularMultivariateTable
			err := recipe.ValidateAddRecipe(ctx)
			So(err, ShouldBeNil)
		})

		Convey("when output-instances.code-lists.href is correctly entered", func() {
			recipe := createRecipeData()
			recipe.OutputInstances[0].CodeLists[0].HRef = "http://localhost:22400/code-lists/789"
			err := recipe.ValidateAddRecipe(ctx)
			So(err, ShouldBeNil)
		})

	})

}

func TestValidateUpdateRecipe(t *testing.T) {
	t.Parallel()

	Convey("Error returned with missing field", t, func() {

		Convey("when any one field of output-instance update is missing", func() {
			recipe := Recipe{OutputInstances: []Instance{createInstance()}}
			recipe.OutputInstances[0].Title = ""
			err := recipe.ValidateUpdateRecipe(ctx)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldResemble, errors.New("missing mandatory fields: [output-instances[0].title]").Error())

			recipe.OutputInstances[0].Title = "test" // Reset
		})

		// test fix: non-existant instanceMissingFields[1] was assigned to, instead of [0] - causing panic
		Convey("when any one field of second output-instance update is missing", func() {
			recipe := Recipe{OutputInstances: []Instance{createInstance(), createInstance()}}
			recipe.OutputInstances[1].Title = ""
			err := recipe.ValidateUpdateRecipe(ctx)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldResemble, errors.New("missing mandatory fields: [output-instances[1].title]").Error())

			recipe.OutputInstances[1].Title = "test" // Reset
		})

	})

	Convey("Successful with no missing fields (nil error returned)", t, func() {

		Convey("when input-files.description is not missing", func() {
			recipe := Recipe{InputFiles: []file{{Description: "test files"}}}
			err := recipe.ValidateUpdateRecipe(ctx)
			So(err, ShouldBeNil)
		})

		Convey("when all fields of output-instance update is given", func() {
			recipe := Recipe{OutputInstances: []Instance{createInstance()}}
			err := recipe.ValidateUpdateRecipe(ctx)
			So(err, ShouldBeNil)
		})

	})

	Convey("Error returned with invalid field", t, func() {

		Convey("when no recipe fields updates is given", func() {
			recipe := Recipe{}
			err := recipe.ValidateUpdateRecipe(ctx)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldResemble, errors.New("invalid fields: [no recipe fields updates given]").Error())
		})

		Convey("when id update given when it should not be changed", func() {
			recipe := Recipe{ID: "123"}
			err := recipe.ValidateUpdateRecipe(ctx)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldResemble, errors.New("invalid fields: [id cannot be changed]").Error())
		})

		Convey("when format is not valid", func() {
			recipe := Recipe{Format: "v1"}
			err := recipe.ValidateUpdateRecipe(ctx)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldResemble, errors.New("invalid fields: [format is not valid]").Error())
		})

		Convey("when empty input-files is given", func() {
			recipe := Recipe{InputFiles: []file{}}
			err := recipe.ValidateUpdateRecipe(ctx)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldResemble, errors.New("invalid fields: [empty input-files update given]").Error())
		})

		Convey("when empty input-files.description is given", func() {
			recipe := Recipe{InputFiles: []file{{Description: ""}}}
			err := recipe.ValidateUpdateRecipe(ctx)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldResemble, errors.New("invalid fields: [empty input-files[0].description given]").Error())

			recipe.InputFiles = []file{{Description: "test files"}} // Reset
		})

		Convey("when empty output-instances is given", func() {
			recipe := Recipe{OutputInstances: []Instance{}}
			err := recipe.ValidateUpdateRecipe(ctx)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldResemble, errors.New("invalid fields: [empty output-instances update given]").Error())
		})

		Convey("when any field of output-instances update is invalid", func() {
			recipe := Recipe{OutputInstances: []Instance{createInstance()}}
			recipe.OutputInstances[0].CodeLists[0].HRef = "incorrect-href"
			err := recipe.ValidateUpdateRecipe(ctx)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldResemble, errors.New("invalid fields: [output-instances[0].code-lists[0].href should be in format (URL/id)]").Error())

			recipe.OutputInstances[0].CodeLists[0].HRef = "http://localhost:22400/code-lists/789" // Reset to original
		})

	})

	Convey("Successful with no invalid fields (nil error returned)", t, func() {

		Convey("when at least one recipe field update is given", func() {

			Convey("and id update is not given", func() {

				Convey("and format is valid", func() {
					recipe := Recipe{Format: v4}
					err := recipe.ValidateUpdateRecipe(ctx)
					So(err, ShouldBeNil)
				})

			})

		})

		Convey("when complete input-files is given", func() {
			recipe := Recipe{InputFiles: []file{{Description: "test files"}}}
			err := recipe.ValidateUpdateRecipe(ctx)
			So(err, ShouldBeNil)
		})

		Convey("when complete output-instances is given", func() {
			recipe := Recipe{OutputInstances: []Instance{createInstance()}}
			err := recipe.ValidateUpdateRecipe(ctx)
			So(err, ShouldBeNil)
		})

		Convey("when any field of output-instances update is valid", func() {
			recipe := Recipe{OutputInstances: []Instance{createInstance()}}
			recipe.OutputInstances[0].CodeLists[0].HRef = "http://localhost:22400/code-lists/789" // Reset
			err := recipe.ValidateUpdateRecipe(ctx)
			So(err, ShouldBeNil)
		})

	})

}

func TestValidateAddInstance(t *testing.T) {
	t.Parallel()

	Convey("Non-empty missing field successfully returned", t, func() {

		Convey("when datasetID is missing", func() {
			recipe := createRecipeData()
			instance := createInstance()
			instance.DatasetID = ""
			err := instance.ValidateAddInstance(ctx, &recipe)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldResemble, errors.New("missing mandatory fields: [dataset_id]").Error())
		})

		Convey("when editions is missing", func() {
			recipe := createRecipeData()
			instance := createInstance()
			instance.Editions = nil
			err := instance.ValidateAddInstance(ctx, &recipe)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldResemble, errors.New("missing mandatory fields: [editions]").Error())
		})

		Convey("when an edition of editions is missing", func() {
			recipe := createRecipeData()
			instance := createInstance()
			instance.Editions = []string{""}
			err := instance.ValidateAddInstance(ctx, &recipe)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldResemble, errors.New("missing mandatory fields: [editions[0]]").Error())

			instance.Editions = []string{"editions"} // Reset to original
		})

		Convey("when title is missing", func() {
			recipe := createRecipeData()
			instance := createInstance()
			instance.Title = ""
			err := instance.ValidateAddInstance(ctx, &recipe)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldResemble, errors.New("missing mandatory fields: [title]").Error())
		})

		Convey("when code lists is missing", func() {
			recipe := createRecipeData()
			instance := createInstance()
			instance.CodeLists = nil
			err := instance.ValidateAddInstance(ctx, &recipe)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldResemble, errors.New("missing mandatory fields: [code-lists]").Error())
		})

		Convey("when any field of code lists is missing", func() {
			recipe := createRecipeData()
			instance := createInstance()
			instance.CodeLists[0].Name = ""
			err := instance.ValidateAddInstance(ctx, &recipe)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldResemble, errors.New("missing mandatory fields: [code-lists[0].name]").Error())

			instance.CodeLists[0].Name = "codelist-test" // Reset to original
		})

	})

	Convey("Empty missing field successfully returned", t, func() {

		Convey("when all fields are output-instances are given", func() {
			recipe := createRecipeData()
			instance := createInstance()
			err := instance.ValidateAddInstance(ctx, &recipe)
			So(err, ShouldBeNil)
		})

	})

	Convey("Non-empty invalid field successfully returned", t, func() {

		Convey("when code-lists.href is incorrectly entered", func() {
			recipe := createRecipeData()
			instance := createInstance()
			instance.CodeLists[0].HRef = "incorrect-href"
			err := instance.ValidateAddInstance(ctx, &recipe)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldResemble, errors.New("invalid fields: [code-lists[0].href should be in format (URL/id)]").Error())

			instance.CodeLists[0].HRef = "http://localhost:22400/code-lists/789" // Reset
		})

	})

	Convey("Empty invalid field successfully returned", t, func() {

		Convey("when code-lists.href is correctly entered", func() {
			recipe := createRecipeData()
			instance := createInstance()
			instance.CodeLists[0].HRef = "http://localhost:22400/code-lists/789"
			err := instance.ValidateAddInstance(ctx, &recipe)
			So(err, ShouldBeNil)
		})

	})

}

func TestValidateUpdateInstance(t *testing.T) {
	t.Parallel()

	Convey("Error returned with missing field", t, func() {

		Convey("when empty editions update is given", func() {
			recipe := createRecipeData()
			instance := Instance{Editions: []string{}}
			err := instance.ValidateUpdateInstance(ctx, &recipe)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldResemble, errors.New("missing mandatory fields: [editions]").Error())

		})

		Convey("when any edition update of editions is incomplete", func() {
			recipe := createRecipeData()
			instance := Instance{Editions: []string{""}}
			err := instance.ValidateUpdateInstance(ctx, &recipe)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldResemble, errors.New("missing mandatory fields: [editions[0]]").Error())

			instance.Editions = []string{"editions"} // Reset
		})

		Convey("when any code lists fields is missing", func() {
			recipe := createRecipeData()
			instance := Instance{CodeLists: []CodeList{createCodeList()}}
			instance.CodeLists[0].Name = ""
			err := instance.ValidateUpdateInstance(ctx, &recipe)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldResemble, errors.New("missing mandatory fields: [code-lists[0].name]").Error())

			instance.CodeLists[0].Name = "codelist-test" // Reset
		})

		// test fix: non-existant codelistMissingFields[1] was assigned to, instead of [0] - causing panic
		Convey("when there are two code lists and second has error", func() {
			recipe := createRecipeData()
			instance := Instance{CodeLists: []CodeList{createCodeList(), createCodeList()}}
			instance.CodeLists[1].Name = ""
			err := instance.ValidateUpdateInstance(ctx, &recipe)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldResemble, errors.New("missing mandatory fields: [code-lists[1].name]").Error())

			instance.CodeLists[1].Name = "codelist-test" // Reset
		})

	})

	Convey("Successful with no missing fields (nil error returned)", t, func() {

		Convey("when complete editions update is given", func() {
			recipe := createRecipeData()
			instance := Instance{Editions: []string{"editions"}}
			err := instance.ValidateUpdateInstance(ctx, &recipe)
			So(err, ShouldBeNil)
		})

		Convey("when all code lists fields are not missing", func() {
			recipe := createRecipeData()
			instance := Instance{CodeLists: []CodeList{createCodeList()}}
			err := instance.ValidateUpdateInstance(ctx, &recipe)
			So(err, ShouldBeNil)
		})

	})

	Convey("Error returned with invalid field", t, func() {

		Convey("when no instance fields updates is given", func() {
			recipe := createRecipeData()
			instance := Instance{}
			err := instance.ValidateUpdateInstance(ctx, &recipe)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldResemble, errors.New("invalid fields: [no instance fields updates given]").Error())
		})

		Convey("when any code lists fields is invalid", func() {
			recipe := createRecipeData()
			instance := Instance{CodeLists: []CodeList{createCodeList()}}
			instance.CodeLists[0].HRef = "incorrect-href"
			err := instance.ValidateUpdateInstance(ctx, &recipe)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldResemble, errors.New("invalid fields: [code-lists[0].href should be in format (URL/id)]").Error())

			instance.CodeLists[0].HRef = "http://localhost:22400/code-lists/789" // Reset
		})

	})

	Convey("Successful with no invalid field (nil error returned)", t, func() {

		Convey("when at least one instance field update is given", func() {

			Convey("and dataset-id is not given", func() {

				Convey("and any code lists fields is valid", func() {
					// createCodeList() is a complete and valid codelist
					// instance just updating code lists of instance
					recipe := createRecipeData()
					instance := Instance{CodeLists: []CodeList{createCodeList()}}
					err := instance.ValidateUpdateInstance(ctx, &recipe)
					So(err, ShouldBeNil)
				})

			})

		})

	})

}

func TestValidateAddCodelists(t *testing.T) {
	t.Parallel()

	Convey("Non-empty missing field successfully returned", t, func() {

		Convey("when id is missing", func() {
			recipe := createRecipeData()
			codelist := createCodeList()
			codelist.ID = ""
			// HRef Updated as the format of HRef follows the value from ID
			codelist.HRef = "http://localhost:22400/code-lists/"
			err := codelist.ValidateAddCodelist(ctx, &recipe)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldResemble, errors.New("missing mandatory fields: [id]").Error())
		})

		Convey("when href is missing", func() {
			recipe := createRecipeData()
			codelist := createCodeList()
			codelist.HRef = ""
			err := codelist.ValidateAddCodelist(ctx, &recipe)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldResemble, errors.New("missing mandatory fields: [href]").Error())
		})

		Convey("when name is missing", func() {
			recipe := createRecipeData()
			codelist := createCodeList()
			codelist.Name = ""
			err := codelist.ValidateAddCodelist(ctx, &recipe)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldResemble, errors.New("missing mandatory fields: [name]").Error())
		})

		Convey("when ishierarchy is missing", func() {
			recipe := createRecipeData()
			codelist := createCodeList()
			codelist.IsHierarchy = nil
			err := codelist.ValidateAddCodelist(ctx, &recipe)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldResemble, errors.New("missing mandatory fields: [isHierarchy]").Error())
		})

		Convey("when isCantabularGeography is missing", func() {
			recipe := createRecipeData()
			codelist := createCodeList()
			codelist.IsCantabularGeography = nil
			recipe.Format = "cantabular_flexible_table"
			err := codelist.ValidateAddCodelist(ctx, &recipe)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldResemble, errors.New("missing mandatory fields: [isCantabularGeography]").Error())
		})

		Convey("when isCantabularDefaultGeography is missing", func() {
			recipe := createRecipeData()
			codelist := createCodeList()
			codelist.IsCantabularDefaultGeography = nil
			recipe.Format = "cantabular_flexible_table"
			err := codelist.ValidateAddCodelist(ctx, &recipe)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldResemble, errors.New("missing mandatory fields: [isCantabularDefaultGeography]").Error())
		})

	})

	Convey("Empty missing field successfully returned", t, func() {

		Convey("when all fields are codelist are given", func() {
			recipe := createRecipeData()
			codelist := createCodeList()
			err := codelist.ValidateAddCodelist(ctx, &recipe)
			So(err, ShouldBeNil)
		})

	})

	Convey("Non-empty invalid field successfully returned", t, func() {

		Convey("when href is incorrectly entered", func() {
			recipe := createRecipeData()
			codelist := createCodeList()
			codelist.HRef = "incorrect-href"
			err := codelist.ValidateAddCodelist(ctx, &recipe)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldResemble, errors.New("invalid fields: [href should be in format (URL/id)]").Error())
		})

	})

	Convey("Empty invalid field successfully returned", t, func() {

		Convey("when href is correctly entered", func() {
			recipe := createRecipeData()
			codelist := createCodeList()
			codelist.HRef = "http://localhost:22400/code-lists/789"
			err := codelist.ValidateAddCodelist(ctx, &recipe)
			So(err, ShouldBeNil)
		})

	})

}

func TestValidateUpdateCodelist(t *testing.T) {
	t.Parallel()

	Convey("Error returned with invalid field", t, func() {

		Convey("when no codelist fields updates are given", func() {
			codelist := CodeList{}
			err := codelist.ValidateUpdateCodeList(ctx)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldResemble, errors.New("invalid fields: [no codelist fields updates given]").Error())
		})

		Convey("when href update is given", func() {
			codelist := CodeList{HRef: "incorrect-href"}
			err := codelist.ValidateUpdateCodeList(ctx)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldResemble, errors.New("invalid fields: [href should be in format (URL/id)]").Error())
		})

	})

	Convey("Successful with no invalid field (nil error returned)", t, func() {
		Convey("when at least one codelist field is given", func() {
			Convey("when id update is not given", func() {
				Convey("when href update is not given", func() {
					codelist := CodeList{Name: "test"}
					err := codelist.ValidateUpdateCodeList(ctx)
					So(err, ShouldBeNil)
				})
			})
		})
	})
}
