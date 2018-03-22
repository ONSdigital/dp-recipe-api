package recipe

//To add a new recipe:
// - add a recipeResponse{} document with an appropriate var name
// - add the var name to the list of Items
// - increment the Count and Total count in the RecipeList

//FullList of recipes available via this API
var FullList = List{
	Items: []Response{CPI, CPIH, MidYearPopEst, ASHE7Hours, ASHE7Earnings, ASHE8Hours, ASHE8Earnings, OPSSMembership, OPSSRates,
		CrimeAccommodation, NPP, CrimeOffences, Migration401AGQ, Migration401AG1, Migration401AG2, Migration402},
	Start: 0,
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
			Editions:  []string{"time-series"},
			Title:     "UK consumer price inflation",
			CodeLists: []CodeList{
				{
					ID:          "64d384f1-ea3b-445c-8fb8-aa453f96e58a",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/64d384f1-ea3b-445c-8fb8-aa453f96e58a",
					IsHierarchy: false,
				}, {
					ID:          "65107A9F-7DA3-4B41-A410-6F6D9FBD68C3",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/65107A9F-7DA3-4B41-A410-6F6D9FBD68C3",
					IsHierarchy: false,
				}, {
					ID:          "e44de4c4-d39e-4e2f-942b-3ca10584d078",
					Name:        "aggregate",
					HRef:        "http://localhost:22400/code-lists/e44de4c4-d39e-4e2f-942b-3ca10584d078",
					IsHierarchy: true,
				},
			},
		},
	},
}

//CPIH recipe for transforming a given input to a CPIH dataset
var CPIH = Response{
	ID:     "2943f3c5-c3f1-4a9a-aa6e-14d21c33524c",
	Alias:  "CPIH",
	Format: "v4",
	InputFiles: []file{
		{"CPIH v4"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "cpih01",
			Editions:  []string{"time-series"},
			Title:     "Consumer Prices Index including owner occupiers’ housing costs (CPIH)",
			CodeLists: []CodeList{
				{
					ID:          "time",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/time",
					IsHierarchy: false,
				}, {
					ID:          "uk-only",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/uk-only",
					IsHierarchy: false,
				}, {
					ID:          "cpih1dim1aggid",
					Name:        "aggregate",
					HRef:        "http://localhost:22400/code-lists/cpih1dim1aggid",
					IsHierarchy: true,
				},
			},
		},
	},
}

//MidYearPopEst recipe for transforming a given input to a mid year population estimate dataset
var MidYearPopEst = Response{
	ID:     "40AA070E-7A43-4EC5-A1FC-84CEA2BC4461",
	Alias:  "Mid-year Population Estimates",
	Format: "v4",
	InputFiles: []file{
		{"Mid-year Population Estimates v4"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "mid-year-pop-est",
			Editions:  []string{"time-series"},
			Title:     "Population Estimates for UK, England and Wales, Scotland and Northern Ireland",
			CodeLists: []CodeList{
				{
					ID:          "calendar-years",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/calendar-years",
					IsHierarchy: false,
				}, {
					ID:          "mid-year-pop-geography",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/mid-year-pop-geography",
					IsHierarchy: true,
				}, {
					ID:          "mid-year-pop-sex",
					Name:        "sex",
					HRef:        "http://localhost:22400/code-lists/mid-year-pop-sex",
					IsHierarchy: false,
				}, {
					ID:          "mid-year-pop-age",
					Name:        "age",
					HRef:        "http://localhost:22400/code-lists/mid-year-pop-age",
					IsHierarchy: false,
				},
			},
		},
	},
}

// ASHE7Hours recipe for transforming a given input to an ASHE table 7 hours dataset
var ASHE7Hours = Response{
	ID:     "613C1384-01FF-4E3E-A24A-3A98A75BBAD8",
	Alias:  "ASHE Table 7 (hours)",
	Format: "v4",
	InputFiles: []file{
		{"Ashe Table 7 Hours Dataset v4"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "ashe-table-7-hours",
			Editions:  []string{"time-series"},
			Title:     "Annual Summary of Hours Worked, Place of Work by Local Authority",
			CodeLists: []CodeList{
				{
					ID:          "calendar-years",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/calendar-years",
					IsHierarchy: false,
				}, {
					ID:          "ashe-sex",
					Name:        "sex",
					HRef:        "http://localhost:22400/code-lists/ashe-sex",
					IsHierarchy: false,
				}, {
					ID:          "ashe-working-pattern",
					Name:        "workingpattern",
					HRef:        "http://localhost:22400/code-lists/ashe-working-pattern",
					IsHierarchy: false,
				}, {
					ID:          "ashe-hours",
					Name:        "hours",
					HRef:        "http://localhost:22400/code-lists/ashe-hours",
					IsHierarchy: false,
				}, {
					ID:          "ashe-statistics",
					Name:        "statistics",
					HRef:        "http://localhost:22400/code-lists/ashe-statistics",
					IsHierarchy: false,
				}, {
					ID:          "ashe-table-7-geography",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/ashe-table-7-geography",
					IsHierarchy: true,
				},
			},
		},
	},
}

