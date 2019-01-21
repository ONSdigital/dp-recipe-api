package recipe

//To add a new recipe:
// - add a recipeResponse{} document with an appropriate var name
// - add the var name to the list of Items
// - increment the Count and Total count in the RecipeList

//FullList of recipes available via this API
var FullList = List{
	Items: []Response{CPI, CPIH, MidYearPopEst, ASHE7Hours, ASHE7Earnings, ASHE8Hours, ASHE8Earnings, OPSSMembership, OPSSRates,
		CrimeAccommodation, CrimeOffences, Migration401AGQ, Migration401AG1, Migration401AG2, Migration402, WellbeingYearEnding,
		BuisInvestGFCG, BuisInvestCapitalFormation, NppPopulationNumbers, NppMortalityAssumptions, NppMigration, NppFertility,
		NppDeaths, NppBirths, NppCrossBorderRates, Trade, WellbeingLocalAuthority, OverseasTravelTourism, ASHE7and8, ASHE9and10, 
		Construction, UKBusinessIndustryGeography, LabourMarketStatistics, Suicides, LifeExpectancy, ParentsCountryOfBirth, DrugRelatedDeaths},
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
					ID:          "mmm-yy",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/mmm-yy",
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
					ID:          "admin-geography",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/admin-geography",
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
					ID:          "ashe-geography",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/ashe-geography",
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
					ID:          "ashe-geography",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/ashe-geography",
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
					ID:          "ashe-geography",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/ashe-geography",
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
					ID:          "ashe-geography",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/ashe-geography",
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
			Editions:  []string{"time-series", "2017-2018"},
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
					IsHierarchy: true,
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

// WellbeingYearEnding is wellbeing data with a time codelist of "Year Ending"
var WellbeingYearEnding = Response{
	ID:     "84E3A8F0-F48A-482F-8B55-6240EF717F2C",
	Alias:  "WellbeingYearEnding",
	Format: "v4",
	InputFiles: []file{
		{"wellbeingExcel"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "wellbeing-year-ending",
			Editions:  []string{"time-series"},
			Title:     "Well-being by year ending",
			CodeLists: []CodeList{
				{
					ID:          "year-ending",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/year-ending",
					IsHierarchy: false,
				}, {
					ID:          "wellbeing-geography",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/wellbeing-geography",
					IsHierarchy: false,
				}, {
					ID:          "wellbeing-estimate",
					Name:        "estimate",
					HRef:        "http://localhost:22400/code-lists/wellbeing-estimate",
					IsHierarchy: false,
				}, {
					ID:          "wellbeing-measureofwellbeing",
					Name:        "allmeasuresofwellbeing",
					HRef:        "http://localhost:22400/code-lists/wellbeing-measureofwellbeing",
					IsHierarchy: false,
				},
			},
		},
	},
}

// BuisInvestGFCG recipe for transforming a given SDMX input to a business investment 'Cross Classification of GFCF by Industry and Asset' dataset
var BuisInvestGFCG = Response{
	ID:     "05F4247E-CC85-4AFE-A3E6-1E1B0B1CC5A5",
	Alias:  "Business Investment GFCF, file 2200",
	Format: "v4",
	InputFiles: []file{
		{"Business Investment GFCF, file 2200"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "buisness-investment-cross-classification-gfcf",
			Editions:  []string{"time-series"},
			Title:     "Business Investment, Cross Classification of GFCF by Industry and Asset",
			CodeLists: []CodeList{
				{
					ID:          "yyyy",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/yyyy",
					IsHierarchy: false,
				}, {
					ID:          "uk-only",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/uk-only",
					IsHierarchy: false,
				}, {
					ID:          "business-investment-activity",
					Name:        "activity",
					HRef:        "http://localhost:22400/code-lists/business-investment-activity",
					IsHierarchy: true,
				}, {
					ID:          "business-investment-instrument-asset",
					Name:        "instrumentasset",
					HRef:        "http://localhost:22400/code-lists/business-investment-instrument-asset",
					IsHierarchy: false,
				}, {
					ID:          "business-investment-prices",
					Name:        "prices",
					HRef:        "http://localhost:22400/code-lists/business-investment-prices",
					IsHierarchy: false,
				},
			},
		},
	},
}

// BuisInvestCapitalFormation recipe for transforming a given SDMX input to a business investment 'Capital Formation' dataset
var BuisInvestCapitalFormation = Response{
	ID:     "2E52DCF1-2C31-4DCC-819A-184C7398F902",
	Alias:  "Business Investment Capital Formation, file 0302",
	Format: "v4",
	InputFiles: []file{
		{"Business Investment Capital Formation, file 0302"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "business-investment-capital-formation",
			Editions:  []string{"time-series"},
			Title:     "Business Investment Capital Formation",
			CodeLists: []CodeList{
				{
					ID:          "yyyy",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/yyyy",
					IsHierarchy: false,
				}, {
					ID:          "uk-only",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/uk-only",
					IsHierarchy: false,
				}, {
					ID:          "business-investment-activity",
					Name:        "activity",
					HRef:        "http://localhost:22400/code-lists/business-investment-activity",
					IsHierarchy: true,
				}, {
					ID:          "business-investment-instrument-asset",
					Name:        "instrumentasset",
					HRef:        "http://localhost:22400/code-lists/business-investment-instrument-asset",
					IsHierarchy: false,
				}, {
					ID:          "business-investment-prices",
					Name:        "prices",
					HRef:        "http://localhost:22400/code-lists/business-investment-prices",
					IsHierarchy: false,
				},
			},
		},
	},
}

// National Population Projections: PopulationNumbers
var NppPopulationNumbers = Response{
	ID:     "9BCE6F29-5FD6-438A-A4DC-92B697D61A33",
	Alias:  "National Population Projections - Population Numbers",
	Format: "v4",
	InputFiles: []file{
		{"National Population Projections - Population Numbers"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "national-population-projections-population",
			Editions:  []string{"time-series"},
			Title:     "National Population Projections: Population Numbers",
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
					ID:          "npp-all-projectiontype",
					Name:        "projectiontype",
					HRef:        "http://localhost:22400/code-lists/npp-all-projectiontype",
					IsHierarchy: false,
				}, {
					ID:          "npp-population-age",
					Name:        "age",
					HRef:        "http://localhost:22400/code-lists/npp-population-age",
					IsHierarchy: false,
				}, {
					ID:          "npp-all-sex",
					Name:        "sex",
					HRef:        "http://localhost:22400/code-lists/npp-all-sex",
					IsHierarchy: false,
				},
			},
		},
	},
}

