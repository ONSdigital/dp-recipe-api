package status

import (
	"context"
	"sync"

	"github.com/ONSdigital/dp-recipe-api/recipe"
	"github.com/ONSdigital/go-ns/clients/dataset"
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

func CheckRecipe(devURL, betaURL, datasetID string, codelists []recipe.CodeList) *Output {
	o := &Output{
		DevStatus:  2,
		BetaStatus: 2,
	}
	var wg sync.WaitGroup
	wg.Add(2)
	go o.devstatus(&wg, devURL, datasetID, codelists)
	go o.betastatus(&wg, betaURL, datasetID, codelists)
	wg.Wait()

	return o
}

func (o *Output) devstatus(w *sync.WaitGroup, url, datasetID string, codelists []recipe.CodeList) {
	defer w.Done()
	o.DevStatus = status(url, datasetID, codelists)
}

func (o *Output) betastatus(w *sync.WaitGroup, url, datasetID string, codelists []recipe.CodeList) {
	defer w.Done()
	o.BetaStatus = status(url, datasetID, codelists)
}

func status(url, datasetID string, codelists []recipe.CodeList) int {
	s := 2
	if datasetExists(url, datasetID) {
		return s
	}

	s--

	goal := len(codelists)
	imported := 0
	var mu sync.Mutex
	var wg sync.WaitGroup

	for _, c := range codelists {
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

func datasetExists(url, id string) bool {
	cli := dataset.NewAPIClient(url, "", "")
	ds, err := cli.Get(context.Background(), id)
	if err != nil || &ds == nil {
		return false
	}

	return true
}
