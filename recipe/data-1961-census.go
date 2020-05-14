package recipe

//Census1961 - 1961 Census Data
var Census1961 = Response{
	ID:     "4fff878f-f642-4113-9bed-85ae19a19ee7",
	Alias:  "1961 Census",
	Format: "v4",
	InputFiles: []file{
		{"1961Census"},
	},
	OutputInstances: []Instance{
		{
			DatasetID: "census-1961",
			Editions:  []string{"time-series"},
			Title:     "1961 Census Data",
			CodeLists: []CodeList{
				{
					ID:          "calendar-years",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/calendar-years",
					IsHierarchy: falseValPtr,
				}, {
					ID:          "1961-geography",
					Name:        "1961geography",
					HRef:        "http://localhost:22400/code-lists/1961-geography",
					IsHierarchy: trueValPtr,
				}, {
					ID:          "ashe-sex",
					Name:        "sex",
					HRef:        "http://localhost:22400/code-lists/ashe-sex",
					IsHierarchy: falseValPtr,
				}, {
					ID:          "census-1961-marital-status",
					Name:        "maritalstatus",
					HRef:        "http://localhost:22400/code-lists/census-1961-marital-status",
					IsHierarchy: falseValPtr,
				}, {
					ID:          "census-1961-age-groups",
					Name:        "agegroup",
					HRef:        "http://localhost:22400/code-lists/census-1961-age-groups",
					IsHierarchy: falseValPtr,
				},
			},
		},
	},
}

//Census1961SH01 - recipe for the SH01 table for the 1961 Census
var Census1961SH01 = Response{
	ID:     "0cbcf92b-4b72-4ecf-bac0-4d9227a8ce12",
	Alias:  "1961 Census: SH01",
	Format: "v4",
	InputFiles: []file{
		{"Census1961SH01"},
	},
	OutputInstances: []Instance{
		{
			DatasetID: "census-1961-sh01",
			Editions:  []string{"time-series"},
			Title:     "1961 Census: SH01",
			CodeLists: []CodeList{
				{
					ID:          "calendar-years",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/calendar-years",
					IsHierarchy: falseValPtr,
				}, {
					ID:          "1961-geography",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/1961-geography",
					IsHierarchy: trueValPtr,
				}, {
					ID:          "census-1961-tenure",
					Name:        "tenure",
					HRef:        "http://localhost:22400/code-lists/census-1961-tenure",
					IsHierarchy: falseValPtr,
				}, {
					ID:          "census-1961-tenure-variable",
					Name:        "tenurevariable",
					HRef:        "http://localhost:22400/code-lists/census-1961-tenure-variable",
					IsHierarchy: falseValPtr,
				},
			},
		},
	},
}

//Census1961SH02 - recipe for the SH02 table for the 1961 Census
var Census1961SH02 = Response{
	ID:     "50f91aa5-4c21-4892-aa63-d6f00cc5466d",
	Alias:  "1961 Census: SH02",
	Format: "v4",
	InputFiles: []file{
		{"Census1961SH02"},
	},
	OutputInstances: []Instance{
		{
			DatasetID: "census-1961-sh02",
			Editions:  []string{"time-series"},
			Title:     "1961 Census: SH02",
			CodeLists: []CodeList{
				{
					ID:          "calendar-years",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/calendar-years",
					IsHierarchy: falseValPtr,
				}, {
					ID:          "1961-geography",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/1961-geography",
					IsHierarchy: trueValPtr,
				}, {
					ID:          "census-1961-dwellings",
					Name:        "dwellings",
					HRef:        "http://localhost:22400/code-lists/census-1961-dwellings",
					IsHierarchy: falseValPtr,
				}, {
					ID:          "census-1961-building-type-variable",
					Name:        "buildingtypevariable",
					HRef:        "http://localhost:22400/code-lists/census-1961-building-type-variable",
					IsHierarchy: falseValPtr,
				},
			},
		},
	},
}