// National Population Projections: Mortality Assumptions
var NppMortalityAssumptions = Response{
	ID:     "85D985D2-A54C-4289-B720-5A571ABD00C1",
	Alias:  "National Population Projections - Mortality Assumptions",
	Format: "v4",
	InputFiles: []file{
		{"National Population Projections - Mortality Assumptions"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "national-population-projections-mortality-assumptions",
			Editions:  []string{"time-series"},
			Title:     "National Population Projections: Mortality Assumptions",
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
					ID:          "npp-all-projectiontype",
					Name:        "projectiontype",
					HRef:        "http://localhost:22400/code-lists/npp-all-projectiontype",
					IsHierarchy: false,
				}, {
					ID:          "npp-mortalityandcrossborderrates-age",
					Name:        "age",
					HRef:        "http://localhost:22400/code-lists/npp-mortalityandcrossborderrates-age",
					IsHierarchy: false,
				}, {
					ID:          "npp-all-sex",
					Name:        "sex",
					HRef:        "http://localhost:22400/code-lists/npp-all-sex",
					IsHierarchy: false,
				},
			},
		},
	},
}

// National Population Projections: Migration
var NppMigration = Response{
	ID:     "0FB0634E-C722-4B84-A3D9-B00F6964742B",
	Alias:  "National Population Projections - Migration",
	Format: "v4",
	InputFiles: []file{
		{"National Population Projections - Migration"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "national-population-projections-migration",
			Editions:  []string{"time-series"},
			Title:     "National Population Projections: Migration",
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
					ID:          "npp-all-projectiontype",
					Name:        "projectiontype",
					HRef:        "http://localhost:22400/code-lists/npp-all-projectiontype",
					IsHierarchy: false,
				}, {
					ID:          "npp-migration-age",
					Name:        "age",
					HRef:        "http://localhost:22400/code-lists/npp-migration-age",
					IsHierarchy: false,
				}, {
					ID:          "npp-all-sex",
					Name:        "sex",
					HRef:        "http://localhost:22400/code-lists/npp-all-sex",
					IsHierarchy: false,
				}, {
					ID:          "npp-migration-populationmeasure",
					Name:        "populationmeasure",
					HRef:        "http://localhost:22400/code-lists/npp-migration-populationmeasure",
					IsHierarchy: false,
				},
			},
		},
	},
}

