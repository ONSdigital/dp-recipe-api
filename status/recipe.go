package status

import (
	"context"
	"sync"

	"github.com/ONSdigital/dp-api-clients-go/dataset"
	"github.com/ONSdigital/dp-recipe-api/recipe"
)

type RecipeList struct {
	Count        int      `json:"count"`
	Start        int      `json:"start_index"`
	ItemsPerPage int      `json:"items_per_page"`
	Items        []Recipe `json:"items"`
	TotalCount   int      `json:"total_count"`
}

type Recipe struct {
	ID      string   `json:"recipe_id"`
	Alias   string   `json:"alias"`
	Outputs []Output `json:"outputs"`
}

type Output struct {
	DatasetName string `json:"dataset_name"`
	DevStatus   int    `json:"dev_status"`
	BetaStatus  int    `json:"beta_status"`
}

type CheckRequest struct {
	UserAuthToken    string
	ServiceAuthToken string
	BetaURL          string
	DevURL           string
	CollectionID     string
	DatasetID        string
	CodeLists        []recipe.CodeList
}

func CheckRecipe(ctx context.Context, req CheckRequest) *Output {
	o := &Output{
		DevStatus:  2,
		BetaStatus: 2,
	}
	var wg sync.WaitGroup
	wg.Add(2)
	go o.devstatus(ctx, &wg, req)
	go o.betastatus(ctx, &wg, req)
	wg.Wait()

	return o
}

func (o *Output) devstatus(ctx context.Context, w *sync.WaitGroup, req CheckRequest) {
	defer w.Done()
	o.DevStatus = status(ctx, req, req.DevURL)
}

func (o *Output) betastatus(ctx context.Context, w *sync.WaitGroup, req CheckRequest) {
	defer w.Done()
	o.BetaStatus = status(ctx, req, req.BetaURL)
}

func status(ctx context.Context, req CheckRequest, url string) int {
	s := 2
	if datasetExists(ctx, req, url) {
		return s
	}

	s--

	goal := len(req.CodeLists)
	imported := 0
	var mu sync.Mutex
	var wg sync.WaitGroup

	for _, c := range req.CodeLists {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if codelistExists(url, c.ID) {
				mu.Lock()
				imported++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()

	if goal > imported {
		s--
	}

	return s
}

func datasetExists(ctx context.Context, req CheckRequest, url string) bool {
	cli := dataset.NewAPIClient(url)
	ds, err := cli.Get(ctx, req.UserAuthToken, req.ServiceAuthToken, req.CollectionID, req.DatasetID)
	if err != nil || &ds == nil {
		return false
	}

	return true
}