//Census1961SH03 - recipe for the SH03 table for the 1961 Census
var Census1961SH03 = Response{
	ID:     "638e67a5-189b-491a-9f81-1596d7cae5c3",
	Alias:  "1961 Census: SH03",
	Format: "v4",
	InputFiles: []file{
		{"Census1961SH03"},
	},
	OutputInstances: []Instance{
		{
			DatasetID: "census-1961-sh03",
			Editions:  []string{"time-series"},
			Title:     "1961 Census: SH03",
			CodeLists: []CodeList{
				{
					ID:          "calendar-years",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/calendar-years",
					IsHierarchy: falseValPtr,
				}, {
					ID:          "1961-geography",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/1961-geography",
					IsHierarchy: trueValPtr,
				}, {
					ID:          "census-1961-birthplace-outside-uk",
					Name:        "birthplace",
					HRef:        "http://localhost:22400/code-lists/census-1961-birthplace-outside-uk",
					IsHierarchy: falseValPtr,
				}, {
					ID:          "ashe-sex",
					Name:        "sex",
					HRef:        "http://localhost:22400/code-lists/ashe-sex",
					IsHierarchy: falseValPtr,
				},
			},
		},
	},
}

//Census1961SH04 - recipe for the SH04 table for the 1961 Census
var Census1961SH04 = Response{
	ID:     "ca1cf317-2fc3-4d8e-bd87-0d2f7238996b",
	Alias:  "1961 Census: SH04",
	Format: "v4",
	InputFiles: []file{
		{"Census1961SH04"},
	},
	OutputInstances: []Instance{
		{
			DatasetID: "census-1961-sh04",
			Editions:  []string{"time-series"},
			Title:     "1961 Census: SH04",
			CodeLists: []CodeList{
				{
					ID:          "calendar-years",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/calendar-years",
					IsHierarchy: falseValPtr,
				}, {
					ID:          "1961-geography",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/1961-geography",
					IsHierarchy: trueValPtr,
				}, {
					ID:          "census-1961-nationality",
					Name:        "nationality",
					HRef:        "http://localhost:22400/code-lists/census-1961-nationality",
					IsHierarchy: falseValPtr,
				}, {
					ID:          "ashe-sex",
					Name:        "sex",
					HRef:        "http://localhost:22400/code-lists/ashe-sex",
					IsHierarchy: falseValPtr,
				},
			},
		},
	},
}

//Census1961SH05 - recipe for the SH05 table for the 1961 Census
var Census1961SH05 = Response{
	ID:     "57c287e9-424d-49a0-ab58-b7cb98bf19fc",
	Alias:  "1961 Census: SH05",
	Format: "v4",
	InputFiles: []file{
		{"Census1961SH05"},
	},
	OutputInstances: []Instance{
		{
			DatasetID: "census-1961-sh05",
			Editions:  []string{"time-series"},
			Title:     "1961 Census: SH05",
			CodeLists: []CodeList{
				{
					ID:          "calendar-years",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/calendar-years",
					IsHierarchy: falseValPtr,
				}, {
					ID:          "1961-geography",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/1961-geography",
					IsHierarchy: trueValPtr,
				}, {
					ID:          "census-1961-non-private-population-establishment-type",
					Name:        "establishmenttype",
					HRef:        "http://localhost:22400/code-lists/census-1961-non-private-population-establishment-type",
					IsHierarchy: falseValPtr,
				}, {
					ID:          "census-1961-non-private-population-variable",
					Name:        "nonprivatepopulationvariable",
					HRef:        "http://localhost:22400/code-lists/census-1961-non-private-population-variable",
					IsHierarchy: falseValPtr,
				},
			},
		},
	},
}

//Census1961SH06 - recipe for the SH06 table for the 1961 Census
var Census1961SH06 = Response{
	ID:     "86ff18f0-8aa0-4073-b8ad-c0f15b9ed75a",
	Alias:  "1961 Census: SH06",
	Format: "v4",
	InputFiles: []file{
		{"Census1961SH06"},
	},
	OutputInstances: []Instance{
		{
			DatasetID: "census-1961-sh06",
			Editions:  []string{"time-series"},
			Title:     "1961 Census: SH06",
			CodeLists: []CodeList{
				{
					ID:          "calendar-years",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/calendar-years",
					IsHierarchy: falseValPtr,
				}, {
					ID:          "1961-geography",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/1961-geography",
					IsHierarchy: trueValPtr,
				}, {
					ID:          "census-1961-amenities",
					Name:        "amenities",
					HRef:        "http://localhost:22400/code-lists/census-1961-amenities",
					IsHierarchy: falseValPtr,
				}, {
					ID:          "census-1961-household-type",
					Name:        "householdtype",
					HRef:        "http://localhost:22400/code-lists/census-1961-household-type",
					IsHierarchy: falseValPtr,
				},
			},
		},
	},
}