// National Population Projections: Fertility
var NppFertility = Response{
	ID:     "B92138FC-37F3-47AC-9C2B-E13C7D61A8AD",
	Alias:  "National Population Projections - Fertility",
	Format: "v4",
	InputFiles: []file{
		{"National Population Projections - Fertility"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "national-population-projections-fertility",
			Editions:  []string{"time-series"},
			Title:     "National Population Projections: Fertility",
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
					ID:          "npp-all-projectiontype",
					Name:        "projectiontype",
					HRef:        "http://localhost:22400/code-lists/npp-all-projectiontype",
					IsHierarchy: false,
				}, {
					ID:          "npp-birthsandfertility-age",
					Name:        "age",
					HRef:        "http://localhost:22400/code-lists/npp-birthsandfertility-age",
					IsHierarchy: false,
				}, {
					ID:          "npp-migration-populationmeasure",
					Name:        "populationmeasure",
					HRef:        "http://localhost:22400/code-lists/npp-migration-populationmeasure",
					IsHierarchy: false,
				}, {
					ID:          "npp-all-sex",
					Name:        "sex",
					HRef:        "http://localhost:22400/code-lists/npp-all-sex",
					IsHierarchy: false,
				},
			},
		},
	},
}

// National Population Projections: Deaths
var NppDeaths = Response{
	ID:     "999EB1B2-C779-4502-9F52-CCC99A1FED08",
	Alias:  "National Population Projections - Deaths",
	Format: "v4",
	InputFiles: []file{
		{"National Population Projections - Deaths"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "national-population-projections-deaths",
			Editions:  []string{"time-series"},
			Title:     "National Population Projections: Deaths",
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
					ID:          "npp-all-projectiontype",
					Name:        "projectiontype",
					HRef:        "http://localhost:22400/code-lists/npp-all-projectiontype",
					IsHierarchy: false,
				}, {
					ID:          "npp-deaths-age",
					Name:        "age",
					HRef:        "http://localhost:22400/code-lists/npp-deaths-age",
					IsHierarchy: false,
				}, {
					ID:          "npp-all-sex",
					Name:        "sex",
					HRef:        "http://localhost:22400/code-lists/npp-all-sex",
					IsHierarchy: false,
				},
			},
		},
	},
}

// National Population Projections: Births
var NppBirths = Response{
	ID:     "3B1172A6-9EA4-4AD6-8BDD-221794A18A5A",
	Alias:  "National Population Projections - Births",
	Format: "v4",
	InputFiles: []file{
		{"National Population Projections - Births"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "national-population-projections-births",
			Editions:  []string{"time-series"},
			Title:     "National Population Projections: Births",
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
					ID:          "npp-all-projectiontype",
					Name:        "projectiontype",
					HRef:        "http://localhost:22400/code-lists/npp-all-projectiontype",
					IsHierarchy: false,
				}, {
					ID:          "npp-birthsandfertility-age",
					Name:        "age",
					HRef:        "http://localhost:22400/code-lists/npp-birthsandfertility-age",
					IsHierarchy: false,
				}, {
					ID:          "npp-all-sex",
					Name:        "sex",
					HRef:        "http://localhost:22400/code-lists/npp-all-sex",
					IsHierarchy: false,
				},
			},
		},
	},
}

