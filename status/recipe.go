package status

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"

	"github.com/ONSdigital/dp-recipe-api/recipe"
	"github.com/ONSdigital/go-ns/clients/codelist"
	"github.com/ONSdigital/go-ns/clients/dataset"
	"github.com/ONSdigital/go-ns/log"
	"github.com/ONSdigital/go-ns/rchttp"
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

func Check(datasetID string, codelists []recipe.CodeList) *Output {
	o := &Output{
		//DevStatus:  2,
		BetaStatus: 2,
	}
	var wg sync.WaitGroup
	wg.Add(1) // +1
	//	go o.devstatus(&wg, "https://api.cmd-dev.onsdigital.co.uk/v1", datasetID, codelists)
	go o.betastatus(&wg, "https://api.beta.ons.gov.uk/v1", datasetID, codelists)
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
		log.Info("returning greenlight", nil)
		return s
	}

	s--

	goal := len(codelists)
	imported := 0
	//var wg sync.WaitGroup
	for _, c := range codelists {
		go func() {
			if codelistExists(url, c.ID) {
				imported++
			}
		}()
	}
	//wg.Wait()
	if goal > imported {
		s--
	}

	return s
}

func datasetExists(url, id string) bool {
	cli := dataset.NewAPIClient(url, "", "")
	ds, err := cli.Get(context.Background(), id)
	if err != nil {
		log.Error(err, log.Data{"couldnt check for published": url})
		return false
	}

	if &ds == nil {
		return false
	}

	return true
}

func (c Codelist) CheckCodelist(wg *sync.WaitGroup) {
	defer wg.Done()
	//c.InDev = codelistExists(url, c.ID)
	c.InBeta = codelistExists("https://api.beta.ons.gov.uk/v1", c.ID)
}

func codelistExists(url, id string) bool {
	client := rchttp.NewClient()

	uri := fmt.Sprintf("%s/code-lists/%s", url, id)
	resp, err := client.Get(context.Background(), uri)
	if err != nil {
		log.Error(err, log.Data{"url": uri, "id": id})
		return false
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false
	}

	var result *codelist.CodeList
	err = json.Unmarshal(b, &result)
	if err != nil {
		return false
	}

	if &result == nil {
		return false
	}

	return true
}

type CodelistList struct {
	Count        int        `json:"count"`
	Start        int        `json:"start_index"`
	ItemsPerPage int        `json:"items_per_page"`
	Items        []Codelist `json:"items"`
	TotalCount   int        `json:"total_count"`
}

type Codelist struct {
	ID   string
	Name string
	//	NumberOfCodes int
	IsHierarchy bool
	InDev       bool
	InBeta      bool
}