//Census1961SH07 - recipe for the SH07 table for the 1961 Census
var Census1961SH07 = Response{
	ID:     "c1960f40-b988-4875-8b85-05878afa8e9c",
	Alias:  "1961 Census: SH07",
	Format: "v4",
	InputFiles: []file{
		{"Census1961SH07"},
	},
	OutputInstances: []Instance{
		{
			DatasetID: "census-1961-sh07",
			Editions:  []string{"time-series"},
			Title:     "1961 Census: Old Persons Alone",
			CodeLists: []CodeList{
				{
					ID:          "calendar-years",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/calendar-years",
					IsHierarchy: falseValPtr,
				}, {
					ID:          "1961-geography",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/1961-geography",
					IsHierarchy: trueValPtr,
				}, {
					ID:          "ashe-sex",
					Name:        "sex",
					HRef:        "http://localhost:22400/code-lists/ashe-sex",
					IsHierarchy: falseValPtr,
				},
			},
		},
	},
}

//Census1961SH08 - recipe for the SH08 table for the 1961 Census
var Census1961SH08 = Response{
	ID:     "a97737de-7cee-4bff-bff6-21f2f85191ca",
	Alias:  "1961 Census: SH08",
	Format: "v4",
	InputFiles: []file{
		{"Census1961SH08"},
	},
	OutputInstances: []Instance{
		{
			DatasetID: "census-1961-sh08",
			Editions:  []string{"time-series"},
			Title:     "1961 Census: SH08",
			CodeLists: []CodeList{
				{
					ID:          "calendar-years",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/calendar-years",
					IsHierarchy: falseValPtr,
				}, {
					ID:          "1961-geography",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/1961-geography",
					IsHierarchy: trueValPtr,
				}, {
					ID:          "census-1961-older-households",
					Name:        "olderhouseholds",
					HRef:        "http://localhost:22400/code-lists/census-1961-older-households",
					IsHierarchy: falseValPtr,
				}, {
					ID:          "census-1961-of-pensionable-age",
					Name:        "ofpensionableage",
					HRef:        "http://localhost:22400/code-lists/census-1961-of-pensionable-age",
					IsHierarchy: falseValPtr,
				},
			},
		},
	},
}

//Census1961SH09 - recipe for the SH09 table for the 1961 Census
var Census1961SH09 = Response{
	ID:     "0c691253-9c47-4e99-ab80-353337a6ab0a",
	Alias:  "1961 Census: SH09",
	Format: "v4",
	InputFiles: []file{
		{"Census1961SH09"},
	},
	OutputInstances: []Instance{
		{
			DatasetID: "census-1961-sh09",
			Editions:  []string{"time-series"},
			Title:     "1961 Census: SH09",
			CodeLists: []CodeList{
				{
					ID:          "calendar-years",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/calendar-years",
					IsHierarchy: falseValPtr,
				}, {
					ID:          "1961-geography",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/1961-geography",
					IsHierarchy: trueValPtr,
				}, {
					ID:          "census-1961-hotel-variable",
					Name:        "hotelvariable",
					HRef:        "http://localhost:22400/code-lists/census-1961-hotel-variable",
					IsHierarchy: falseValPtr,
				}, {
					ID:          "census-1961-number-of-rooms",
					Name:        "numberofrooms",
					HRef:        "http://localhost:22400/code-lists/census-1961-number-of-rooms",
					IsHierarchy: falseValPtr,
				},
			},
		},
	},
}

