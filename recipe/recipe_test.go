package recipe

import (
	"errors"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func createCodeList() CodeList {
	return CodeList{
		ID:          "789",
		Name:        "codelist-test",
		HRef:        "http://localhost:22400/code-lists/789",
		IsHierarchy: falseValPtr,
	}
}

func createRecipeData() Response {
	return Response{
		ID:              "123",
		Alias:           "test",
		Format:          "v4",
		InputFiles:      []file{{Description: "test files"}},
		OutputInstances: []Instance{createInstance()},
	}
}

func createInstance() Instance {
	return Instance{
		DatasetID: "456",
		Editions:  []string{"editions"},
		Title:     "test",
		CodeLists: []CodeList{createCodeList()},
	}
}

//ValidateAddInstance uses validateInstance so this test covers the former function as well
func TestValidateInstance(t *testing.T) {
	t.Parallel()

	Convey("Non-empty missing field successfully returned", t, func() {

		Convey("when datasetID is missing", func() {
			instance := createInstance()
			instance.DatasetID = ""
			missingFields, invalidFields := instance.validateInstance()
			So(missingFields, ShouldNotBeNil)
			So(missingFields, ShouldResemble, []string{"dataset-id"})
			So(invalidFields, ShouldBeNil)
		})

		Convey("when editions is missing", func() {
			instance := createInstance()
			instance.Editions = nil
			missingFields, invalidFields := instance.validateInstance()
			So(missingFields, ShouldNotBeNil)
			So(missingFields, ShouldResemble, []string{"editions"})
			So(invalidFields, ShouldBeNil)
		})

		Convey("when an edition of editions is missing", func() {
			instance := createInstance()
			instance.Editions = []string{""}
			missingFields, invalidFields := instance.validateInstance()
			So(missingFields, ShouldNotBeNil)
			So(missingFields, ShouldResemble, []string{"editions[0]"})
			So(invalidFields, ShouldBeNil)

			instance.Editions = []string{"editions"} //Reset to original
		})

		Convey("when title is missing", func() {
			instance := createInstance()
			instance.Title = ""
			missingFields, invalidFields := instance.validateInstance()
			So(missingFields, ShouldNotBeNil)
			So(missingFields, ShouldResemble, []string{"title"})
			So(invalidFields, ShouldBeNil)
		})

		Convey("when codelists is missing", func() {
			instance := createInstance()
			instance.CodeLists = nil
			missingFields, invalidFields := instance.validateInstance()
			So(missingFields, ShouldNotBeNil)
			So(missingFields, ShouldResemble, []string{"codelists"})
			So(invalidFields, ShouldBeNil)
		})

		Convey("when any field of codelists is missing", func() {
			instance := createInstance()
			instance.CodeLists[0].Name = ""
			missingFields, invalidFields := instance.validateInstance()
			So(missingFields, ShouldNotBeNil)
			So(missingFields, ShouldResemble, []string{"codelists[0].name"})
			So(invalidFields, ShouldBeNil)

			instance.CodeLists[0].Name = "codelist-test" //Reset to original
		})

	})

	Convey("Empty missing field successfully returned", t, func() {

		Convey("when all fields are output-instances are given", func() {
			instance := createInstance()
			missingFields, invalidFields := instance.validateInstance()
			So(missingFields, ShouldBeNil)
			So(invalidFields, ShouldBeNil)
		})

	})

	Convey("Non-empty invalid field successfully returned", t, func() {

		Convey("when codelists.href is incorrectly entered", func() {
			instance := createInstance()
			instance.CodeLists[0].HRef = "incorrect-href"
			missingFields, invalidFields := instance.validateInstance()
			So(missingFields, ShouldBeNil)
			So(invalidFields, ShouldNotBeNil)
			So(invalidFields, ShouldResemble, []string{"codelists[0].href should be in format (URL/id)"})

			instance.CodeLists[0].HRef = "http://localhost:22400/code-lists/789" //Reset
		})

	})

	Convey("Empty invalid field successfully returned", t, func() {

		Convey("when codelists.href is correctly entered", func() {
			instance := createInstance()
			instance.CodeLists[0].HRef = "http://localhost:22400/code-lists/789"
			missingFields, invalidFields := instance.validateInstance()
			So(missingFields, ShouldBeNil)
			So(invalidFields, ShouldBeNil)
		})

	})

}

//ValidateAddCodelist uses validateCodelist so this test covers the former function as well
func TestValidateCodelists(t *testing.T) {
	t.Parallel()

	Convey("Non-empty missing field successfully returned", t, func() {

		Convey("when id is missing", func() {
			codelist := createCodeList()
			codelist.ID = ""
			//HRef Updated as the format of HRef follows the value from ID
			codelist.HRef = "http://localhost:22400/code-lists/"
			missingFields, invalidFields := codelist.validateCodelists()
			So(missingFields, ShouldNotBeNil)
			So(missingFields, ShouldResemble, []string{"id"})
			So(invalidFields, ShouldBeNil)
		})

		Convey("when href is missing", func() {
			codelist := createCodeList()
			codelist.HRef = ""
			missingFields, invalidFields := codelist.validateCodelists()
			So(missingFields, ShouldNotBeNil)
			So(missingFields, ShouldResemble, []string{"href"})
			So(invalidFields, ShouldBeNil)
		})

		Convey("when name is missing", func() {
			codelist := createCodeList()
			codelist.Name = ""
			missingFields, invalidFields := codelist.validateCodelists()
			So(missingFields, ShouldNotBeNil)
			So(missingFields, ShouldResemble, []string{"name"})
			So(invalidFields, ShouldBeNil)
		})

		Convey("when ishierarchy is missing", func() {
			codelist := createCodeList()
			codelist.IsHierarchy = nil
			missingFields, invalidFields := codelist.validateCodelists()
			So(missingFields, ShouldNotBeNil)
			So(missingFields, ShouldResemble, []string{"isHierarchy"})
			So(invalidFields, ShouldBeNil)
		})

	})

	Convey("Empty missing field successfully returned", t, func() {

		Convey("when all fields are codelist are given", func() {
			codelist := createCodeList()
			missingFields, invalidFields := codelist.validateCodelists()
			So(missingFields, ShouldBeNil)
			So(invalidFields, ShouldBeNil)
		})

	})

	Convey("Non-empty invalid field successfully returned", t, func() {

		Convey("when href is incorrectly entered", func() {
			codelist := createCodeList()
			codelist.HRef = "incorrect-href"
			missingFields, invalidFields := codelist.validateCodelists()
			So(missingFields, ShouldBeNil)
			So(invalidFields, ShouldNotBeNil)
			So(invalidFields, ShouldResemble, []string{"href should be in format (URL/id)"})
		})

	})

	Convey("Empty invalid field successfully returned", t, func() {

		Convey("when href is correctly entered", func() {
			codelist := createCodeList()
			codelist.HRef = "http://localhost:22400/code-lists/789"
			missingFields, invalidFields := codelist.validateCodelists()
			So(missingFields, ShouldBeNil)
			So(invalidFields, ShouldBeNil)
		})

	})

}

func TestValidateAddRecipe(t *testing.T) {
	t.Parallel()

	Convey("Error returned with missing field", t, func() {

		Convey("when alias is missing", func() {
			recipe := createRecipeData()
			recipe.Alias = ""
			err := recipe.ValidateAddRecipe()
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldResemble, errors.New("missing mandatory fields: [alias]").Error())
		})

		Convey("when format is missing", func() {
			recipe := createRecipeData()
			recipe.Format = ""
			err := recipe.ValidateAddRecipe()
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldResemble, errors.New("missing mandatory fields: [format]").Error())
		})

		Convey("when input-files is missing", func() {
			recipe := createRecipeData()
			recipe.InputFiles = nil
			err := recipe.ValidateAddRecipe()
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldResemble, errors.New("missing mandatory fields: [input-files]").Error())
		})

		Convey("when input-files.description is missing", func() {
			recipe := createRecipeData()
			recipe.InputFiles[0].Description = ""
			err := recipe.ValidateAddRecipe()
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldResemble, errors.New("missing mandatory fields: [input-files[0].description]").Error())

			recipe.InputFiles[0].Description = "test files" //Reset to original
		})

		Convey("when output-instances is missing", func() {
			recipe := createRecipeData()
			recipe.OutputInstances = nil
			err := recipe.ValidateAddRecipe()
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldResemble, errors.New("missing mandatory fields: [output-instances]").Error())
		})

		Convey("when any field of output-instances is missing", func() {
			recipe := createRecipeData()
			recipe.OutputInstances[0].Title = ""
			err := recipe.ValidateAddRecipe()
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldResemble, errors.New("missing mandatory fields: [output-instances[0].title]").Error())

			recipe.OutputInstances[0].Title = "test" //Reset to original
		})

	})

	Convey("Successful with no missing fields (nil error returned)", t, func() {

		Convey("when all fields of recipe are given", func() {
			recipe := createRecipeData()
			err := recipe.ValidateAddRecipe()
			So(err, ShouldBeNil)
		})

	})

	Convey("Error returned with invalid field", t, func() {

		Convey("when output-instances.codelists.href is incorrectly entered", func() {
			recipe := createRecipeData()
			recipe.OutputInstances[0].CodeLists[0].HRef = "incorrect-href"
			err := recipe.ValidateAddRecipe()
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldResemble, errors.New("invalid fields: [output-instances[0].codelists[0].href should be in format (URL/id)]").Error())

			recipe.OutputInstances[0].CodeLists[0].HRef = "http://localhost:22400/code-lists/789" //Reset to original
		})

		Convey("when format is not valid", func() {
			recipe := createRecipeData()
			recipe.Format = "v1"
			err := recipe.ValidateAddRecipe()
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldResemble, errors.New("invalid fields: [format is not valid]").Error())
		})

	})

	Convey("Successful with no invalid fields (nil error returned)", t, func() {

		Convey("when format is valid", func() {
			recipe := createRecipeData()
			recipe.Format = "v4"
			err := recipe.ValidateAddRecipe()
			So(err, ShouldBeNil)
		})

		Convey("when output-instances.codelists.href is correctly entered", func() {
			recipe := createRecipeData()
			recipe.OutputInstances[0].CodeLists[0].HRef = "http://localhost:22400/code-lists/789"
			err := recipe.ValidateAddRecipe()
			So(err, ShouldBeNil)
		})

	})

}

