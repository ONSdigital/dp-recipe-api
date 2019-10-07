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
					HRef:        "http://localhost:22400/code-lists/census-1961-tenure-variable",
					IsHierarchy: false,
				},
			},
		},
	},
}

//recipe for the SH02 table for the 1961 Census
var Census1961SH02 = Response{
	ID:     "50f91aa5-4c21-4892-aa63-d6f00cc5466d",
	Alias:  "1961 Census: SH02",
	Format: "v4",
	InputFiles: []file{
		{"Census1961SH02"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "census-1961-sh02",
			Editions:  []string{"time-series"},
			Title:     "1961 Census: SH02",
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
					ID:          "census-1961-building-type",
					Name:        "buildingtype",
					HRef:        "http://localhost:22400/code-lists/census-1961-building-type",
					IsHierarchy: false,
				}, {
					ID:          "census-1961-building-type-variable",
					Name:        "buildingtypevariable",
					HRef:        "http://localhost:22400/code-lists/census-1961-building-type-variable",
					IsHierarchy: false,
				},
			},
		},
	},
}

//recipe for the SH03 table for the 1961 Census
var Census1961SH03 = Response{
	ID:     "638e67a5-189b-491a-9f81-1596d7cae5c3",
	Alias:  "1961 Census: SH03",
	Format: "v4",
	InputFiles: []file{
		{"Census1961SH03"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "census-1961-sh03",
			Editions:  []string{"time-series"},
			Title:     "1961 Census: SH03",
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
					ID:          "census-1961-birthplace-outside-uk",
					Name:        "birthplace",
					HRef:        "http://localhost:22400/code-lists/census-1961-birthplace-outside-uk",
					IsHierarchy: false,
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

//recipe for the SH04 table for the 1961 Census
var Census1961SH04 = Response{
	ID:     "ca1cf317-2fc3-4d8e-bd87-0d2f7238996b",
	Alias:  "1961 Census: SH04",
	Format: "v4",
	InputFiles: []file{
		{"Census1961SH04"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "census-1961-sh04",
			Editions:  []string{"time-series"},
			Title:     "1961 Census: SH04",
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
					ID:          "census-1961-nationality",
					Name:        "nationality",
					HRef:        "http://localhost:22400/code-lists/census-1961-nationality",
					IsHierarchy: false,
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

//recipe for the SH05 table for the 1961 Census
var Census1961SH05 = Response{
	ID:     "57c287e9-424d-49a0-ab58-b7cb98bf19fc",
	Alias:  "1961 Census: SH05",
	Format: "v4",
	InputFiles: []file{
		{"Census1961SH05"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "census-1961-sh05",
			Editions:  []string{"time-series"},
			Title:     "1961 Census: SH05",
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
					ID:          "census-1961-non-private-population-establishment-type",
					Name:        "establishmenttype",
					HRef:        "http://localhost:22400/code-lists/census-1961-non-private-population-establishment-type",
					IsHierarchy: false,
				}, {
					ID:          "census-1961-non-private-population-variable",
					Name:        "nonprivatepopulationvariable",
					HRef:        "http://localhost:22400/code-lists/census-1961-non-private-population-variable",
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
					Name:        "geography",
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

//recipe for the SH08 table for the 1961 Census
var Census1961SH08 = Response{
	ID:     "a97737de-7cee-4bff-bff6-21f2f85191ca",
	Alias:  "1961 Census: SH08",
	Format: "v4",
	InputFiles: []file{
		{"Census1961SH08"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "census-1961-sh08",
			Editions:  []string{"time-series"},
			Title:     "1961 Census: SH08",
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
					ID:          "census-1961-older-households",
					Name:        "olderhouseholds",
					HRef:        "http://localhost:22400/code-lists/census-1961-older-households",
					IsHierarchy: false,
				}, {
					ID:          "census-1961-of-pensionable-age",
					Name:        "ofpensionableage",
					HRef:        "http://localhost:22400/code-lists/census-1961-of-pensionable-age",
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
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/1961-geography",
					IsHierarchy: true,
				}, {
					ID:          "five-year-age-brackets",
					Name:        "agebrackets",
					HRef:        "http://localhost:22400/code-lists/five-year-age-brackets",
					IsHierarchy: false,
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

//recipe for the SH12 table for the 1961 Census
var Census1961SH12 = Response{
	ID:     "c2ec6e78-d130-44cd-bc86-2b9f9b897117",
	Alias:  "1961 Census: SH12",
	Format: "v4",
	InputFiles: []file{
		{"Census1961SH12"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "census-1961-sh12",
			Editions:  []string{"time-series"},
			Title:     "1961 Census: SH12",
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
					ID:          "census-1961-usual-residence-outside-local-authority",
					Name:        "usualresidenceoutsidelocalauthority",
					HRef:        "http://localhost:22400/code-lists/census-1961-usual-residence-outside-local-authority",
					IsHierarchy: false,
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
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/1961-geography",
					IsHierarchy: true,
				}, {
					ID:          "mid-year-pop-age",
					Name:        "age",
					HRef:        "http://localhost:22400/code-lists/mid-year-pop-age",
					IsHierarchy: false,
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

//recipe for the SH15 table for the 1961 Census
var Census1961SH15 = Response{
	ID:     "8314c806-453d-4c80-b0f7-dde5f40bcf66",
	Alias:  "1961 Census: SH15",
	Format: "v4",
	InputFiles: []file{
		{"Census1961SH15"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "census-1961-sh15",
			Editions:  []string{"time-series"},
			Title:     "1961 Census: SH15",
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
					ID:          "census-1961-households-sharing-dwelling",
					Name:        "numberofhouseholds",
					HRef:        "http://localhost:22400/code-lists/census-1961-households-sharing-dwelling",
					IsHierarchy: false,
				}, 
			},
		},
	},
}