// National Population Projections: Cross Border Rates
var NppCrossBorderRates = Response{
	ID:     "182DF2D1-9E9D-475C-949C-365CAC6A9834",
	Alias:  "National Population Projections - Cross Border Rates",
	Format: "v4",
	InputFiles: []file{
		{"National Population Projections - Cross Border Rates"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "national-population-projections-cross-border-rates",
			Editions:  []string{"time-series"},
			Title:     "National Population Projections: Cross Border Rates",
			CodeLists: []CodeList{
				{
					ID:          "calendar-years",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/financial-years",
					IsHierarchy: false,
				}, {
					ID:          "uk-only",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/uk-only",
					IsHierarchy: false,
				}, {
					ID:          "npp-all-sex",
					Name:        "sex",
					HRef:        "http://localhost:22400/code-lists/npp-all-sex",
					IsHierarchy: false,
				}, {
					ID:          "npp-all-projectiontype",
					Name:        "projectiontype",
					HRef:        "http://localhost:22400/code-lists/npp-all-projectiontype",
					IsHierarchy: false,
				}, {
					ID:          "npp-birthsandfertility-age",
					Name:        "age",
					HRef:        "http://localhost:22400/code-lists/npp-birthsandfertility-age",
					IsHierarchy: false,
				}, {
					ID:          "npp-border",
					Name:        "border",
					HRef:        "http://localhost:22400/code-lists/npp-border",
					IsHierarchy: false,
				},
			},
		},
	},
}

// Trade recipe
var Trade = Response{
	ID:     "9D2BE73A-F3E4-46FD-9495-8AA7E7C7DAE9",
	Alias:  "Trade",
	Format: "v4",
	InputFiles: []file{
		{"Trade"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "trade",
			Editions:  []string{"time-series"},
			Title:     "UK Trade in Goods by Country",
			CodeLists: []CodeList{
				{
					ID:          "mmm-yy",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/mmm-yy",
					IsHierarchy: false,
				}, {
					ID:          "uk-only",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/uk-only",
					IsHierarchy: false,
				}, {
					ID:          "trade-commodity",
					Name:        "commodity",
					HRef:        "http://localhost:22400/code-lists/trade-commodity",
					IsHierarchy: true,
				}, {
					ID:          "trade-country",
					Name:        "country",
					HRef:        "http://localhost:22400/code-lists/trade-country",
					IsHierarchy: true,
				}, {
					ID:          "trade-direction",
					Name:        "direction",
					HRef:        "http://localhost:22400/code-lists/trade-direction",
					IsHierarchy: false,
				},
			},
		},
	},
}

// WellbeingLocalAuthority is wellbeing data with a time codelist of "yyyy-yy"
var WellbeingLocalAuthority = Response{
	ID:     "e8238bad-e248-4008-92cf-d29c087741b7",
	Alias:  "Wellbeing Local Authority",
	Format: "v4",
	InputFiles: []file{
		{"wellbeingExcel"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "wellbeing-local-authority",
			Editions:  []string{"time-series"},
			Title:     "Well-being by local authority",
			CodeLists: []CodeList{
				{
					ID:          "yyyy-yy",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/yyyy-yy",
					IsHierarchy: false,
				}, {
					ID:          "admin-geography",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/admin-geography",
					IsHierarchy: true,
				}, {
					ID:          "wellbeing-estimate",
					Name:        "estimate",
					HRef:        "http://localhost:22400/code-lists/wellbeing-estimate",
					IsHierarchy: false,
				}, {
					ID:          "wellbeing-measureofwellbeing",
					Name:        "allmeasuresofwellbeing",
					HRef:        "http://localhost:22400/code-lists/wellbeing-measureofwellbeing",
					IsHierarchy: false,
				},
			},
		},
	},
}

	
// Overseas travel and tourism recipe
var OverseasTravelTourism = Response{
	ID:     "35b7fe99-b7db-4237-9af0-f8c2c6c0935c",
	Alias:  "Overseas Travel and Tourism",
	Format: "v4",
	InputFiles: []file{
		{"Overseas v4"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "overseas-travel-and-tourism",
			Editions:  []string{"time-series"},
			Title:     "Overseas Travel and Tourism",
			CodeLists: []CodeList{
				{
					ID:          "mmm-yy",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/mmm-yy",
					IsHierarchy: false,
				}, {
					ID:          "ott-countries-in-groups",
					Name:        "regions",
					HRef:        "http://localhost:22400/code-lists/ott-countries-in-groups",
					IsHierarchy: true,
				}, {
					ID:          "ott-residence",
					Name:        "residence",
					HRef:        "http://localhost:22400/code-lists/ott-residence",
					IsHierarchy: false,
				},{
					ID:          "ott-purpose",
					Name:        "purpose",
					HRef:        "http://localhost:22400/code-lists/ott-purpose",
					IsHierarchy: false,
				}, {
					ID:          "ott-cost",
					Name:        "cost",
					HRef:        "http://localhost:22400/code-lists/ott-cost",
					IsHierarchy: false,
				},
			},
		},
	},
}

