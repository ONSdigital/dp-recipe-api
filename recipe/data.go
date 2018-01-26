package recipe

//To add a new recipe:
// - add a recipeResponse{} document with an appropriate var name
// - add the var name to the list of Items
// - increment the Count and Total count in the RecipeList

//FullList of recipes available via this API
var FullList = List{
	Items:        []Response{CPI, CPIH, MidYearPopEst, ASHE7Hours, ASHE7Earnings, ASHE8Hours, ASHE8Earnings, OPSSMembership, OPSSRates},
	Count:        8,
	TotalCount:   8,
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
			Editions:  []string{"time-series"},
			Title:     "UK consumer price inflation",
			CodeLists: []CodeList{
				{
					ID:   "64d384f1-ea3b-445c-8fb8-aa453f96e58a",
					Name: "time",
					HRef: "http://localhost:22400/code-lists/64d384f1-ea3b-445c-8fb8-aa453f96e58a",
					IsHierarchy:false,
				}, {
					ID:   "65107A9F-7DA3-4B41-A410-6F6D9FBD68C3",
					Name: "geography",
					HRef: "http://localhost:22400/code-lists/65107A9F-7DA3-4B41-A410-6F6D9FBD68C3",
					IsHierarchy:false,
				}, {
					ID:   "e44de4c4-d39e-4e2f-942b-3ca10584d078",
					Name: "aggregate",
					HRef: "http://localhost:22400/code-lists/e44de4c4-d39e-4e2f-942b-3ca10584d078",
					IsHierarchy:true,
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
					ID:   "time",
					Name: "time",
					HRef: "http://localhost:22400/code-lists/time",
					IsHierarchy:false,
				}, {
					ID:   "uk-only",
					Name: "geography",
					HRef: "http://localhost:22400/code-lists/uk-only",
					IsHierarchy:false,
				}, {
					ID:   "cpih1dim1aggid",
					Name: "aggregate",
					HRef: "http://localhost:22400/code-lists/cpih1dim1aggid",
					IsHierarchy:true,
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
					ID:   "calendar-years",
					Name: "time",
					HRef: "http://localhost:22400/code-lists/calendar-years",
					IsHierarchy:false,
				}, {
					ID:   "mid-year-pop-geography",
					Name: "geography",
					HRef: "http://localhost:22400/code-lists/mid-year-pop-geography",
					IsHierarchy:true,
				}, {
					ID:   "mid-year-pop-sex",
					Name: "sex",
					HRef: "http://localhost:22400/code-lists/mid-year-pop-sex",
					IsHierarchy:false,
				}, {
					ID:   "mid-year-pop-age",
					Name: "age",
					HRef: "http://localhost:22400/code-lists/mid-year-pop-age",
					IsHierarchy:false,
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
			CodeLists: []CodeList{{ID: "calendar-years", Name: "time", HRef: "http://localhost:22400/code-lists/calendar-years"},
				{ID: "ashe-sex", Name: "sex", HRef: "http://localhost:22400/code-lists/ashe-sex"},
				{ID: "ashe-working-pattern", Name: "workingpattern", HRef: "http://localhost:22400/code-lists/ashe-working-pattern"},
				{ID: "ashe-hours", Name: "hours", HRef: "http://localhost:22400/code-lists/ashe-hours"},
				{ID: "ashe-statistics", Name: "statistics", HRef: "http://localhost:22400/code-lists/ashe-statistics"},
				{ID: "ashe-table-7-geography", Name: "geography", HRef: "http://localhost:22400/code-lists/ashe-table-7-geography"}},
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
			CodeLists: []CodeList{{ID: "calendar-years", Name: "time", HRef: "http://localhost:22400/code-lists/calendar-years"},
				{ID: "ashe-sex", Name: "sex", HRef: "http://localhost:22400/code-lists/ashe-sex"},
				{ID: "ashe-working-pattern", Name: "workingpattern", HRef: "http://localhost:22400/code-lists/ashe-working-pattern"},
				{ID: "ashe-earnings", Name: "earnings", HRef: "http://localhost:22400/code-lists/ashe-earnings"},
				{ID: "ashe-statistics", Name: "statistics", HRef: "http://localhost:22400/code-lists/ashe-statistics"},
				{ID: "ashe-table-7-geography", Name: "geography", HRef: "http://localhost:22400/code-lists/ashe-table-7-geography"}},
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
			CodeLists: []CodeList{{ID: "calendar-years", Name: "time", HRef: "http://localhost:22400/code-lists/calendar-years"},
				{ID: "ashe-sex", Name: "sex", HRef: "http://localhost:22400/code-lists/ashe-sex"},
				{ID: "ashe-working-pattern", Name: "workingpattern", HRef: "http://localhost:22400/code-lists/ashe-working-pattern"},
				{ID: "ashe-hours", Name: "hours", HRef: "http://localhost:22400/code-lists/ashe-hours"},
				{ID: "ashe-statistics", Name: "statistics", HRef: "http://localhost:22400/code-lists/ashe-statistics"},
				{ID: "ashe-table-8-geography", Name: "geography", HRef: "http://localhost:22400/code-lists/ashe-table-8-geography"}},
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
			CodeLists: []CodeList{{ID: "calendar-years", Name: "time", HRef: "http://localhost:22400/code-lists/calendar-years"},
				{ID: "ashe-sex", Name: "sex", HRef: "http://localhost:22400/code-lists/ashe-sex"},
				{ID: "ashe-working-pattern", Name: "workingpattern", HRef: "http://localhost:22400/code-lists/ashe-working-pattern"},
				{ID: "ashe-earnings", Name: "earnings", HRef: "http://localhost:22400/code-lists/ashe-earnings"},
				{ID: "ashe-statistics", Name: "statistics", HRef: "http://localhost:22400/code-lists/ashe-statistics"},
				{ID: "ashe-table-8-geography", Name: "geography", HRef: "http://localhost:22400/code-lists/ashe-table-8-geography"}},
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
			Editions:  []string{"time-series"},
			Title:     "Occupational Pension Schemes Survey, Membership",
			CodeLists: []CodeList{{ID: "calendar-years", Name: "time", HRef: "http://localhost:22400/code-lists/calendar-years"},
				{ID: "uk-only", Name: "geography", HRef: "http://localhost:22400/code-lists/uk-only"},
				{ID: "opss-scheme-membership-sizeband", Name: "schememembershipsizeband", HRef: "http://localhost:22400/code-lists/opss-scheme-membership-sizeband"},
				{ID: "opss-status", Name: "status", HRef: "http://localhost:22400/code-lists/opss-status"},
				{ID: "opss-benefit-type", Name: "benefittype", HRef: "http://localhost:22400/code-lists/opss-benefit-type"},
				{ID: "opss-public-private-sector", Name: "publicprivatesector", HRef: "http://localhost:22400/code-lists/opss-public-private-sector"},
				{ID: "opss-membership-type", Name: "membershiptype", HRef: "http://localhost:22400/code-lists/opss-membership-type"}},
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
			Editions:  []string{"time-series"},
			Title:     "Occupational Pension Schemes Survey, Rates",
			CodeLists: []CodeList{{ID: "calendar-years", Name: "time", HRef: "http://localhost:22400/code-lists/calendar-years"},
				{ID: "uk-only", Name: "geography", HRef: "http://localhost:22400/code-lists/uk-only"},
				{ID: "opss-scheme-membership-sizeband", Name: "schememembershipsizeband", HRef: "http://localhost:22400/code-lists/opss-scheme-membership-sizeband"},
				{ID: "opss-status", Name: "status", HRef: "http://localhost:22400/code-lists/opss-status"},
				{ID: "opss-benefit-type", Name: "benefittype", HRef: "http://localhost:22400/code-lists/opss-benefit-type"},
				{ID: "opss-contributor", Name: "contributor", HRef: "http://localhost:22400/code-lists/opss-contributor"}},
		},
	},
}