//Census1961SH10 - recipe for the SH10 table for the 1961 Census
var Census1961SH10 = Response{
	ID:     "00811186-299d-40db-a2b3-0714eb48b880",
	Alias:  "1961 Census: SH10",
	Format: "v4",
	InputFiles: []file{
		{"Census1961SH10"},
	},
	OutputInstances: []Instance{
		{
			DatasetID: "census-1961-sh10",
			Editions:  []string{"time-series"},
			Title:     "1961 Census: Population outside Private Households",
			CodeLists: []CodeList{
				{
					ID:          "calendar-years",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/calendar-years",
					IsHierarchy: falseValPtr,
				}, {
					ID:          "1961-geography",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/1961-geography",
					IsHierarchy: trueValPtr,
				}, {
					ID:          "five-year-age-brackets",
					Name:        "agebrackets",
					HRef:        "http://localhost:22400/code-lists/five-year-age-brackets",
					IsHierarchy: falseValPtr,
				}, {
					ID:          "ashe-sex",
					Name:        "sex",
					HRef:        "http://localhost:22400/code-lists/ashe-sex",
					IsHierarchy: falseValPtr,
				},
			},
		},
	},
}

//Census1961SH11 - recipe for the SH11 table for the 1961 Census
var Census1961SH11 = Response{
	ID:     "f655b71f-ee31-4c6d-b57e-0a06d8793788",
	Alias:  "1961 Census: SH11",
	Format: "v4",
	InputFiles: []file{
		{"Census1961SH11"},
	},
	OutputInstances: []Instance{
		{
			DatasetID: "census-1961-sh11",
			Editions:  []string{"time-series"},
			Title:     "1961 Census: SH11",
			CodeLists: []CodeList{
				{
					ID:          "calendar-years",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/calendar-years",
					IsHierarchy: falseValPtr,
				}, {
					ID:          "1961-geography",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/1961-geography",
					IsHierarchy: trueValPtr,
				}, {
					ID:          "census-1961-household-type",
					Name:        "householdtype",
					HRef:        "http://localhost:22400/code-lists/census-1961-household-type",
					IsHierarchy: falseValPtr,
				}, {
					ID:          "census-1961-persons-per-room",
					Name:        "personsperroom",
					HRef:        "http://localhost:22400/code-lists/census-1961-persons-per-room",
					IsHierarchy: falseValPtr,
				},
			},
		},
	},
}

//Census1961SH12 - recipe for the SH12 table for the 1961 Census
var Census1961SH12 = Response{
	ID:     "c2ec6e78-d130-44cd-bc86-2b9f9b897117",
	Alias:  "1961 Census: SH12",
	Format: "v4",
	InputFiles: []file{
		{"Census1961SH12"},
	},
	OutputInstances: []Instance{
		{
			DatasetID: "census-1961-sh12",
			Editions:  []string{"time-series"},
			Title:     "1961 Census: SH12",
			CodeLists: []CodeList{
				{
					ID:          "calendar-years",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/calendar-years",
					IsHierarchy: falseValPtr,
				}, {
					ID:          "1961-geography",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/1961-geography",
					IsHierarchy: trueValPtr,
				}, {
					ID:          "census-1961-usual-residence-outside-local-authority",
					Name:        "usualresidenceoutsidelocalauthority",
					HRef:        "http://localhost:22400/code-lists/census-1961-usual-residence-outside-local-authority",
					IsHierarchy: falseValPtr,
				}, {
					ID:          "ashe-sex",
					Name:        "sex",
					HRef:        "http://localhost:22400/code-lists/ashe-sex",
					IsHierarchy: falseValPtr,
				},
			},
		},
	},
}

//Census1961SH13 - recipe for the SH13 table for the 1961 Census
var Census1961SH13 = Response{
	ID:     "c4ffac95-f103-44b0-b3a3-1a1a18c1a5f2",
	Alias:  "1961 Census: SH13",
	Format: "v4",
	InputFiles: []file{
		{"Census1961SH13"},
	},
	OutputInstances: []Instance{
		{
			DatasetID: "census-1961-sh13",
			Editions:  []string{"time-series"},
			Title:     "1961 Census: SH13",
			CodeLists: []CodeList{
				{
					ID:          "calendar-years",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/calendar-years",
					IsHierarchy: falseValPtr,
				}, {
					ID:          "1961-geography",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/1961-geography",
					IsHierarchy: trueValPtr,
				}, {
					ID:          "ashe-sex",
					Name:        "sex",
					HRef:        "http://localhost:22400/code-lists/ashe-sex",
					IsHierarchy: falseValPtr,
				}, {
					ID:          "census-1961-marital-status",
					Name:        "maritalstatus",
					HRef:        "http://localhost:22400/code-lists/census-1961-marital-status",
					IsHierarchy: falseValPtr,
				}, {
					ID:          "five-year-age-brackets",
					Name:        "agebrackets",
					HRef:        "http://localhost:22400/code-lists/five-year-age-brackets",
					IsHierarchy: falseValPtr,
				},
			},
		},
	},
}