// ASHE7Earnings recipe for transforming a given input to an ASHE table 7 hours dataset
var ASHE7Earnings = Response{
	ID:     "18FFF4C4-1A2A-466A-B157-8CA872F6FCF0",
	Alias:  "ASHE Table 7 (earnings)",
	Format: "v4",
	InputFiles: []file{
		{"Ashe Table 7 Earnings Dataset v4"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "ashe-table-7-earnings",
			Editions:  []string{"time-series"},
			Title:     "Annual Summary of Earnings, Place of Work by Local Authority",
			CodeLists: []CodeList{
				{
					ID:          "calendar-years",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/calendar-years",
					IsHierarchy: false,
				}, {
					ID:          "ashe-sex",
					Name:        "sex",
					HRef:        "http://localhost:22400/code-lists/ashe-sex",
					IsHierarchy: false,
				}, {
					ID:          "ashe-working-pattern",
					Name:        "workingpattern",
					HRef:        "http://localhost:22400/code-lists/ashe-working-pattern",
					IsHierarchy: false,
				}, {
					ID:          "ashe-earnings",
					Name:        "earnings",
					HRef:        "http://localhost:22400/code-lists/ashe-earnings",
					IsHierarchy: false,
				}, {
					ID:          "ashe-statistics",
					Name:        "statistics",
					HRef:        "http://localhost:22400/code-lists/ashe-statistics",
					IsHierarchy: false,
				}, {
					ID:          "ashe-table-7-geography",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/ashe-table-7-geography",
					IsHierarchy: true,
				},
			},
		},
	},
}

// ASHE8Hours recipe for transforming a given input to an ASHE table 8 hours dataset
var ASHE8Hours = Response{
	ID:     "15340072-5D80-4581-A29A-E61BE1B2D815",
	Alias:  "ASHE Table 8 (hours)",
	Format: "v4",
	InputFiles: []file{
		{"Ashe Table 8 Hours Dataset v4"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "ashe-table-8-hours",
			Editions:  []string{"time-series"},
			Title:     "Annual Summary of Hours Worked, Place of Residence by Local Authority",
			CodeLists: []CodeList{
				{
					ID:          "calendar-years",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/calendar-years",
					IsHierarchy: false,
				}, {
					ID:          "ashe-sex",
					Name:        "sex",
					HRef:        "http://localhost:22400/code-lists/ashe-sex",
					IsHierarchy: false,
				}, {
					ID:          "ashe-working-pattern",
					Name:        "workingpattern",
					HRef:        "http://localhost:22400/code-lists/ashe-working-pattern",
					IsHierarchy: false,
				}, {
					ID:          "ashe-hours",
					Name:        "hours",
					HRef:        "http://localhost:22400/code-lists/ashe-hours",
					IsHierarchy: false,
				}, {
					ID:          "ashe-statistics",
					Name:        "statistics",
					HRef:        "http://localhost:22400/code-lists/ashe-statistics",
					IsHierarchy: false,
				}, {
					ID:          "ashe-table-8-geography",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/ashe-table-8-geography",
					IsHierarchy: true,
				},
			},
		},
	},
}

// ASHE8Earnings recipe for transforming a given input to an ASHE table 8 hours dataset
var ASHE8Earnings = Response{
	ID:     "2CE32F03-C2B3-4582-9989-1832935D4BA0",
	Alias:  "ASHE Table 8 (earnings)",
	Format: "v4",
	InputFiles: []file{
		{"Ashe Table 8 Earnings Dataset v4"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "ashe-table-8-earnings",
			Editions:  []string{"time-series"},
			Title:     "Annual Summary of Earnings, Place of Residence by Local Authority",
			CodeLists: []CodeList{
				{
					ID:          "calendar-years",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/calendar-years",
					IsHierarchy: false,
				}, {
					ID:          "ashe-sex",
					Name:        "sex",
					HRef:        "http://localhost:22400/code-lists/ashe-sex",
					IsHierarchy: false,
				}, {
					ID:          "ashe-working-pattern",
					Name:        "workingpattern",
					HRef:        "http://localhost:22400/code-lists/ashe-working-pattern",
					IsHierarchy: false,
				}, {
					ID:          "ashe-earnings",
					Name:        "earnings",
					HRef:        "http://localhost:22400/code-lists/ashe-earnings",
					IsHierarchy: false,
				}, {
					ID:          "ashe-statistics",
					Name:        "statistics",
					HRef:        "http://localhost:22400/code-lists/ashe-statistics",
					IsHierarchy: false,
				}, {
					ID:          "ashe-table-8-geography",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/ashe-table-8-geography",
					IsHierarchy: true,
				},
			},
		},
	},
}

