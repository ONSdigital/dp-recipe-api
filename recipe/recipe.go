package recipe

//List - struct for list of recipes
type List struct {
	Count        int        `bson:"count,omitempty" json:"count,omitempty"`
	Start        int        `bson:"start_index,omitempty" json:"start_index,omitempty"`
	ItemsPerPage int        `bson:"items_per_page,omitempty" json:"items_per_page,omitempty"`
	Items        []Response `bson:"items,omitempty" json:"items,omitempty"`
	TotalCount   int        `bson:"total_count,omitempty" json:"total_count,omitempty"`
}

//Response - struct for individual recipe
type Response struct {
	ID              string     `bson:"_id" json:"id"`
	Alias           string     `bson:"alias,omitempty" json:"alias,omitempty"`
	Format          string     `bson:"format,omitempty" json:"format,omitempty"`
	InputFiles      []file     `bson:"files,omitempty" json:"files,omitempty"`
	OutputInstances []instance `bson:"output_instances,omitempty" json:"output_instances,omitempty"`
}

//CodeList - Code lists for instance
type CodeList struct {
	ID          string `bson:"id,omitempty" json:"id,omitempty"`
	HRef        string `bson:"href,omitempty" json:"href,omitempty"`
	Name        string `bson:"name,omitempty" json:"name,omitempty"`
	IsHierarchy bool   `bson:"is_hierarchy,omitempty" json:"is_hierarchy,omitempty"`
}

type instance struct {
	DatasetID string     `bson:"dataset_id,omitempty" json:"dataset_id,omitempty"`
	Editions  []string   `bson:"editions,omitempty" json:"editions,omitempty"`
	Title     string     `bson:"title,omitempty" json:"title,omitempty"`
	CodeLists []CodeList `bson:"code_lists,omitempty" json:"code_lists,omitempty"`
}

type file struct {
	Description string `bson:"description,omitempty" json:"description,omitempty"`
}
