package recipe

//To add a new recipe:
// - add a recipeResponse{} document with an appropriate var name
// - add the var name to the list of Items
// - increment the Count and Total count in the RecipeList

//FullList of recipes available via this API
var FullList = List{
	Items:        []Response{CPI, SAPE},
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
			Title:     "UK consumer price inflation",
			CodeLists: []CodeList{{ID: "64d384f1-ea3b-445c-8fb8-aa453f96e58a", Name: "time", HRef: "http://localhost:22400/code-lists/64d384f1-ea3b-445c-8fb8-aa453f96e58a"},
				{ID: "65107A9F-7DA3-4B41-A410-6F6D9FBD68C3", Name: "geography", HRef: "http://localhost:22400/code-lists/65107A9F-7DA3-4B41-A410-6F6D9FBD68C3"},
				{ID: "e44de4c4-d39e-4e2f-942b-3ca10584d078", Name: "aggregate", HRef: "http://localhost:22400/code-lists/e44de4c4-d39e-4e2f-942b-3ca10584d078"}},
		},
	},
}

var SAPE = Response{
	ID:     "8ef16b08-fef5-4bd3-8300-8066b0c777ce",
	Alias:  "sape",
	Format: "v4",
	InputFiles: []file{
		{"sape"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "4e1a21a9-3fb9-4a71-b2ad-9be29457236b",
			Editions:  []string{"Time-series"},
			CodeLists: []CodeList{{ID: "64d384f1-ea3b-445c-8fb8-aa453f96e58a", Name: "time", HRef: "http://localhost:22400/code-lists/64d384f1-ea3b-445c-8fb8-aa453f96e58a"},
				{ID: "65107A9f-7da3-4b41-a410-6f6d9fbd68c3", Name: "geography", HRef: "http://localhost:22400/code-lists/65107A9f-7da3-4b41-a410-6f6d9fbd68c3"},
				{ID: "4e1a21a9-3fb9-4a71-b2ad-9be29457236b", Name: "age", HRef: "http://localhost:22400/code-lists/4e1a21a9-3fb9-4a71-b2ad-9be29457236b"},
				{ID: "54ff5089-ea78-45ef-afa2-0dfe58f89497", Name: "sex", HRef: "http://localhost:22400/code-lists/54ff5089-ea78-45ef-afa2-0dfe58f89497"}},
		},
	},
}