// ASHE tables 7 and 8 combined
var ASHE7and8 = Response{
	ID:     "8dd6cc0f-54d5-42e2-b186-4c00b00dcb02",
	Alias:  "ASHE Tables 7 and 8",
	Format: "v4",
	InputFiles: []file{
		{"ASHE78Excel"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "ashe-tables-7-and-8",
			Editions:  []string{"time-series"},
			Title:     "ASHE Tables 7 and 8",
			CodeLists: []CodeList{
				{
					ID:          "calendar-years",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/calendar-years",
					IsHierarchy: false,
				}, {
					ID:          "admin-geography",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/admin-geography",
					IsHierarchy: true,
				}, {
					ID:          "ashe-statistics",
					Name:        "statistics",
					HRef:        "http://localhost:22400/code-lists/ashe-statistics",
					IsHierarchy: false,
				}, {
					ID:          "ashe-sex",
					Name:        "sex",
					HRef:        "http://localhost:22400/code-lists/ashe-sex",
					IsHierarchy: false,
				},{
					ID:          "ashe-working-pattern",
					Name:        "workingpattern",
					HRef:        "http://localhost:22400/code-lists/ashe-working-pattern",
					IsHierarchy: false,
				}, {
					ID:          "ashe-hours-and-earnings",
					Name:        "hoursandearnings",
					HRef:        "http://localhost:22400/code-lists/ashe-hours-and-earnings",
					IsHierarchy: false,
				}, {
					ID:          "ashe-workplace-or-residence",
					Name:        "workplaceorresidence",
					HRef:        "http://localhost:22400/code-lists/ashe-workplace-or-residence",
					IsHierarchy: false,
				},
			},
		},
	},
}

// ASHE tables 9 and 10 combined
var ASHE9and10 = Response{
	ID:     "e8b95232-1135-4bc3-a9f4-ae6551669dba",
	Alias:  "ASHE Tables 9 and 10",
	Format: "v4",
	InputFiles: []file{
		{"ASHE910Excel"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "ashe-tables-9-and-10",
			Editions:  []string{"time-series"},
			Title:     "ASHE Tables 9 and 10",
			CodeLists: []CodeList{
				{
					ID:          "calendar-years",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/calendar-years",
					IsHierarchy: false,
				}, {
					ID:          "parliamentary-geography",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/parliamentary-geography",
					IsHierarchy: true,
				}, {
					ID:          "ashe-statistics",
					Name:        "statistics",
					HRef:        "http://localhost:22400/code-lists/ashe-statistics",
					IsHierarchy: false,
				}, {
					ID:          "ashe-sex",
					Name:        "sex",
					HRef:        "http://localhost:22400/code-lists/ashe-sex",
					IsHierarchy: false,
				},{
					ID:          "ashe-working-pattern",
					Name:        "workingpattern",
					HRef:        "http://localhost:22400/code-lists/ashe-working-pattern",
					IsHierarchy: false,
				}, {
					ID:          "ashe-hours-and-earnings",
					Name:        "hoursandearnings",
					HRef:        "http://localhost:22400/code-lists/ashe-hours-and-earnings",
					IsHierarchy: false,
				}, {
					ID:          "ashe-workplace-or-residence",
					Name:        "workplaceorresidence",
					HRef:        "http://localhost:22400/code-lists/ashe-workplace-or-residence",
					IsHierarchy: false,
				},
			},
		},
	},
}

