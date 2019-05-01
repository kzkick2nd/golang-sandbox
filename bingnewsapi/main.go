package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/davecgh/go-spew/spew"
)

// This struct formats the answer provided by the Bing News Search API.
type NewsAnswer struct {
	ReadLink     string `json: "readLink"`
	QueryContext struct {
		OriginalQuery string `json: "originalQuery"`
		AdultIntent   bool   `json: "adultIntent"`
	} `json: "queryContext"`
	TotalEstimatedMatches int `json: totalEstimatedMatches"`
	Sort                  []struct {
		Name       string `json: "name"`
		ID         string `json: "id"`
		IsSelected bool   `json: "isSelected"`
		URL        string `json: "url"`
	} `json: "sort"`
	Value []struct {
		Name  string `json: "name"`
		URL   string `json: "url"`
		Image struct {
			Thumbnail struct {
				ContentUrl string `json: "thumbnail"`
				Width      int    `json: "width"`
				Height     int    `json: "height"`
			} `json: "thumbnail"`
			Description string `json: "description"`
			Provider    []struct {
				Type string `json: "_type"`
				Name string `json: "name"`
			} `json: "provider"`
			DatePublished string `json: "datePublished"`
		} `json: "image"`
	} `json: "value"`
}

func main() {
	// Verify the endpoint URI and replace the token string with a valid subscription key.
	const endpoint = "https://japaneast.api.cognitive.microsoft.com/bing/v7.0/news/search"
	token := os.Getenv("AZURE_COGNITIVE_KEY")
	searchTerm := "Microsoft Cognitive Services"

	// Declare a new GET request.
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		panic(err)
	}

	// The rest of the code in this example goes here in the main function.

	// Add the query to the request.
	param := req.URL.Query()
	param.Add("q", searchTerm)
	req.URL.RawQuery = param.Encode()

	// Insert the subscription-key header.
	req.Header.Add("Ocp-Apim-Subscription-Key", token)

	// Instantiate a client.
	client := new(http.Client)

	// Send the request to Bing.
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	// Close the connection.
	defer resp.Body.Close()

	// Read the results
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	// Create a new answer object
	ans := new(NewsAnswer)
	err = json.Unmarshal(body, &ans)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print("Output of BingAnswer: \r\n\r\n")

	// Iterate over search results and print the result name and URL.
	for _, result := range ans.Value {
		spew.Dump(result.Name, result.URL)
	}
}