//Census1961SH14 - recipe for the SH14 table for the 1961 Census
var Census1961SH14 = Response{
	ID:     "a8614752-df5d-4f4a-9bdc-8de9beaa3d2d",
	Alias:  "1961 Census: SH14",
	Format: "v4",
	InputFiles: []file{
		{"Census1961SH14"},
	},
	OutputInstances: []Instance{
		{
			DatasetID: "census-1961-under-21",
			Editions:  []string{"time-series"},
			Title:     "1961 Census: Population under 21 Years of Age",
			CodeLists: []CodeList{
				{
					ID:          "calendar-years",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/calendar-years",
					IsHierarchy: falseValPtr,
				}, {
					ID:          "1961-geography",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/1961-geography",
					IsHierarchy: trueValPtr,
				}, {
					ID:          "mid-year-pop-age",
					Name:        "age",
					HRef:        "http://localhost:22400/code-lists/mid-year-pop-age",
					IsHierarchy: falseValPtr,
				}, {
					ID:          "ashe-sex",
					Name:        "sex",
					HRef:        "http://localhost:22400/code-lists/ashe-sex",
					IsHierarchy: falseValPtr,
				},
			},
		},
	},
}

//Census1961SH15 - recipe for the SH15 table for the 1961 Census
var Census1961SH15 = Response{
	ID:     "8314c806-453d-4c80-b0f7-dde5f40bcf66",
	Alias:  "1961 Census: SH15",
	Format: "v4",
	InputFiles: []file{
		{"Census1961SH15"},
	},
	OutputInstances: []Instance{
		{
			DatasetID: "census-1961-sh15",
			Editions:  []string{"time-series"},
			Title:     "1961 Census: SH15",
			CodeLists: []CodeList{
				{
					ID:          "calendar-years",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/calendar-years",
					IsHierarchy: falseValPtr,
				}, {
					ID:          "1961-geography",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/1961-geography",
					IsHierarchy: trueValPtr,
				}, {
					ID:          "census-1961-households-sharing-dwelling",
					Name:        "numberofhouseholds",
					HRef:        "http://localhost:22400/code-lists/census-1961-households-sharing-dwelling",
					IsHierarchy: falseValPtr,
				},
			},
		},
	},
}

//Census1961ST01 - recipe for the ST01 table for the 1961 Census
var Census1961ST01 = Response{
	ID:     "8b120c65-9ad7-4a85-82bd-d6dc8862a4eb",
	Alias:  "1961 Census: ST01",
	Format: "v4",
	InputFiles: []file{
		{"Census1961ST01"},
	},
	OutputInstances: []Instance{
		{
			DatasetID: "census-1961-st01",
			Editions:  []string{"time-series"},
			Title:     "1961 Census: ST01",
			CodeLists: []CodeList{
				{
					ID:          "calendar-years",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/calendar-years",
					IsHierarchy: falseValPtr,
				}, {
					ID:          "1961-geography",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/1961-geography",
					IsHierarchy: trueValPtr,
				}, {
					ID:          "census-1961-sex-activity",
					Name:        "sexactivity",
					HRef:        "http://localhost:22400/code-lists/census-1961-sex-activity",
					IsHierarchy: falseValPtr,
				},
			},
		},
	},
}