//OPSSMembership recipe for transforming a given input to a OPSSMembership dataset
var OPSSMembership = Response{
	ID:     "F09C8D91-5A1E-4A1E-B2E9-A029C784E1F1",
	Alias:  "OPSSMembership",
	Format: "v4",
	InputFiles: []file{
		{"OPSSMembership v4"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "opss-membership",
			Editions:  []string{"2016"},
			Title:     "Occupational Pension Schemes Survey, Membership",
			CodeLists: []CodeList{
				{
					ID:          "calendar-years",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/calendar-years",
					IsHierarchy: false,
				}, {
					ID:          "uk-only",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/uk-only",
					IsHierarchy: false,
				}, {
					ID:          "opss-scheme-membership-sizeband",
					Name:        "schememembershipsizeband",
					HRef:        "http://localhost:22400/code-lists/opss-scheme-membership-sizeband",
					IsHierarchy: false,
				}, {
					ID:          "opss-status",
					Name:        "status",
					HRef:        "http://localhost:22400/code-lists/opss-status",
					IsHierarchy: false,
				}, {
					ID:          "opss-benefit-type",
					Name:        "benefittype",
					HRef:        "http://localhost:22400/code-lists/opss-benefit-type",
					IsHierarchy: false,
				}, {
					ID:          "opss-public-private-sector",
					Name:        "publicprivatesector",
					HRef:        "http://localhost:22400/code-lists/opss-public-private-sector",
					IsHierarchy: false,
				}, {
					ID:          "opss-membership-type",
					Name:        "membershiptype",
					HRef:        "http://localhost:22400/code-lists/opss-membership-type",
					IsHierarchy: false,
				},
			},
		},
	},
}

//OPSSRates recipe for transforming a given input to a OPSSMembership dataset
var OPSSRates = Response{
	ID:     "84FBAFE3-C40C-428E-A6C6-7717DAB89900",
	Alias:  "OPSSRates",
	Format: "v4",
	InputFiles: []file{
		{"OPSSRates v4"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "opss-rates",
			Editions:  []string{"2016"},
			Title:     "Occupational Pension Schemes Survey, Rates",
			CodeLists: []CodeList{
				{
					ID:          "calendar-years",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/calendar-years",
					IsHierarchy: false,
				}, {
					ID:          "uk-only",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/uk-only",
					IsHierarchy: false,
				}, {
					ID:          "opss-scheme-membership-sizeband",
					Name:        "schememembershipsizeband",
					HRef:        "http://localhost:22400/code-lists/opss-scheme-membership-sizeband",
					IsHierarchy: false,
				}, {
					ID:          "opss-status",
					Name:        "status",
					HRef:        "http://localhost:22400/code-lists/opss-status",
					IsHierarchy: false,
				}, {
					ID:          "opss-benefit-type",
					Name:        "benefittype",
					HRef:        "http://localhost:22400/code-lists/opss-benefit-type",
					IsHierarchy: false,
				}, {
					ID:          "opss-contributor",
					Name:        "contributor",
					HRef:        "http://localhost:22400/code-lists/opss-contributor",
					IsHierarchy: false,
				},
			},
		},
	},
}

//CrimeAccommodation recipe for transforming a given input to a Personnal Crime by Accommodation dataset
var CrimeAccommodation = Response{
	ID:     "171708A8-27CC-4ACD-B3D6-E0FE5131D9F8",
	Alias:  "CrimeAccommodation",
	Format: "v4",
	InputFiles: []file{
		{"Personal Crime by Accomodation v4"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "crime-accommodation",
			Editions:  []string{"time-series"},
			Title:     "Personal Crime by Accommodation",
			CodeLists: []CodeList{
				{
					ID:          "calendar-years",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/calendar-years",
					IsHierarchy: false,
				}, {
					ID:          "crime-accommodation-type",
					Name:        "accommodationtype",
					HRef:        "http://localhost:22400/code-lists/crime-accommodation-type",
					IsHierarchy: false,
				}, {
					ID:          "crime-age",
					Name:        "age",
					HRef:        "http://localhost:22400/code-lists/crime-age",
					IsHierarchy: false,
				}, {
					ID:          "england-and-wales-only",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/england-and-wales-only",
					IsHierarchy: false,
				}, {
					ID:          "crime-type",
					Name:        "crimetype",
					HRef:        "http://localhost:22400/code-lists/crime-type",
					IsHierarchy: false,
				}, {
					ID:          "crime-measurement-type",
					Name:        "measurementtype",
					HRef:        "http://localhost:22400/code-lists/crime-measurement-type",
					IsHierarchy: false,
				}, {
					ID:          "crime-sex",
					Name:        "sex",
					HRef:        "http://localhost:22400/code-lists/crime-sex",
					IsHierarchy: false,
				},
			},
		},
	},
}

