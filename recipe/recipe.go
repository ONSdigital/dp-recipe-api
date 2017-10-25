package recipe

type List struct {
	Count        int        `json:"count"`
	Start        int        `json:"start_index"`
	ItemsPerPage int        `json:"items_per_page"`
	Items        []Response `json:"items"`
	TotalCount   int        `json:"total_count"`
}

type Response struct {
	ID              string     `json:"id"`
	Alias           string     `json:"alias"`
	Format          string     `json:"format"`
	InputFiles      []file     `json:"files"`
	OutputInstances []instance `json:"output_instances"`
}

type CodeList struct {
	ID   string `json:"id"`
	HRef string `json:"href"`
	Name string `json:"name"`
}

type instance struct {
	DatasetID string     `json:"dataset_id"`
	Editions  []string   `json:"editions"`
	Title     string     `json:"title"`
	CodeLists []CodeList `json:"code_lists"`
}

type file struct {
	Description string `json:"description"`
}