//Census1961ST02 - recipe for the ST02 table for the 1961 Census
var Census1961ST02 = Response{
	ID:     "5129ad6e-3bea-4c0d-8092-b8dd617b7c19",
	Alias:  "1961 Census: ST02",
	Format: "v4",
	InputFiles: []file{
		{"Census1961ST02"},
	},
	OutputInstances: []Instance{
		{
			DatasetID: "census-1961-st02",
			Editions:  []string{"time-series"},
			Title:     "1961 Census: ST02",
			CodeLists: []CodeList{
				{
					ID:          "calendar-years",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/calendar-years",
					IsHierarchy: falseValPtr,
				}, {
					ID:          "1961-geography",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/1961-geography",
					IsHierarchy: trueValPtr,
				}, {
					ID:          "census-1961-work-location",
					Name:        "worklocation",
					HRef:        "http://localhost:22400/code-lists/census-1961-work-location",
					IsHierarchy: falseValPtr,
				},
			},
		},
	},
}

//Census1961ST03 - recipe for the ST03 table for the 1961 Census
var Census1961ST03 = Response{
	ID:     "1f628966-019d-42fd-b1a8-bbb22e82714a",
	Alias:  "1961 Census: ST03",
	Format: "v4",
	InputFiles: []file{
		{"Census1961ST03"},
	},
	OutputInstances: []Instance{
		{
			DatasetID: "census-1961-st03",
			Editions:  []string{"time-series"},
			Title:     "1961 Census: ST03",
			CodeLists: []CodeList{
				{
					ID:          "calendar-years",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/calendar-years",
					IsHierarchy: falseValPtr,
				}, {
					ID:          "1961-geography",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/1961-geography",
					IsHierarchy: trueValPtr,
				}, {
					ID:          "census-1961-sex-activity",
					Name:        "sexactivity",
					HRef:        "http://localhost:22400/code-lists/census-1961-sex-activity",
					IsHierarchy: falseValPtr,
				}, {
					ID:          "census-1961-economic-activity-age-group",
					Name:        "economicactivityagegroup",
					HRef:        "http://localhost:22400/code-lists/census-1961-economic-activity-age-group",
					IsHierarchy: falseValPtr,
				},
			},
		},
	},
}

//Census1961ST04 - recipe for the ST04 table for the 1961 Census
var Census1961ST04 = Response{
	ID:     "30e03190-0f80-4586-b7b8-e56443d92717",
	Alias:  "1961 Census: ST04",
	Format: "v4",
	InputFiles: []file{
		{"Census1961ST04"},
	},
	OutputInstances: []Instance{
		{
			DatasetID: "census-1961-st04",
			Editions:  []string{"time-series"},
			Title:     "1961 Census: ST04",
			CodeLists: []CodeList{
				{
					ID:          "calendar-years",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/calendar-years",
					IsHierarchy: falseValPtr,
				}, {
					ID:          "1961-geography",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/1961-geography",
					IsHierarchy: trueValPtr,
				}, {
					ID:          "ashe-sex",
					Name:        "sex",
					HRef:        "http://localhost:22400/code-lists/ashe-sex",
					IsHierarchy: falseValPtr,
				},
			},
		},
	},
}

//Census1961ST05 - recipe for the ST05 table for the 1961 Census
var Census1961ST05 = Response{
	ID:     "c4949ddd-63fb-4b29-8f7c-ba9209df59a2",
	Alias:  "1961 Census: ST05",
	Format: "v4",
	InputFiles: []file{
		{"Census1961ST05"},
	},
	OutputInstances: []Instance{
		{
			DatasetID: "census-1961-st05",
			Editions:  []string{"time-series"},
			Title:     "1961 Census: ST05",
			CodeLists: []CodeList{
				{
					ID:          "calendar-years",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/calendar-years",
					IsHierarchy: falseValPtr,
				}, {
					ID:          "1961-geography",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/1961-geography",
					IsHierarchy: trueValPtr,
				}, {
					ID:          "ashe-sex",
					Name:        "sex",
					HRef:        "http://localhost:22400/code-lists/ashe-sex",
					IsHierarchy: falseValPtr,
				}, {
					ID:          "census-1961-industry",
					Name:        "industry",
					HRef:        "http://localhost:22400/code-lists/census-1961-industry",
					IsHierarchy: falseValPtr,
				},
			},
		},
	},
}

