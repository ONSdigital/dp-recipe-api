package recipe

// 1961 Census Data
var Census1961 = Response{
	ID:     "4fff878f-f642-4113-9bed-85ae19a19ee7",
	Alias:  "1961 Census",
	Format: "v4",
	InputFiles: []file{
		{"1961Census"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "census-1961",
			Editions:  []string{"time-series"},
			Title:     "1961 Census Data",
			CodeLists: []CodeList{
				{
					ID:          "calendar-years",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/calendar-years",
					IsHierarchy: false,
				}, {
					ID:          "1961-geography",
					Name:        "1961geography",
					HRef:        "http://localhost:22400/code-lists/1961-geography",
					IsHierarchy: true,
				}, {
					ID:          "ashe-sex",
					Name:        "sex",
					HRef:        "http://localhost:22400/code-lists/ashe-sex",
					IsHierarchy: false,
				}, {
					ID:          "census-1961-marital-status",
					Name:        "maritalstatus",
					HRef:        "http://localhost:22400/code-lists/census-1961-marital-status",
					IsHierarchy: false,
				}, {
					ID:          "census-1961-age-groups",
					Name:        "agegroup",
					HRef:        "http://localhost:22400/code-lists/census-1961-age-groups",
					IsHierarchy: false,
				},
			},
		},
	},
}

//recipe for the SH01 table for the 1961 Census
var Census1961SH01 = Response{
	ID:     "0cbcf92b-4b72-4ecf-bac0-4d9227a8ce12",
	Alias:  "1961 Census: SH01",
	Format: "v4",
	InputFiles: []file{
		{"Census1961SH01"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "census-1961-sh01",
			Editions:  []string{"time-series"},
			Title:     "1961 Census: SH01",
			CodeLists: []CodeList{
				{
					ID:          "calendar-years",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/calendar-years",
					IsHierarchy: false,
				}, {
					ID:          "1961-geography",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/1961-geography",
					IsHierarchy: true,
				}, {
					ID:          "census-1961-tenure",
					Name:        "tenure",
					HRef:        "http://localhost:22400/code-lists/census-1961-tenure",
					IsHierarchy: false,
				}, {
					ID:          "census-1961-tenure-variable",
					Name:        "tenurevariable",
					HRef:        "http://localhost:22400/code-lists/ccensus-1961-tenure-variable",
					IsHierarchy: false,
				},
			},
		},
	},
}

//recipe for the SH07 table for the 1961 Census
var Census1961SH07 = Response{
	ID:     "c1960f40-b988-4875-8b85-05878afa8e9c",
	Alias:  "1961 Census: Old Persons Alone",
	Format: "v4",
	InputFiles: []file{
		{"Census1961SH07"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "census-1961-old-persons-alone",
			Editions:  []string{"time-series"},
			Title:     "1961 Census: Old Persons Alone",
			CodeLists: []CodeList{
				{
					ID:          "calendar-years",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/calendar-years",
					IsHierarchy: false,
				}, {
					ID:          "1961-geography",
					Name:        "1961geography",
					HRef:        "http://localhost:22400/code-lists/1961-geography",
					IsHierarchy: true,
				}, {
					ID:          "ashe-sex",
					Name:        "sex",
					HRef:        "http://localhost:22400/code-lists/ashe-sex",
					IsHierarchy: false,
				},
			},
		},
	},
}

//recipe for the SH10 table for the 1961 Census
var Census1961SH10 = Response{
	ID:     "00811186-299d-40db-a2b3-0714eb48b880",
	Alias:  "1961 Census: Population outside Private Households",
	Format: "v4",
	InputFiles: []file{
		{"Census1961SH10"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "census-1961-outside-private-households",
			Editions:  []string{"time-series"},
			Title:     "1961 Census: Population outside Private Households",
			CodeLists: []CodeList{
				{
					ID:          "calendar-years",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/calendar-years",
					IsHierarchy: false,
				}, {
					ID:          "1961-geography",
					Name:        "1961geography",
					HRef:        "http://localhost:22400/code-lists/1961-geography",
					IsHierarchy: true,
				}, {
					ID:          "ashe-sex",
					Name:        "sex",
					HRef:        "http://localhost:22400/code-lists/ashe-sex",
					IsHierarchy: false,
				},
			},
		},
	},
}

//recipe for the SH14 table for the 1961 Census
var Census1961SH14 = Response{
	ID:     "a8614752-df5d-4f4a-9bdc-8de9beaa3d2d",
	Alias:  "1961 Census: Population under 21 Years of Age",
	Format: "v4",
	InputFiles: []file{
		{"Census1961SH14"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "census-1961-under-21",
			Editions:  []string{"time-series"},
			Title:     "1961 Census: Population under 21 Years of Age",
			CodeLists: []CodeList{
				{
					ID:          "calendar-years",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/calendar-years",
					IsHierarchy: false,
				}, {
					ID:          "1961-geography",
					Name:        "1961geography",
					HRef:        "http://localhost:22400/code-lists/1961-geography",
					IsHierarchy: true,
				}, {
					ID:          "ashe-sex",
					Name:        "sex",
					HRef:        "http://localhost:22400/code-lists/ashe-sex",
					IsHierarchy: false,
				},
			},
		},
	},
}