//NPP recipe for transforming a given input to a npp dataset
var NPP = Response{
	ID:     "d5626f70-8fff-4538-88f1-764870f9e1ee",
	Alias:  "NPP",
	Format: "v4",
	InputFiles: []file{
		{"NPP_Experimental"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "npp",
			Editions:  []string{"2012-based", "2014-based", "2016-based", "2018-based", "2020-based", "2022-based"},
			Title:     "National Population Projection",
			CodeLists: []CodeList{
				{
					ID:          "calendar-years",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/calendar-years",
					IsHierarchy: false,
				}, {
					ID:          "uk-only",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/uk-only",
					IsHierarchy: false,
				}, {
					ID:          "npp-sex",
					Name:        "sex",
					HRef:        "http://localhost:22400/code-lists/npp-sex",
					IsHierarchy: false,
				}, {
					ID:          "npp-age",
					Name:        "age",
					HRef:        "http://localhost:22400/code-lists/npp-age",
					IsHierarchy: false,
				}, {
					ID:          "npp-projection-type",
					Name:        "projectiontype",
					HRef:        "http://localhost:22400/code-lists/npp-projection-type",
					IsHierarchy: false,
				}, {
					ID:          "npp-population-measure",
					Name:        "populationmeasure",
					HRef:        "http://localhost:22400/code-lists/npp-population-measure",
					IsHierarchy: false,
				},
			},
		},
	},
}

//CrimeOffences recipe for transforming a given input to a Crime offences dataset
var CrimeOffences = Response{
	ID:     "5d716747-0f45-4f55-a228-24e54a25bc57",
	Alias:  "crime-offences",
	Format: "v4",
	InputFiles: []file{
		{"Crime with home office, offence data"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "crime-with-home-office-offences",
			Editions:  []string{"time-series"},
			Title:     "Crime with home Office: Offences",
			CodeLists: []CodeList{
				{
					ID:          "financial-years",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/financial-years",
					IsHierarchy: false,
				}, {
					ID:          "police-force-geography",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/police-force-geography",
					IsHierarchy: false,
				}, {
					ID:          "crime-offence",
					Name:        "offence",
					HRef:        "http://localhost:22400/code-lists/crime-offence",
					IsHierarchy: true,
				},
			},
		},
	},
}

//Migration401AGQ recipe for transforming a given input to a Migration estimates AGQ dataset
var Migration401AGQ = Response{
	ID:     "b0756977-36ab-4d2f-8038-257c18233d8d",
	Alias:  "migration401agq",
	Format: "v4",
	InputFiles: []file{
		{"AGQcombinedYears"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "migration-401-5-year-age-groups",
			Editions:  []string{"time-series"},
			Title:     "Migration estimates by Citizen Group, Quinary Age Groups",
			CodeLists: []CodeList{
				{
					ID:          "financial-years",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/financial-years",
					IsHierarchy: false,
				}, {
					ID:          "uk-only",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/uk-only",
					IsHierarchy: false,
				}, {
					ID:          "migration-flow",
					Name:        "flow",
					HRef:        "http://localhost:22400/code-lists/migration-flow",
					IsHierarchy: false,
				}, {
					ID:          "migration-citizenship-group",
					Name:        "citizenshipgroup",
					HRef:        "http://localhost:22400/code-lists/migration-citizenship-group",
					IsHierarchy: false,
				}, {
					ID:          "migration-sex",
					Name:        "sex",
					HRef:        "http://localhost:22400/code-lists/migration-sex",
					IsHierarchy: false,
				}, {
					ID:          "migration-age",
					Name:        "age",
					HRef:        "http://localhost:22400/code-lists/migration-age",
					IsHierarchy: false,
				}, {
					ID:          "migration-country",
					Name:        "country",
					HRef:        "http://localhost:22400/code-lists/migration-country",
					IsHierarchy: false,
				},
			},
		},
	},
}