//Census1961ST06 - recipe for the ST06 table for the 1961 Census
var Census1961ST06 = Response{
	ID:     "36a977da-c536-4bbc-a52c-aef8f9482a26",
	Alias:  "1961 Census: ST06",
	Format: "v4",
	InputFiles: []file{
		{"Census1961ST06"},
	},
	OutputInstances: []Instance{
		{
			DatasetID: "census-1961-st06",
			Editions:  []string{"time-series"},
			Title:     "1961 Census: ST06",
			CodeLists: []CodeList{
				{
					ID:          "calendar-years",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/calendar-years",
					IsHierarchy: falseValPtr,
				}, {
					ID:          "1961-geography",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/1961-geography",
					IsHierarchy: trueValPtr,
				}, {
					ID:          "ashe-sex",
					Name:        "sex",
					HRef:        "http://localhost:22400/code-lists/ashe-sex",
					IsHierarchy: falseValPtr,
				}, {
					ID:          "census-1961-terminal-education-age",
					Name:        "terminaleducationage",
					HRef:        "http://localhost:22400/code-lists/census-1961-terminal-education-age",
					IsHierarchy: falseValPtr,
				},
			},
		},
	},
}

//Census1961ST07 - recipe for the ST07 table for the 1961 Census
var Census1961ST07 = Response{
	ID:     "1f61208f-3121-453b-aa89-8a5d9f97581d",
	Alias:  "1961 Census: ST07",
	Format: "v4",
	InputFiles: []file{
		{"Census1961ST07"},
	},
	OutputInstances: []Instance{
		{
			DatasetID: "census-1961-st07",
			Editions:  []string{"time-series"},
			Title:     "1961 Census: ST07",
			CodeLists: []CodeList{
				{
					ID:          "calendar-years",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/calendar-years",
					IsHierarchy: falseValPtr,
				}, {
					ID:          "1961-geography",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/1961-geography",
					IsHierarchy: trueValPtr,
				}, {
					ID:          "census-1961-sex-activity",
					Name:        "sexactivity",
					HRef:        "http://localhost:22400/code-lists/census-1961-sex-activity",
					IsHierarchy: falseValPtr,
				}, {
					ID:          "census-1961-occupation-classification",
					Name:        "occupationclassification",
					HRef:        "http://localhost:22400/code-lists/census-1961-occupation-classification",
					IsHierarchy: falseValPtr,
				},
			},
		},
	},
}

//Census1961ST08 - recipe for the ST08 table for the 1961 Census
var Census1961ST08 = Response{
	ID:     "6c744f10-a92e-43f1-97ad-463ff0d6a0ae",
	Alias:  "1961 Census: ST08",
	Format: "v4",
	InputFiles: []file{
		{"Census1961ST08"},
	},
	OutputInstances: []Instance{
		{
			DatasetID: "census-1961-st08",
			Editions:  []string{"time-series"},
			Title:     "1961 Census: ST08",
			CodeLists: []CodeList{
				{
					ID:          "calendar-years",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/calendar-years",
					IsHierarchy: falseValPtr,
				}, {
					ID:          "1961-geography",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/1961-geography",
					IsHierarchy: trueValPtr,
				}, {
					ID:          "census-1961-movers",
					Name:        "movers",
					HRef:        "http://localhost:22400/code-lists/census-1961-movers",
					IsHierarchy: falseValPtr,
				}, {
					ID:          "census-1961-movers-age-sex",
					Name:        "moversagesex",
					HRef:        "http://localhost:22400/code-lists/census-1961-movers-age-sex",
					IsHierarchy: falseValPtr,
				},
			},
		},
	},
}

//Census1961ST09 - recipe for the ST09 table for the 1961 Census
var Census1961ST09 = Response{
	ID:     "a57378d4-8034-4c52-b68f-224e17b136a0",
	Alias:  "1961 Census: ST09",
	Format: "v4",
	InputFiles: []file{
		{"Census1961ST09"},
	},
	OutputInstances: []Instance{
		{
			DatasetID: "census-1961-st09",
			Editions:  []string{"time-series"},
			Title:     "1961 Census: ST09",
			CodeLists: []CodeList{
				{
					ID:          "calendar-years",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/calendar-years",
					IsHierarchy: falseValPtr,
				}, {
					ID:          "1961-geography",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/1961-geography",
					IsHierarchy: trueValPtr,
				}, {
					ID:          "census-1961-family-count",
					Name:        "familycount",
					HRef:        "http://localhost:22400/code-lists/census-1961-family-count",
					IsHierarchy: falseValPtr,
				},
			},
		},
	},
}

