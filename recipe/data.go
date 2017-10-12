package recipe

//To add a new recipe:
// - add a recipeResponse{} document with an appropriate var name
// - add the var name to the list of Items
// - increment the Count and Total count in the RecipeList

//FullLit of recipes available via this API
var FullList = List{
	Items:        []Response{CPI},
	Count:        1,
	TotalCount:   1,
	ItemsPerPage: 10,
	Start:        0,
}

//CPI recipe for transforming a given input to a CPI COICOP dataset
var CPI = Response{
	ID:     "b944be78-f56d-409b-9ebd-ab2b77ffe187",
	Alias:  "CPI COICOP",
	Format: "v4",
	InputFiles: []file{
		{"CPI COICOP v4"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "931a8a2a-0dc8-42b6-a884-7b6054ed3b68",
			Editions:  []string{"Time-series"},
			CodeLists: []CodeList{{ID: "64d384f1-ea3b-445c-8fb8-aa453f96e58a", Name: "time", HRef: "http://localhost:22400/code-lists/64d384f1-ea3b-445c-8fb8-aa453f96e58a"},
				{ID: "65107A9F-7DA3-4B41-A410-6F6D9FBD68C3", Name: "geography", HRef: "http://localhost:22400/code-lists/65107A9F-7DA3-4B41-A410-6F6D9FBD68C3"},
				{ID: "e44de4c4-d39e-4e2f-942b-3ca10584d078", Name: "aggregate", HRef: "http://localhost:22400/code-lists/e44de4c4-d39e-4e2f-942b-3ca10584d078"}},
		},
	},
}