func TestValidateUpdateRecipe(t *testing.T) {
	t.Parallel()

	Convey("Error returned with missing field", t, func() {

		Convey("when input-files.description is missing", func() {
			recipe := Response{InputFiles: []file{{Description: ""}}}
			err := recipe.ValidateUpdateRecipe()
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldResemble, errors.New("missing mandatory fields: [input-files[0].description]").Error())

			recipe.InputFiles = []file{{Description: "test files"}} //Reset
		})

		Convey("when any one field of output-instance update is missing", func() {
			recipe := Response{OutputInstances: []Instance{createInstance()}}
			recipe.OutputInstances[0].Title = ""
			err := recipe.ValidateUpdateRecipe()
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldResemble, errors.New("missing mandatory fields: [output-instances[0].title]").Error())

			recipe.OutputInstances[0].Title = "test" //Reset
		})

	})

	Convey("Successful with no missing fields (nil error returned)", t, func() {

		Convey("when input-files.description is not missing", func() {
			recipe := Response{InputFiles: []file{{Description: "test files"}}}
			err := recipe.ValidateUpdateRecipe()
			So(err, ShouldBeNil)
		})

		Convey("when all fields of output-instance update is given", func() {
			recipe := Response{OutputInstances: []Instance{createInstance()}}
			err := recipe.ValidateUpdateRecipe()
			So(err, ShouldBeNil)
		})

	})

	Convey("Error returned with invalid field", t, func() {

		Convey("when no recipe fields updates is given", func() {
			recipe := Response{}
			err := recipe.ValidateUpdateRecipe()
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldResemble, errors.New("invalid fields: [no recipe fields updates given]").Error())
		})

		Convey("when id update given when it should not be changed", func() {
			recipe := Response{ID: "123"}
			err := recipe.ValidateUpdateRecipe()
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldResemble, errors.New("invalid fields: [id cannot be changed]").Error())
		})

		Convey("when format is not valid", func() {
			recipe := Response{Format: "v1"}
			err := recipe.ValidateUpdateRecipe()
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldResemble, errors.New("invalid fields: [format is not valid]").Error())
		})

		Convey("when empty input-files is given", func() {
			recipe := Response{InputFiles: []file{}}
			err := recipe.ValidateUpdateRecipe()
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldResemble, errors.New("invalid fields: [empty input-files update given]").Error())
		})

		Convey("when empty output-instances is given", func() {
			recipe := Response{OutputInstances: []Instance{}}
			err := recipe.ValidateUpdateRecipe()
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldResemble, errors.New("invalid fields: [empty output-instances update given]").Error())
		})

		Convey("when any field of output-instances update is invalid", func() {
			recipe := Response{OutputInstances: []Instance{createInstance()}}
			recipe.OutputInstances[0].CodeLists[0].HRef = "incorrect-href"
			err := recipe.ValidateUpdateRecipe()
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldResemble, errors.New("invalid fields: [output-instances[0].codelists[0].href should be in format (URL/id)]").Error())

			recipe.OutputInstances[0].CodeLists[0].HRef = "http://localhost:22400/code-lists/789" //Reset to original
		})

	})

	Convey("Successful with no invalid fields (nil error returned)", t, func() {

		Convey("when at least one recipe field update is given", func() {

			Convey("and id update is not given", func() {

				Convey("and format is valid", func() {
					recipe := Response{Format: "v4"}
					err := recipe.ValidateUpdateRecipe()
					So(err, ShouldBeNil)
				})

			})

		})

		Convey("when complete input-files is given", func() {
			recipe := Response{InputFiles: []file{{Description: "test files"}}}
			err := recipe.ValidateUpdateRecipe()
			So(err, ShouldBeNil)
		})

		Convey("when complete output-instances is given", func() {
			recipe := Response{OutputInstances: []Instance{createInstance()}}
			err := recipe.ValidateUpdateRecipe()
			So(err, ShouldBeNil)
		})

		Convey("when any field of output-instances update is valid", func() {
			recipe := Response{OutputInstances: []Instance{createInstance()}}
			recipe.OutputInstances[0].CodeLists[0].HRef = "http://localhost:22400/code-lists/789" //Reset
			err := recipe.ValidateUpdateRecipe()
			So(err, ShouldBeNil)
		})

	})

}