//Census1961SC11 - recipe for the SC11 table for the 1961 Census
var Census1961SC11 = Response{
	ID:     "49f7ab00-35d9-4f79-b55c-d2c99936a68f",
	Alias:  "1961 Census: SC11",
	Format: "v4",
	InputFiles: []file{
		{"Census1961SC11"},
	},
	OutputInstances: []Instance{
		{
			DatasetID: "census-1961-sc11",
			Editions:  []string{"time-series"},
			Title:     "1961 Census: SC11",
			CodeLists: []CodeList{
				{
					ID:          "calendar-years",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/calendar-years",
					IsHierarchy: falseValPtr,
				}, {
					ID:          "1961-geography",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/1961-geography",
					IsHierarchy: trueValPtr,
				}, {
					ID:          "census-1961-dwellings",
					Name:        "dwellings",
					HRef:        "http://localhost:22400/code-lists/census-1961-dwellings",
					IsHierarchy: falseValPtr,
				}, {
					ID:          "census-1961-number-of-rooms",
					Name:        "numberofrooms",
					HRef:        "http://localhost:22400/code-lists/census-1961-number-of-rooms",
					IsHierarchy: falseValPtr,
				},
			},
		},
	},
}

//Census1961SC13 - recipe for the SC13 table for the 1961 Census
var Census1961SC13 = Response{
	ID:     "4cb260b8-8892-4710-af30-549c28cab6e8",
	Alias:  "1961 Census: SC13",
	Format: "v4",
	InputFiles: []file{
		{"Census1961SC13"},
	},
	OutputInstances: []Instance{
		{
			DatasetID: "census-1961-sc13",
			Editions:  []string{"time-series"},
			Title:     "1961 Census: SC13",
			CodeLists: []CodeList{
				{
					ID:          "calendar-years",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/calendar-years",
					IsHierarchy: falseValPtr,
				}, {
					ID:          "1961-geography",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/1961-geography",
					IsHierarchy: trueValPtr,
				}, {
					ID:          "census-1961-dwellings",
					Name:        "dwellings",
					HRef:        "http://localhost:22400/code-lists/census-1961-dwellings",
					IsHierarchy: falseValPtr,
				}, {
					ID:          "census-1961-number-of-rooms",
					Name:        "numberofrooms",
					HRef:        "http://localhost:22400/code-lists/census-1961-number-of-rooms",
					IsHierarchy: falseValPtr,
				},
			},
		},
	},
}

//Census1961SC22 - recipe for the SC22 table for the 1961 Census
var Census1961SC22 = Response{
	ID:     "5b300852-c119-492b-bb4b-098a83ddb1b3",
	Alias:  "1961 Census: SC22",
	Format: "v4",
	InputFiles: []file{
		{"Census1961SC22"},
	},
	OutputInstances: []Instance{
		{
			DatasetID: "census-1961-sc22",
			Editions:  []string{"time-series"},
			Title:     "1961 Census: SC22",
			CodeLists: []CodeList{
				{
					ID:          "calendar-years",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/calendar-years",
					IsHierarchy: falseValPtr,
				}, {
					ID:          "1961-geography",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/1961-geography",
					IsHierarchy: trueValPtr,
				}, {
					ID:          "census-1961-dwellings",
					Name:        "dwellings",
					HRef:        "http://localhost:22400/code-lists/census-1961-dwellings",
					IsHierarchy: falseValPtr,
				}, {
					ID:          "census-1961-amenities",
					Name:        "amenities",
					HRef:        "http://localhost:22400/code-lists/census-1961-amenities",
					IsHierarchy: falseValPtr,
				}, {
					ID:          "unit-of-measure",
					Name:        "unitofmeasure",
					HRef:        "http://localhost:22400/code-lists/unit-of-measure",
					IsHierarchy: falseValPtr,
				},
			},
		},
	},
}