// Output in the construction industry
var Construction = Response{
	ID:     "5c1aab52-7538-4105-8305-c63a4ba37cab",
	Alias:  "Output in the Construction Industry",
	Format: "v4",
	InputFiles: []file{
		{"Construction"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "output-in-the-construction-industry",
			Editions:  []string{"time-series"},
			Title:     "Output in the Construction Industry",
			CodeLists: []CodeList{
				{
					ID:          "mmm-yy",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/mmm-yy",
					IsHierarchy: false,
				}, {
					ID:          "type-of-prices",
					Name:        "prices",
					HRef:        "http://localhost:22400/code-lists/type-of-prices",
					IsHierarchy: false,
				}, {
					ID:          "seasonal-adjustment",
					Name:        "seasonaladjustment",
					HRef:        "http://localhost:22400/code-lists/seasonal-adjustment",
					IsHierarchy: false,
				},{
					ID:          "construction-series-type",
					Name:        "seriestype",
					HRef:        "http://localhost:22400/code-lists/construction-series-type",
					IsHierarchy: false,
				}, {
					ID:          "construction-classifications",
					Name:        "typeofwork",
					HRef:        "http://localhost:22400/code-lists/construction-classifications",
					IsHierarchy: false,
				}, 
			},
		},
	},
}

// UK business - broad industry group by geography
var UKBusinessIndustryGeography = Response{
	ID:     "a6e80e1e-2011-48ab-8c1d-35a7833b38db",
	Alias:  "UK Business by Broad Industry Group",
	Format: "v4",
	InputFiles: []file{
		{"UKbizIndustry"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "uk-business-by-industry-group",
			Editions:  []string{"time-series"},
			Title:     "UK Business by Broad Industry Group",
			CodeLists: []CodeList{
				{
					ID:          "calendar-years",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/calendar-years",
					IsHierarchy: false,
				}, {
					ID:          "admin-geography",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/admin-geography",
					IsHierarchy: true,
				}, {
					ID:          "uk-business-broad-industry-group",
					Name:        "broadindustrygroup",
					HRef:        "http://localhost:22400/code-lists/uk-business-broad-industry-group",
					IsHierarchy: false,
				}, {
					ID:          "uk-business-unit",
					Name:        "unit",
					HRef:        "http://localhost:22400/code-lists/uk-business-unit",
					IsHierarchy: false,
				},
			},
		},
	},
}

// Labour Market  
var LabourMarketStatistics = Response{
	ID:     "daf08e97-0a21-4800-9a2f-d7c90c88519b",
	Alias:  "Labour Market",
	Format: "v4",
	InputFiles: []file{
		{"LMS"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "labour-market",
			Editions:  []string{"time-series"},
			Title:     "Labour Market",
			CodeLists: []CodeList{
				{
					ID:          "mmm-mmm-yyyy",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/mmm-mmm-yyyy",
					IsHierarchy: false,
				}, {
					ID:          "lms-unit-of-measure",
					Name:        "unitofmeasure",
					HRef:        "http://localhost:22400/code-lists/lms-unit-of-measure",
					IsHierarchy: false,
				}, {
					ID:          "lms-economic-status",
					Name:        "economicstatus",
					HRef:        "http://localhost:22400/code-lists/lms-economic-status",
					IsHierarchy: false,
				}, {
					ID:          "lms-age-bracket",
					Name:        "age",
					HRef:        "http://localhost:22400/code-lists/lms-age-bracket",
					IsHierarchy: false,
				}, {
					ID:          "ashe-sex",
					Name:        "sex",
					HRef:        "http://localhost:22400/code-lists/ashe-sex",
					IsHierarchy: false,
				}, {
					ID:          "seasonal-adjustment",
					Name:        "seasonaladjustment",
					HRef:        "http://localhost:22400/code-lists/seasonal-adjustment",
					IsHierarchy: false,
				}, 
			},
		},
	},
}

