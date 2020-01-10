package recipe

//To add a new recipe:
// - add a recipeResponse{} document with an appropriate var name
// - add the var name to the list of Items
// - increment the Count and Total count in the RecipeList

//FullList of recipes available via this API
var FullList = List{
	Items: []Response{CPI, CPIH, Trade, MidYearPopEst, MidYearPopEstPCON, MidYearPopEstCCG, ASHE3, ASHE5, ASHE7Hours, ASHE7Earnings, 
		ASHE8Hours, ASHE8Earnings, ASHE7and8, ASHE9and10, ASHE11and12, ASHE20, ASHE25, ASHE26, ASHE27and28, WellbeingYearEnding, WellbeingLocalAuthority, 
		WellbeingQuarterly, WellbeingChildrens, NppPopulationNumbers, NppMortalityAssumptions, NppMigration, NppFertility, NppDeaths, NppBirths, NppCrossBorderRates, 
		QuarterlyDiffusion, MonthlyDiffusion, ReportingBehaviour, ReportingBehaviour2, AgeingPopProj, AgeingSingleHouseholds, AgeingSexRatios, 
		AgeingNetFlows, AgeingProspectiveMeasures, LabourMarketStatistics, LMSEconomicByAge, LMSWorkType, LMSActualHoursWork, LMSClaimantCount, 
		LMSJobseekersByAgeDuration, LMSEconomicInactivity, LMSJobsByIndustry, AWE, AWEIndex, CancerRegEng, CancerRegRegions, RegionalGDPYear, 
		RegionalGDPQuarter, TaxBenefitStats, GenerationalIncome, HousePrices, PrivateHousingRentalPrices, OPSSMembership, OPSSRates, CrimeAccommodation, 
		CrimeOffences, KnifeCrime, InternalMigrationLA, Migration401AGQ, Migration401AG1, Migration401AG2, Migration402, BuisInvestGFCG, 
		BuisInvestCapitalFormation, UKBusinessIndustryGeography, OverseasTravelTourism, Construction, FamiliesAndHouseholds, ParentsCountryOfBirth, LifeExpectancy, 
		ChildMortality, Suicides, DrugRelatedDeaths, MonthlyDeaths, SexualOrientation, Census1961, Census1961SH01, Census1961SH02, Census1961SH03, Census1961SH04, 
		Census1961SH05, Census1961SH06, Census1961SH07, Census1961SH08, Census1961SH09, Census1961SH10, Census1961SH11, Census1961SH12, Census1961SH13, 
		Census1961SH14, Census1961SH15, Census1961ST01, Census1961ST02, Census1961ST03, Census1961ST04, Census1961ST05, Census1961ST06, Census1961ST07, 
		Census1961ST08, Census1961ST09,Census1961SC11, Census1961SC13, Census1961SC22},
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
			Editions:  []string{"2012", "2013", "2014", "2015", "2016", "2017", "2018", "2019", "time-series"},
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

//MidYearPopEst recipe for transforming a given input to a mid year population estimate dataset using the parliamentary geography
var MidYearPopEstPCON = Response{
	ID:     "b83ffdaf-c5fa-451e-bbe0-6137c9f90b5a",
	Alias:  "Mid-year Population Estimates by Parliamentary Constituencies",
	Format: "v4",
	InputFiles: []file{
		{"Mid-year Population Estimates Parliamentary Constituencies v4"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "mid-year-pop-est-parliamentary-constituencies",
			Editions:  []string{"time-series"},
			Title:     "Population Estimates for UK, England and Wales, Scotland and Northern Ireland by Parliamentary Constituencies",
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

//MidYearPopEst recipe for transforming a given input to a mid year population estimate dataset using the CCG geography
var MidYearPopEstCCG = Response{
	ID:     "345f39e9-ef90-415e-b73e-a68a826f41cc",
	Alias:  "Mid-year Population Estimates by Clinical Commissioning Group",
	Format: "v4",
	InputFiles: []file{
		{"Mid-year Population Estimates Clinical Commissioning Group v4"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "mid-year-pop-est-ccg",
			Editions:  []string{"time-series"},
			Title:     "Population Estimates for England by Clinical Commissioning Group",
			CodeLists: []CodeList{
				{
					ID:          "calendar-years",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/calendar-years",
					IsHierarchy: false,
				}, {
					ID:          "ccg-geography",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/ccg-geography",
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

// ASHE tables 3
var ASHE3 = Response{
	ID:     "55c24ef5-82fe-4752-9455-a119bac39293",
	Alias:  "ASHE Tables 3",
	Format: "v4",
	InputFiles: []file{
		{"ASHE3"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "ashe-tables-3",
			Editions:  []string{"time-series"},
			Title:     "ASHE Tables 3",
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
					IsHierarchy: false,
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
				}, {
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
					ID:          "soc",
					Name:        "standardoccupationalclassification",
					HRef:        "http://localhost:22400/code-lists/soc",
					IsHierarchy: true,
				},
			},
		},
	},
}

// ASHE table 5
var ASHE5 = Response{
	ID:     "edf7d98c-fd59-4901-a813-87d6aed077d0",
	Alias:  "ASHE Table 5",
	Format: "v4",
	InputFiles: []file{
		{"ASHE5Excel"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "ashe-table-5",
			Editions:  []string{"time-series"},
			Title:     "ASHE Table 5",
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
				}, {
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
					ID:          "sic",
					Name:        "standardindustrialclassification",
					HRef:        "http://localhost:22400/code-lists/sic",
					IsHierarchy: true,
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
			Editions:  []string{"2012", "2013", "2014", "2015", "2016", "2017", "2018", "2019", "2020", "time-series"},
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

// ASHE tables 11 and 12 combined
var ASHE11and12 = Response{
	ID:     "5d3a275f-0ad3-4181-93a5-332f936354f9",
	Alias:  "ASHE Tables 11 and 12",
	Format: "v4",
	InputFiles: []file{
		{"ASHE1112Excel"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "ashe-tables-11-and-12",
			Editions:  []string{"time-series"},
			Title:     "ASHE Tables 11 and 12",
			CodeLists: []CodeList{
				{
					ID:          "calendar-years",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/calendar-years",
					IsHierarchy: false,
				}, {
					ID:          "travel-to-work-geography",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/travel-to-work-geography",
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
				}, {
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

// ASHE tables 20
var ASHE20 = Response{
	ID:     "70cab339-c01a-477a-8e4c-96a261e3595e",
	Alias:  "ASHE Tables 20",
	Format: "v4",
	InputFiles: []file{
		{"ASHE20"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "ashe-tables-20",
			Editions:  []string{"time-series"},
			Title:     "ASHE Tables 20",
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
					ID:          "ashe-statistics",
					Name:        "statistics",
					HRef:        "http://localhost:22400/code-lists/ashe-statistics",
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
					ID:          "ashe-hours-and-earnings",
					Name:        "hoursandearnings",
					HRef:        "http://localhost:22400/code-lists/ashe-hours-and-earnings",
					IsHierarchy: false,
				}, {
					ID:          "soc",
					Name:        "standardoccupationalclassification",
					HRef:        "http://localhost:22400/code-lists/soc",
					IsHierarchy: true,
				}, {
					ID:          "age-groups",
					Name:        "agegroups",
					HRef:        "http://localhost:22400/code-lists/age-groups",
					IsHierarchy: false,
				},
			},
		},
	},
}

// ASHE tables 25
var ASHE25 = Response{
	ID:     "3669408a-6705-445c-a1a4-a4c5c75a0346",
	Alias:  "ASHE Tables 25",
	Format: "v4",
	InputFiles: []file{
		{"ASHE25"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "ashe-tables-25",
			Editions:  []string{"time-series"},
			Title:     "ASHE Tables 25",
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
					IsHierarchy: false,
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
					ID:          "sector",
					Name:        "sectortype",
					HRef:        "http://localhost:22400/code-lists/sector",
					IsHierarchy: false,
				},
			},
		},
	},
}

// ASHE tables 26
var ASHE26 = Response{
	ID:     "e7fc556a-a71b-44a6-9843-c38e6d4b4e62",
	Alias:  "ASHE Tables 26",
	Format: "v4",
	InputFiles: []file{
		{"ASHE26"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "ashe-tables-26",
			Editions:  []string{"time-series"},
			Title:     "ASHE Tables 26",
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
					ID:          "ashe-statistics",
					Name:        "statistics",
					HRef:        "http://localhost:22400/code-lists/ashe-statistics",
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
					ID:          "ashe-hours-and-earnings",
					Name:        "hoursandearnings",
					HRef:        "http://localhost:22400/code-lists/ashe-hours-and-earnings",
					IsHierarchy: false,
				}, 
			},
		},
	},
}

// ASHE tables 27 and 28 combined
var ASHE27and28 = Response{
	ID:     "10213937-4076-4fd9-ba1a-506b2206a71f",
	Alias:  "ASHE Tables 27 and 28",
	Format: "v4",
	InputFiles: []file{
		{"ASHE27-28"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "ashe-tables-27-and-28",
			Editions:  []string{"time-series"},
			Title:     "ASHE Tables 27 and 28",
			CodeLists: []CodeList{
				{
					ID:          "calendar-years",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/calendar-years",
					IsHierarchy: false,
				}, {
					ID:          "local-enterprise-partnership-geography",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/local-enterprise-partnership-geography",
					IsHierarchy: false,
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
				}, {
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
					ID:          "admin-geography",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/admin-geography",
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

// Wellbeing quarterly is wellbeing data with a time codelist of "yyyy-qq"
var WellbeingQuarterly = Response{
	ID:     "cc738a86-e14f-4181-93c4-9aab2dac467a",
	Alias:  "Wellbeing Quarterly",
	Format: "v4",
	InputFiles: []file{
		{"wellbeingquarterly"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "wellbeing-quarterly",
			Editions:  []string{"time-series"},
			Title:     "Well-being by quarters",
			CodeLists: []CodeList{
				{
					ID:          "yyyy-qq",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/yyyy-qq",
					IsHierarchy: false,
				}, {
					ID:          "uk-only",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/uk-only",
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

// Children's Well-being Measures
var WellbeingChildrens = Response{
	ID:     "1a56119b-584b-4455-aa2c-81635d1aec56",
	Alias:  "Children's Well-being Measures",
	Format: "v4",
	InputFiles: []file{
		{"WellbeingChildrens"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "childrens-wellbeing",
			Editions:  []string{"time-series"},
			Title:     "Children's Well-being Measures",
			CodeLists: []CodeList{
				{
					ID:          "calendar-years",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/calendar-years",
					IsHierarchy: false,
				}, {
					ID:          "countries",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/countries",
					IsHierarchy: false,
				}, {
					ID:          "children-sex",
					Name:        "sex",
					HRef:        "http://localhost:22400/code-lists/children-sex",
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
					ID:          "npp-age",
					Name:        "age",
					HRef:        "http://localhost:22400/code-lists/npp-age",
					IsHierarchy: false,
				}, {
					ID:          "npp-sex",
					Name:        "sex",
					HRef:        "http://localhost:22400/code-lists/npp-sex",
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
					ID:          "yyyy-yy",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/yyyy-yy",
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
					ID:          "npp-age",
					Name:        "age",
					HRef:        "http://localhost:22400/code-lists/npp-age",
					IsHierarchy: false,
				}, {
					ID:          "npp-sex",
					Name:        "sex",
					HRef:        "http://localhost:22400/code-lists/npp-sex",
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
					ID:          "yyyy-yy",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/yyyy-yy",
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
					ID:          "npp-age",
					Name:        "age",
					HRef:        "http://localhost:22400/code-lists/npp-age",
					IsHierarchy: false,
				}, {
					ID:          "npp-sex",
					Name:        "sex",
					HRef:        "http://localhost:22400/code-lists/npp-sex",
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
					ID:          "yyyy-yy",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/yyyy-yy",
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
					ID:          "npp-age",
					Name:        "age",
					HRef:        "http://localhost:22400/code-lists/npp-age",
					IsHierarchy: false,
				}, {
					ID:          "npp-sex",
					Name:        "sex",
					HRef:        "http://localhost:22400/code-lists/npp-sex",
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
					ID:          "yyyy-yy",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/yyyy-yy",
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
					ID:          "npp-age",
					Name:        "age",
					HRef:        "http://localhost:22400/code-lists/npp-age",
					IsHierarchy: false,
				}, {
					ID:          "npp-sex",
					Name:        "sex",
					HRef:        "http://localhost:22400/code-lists/npp-sex",
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
					ID:          "yyyy-yy",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/yyyy-yy",
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
					ID:          "npp-age",
					Name:        "age",
					HRef:        "http://localhost:22400/code-lists/npp-age",
					IsHierarchy: false,
				}, {
					ID:          "npp-sex",
					Name:        "sex",
					HRef:        "http://localhost:22400/code-lists/npp-sex",
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
					ID:          "npp-sex",
					Name:        "sex",
					HRef:        "http://localhost:22400/code-lists/npp-sex",
					IsHierarchy: false,
				}, {
					ID:          "npp-all-projectiontype",
					Name:        "projectiontype",
					HRef:        "http://localhost:22400/code-lists/npp-all-projectiontype",
					IsHierarchy: false,
				}, {
					ID:          "npp-age",
					Name:        "age",
					HRef:        "http://localhost:22400/code-lists/npp-age",
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

// Faster indicators of UK economic activity, VAT quarterly diffusion indices
var QuarterlyDiffusion = Response{
	ID:     "542b170b-e654-40a7-b357-88031b7a73ec",
	Alias:  "Faster indicators of UK economic activity, VAT quarterly diffusion indices",
	Format: "v4",
	InputFiles: []file{
		{"VATQuarterlyDiffusionIndicesExcel"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "faster-indicators-quarterly-diffusion-indices",
			Editions:  []string{"time-series"},
			Title:     "Faster indicators of UK economic activity, VAT quarterly diffusion indices",
			CodeLists: []CodeList{
				{
					ID:          "yyyy-qq",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/yyyy-qq",
					IsHierarchy: false,
				}, {
					ID:          "uk-only",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/uk-only",
					IsHierarchy: false,
				}, {
					ID:          "faster-indicators-estimate",
					Name:        "estimate",
					HRef:        "http://localhost:22400/code-lists/faster-indicators-estimate",
					IsHierarchy: false,
				}, {
					ID:          "faster-indicators-index",
					Name:        "index",
					HRef:        "http://localhost:22400/code-lists/faster-indicators-index",
					IsHierarchy: false,
				}, {
					ID:          "faster-indicators-metric",
					Name:        "metric",
					HRef:        "http://localhost:22400/code-lists/faster-indicators-metric",
					IsHierarchy: false,
				}, {
					ID:          "seasonal-adjustment",
					Name:        "seasonaladjustment",
					HRef:        "http://localhost:22400/code-lists/seasonal-adjustment",
					IsHierarchy: false,
				}, {
					ID:          "faster-indicators-sic",
					Name:        "sic",
					HRef:        "http://localhost:22400/code-lists/faster-indicators-sic",
					IsHierarchy: false,
				}, {
					ID:          "time-period",
					Name:        "timeperiod",
					HRef:        "http://localhost:22400/code-lists/time-period",
					IsHierarchy: false,
				},
			},
		},
	},
}

// Faster indicators of UK economic activity, VAT monthly diffusion indices
var MonthlyDiffusion = Response{
	ID:     "0ebf96fc-6bbb-4940-8c70-00400b144c87",
	Alias:  "Faster indicators of UK economic activity, VAT monthly diffusion indices",
	Format: "v4",
	InputFiles: []file{
		{"VATMonthlyDiffusionIndicesExcel"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "faster-indicators-monthly-diffusion-indices",
			Editions:  []string{"time-series"},
			Title:     "Faster indicators of UK economic activity, VAT monthly diffusion indices",
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
					ID:          "faster-indicators-estimate",
					Name:        "estimate",
					HRef:        "http://localhost:22400/code-lists/faster-indicators-estimate",
					IsHierarchy: false,
				}, {
					ID:          "faster-indicators-index",
					Name:        "index",
					HRef:        "http://localhost:22400/code-lists/faster-indicators-index",
					IsHierarchy: false,
				}, {
					ID:          "faster-indicators-metric",
					Name:        "metric",
					HRef:        "http://localhost:22400/code-lists/faster-indicators-metric",
					IsHierarchy: false,
				}, {
					ID:          "seasonal-adjustment",
					Name:        "seasonaladjustment",
					HRef:        "http://localhost:22400/code-lists/seasonal-adjustment",
					IsHierarchy: false,
				}, {
					ID:          "faster-indicators-sic",
					Name:        "sic",
					HRef:        "http://localhost:22400/code-lists/faster-indicators-sic",
					IsHierarchy: false,
				}, {
					ID:          "time-period",
					Name:        "timeperiod",
					HRef:        "http://localhost:22400/code-lists/time-period",
					IsHierarchy: false,
				}, 
			},
		},
	},
}

// Faster indicators of UK economic activity, VAT reporting behaviour indices
var ReportingBehaviour = Response{
	ID:     "aa0bd113-110d-47e4-9418-b71ff54ed93f",
	Alias:  "Faster indicators of UK economic activity, VAT reporting behaviour indices",
	Format: "v4",
	InputFiles: []file{
		{"VATReportingBehaviourIndicesExcel"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "faster-indicators-reporting-behaviour-indices",
			Editions:  []string{"time-series"},
			Title:     "Faster indicators of UK economic activity, VAT reporting behaviour indices",
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
					ID:          "faster-indicators-estimate",
					Name:        "estimate",
					HRef:        "http://localhost:22400/code-lists/faster-indicators-estimate",
					IsHierarchy: false,
				}, {
					ID:          "faster-indicators-index",
					Name:        "index",
					HRef:        "http://localhost:22400/code-lists/faster-indicators-index",
					IsHierarchy: false,
				}, {
					ID:          "faster-indicators-metric",
					Name:        "metric",
					HRef:        "http://localhost:22400/code-lists/faster-indicators-metric",
					IsHierarchy: false,
				}, {
					ID:          "seasonal-adjustment",
					Name:        "seasonaladjustment",
					HRef:        "http://localhost:22400/code-lists/seasonal-adjustment",
					IsHierarchy: false,
				}, {
					ID:          "faster-indicators-sic",
					Name:        "sic",
					HRef:        "http://localhost:22400/code-lists/faster-indicators-sic",
					IsHierarchy: false,
				},
			},
		},
	},
}

// Faster indicators of UK economic activity, VAT reporting behaviour indices with record type
var ReportingBehaviour2 = Response{
	ID:     "96a54186-b907-4d94-bc39-22bdfd5ca1f9",
	Alias:  "Faster indicators of UK economic activity, VAT reporting behaviour indices with record type",
	Format: "v4",
	InputFiles: []file{
		{"VATReportingBehaviourIndices2Excel"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "faster-indicators-reporting-behaviour-indices-2",
			Editions:  []string{"time-series"},
			Title:     "Faster indicators of UK economic activity, VAT reporting behaviour indices with record type",
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
					ID:          "faster-indicators-estimate",
					Name:        "estimate",
					HRef:        "http://localhost:22400/code-lists/faster-indicators-estimate",
					IsHierarchy: false,
				}, {
					ID:          "faster-indicators-index",
					Name:        "index",
					HRef:        "http://localhost:22400/code-lists/faster-indicators-index",
					IsHierarchy: false,
				}, {
					ID:          "faster-indicators-record-type",
					Name:        "recordtype",
					HRef:        "http://localhost:22400/code-lists/faster-indicators-record-type",
					IsHierarchy: false,
				}, {
					ID:          "seasonal-adjustment",
					Name:        "seasonaladjustment",
					HRef:        "http://localhost:22400/code-lists/seasonal-adjustment",
					IsHierarchy: false,
				}, {
					ID:          "faster-indicators-sic",
					Name:        "sic",
					HRef:        "http://localhost:22400/code-lists/faster-indicators-sic",
					IsHierarchy: false,
				},
			},
		},
	},
}

// Local authority ageing statistics - principal population projections
var AgeingPopProj = Response{
	ID:     "1c5b0a7b-024c-49fe-8446-a66eb9736ed3",
	Alias:  "Local authority ageing statistics - principal population projections",
	Format: "v4",
	InputFiles: []file{
		{"AgeingPopulationProjections"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "ageing-population-projections",
			Editions:  []string{"time-series"},
			Title:     "Local authority ageing statistics, principal population projections",
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
					ID:          "ashe-sex",
					Name:        "sex",
					HRef:        "http://localhost:22400/code-lists/ashe-sex",
					IsHierarchy: false,
				}, {
					ID:          "age-groups",
					Name:        "agegroups",
					HRef:        "http://localhost:22400/code-lists/age-groups",
					IsHierarchy: false,
				}, 
			},
		},
	},
}

// Local authority ageing statistics - % of single households headed by someone over 65 and over 85
var AgeingSingleHouseholds = Response{
	ID:     "226d6940-ed0c-4556-91f1-9ca61d56c8a0",
	Alias:  "Local authority ageing statistics - % of single households headed by someone over 65 and over 85",
	Format: "v4",
	InputFiles: []file{
		{"AgeingSingleHouseholds"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "ageing-single-households",
			Editions:  []string{"time-series"},
			Title:     "Local authority ageing statistics, percentage of single households headed by someone over 65 and over 85",
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
					ID:          "ashe-sex",
					Name:        "sex",
					HRef:        "http://localhost:22400/code-lists/ashe-sex",
					IsHierarchy: false,
				}, {
					ID:          "age-groups",
					Name:        "agegroups",
					HRef:        "http://localhost:22400/code-lists/age-groups",
					IsHierarchy: false,
				}, 
			},
		},
	},
}

// Local authority ageing statistics - sex ratios for people over 65 and over 85
var AgeingSexRatios = Response{
	ID:     "df8c8df5-484a-4d29-b246-a7720caf2474",
	Alias:  "Local authority ageing statistics - sex ratios for people over 65 and over 85",
	Format: "v4",
	InputFiles: []file{
		{"AgeingSexRatios"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "ageing-sex-ratios",
			Editions:  []string{"time-series"},
			Title:     "Local authority ageing statistics, sex ratios for people over 65 and over 85",
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
					ID:          "age-groups",
					Name:        "agegroups",
					HRef:        "http://localhost:22400/code-lists/age-groups",
					IsHierarchy: false,
				}, 
			},
		},
	},
}

// Local authority ageing statistics - net flows for over 65s and over 85s
var AgeingNetFlows = Response{
	ID:     "c5875623-0874-429d-bbc6-63d2d84e6159",
	Alias:  "Local authority ageing statistics - net flows for over 65s and over 85s",
	Format: "v4",
	InputFiles: []file{
		{"AgeingNetFlows"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "ageing-net-flows",
			Editions:  []string{"time-series"},
			Title:     "Local authority ageing statistics, net flows for over 65s and over 85s",
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
					ID:          "ashe-sex",
					Name:        "sex",
					HRef:        "http://localhost:22400/code-lists/ashe-sex",
					IsHierarchy: false,
				}, {
					ID:          "age-groups",
					Name:        "agegroups",
					HRef:        "http://localhost:22400/code-lists/age-groups",
					IsHierarchy: false,
				}, 
			},
		},
	},
}

// Local authority ageing statistics - Subnational Prospective Ageing Measures
var AgeingProspectiveMeasures = Response{
	ID:     "ebda7264-b7c0-4a6f-98e3-63ebb328008c",
	Alias:  "Local authority ageing statistics - Subnational Prospective Ageing Measures",
	Format: "v4",
	InputFiles: []file{
		{"AgeingProspectiveMeasures"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "ageing-prospective-measures",
			Editions:  []string{"time-series"},
			Title:     "Subnational Prospective Ageing Measures",
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
					ID:          "ashe-sex",
					Name:        "sex",
					HRef:        "http://localhost:22400/code-lists/ashe-sex",
					IsHierarchy: false,
				}, {
					ID:          "remaining-life-expectancy",
					Name:        "remaininglifeexpectancy",
					HRef:        "http://localhost:22400/code-lists/remaining-life-expectancy",
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
					ID:          "uk-only",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/uk-only",
					IsHierarchy: false,
				}, {
					ID:          "unit-of-measure",
					Name:        "unitofmeasure",
					HRef:        "http://localhost:22400/code-lists/unit-of-measure",
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
					ID:          "adult-sex",
					Name:        "sex",
					HRef:        "http://localhost:22400/code-lists/adult-sex",
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

// Regional Labour Market, economic status by age
var LMSEconomicByAge = Response{
	ID:     "1441209b-ca85-4537-87c4-78f2a33eda0d",
	Alias:  "Regional Labour Market, Economic Status by Age",
	Format: "v4",
	InputFiles: []file{
		{"LMSEconomicByAgeExcel"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "regional-lms-economic-status-age",
			Editions:  []string{"time-series"},
			Title:     "Regional Labour Market, Economic Status by Age",
			CodeLists: []CodeList{
				{
					ID:          "mmm-mmm-yyyy",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/mmm-mmm-yyyy",
					IsHierarchy: false,
				}, {
					ID:          "admin-geography",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/admin-geography",
					IsHierarchy: false,
				}, {
					ID:          "lms-age-bracket",
					Name:        "age",
					HRef:        "http://localhost:22400/code-lists/lms-age-bracket",
					IsHierarchy: false,
				}, {
					ID:          "unit-of-measure",
					Name:        "unitofmeasure",
					HRef:        "http://localhost:22400/code-lists/unit-of-measure",
					IsHierarchy: false,
				}, {
					ID:          "ashe-sex",
					Name:        "sex",
					HRef:        "http://localhost:22400/code-lists/ashe-sex",
					IsHierarchy: false,
				}, {
					ID:          "lms-economic-status",
					Name:        "economicstatus",
					HRef:        "http://localhost:22400/code-lists/lms-economic-status",
					IsHierarchy: false,
				},
			},
		},
	},
}

// Regional Labour Market, Full-time, part-time and temporary workers
var LMSWorkType = Response{
	ID:     "6e62de16-6d28-44e0-aca0-674ec9b27430",
	Alias:  "Regional Labour Market, Full-time, part-time and temporary workers",
	Format: "v4",
	InputFiles: []file{
		{"LMSWorkTypeExcel"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "regional-lms-by-work-type",
			Editions:  []string{"time-series"},
			Title:     "Regional Labour Market, Full-time, part-time and temporary workers",
			CodeLists: []CodeList{
				{
					ID:          "mmm-mmm-yyyy",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/mmm-mmm-yyyy",
					IsHierarchy: false,
				}, {
					ID:          "admin-geography",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/admin-geography",
					IsHierarchy: false,
				}, {
					ID:          "ashe-sex",
					Name:        "sex",
					HRef:        "http://localhost:22400/code-lists/ashe-sex",
					IsHierarchy: false,
				}, {
					ID:          "lms-work-type",
					Name:        "worktype",
					HRef:        "http://localhost:22400/code-lists/lms-work-type",
					IsHierarchy: false,
				},
			},
		},
	},
}

// Regional Labour Market, Actual weekly hours of work
var LMSActualHoursWork = Response{
	ID:     "8489d0e5-5edd-4bb1-8aec-3101cbeff2c4",
	Alias:  "Regional Labour Market, Actual weekly hours of work",
	Format: "v4",
	InputFiles: []file{
		{"LMSActualHoursWorkExcel"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "regional-lms-actual-hours-work",
			Editions:  []string{"time-series"},
			Title:     "Regional Labour Market, Actual weekly hours of work",
			CodeLists: []CodeList{
				{
					ID:          "mmm-mmm-yyyy",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/mmm-mmm-yyyy",
					IsHierarchy: false,
				}, {
					ID:          "admin-geography",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/admin-geography",
					IsHierarchy: false,
				}, {
					ID:          "ashe-sex",
					Name:        "sex",
					HRef:        "http://localhost:22400/code-lists/ashe-sex",
					IsHierarchy: false,
				}, {
					ID:          "lms-average-actual-hours-of-work",
					Name:        "averageactualhoursofwork",
					HRef:        "http://localhost:22400/code-lists/lms-average-actual-hours-of-work",
					IsHierarchy: false,
				},
			},
		},
	},
}

// Regional Labour Market, Claimant Count	
var LMSClaimantCount = Response{
	ID:     "064ff64f-70bb-4c12-9bc1-150294f31921",
	Alias:  "Regional Labour Market, Claimant Count",
	Format: "v4",
	InputFiles: []file{
		{"LMSClaimantCountExcel"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "regional-lms-claimant-count",
			Editions:  []string{"time-series"},
			Title:     "Regional Labour Market, Claimant Count",
			CodeLists: []CodeList{
				{
					ID:          "mmm-yy",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/mmm-yy",
					IsHierarchy: false,
				}, {
					ID:          "admin-geography",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/admin-geography",
					IsHierarchy: false,
				}, {
					ID:          "seasonal-adjustment",
					Name:        "seasonaladjustment",
					HRef:        "http://localhost:22400/code-lists/seasonal-adjustment",
					IsHierarchy: false,
				}, {
					ID:          "unit-of-measure",
					Name:        "unitofmeasure",
					HRef:        "http://localhost:22400/code-lists/unit-of-measure",
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

// Regional Labour Market, Jobseeker's Allowance by age and duration - computerised claims only
var LMSJobseekersByAgeDuration = Response{
	ID:     "239bb6e2-220d-4544-8f1f-0183a1a9ba42",
	Alias:  "Regional Labour Market, Jobseeker's Allowance by age and duration - computerised claims only",
	Format: "v4",
	InputFiles: []file{
		{"LMSJobseekersByAgeDurationExcel"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "regional-lms-jobseekers-allowance-age-duration",
			Editions:  []string{"time-series"},
			Title:     "Regional Labour Market, Jobseeker's Allowance by age and duration - computerised claims only",
			CodeLists: []CodeList{
				{
					ID:          "mmm-yy",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/mmm-yy",
					IsHierarchy: false,
				}, {
					ID:          "admin-geography",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/admin-geography",
					IsHierarchy: false,
				}, {
					ID:          "lms-age-bracket",
					Name:        "age",
					HRef:        "http://localhost:22400/code-lists/lms-age-bracket",
					IsHierarchy: false,
				}, {
					ID:          "lms-jobseekers-duration",
					Name:        "duration",
					HRef:        "http://localhost:22400/code-lists/lms-jobseekers-duration",
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

// Regional Labour Market, Economic inactivity: reasons
var LMSEconomicInactivity = Response{
	ID:     "eddc6106-4de5-46c7-b4d9-8bdcd3f33a50",
	Alias:  "Regional Labour Market, Economic inactivity: reasons",
	Format: "v4",
	InputFiles: []file{
		{"LMSEconomicInactivityExcel"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "regional-lms-economic-inactivity-reason",
			Editions:  []string{"time-series"},
			Title:     "Regional Labour Market, Economic inactivity: reasons",
			CodeLists: []CodeList{
				{
					ID:          "mmm-mmm-yyyy",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/mmm-mmm-yyyy",
					IsHierarchy: false,
				}, {
					ID:          "admin-geography",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/admin-geography",
					IsHierarchy: false,
				}, {
					ID:          "unit-of-measure",
					Name:        "unitofmeasure",
					HRef:        "http://localhost:22400/code-lists/unit-of-measure",
					IsHierarchy: false,
				}, {
					ID:          "lms-economic-inactivity-reason",
					Name:        "economicinactivityreason",
					HRef:        "http://localhost:22400/code-lists/lms-economic-inactivity-reason",
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

// Regional Labour Market, Workforce jobs by industry
var LMSJobsByIndustry = Response{
	ID:     "2fb88df2-d438-42d4-baaf-9206c777cc1b",
	Alias:  "Regional Labour Market, Workforce jobs by industry",
	Format: "v4",
	InputFiles: []file{
		{"LMSJobsByIndustryExcel"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "regional-lms-workforce-jobs-industry",
			Editions:  []string{"time-series"},
			Title:     "Regional Labour Market, Workforce jobs by industry",
			CodeLists: []CodeList{
				{
					ID:          "mmm-mmm-yyyy",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/mmm-mmm-yyyy",
					IsHierarchy: false,
				}, {
					ID:          "admin-geography",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/admin-geography",
					IsHierarchy: false,
				}, {
					ID:          "lms-sic",
					Name:        "standardindustrialclassification",
					HRef:        "http://localhost:22400/code-lists/lms-sic",
					IsHierarchy: false,
				}, 
			},
		},
	},
}

// Average weekly earnings
var AWE = Response{
	ID:     "72a8082f-589f-4adc-8915-4448dd3658c2",
	Alias:  "Average Weekly Earnings",
	Format: "v4",
	InputFiles: []file{
		{"AverageWeeklyEarnings"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "average-weekly-earnings",
			Editions:  []string{"time-series"},
			Title:     "Average Weekly Earnings",
			CodeLists: []CodeList{
				{
					ID:          "mmm-yy",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/mmm-yy",
					IsHierarchy: false,
				}, {
					ID:          "countries",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/countries",
					IsHierarchy: false,
				}, {
					ID:          "sic",
					Name:        "standardindustrialclassification",
					HRef:        "http://localhost:22400/code-lists/sic",
					IsHierarchy: false,
				}, {
					ID:          "awe-earnings",
					Name:        "earnings",
					HRef:        "http://localhost:22400/code-lists/awe-earnings",
					IsHierarchy: false,
				}, {
					ID:          "awe-type-of-pay",
					Name:        "typeofpay",
					HRef:        "http://localhost:22400/code-lists/awe-type-of-pay",
					IsHierarchy: false,
				}, 
			},
		},
	},
}

// Average weekly earnings Index
var AWEIndex = Response{
	ID:     "0cab7368-3aa8-48a6-b1fe-7e3aa05f0a6e",
	Alias:  "Average Weekly Earnings Index",
	Format: "v4",
	InputFiles: []file{
		{"AverageWeeklyEarningsIndex"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "average-weekly-earnings-index",
			Editions:  []string{"time-series"},
			Title:     "Average Weekly Earnings Index",
			CodeLists: []CodeList{
				{
					ID:          "mmm-yy",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/mmm-yy",
					IsHierarchy: false,
				}, {
					ID:          "countries",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/countries",
					IsHierarchy: false,
				}, {
					ID:          "sic",
					Name:        "standardindustrialclassification",
					HRef:        "http://localhost:22400/code-lists/sic",
					IsHierarchy: false,
				}, {
					ID:          "awe-type-of-pay",
					Name:        "typeofpay",
					HRef:        "http://localhost:22400/code-lists/awe-type-of-pay",
					IsHierarchy: false,
				}, 
			},
		},
	},
}

// Cancer Registrations for England
var CancerRegEng = Response{
	ID:     "db8d4b24-20e4-4c22-927c-74f4688f6b4c",
	Alias:  "Cancer Registrations England",
	Format: "v4",
	InputFiles: []file{
		{"CancerRegistrationsEnglandExcel"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "cancer-registrations-england",
			Editions:  []string{"time-series"},
			Title:     "Cancer registration statistics, England",
			CodeLists: []CodeList{
				{
					ID:          "calendar-years",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/calendar-years",
					IsHierarchy: false,
				}, {
					ID:          "countries",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/countries",
					IsHierarchy: false,
				}, {
					ID:          "icd-10",
					Name:        "icd10",
					HRef:        "http://localhost:22400/code-lists/icd-10",
					IsHierarchy: true,
				}, {
					ID:          "ashe-sex",
					Name:        "sex",
					HRef:        "http://localhost:22400/code-lists/ashe-sex",
					IsHierarchy: false,
				}, {
					ID:          "five-year-age-brackets",
					Name:        "agebrackets",
					HRef:        "http://localhost:22400/code-lists/five-year-age-brackets",
					IsHierarchy: false,
				}, {
					ID:          "cancer-registrations",
					Name:        "cancerregistrations",
					HRef:        "http://localhost:22400/code-lists/cancer-registrations",
					IsHierarchy: false,
				}, 
			},
		},
	},
}

// Cancer Registrations for Regions, England
var CancerRegRegions = Response{
	ID:     "d07c600d-1511-4e76-aa78-c6ce461c9dc7",
	Alias:  "Cancer Registrations by Regions, England",
	Format: "v4",
	InputFiles: []file{
		{"CancerRegistrationsRegionsExcel"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "cancer-registrations-regions",
			Editions:  []string{"time-series"},
			Title:     "Cancer registration statistics by Regions, England",
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
					ID:          "icd-10",
					Name:        "icd10",
					HRef:        "http://localhost:22400/code-lists/icd-10",
					IsHierarchy: true,
				}, {
					ID:          "ashe-sex",
					Name:        "sex",
					HRef:        "http://localhost:22400/code-lists/ashe-sex",
					IsHierarchy: false,
				}, {
					ID:          "cancer-registrations",
					Name:        "cancerregistrations",
					HRef:        "http://localhost:22400/code-lists/cancer-registrations",
					IsHierarchy: false,
				}, 
			},
		},
	},
}

// Regional GDP by year
var RegionalGDPYear = Response{
	ID:     "eec65efd-da07-4089-8e3e-a356751fa72d",
	Alias:  "Regional GDP by Year",
	Format: "v4",
	InputFiles: []file{
		{"RegionalGDPYear"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "regional-gdp-by-year",
			Editions:  []string{"time-series"},
			Title:     "Regional GDP by Year",
			CodeLists: []CodeList{
				{
					ID:          "calendar-years",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/calendar-years",
					IsHierarchy: false,
				}, {
					ID:          "nuts-geography",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/nuts-geography",
					IsHierarchy: false,
				}, {
					ID:          "gdp-sic",
					Name:        "sic",
					HRef:        "http://localhost:22400/code-lists/gdp-sic",
					IsHierarchy: false,
				}, {
					ID:          "type-of-prices",
					Name:        "prices",
					HRef:        "http://localhost:22400/code-lists/type-of-prices",
					IsHierarchy: false,
				}, {
					ID:          "gdp-measure",
					Name:        "measure",
					HRef:        "http://localhost:22400/code-lists/gdp-measure",
					IsHierarchy: false,
				},
			},
		},
	},
}

// Regional GDP by quarter
var RegionalGDPQuarter = Response{
	ID:     "b89ba5a1-96ec-4311-af3f-d732943eeb2f",
	Alias:  "Regional GDP by Quarter",
	Format: "v4",
	InputFiles: []file{
		{"RegionalGDPQuarter"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "regional-gdp-by-quarter",
			Editions:  []string{"time-series"},
			Title:     "Regional GDP by Quarter",
			CodeLists: []CodeList{
				{
					ID:          "yyyy-qq",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/yyyy-qq",
					IsHierarchy: false,
				}, {
					ID:          "nuts-geography",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/nuts-geography",
					IsHierarchy: false,
				}, {
					ID:          "gdp-sic",
					Name:        "sic",
					HRef:        "http://localhost:22400/code-lists/gdp-sic",
					IsHierarchy: false,
				}, {
					ID:          "type-of-prices",
					Name:        "prices",
					HRef:        "http://localhost:22400/code-lists/type-of-prices",
					IsHierarchy: false,
				}, {
					ID:          "gdp-measure",
					Name:        "measure",
					HRef:        "http://localhost:22400/code-lists/gdp-measure",
					IsHierarchy: false,
				},
			},
		},
	},
}

// Tax Benefits Statistics
var TaxBenefitStats = Response{
	ID:     "6cadb0b9-927d-4fa3-b07f-2654ad7e71b9",
	Alias:  "Tax Benefits Statistics",
	Format: "v4",
	InputFiles: []file{
		{"TaxBenefitsStats"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "tax-benefits-statistics",
			Editions:  []string{"time-series"},
			Title:     "Tax Benefits Statistics",
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
					ID:          "income-quintile",
					Name:        "quintile",
					HRef:        "http://localhost:22400/code-lists/income-quintile",
					IsHierarchy: false,
				}, {
					ID:          "ashe-statistics",
					Name:        "statistics",
					HRef:        "http://localhost:22400/code-lists/ashe-statistics",
					IsHierarchy: false,
				}, {
					ID:          "income-type",
					Name:        "income",
					HRef:        "http://localhost:22400/code-lists/income-type",
					IsHierarchy: false,
				},
			},
		},
	},
}

// Generational Income
var GenerationalIncome = Response{
	ID:     "d4e813a9-dc67-4a66-b0ae-c7e531f3b64b",
	Alias:  "Generational Income",
	Format: "v4",
	InputFiles: []file{
		{"GenerationalIncome"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "generational-income",
			Editions:  []string{"time-series"},
			Title:     "Generational income",
			CodeLists: []CodeList{
				{
					ID:          "yyyy-yy",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/yyyy-years",
					IsHierarchy: false,
				}, {
					ID:          "uk-only",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/uk-only",
					IsHierarchy: false,
				}, {
					ID:          "mid-year-pop-age",
					Name:        "age",
					HRef:        "http://localhost:22400/code-lists/mid-year-pop-age",
					IsHierarchy: false,
				}, {
					ID:          "tax-benefit-type",
					Name:        "typeoftaxorbenefit",
					HRef:        "http://localhost:22400/code-lists/tax-benefit-type",
					IsHierarchy: false,
				}, {
					ID:          "decades",
					Name:        "decade",
					HRef:        "http://localhost:22400/code-lists/decades",
					IsHierarchy: false,
				},
			},
		},
	},
}

// House Prices by Local Authority
var HousePrices = Response{
	ID:     "4a4273d2-85ed-4d8e-9da9-22b42b70d624",
	Alias:  "House Prices",
	Format: "v4",
	InputFiles: []file{
		{"HousePrices"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "house-prices-local-authority",
			Editions:  []string{"time-series"},
			Title:     "UK House Price Index",
			CodeLists: []CodeList{
				{
					ID:          "mmm-yy",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/mmm-yy",
					IsHierarchy: false,
				}, {
					ID:          "admin-geography",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/admin-geography",
					IsHierarchy: true,
				}, {
					ID:          "house-price-variable",
					Name:        "housepricevariable",
					HRef:        "http://localhost:22400/code-lists/house-price-variable",
					IsHierarchy: false,
				}, {
					ID:          "property-type",
					Name:        "propertytype",
					HRef:        "http://localhost:22400/code-lists/property-type",
					IsHierarchy: false,
				}, {
					ID:          "house-price-age",
					Name:        "housepriceage",
					HRef:        "http://localhost:22400/code-lists/house-price-age",
					IsHierarchy: false,
				},
			},
		},
	},
}

// Index of Private Housing Rental Prices
var PrivateHousingRentalPrices = Response{
	ID:     "d5943cff-a8b7-4002-ad1f-1957764aec7b",
	Alias:  "Index of Private Housing Rental Prices",
	Format: "v4",
	InputFiles: []file{
		{"PrivateHousingRentalPrices"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "index-private-housing-rental-prices",
			Editions:  []string{"time-series"},
			Title:     "Index of Private Housing Rental Prices",
			CodeLists: []CodeList{
				{
					ID:          "mmm-yy",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/mmm-yy",
					IsHierarchy: false,
				}, {
					ID:          "admin-geography",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/admin-geography",
					IsHierarchy: true,
				}, {
					ID:          "housing-rental-prices-variable",
					Name:        "variable",
					HRef:        "http://localhost:22400/code-lists/housing-rental-prices-variable",
					IsHierarchy: false,
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
			Editions:  []string{"2016", "2017", "2018", "2019", "2020", "time-series"},
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
			Editions:  []string{"2016", "2017", "2018", "2019", "2020", "time-series"},
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

// Knife crime
var KnifeCrime = Response{
	ID:     "2f967891-b0c9-403e-9747-766aed0dfc2a",
	Alias:  "Knife Crime",
	Format: "v4",
	InputFiles: []file{
		{"KnifeCrimeExcel"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "knife-crime",
			Editions:  []string{"time-series"},
			Title:     "Knife Crime",
			CodeLists: []CodeList{
				{
					ID:          "yyyy-yy",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/yyyy-yy",
					IsHierarchy: false,
				}, {
					ID:          "quarters",
					Name:        "quarter",
					HRef:        "http://localhost:22400/code-lists/quarters",
					IsHierarchy: false,
				}, {
					ID:          "police-force-area-geography",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/police-force-area-geography",
					IsHierarchy: false,
				}, {
					ID:          "crime-offence-code",
					Name:        "crime",
					HRef:        "http://localhost:22400/code-lists/crime-offence-code",
					IsHierarchy: false,
				}, 
			},
		},
	},
}

// Internal migration - Moves by local authorities and regions in England and Wales by 5 year age group and sex
var InternalMigrationLA = Response{
	ID:     "b135f977-6f4d-4f38-821a-d74722e6737c",
	Alias:  "Internal migration - Moves by local authorities",
	Format: "v4",
	InputFiles: []file{
		{"InternalMigrationLA"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "internal-migration-local-authority",
			Editions:  []string{"time-series"},
			Title:     "Internal migration - Moves by local authorities and regions in England and Wales by 5 year age group and sex",
			CodeLists: []CodeList{
				{
					ID:          "year-ending",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/year-ending",
					IsHierarchy: false,
				}, {
					ID:          "admin-geography",
					Name:        "geography",
					HRef:        "http://localhost:22400/code-lists/admin-geography",
					IsHierarchy: true,
				}, {
					ID:          "five-year-age-brackets",
					Name:        "age",
					HRef:        "http://localhost:22400/code-lists/five-year-age-brackets",
					IsHierarchy: false,
				}, {
					ID:          "mid-year-pop-sex",
					Name:        "sex",
					HRef:        "http://localhost:22400/code-lists/mid-year-pop-sex",
					IsHierarchy: false,
				}, {
					ID:          "migration-direction",
					Name:        "migrationdirection",
					HRef:        "http://localhost:22400/code-lists/migration-direction",
					IsHierarchy: false,
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

// Families and Households
var FamiliesAndHouseholds = Response{
	ID:     "bb9dacb0-2b73-4b3c-8198-4b1350dd13c7",
	Alias:  "Families and households",
	Format: "v4",
	InputFiles: []file{
		{"FamiliesAndHouseholds"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "families-and-households",
			Editions:  []string{"time-series"},
			Title:     "Families and households",
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
					IsHierarchy: false,
				}, {
					ID:          "family-household-type",
					Name:        "familyhousehold",
					HRef:        "http://localhost:22400/code-lists/family-household-type",
					IsHierarchy: false,
				}, {
					ID:          "children-in-family",
					Name:        "childreninfamiy",
					HRef:        "http://localhost:22400/code-lists/children-in-family",
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
					ID:          "unit-of-measure",
					Name:        "unitofmeasure",
					HRef:        "http://localhost:22400/code-lists/unit-of-measure",
					IsHierarchy: false,
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

// Child mortality by local authority in England and Wales
var ChildMortality = Response{
	ID:     "f507274d-2f17-4cbc-86e5-bf2c91dcc91d",
	Alias:  "Child mortality by local authority",
	Format: "v4",
	InputFiles: []file{
		{"ChildMortalityLA"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "child-mortality-local-authority",
			Editions:  []string{"time-series"},
			Title:     "Child mortality by local authority in England and Wales",
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
					ID:          "unit-of-measure",
					Name:        "unitofmeasure",
					HRef:        "http://localhost:22400/code-lists/unit-of-measure",
					IsHierarchy: false,
				}, {
					ID:          "births-and-deaths",
					Name:        "birthsanddeaths",
					HRef:        "http://localhost:22400/code-lists/births-and-deaths",
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
			Editions:  []string{"2012", "2013", "2014", "2015", "2016", "2017", "2018", "2019", "2020", "time-series"},
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

// Deaths registered monthly in England and Wales
var MonthlyDeaths = Response{
	ID:     "7e4b9781-2732-4847-a290-e547198acfe2",
	Alias:  "Deaths registered monthly in England and Wales",
	Format: "v4",
	InputFiles: []file{
		{"MonthlyDeaths"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "deaths-registered-monthly-england-wales",
			Editions:  []string{"time-series"},
			Title:     "Deaths registered monthly in England and Wales",
			CodeLists: []CodeList{
				{
					ID:          "mmm-yy",
					Name:        "time",
					HRef:        "http://localhost:22400/code-lists/mmm-yy",
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

// Sexual orientation, UK
var SexualOrientation = Response{
	ID:     "7d79a78e-dd94-4c0a-aeb0-3d9f5d7ff88a",
	Alias:  "Sexual orientation, UK",
	Format: "v4",
	InputFiles: []file{
		{"SexualOrientationExcel"},
	},
	OutputInstances: []instance{
		{
			DatasetID: "sexual-orientation-uk",
			Editions:  []string{"time-series"},
			Title:     "Sexual orientation, UK",
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
					ID:          "ashe-sex",
					Name:        "sex",
					HRef:        "http://localhost:22400/code-lists/ashe-sex",
					IsHierarchy: false,
				}, {
					ID:          "lms-age-bracket",
					Name:        "age",
					HRef:        "http://localhost:22400/code-lists/lms-age-bracket",
					IsHierarchy: false,
				}, {
					ID:          "sexual-identity",
					Name:        "sexualidentity",
					HRef:        "http://localhost:22400/code-lists/sexual-identity",
					IsHierarchy: false,
				}, {
					ID:          "unit-of-measure",
					Name:        "unitofmeasure",
					HRef:        "http://localhost:22400/code-lists/unit-of-measure",
					IsHierarchy: false,
				},
			},
		},
	},
}