func TestValidateUpdateInstance(t *testing.T) {
	t.Parallel()

	Convey("Error returned with missing field", t, func() {

		Convey("when empty editions update is given", func() {
			instance := Instance{Editions: []string{}}
			err := instance.ValidateUpdateInstance()
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldResemble, errors.New("missing mandatory fields: [editions]").Error())

		})

		Convey("when any edition update of editions is incomplete", func() {
			instance := Instance{Editions: []string{""}}
			err := instance.ValidateUpdateInstance()
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldResemble, errors.New("missing mandatory fields: [editions[0]]").Error())

			instance.Editions = []string{"editions"} //Reset
		})

		Convey("when any codelists fields is missing", func() {
			instance := Instance{CodeLists: []CodeList{createCodeList()}}
			instance.CodeLists[0].Name = ""
			err := instance.ValidateUpdateInstance()
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldResemble, errors.New("missing mandatory fields: [codelists[0].name]").Error())

			instance.CodeLists[0].Name = "codelist-test" //Reset
		})

	})

	Convey("Successful with no missing fields (nil error returned)", t, func() {

		Convey("when complete editions update is given", func() {
			instance := Instance{Editions: []string{"editions"}}
			err := instance.ValidateUpdateInstance()
			So(err, ShouldBeNil)
		})

		Convey("when all codelists fields are not missing", func() {
			instance := Instance{CodeLists: []CodeList{createCodeList()}}
			err := instance.ValidateUpdateInstance()
			So(err, ShouldBeNil)
		})

	})

	Convey("Error returned with invalid field", t, func() {

		Convey("when no instance fields updates is given", func() {
			instance := Instance{}
			err := instance.ValidateUpdateInstance()
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldResemble, errors.New("invalid fields: [no instance fields updates given]").Error())
		})

		Convey("when dataset-id is given to be updated", func() {
			instance := Instance{DatasetID: "123"}
			err := instance.ValidateUpdateInstance()
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldResemble, errors.New("invalid fields: [dataset-id cannot be changed]").Error())
		})

		Convey("when any codelists fields is invalid", func() {
			instance := Instance{CodeLists: []CodeList{createCodeList()}}
			instance.CodeLists[0].HRef = "incorrect-href"
			err := instance.ValidateUpdateInstance()
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldResemble, errors.New("invalid fields: [codelists[0].href should be in format (URL/id)]").Error())

			instance.CodeLists[0].HRef = "http://localhost:22400/code-lists/789" //Reset
		})

	})

	Convey("Successful with no invalid field (nil error returned)", t, func() {

		Convey("when at least one instance field update is given", func() {

			Convey("and dataset-id is not given", func() {

				Convey("and any codelists fields is valid", func() {
					//createCodeList() is a complete and valid codelist
					//instance just updating codelists of instance
					instance := Instance{CodeLists: []CodeList{createCodeList()}}
					err := instance.ValidateUpdateInstance()
					So(err, ShouldBeNil)
				})

			})

		})

	})

}

func TestValidateUpdateCodelist(t *testing.T) {
	t.Parallel()

	Convey("Error returned with invalid field", t, func() {

		Convey("when no codelist fields updates are given", func() {
			codelist := CodeList{}
			err := codelist.ValidateUpdateCodeList()
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldResemble, errors.New("invalid fields: [no codelist fields updates given]").Error())
		})

		Convey("when id update is given", func() {
			codelist := CodeList{ID: "789"}
			err := codelist.ValidateUpdateCodeList()
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldResemble, errors.New("invalid fields: [id cannot be changed]").Error())
		})

		Convey("when href update is given", func() {
			codelist := CodeList{HRef: "incorrect-href"}
			err := codelist.ValidateUpdateCodeList()
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldResemble, errors.New("invalid fields: [href cannot be changed as linked to id]").Error())
		})

	})

	Convey("Successful with no invalid field (nil error returned)", t, func() {

		Convey("when at least one codelist field is given", func() {

			Convey("when id update is not given", func() {

				Convey("when href update is not given", func() {
					codelist := CodeList{Name: "test"}
					err := codelist.ValidateUpdateCodeList()
					So(err, ShouldBeNil)
				})

			})

		})

	})

}