// Suicides  
var Suicides = Response{
	ID:     "f78ee223-ac49-450d-b2ae-ee8efeb53b6a",
	Alias:  "Suicides",
	Format: "v4",
	InputFiles: []file{
		{"Suicides"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "suicides-in-the-uk",
			Editions:  []string{"time-series"},
			Title:     "Suicides in the UK",
			CodeLists: []CodeList{
				{
					ID:          "calendar-years",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/calendar-years",
					IsHierarchy: false,
				}, {
					ID:          "admin-geography",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/admin-geography",
					IsHierarchy: true,
				}, 
			},
		},
	},
}

// Life Expectancy  
var LifeExpectancy = Response{
	ID:     "c1a25f33-506e-405e-a9a2-9ebd85b46e6d",
	Alias:  "Life Expecancy",
	Format: "v4",
	InputFiles: []file{
		{"LifeExpectancy"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "life-expectancy-local-authority",
			Editions:  []string{"time-series"},
			Title:     "Life Expectancy in the UK",
			CodeLists: []CodeList{
				{
					ID:          "two-year-intervals",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/two-year-intervals",
					IsHierarchy: false,
				}, {
					ID:          "admin-geography",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/admin-geography",
					IsHierarchy: true,
				}, {
					ID:          "birth-cohort",
					Name:        "birthcohort",
					HRef:        "http://localhost:22400/code-lists/birth-cohort",
					IsHierarchy: false,
				},{
					ID:          "life-expectancy-variable",
					Name:        "lifeexpectancyvariable",
					HRef:        "http://localhost:22400/code-lists/life-expectancy-variable",
					IsHierarchy: false,
				},
			},
		},
	},
}

// Parents Country of Birth 
var ParentsCountryOfBirth = Response{
	ID:     "916475c2-98a8-4d86-9ff5-a6d1c1d4688d",
	Alias:  "Parents Country of Birth",
	Format: "v4",
	InputFiles: []file{
		{"ParentsBirth"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "parents-country-of-birth",
			Editions:  []string{"time-series"},
			Title:     "Parents' Country of Birth",
			CodeLists: []CodeList{
				{
					ID:          "calendar-years",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/calendar-years",
					IsHierarchy: false,
				}, {
					ID:          "admin-geography",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/admin-geography",
					IsHierarchy: true,
				}, {
					ID:          "parents-country-birth",
					Name:        "parentscountryofbirth",
					HRef:        "http://localhost:22400/code-lists/parents-country-birth",
					IsHierarchy: false,
				},{
					ID:          "type-of-number",
					Name:        "typeofnumber",
					HRef:        "http://localhost:22400/code-lists/type-of-number",
					IsHierarchy: false,
				},
			},
		},
	},
}

// Drug-related deaths by local authority, England and Wales
var DrugRelatedDeaths = Response{
	ID:     "b5923749-b3ac-41a0-9720-cefe94d765dc",
	Alias:  "Drug-related deaths by local authority",
	Format: "v4",
	InputFiles: []file{
		{"DrugRelatedDeaths"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "drug-related-deaths-local-authority",
			Editions:  []string{"time-series"},
			Title:     "Drug-related deaths by local authority, England and Wales",
			CodeLists: []CodeList{
				{
					ID:          "two-year-intervals",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/two-year-intervals",
					IsHierarchy: false,
				}, {
					ID:          "admin-geography",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/admin-geography",
					IsHierarchy: true,
				}, {
					ID:          "drug-deaths-mortality",
					Name:        "mortality",
					HRef:        "http://localhost:22400/code-lists/drug-deaths-mortality",
					IsHierarchy: false,
				},{
					ID:          "ashe-sex",
					Name:        "sex",
					HRef:        "http://localhost:22400/code-lists/ashe-sex",
					IsHierarchy: false,
				},{
					ID:          "drug-deaths-type-of-death",
					Name:        "typeofdeath",
					HRef:        "http://localhost:22400/code-lists/drug-deaths-type-of-death",
					IsHierarchy: false,
				},
			},
		},
	},
}