//Migration401AG1 recipe for transforming a given input to a Migration estimates alternative age group 1 dataset
var Migration401AG1 = Response{
	ID:     "82469D1C-7C5D-460E-B1C1-5E03B2193041",
	Alias:  "migration401ag1",
	Format: "v4",
	InputFiles: []file{
		{"AG1combinedYears"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "migration-401-alternative-age-groups-1",
			Editions:  []string{"time-series"},
			Title:     "Migration estimates by Citizen Group, Alternative Age Group 1",
			CodeLists: []CodeList{
				{
					ID:          "calendar-years",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/calendar-years",
					IsHierarchy: false,
				}, {
					ID:          "uk-only",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/uk-only",
					IsHierarchy: false,
				}, {
					ID:          "migration-flow",
					Name:        "flow",
					HRef:        "http://localhost:22400/code-lists/migration-flow",
					IsHierarchy: false,
				}, {
					ID:          "migration-citizenship-group",
					Name:        "citizenshipgroup",
					HRef:        "http://localhost:22400/code-lists/migration-citizenship-group",
					IsHierarchy: false,
				}, {
					ID:          "migration-sex",
					Name:        "sex",
					HRef:        "http://localhost:22400/code-lists/migration-sex",
					IsHierarchy: false,
				}, {
					ID:          "migration-age",
					Name:        "age",
					HRef:        "http://localhost:22400/code-lists/migration-age",
					IsHierarchy: false,
				}, {
					ID:          "migration-country",
					Name:        "country",
					HRef:        "http://localhost:22400/code-lists/migration-country",
					IsHierarchy: false,
				},
			},
		},
	},
}

//Migration401AG2 recipe for transforming a given input to a Migration estimates alternative age group 1 dataset
var Migration401AG2 = Response{
	ID:     "281145C5-AFE1-4B9D-811C-353C2FA2215A",
	Alias:  "migration401ag2",
	Format: "v4",
	InputFiles: []file{
		{"AG2combinedYears"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "migration-401-alternative-age-groups-2",
			Editions:  []string{"time-series"},
			Title:     "Migration estimates by Citizen Group, Alternative Age Groups 2",
			CodeLists: []CodeList{
				{
					ID:          "calendar-years",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/calendar-years",
					IsHierarchy: false,
				}, {
					ID:          "uk-only",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/uk-only",
					IsHierarchy: false,
				}, {
					ID:          "migration-flow",
					Name:        "flow",
					HRef:        "http://localhost:22400/code-lists/migration-flow",
					IsHierarchy: false,
				}, {
					ID:          "migration-citizenship-group",
					Name:        "citizenshipgroup",
					HRef:        "http://localhost:22400/code-lists/migration-citizenship-group",
					IsHierarchy: false,
				}, {
					ID:          "migration-sex",
					Name:        "sex",
					HRef:        "http://localhost:22400/code-lists/migration-sex",
					IsHierarchy: false,
				}, {
					ID:          "migration-age",
					Name:        "age",
					HRef:        "http://localhost:22400/code-lists/migration-age",
					IsHierarchy: false,
				}, {
					ID:          "migration-country",
					Name:        "country",
					HRef:        "http://localhost:22400/code-lists/migration-country",
					IsHierarchy: false,
				},
			},
		},
	},
}

//Migration402 recipe for transforming a given input to a Migration estimates 402 (by reason for migration) dataset
var Migration402 = Response{
	ID:     "89657893-13BB-4E79-8DDA-A38245782B69",
	Alias:  "migration402",
	Format: "v4",
	InputFiles: []file{
		{"migration402combinedYears"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "migration-402",
			Editions:  []string{"time-series"},
			Title:     "Migration estimates by Reason for Migration",
			CodeLists: []CodeList{
				{
					ID:          "calendar-years",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/calendar-years",
					IsHierarchy: false,
				}, {
					ID:          "uk-only",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/uk-only",
					IsHierarchy: false,
				}, {
					ID:          "migration-flow",
					Name:        "flow",
					HRef:        "http://localhost:22400/code-lists/migration-flow",
					IsHierarchy: false,
				}, {
					ID:          "migration-reason-for-migration",
					Name:        "reasonformigration",
					HRef:        "http://localhost:22400/code-lists/migration-reason-for-migration",
					IsHierarchy: false,
				}, {
					ID:          "migration-country",
					Name:        "country",
					HRef:        "http://localhost:22400/code-lists/migration-country",
					IsHierarchy: false,
				},
			},
		},
	},
}
