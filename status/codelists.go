package status

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"

	"github.com/ONSdigital/dp-api-clients-go/codelist"
	"github.com/ONSdigital/dp-rchttp"
)

type CodelistList struct {
	Count        int        `json:"count"`
	Start        int        `json:"start_index"`
	ItemsPerPage int        `json:"items_per_page"`
	Items        []Codelist `json:"items"`
	TotalCount   int        `json:"total_count"`
}

type Codelist struct {
	ID          string
	Name        string
	IsHierarchy bool
	InDev       bool
	InBeta      bool
}

func (c *Codelist) CheckCodelist(devURL, betaURL string, wg *sync.WaitGroup) {
	defer wg.Done()
	c.InDev = codelistExists(devURL, c.ID)
	c.InBeta = codelistExists(betaURL, c.ID)
}

func codelistExists(url, id string) bool {
	client := rchttp.NewClient()

	uri := fmt.Sprintf("%s/code-lists/%s", url, id)
	resp, err := client.Get(context.Background(), uri)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false
	}

	var result *codelist.CodeList
	if err = json.Unmarshal(b, &result); err != nil || &result == nil {
		return false
	}

	return true
}
