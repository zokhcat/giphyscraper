package scrape_gif

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type GiphyResult struct {
	ID       string `json:"id"`
	URL      string `json:"url"`
	Username string `json:"username"`
	Title    string `json:"title"`
}

type GiphyResults struct {
	XMLName xml.Name      `xml:"giphyResults"`
	Results []GiphyResult `xml:"giphyImages"`
}

func Scrape() {
	searchTerm := "cats"
	limit := 25
	JSONoutputFileName := "giphy_results.json"
	XMLoutputFileName := "giphy_results.xml"
	// ProtoBufoutputFileName := "giphy_results.pb"

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("GIPHY_API_KEY")
	if apiKey == "" {
		log.Fatal("GIPHY_API_KEY not set in .env file")
	}

	url := fmt.Sprintf("https://api.giphy.com/v1/gifs/search?api_key=%s&q=%s&limit=%d", apiKey, searchTerm, limit)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to make GET request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatalf("Error parsing JSON: %v\n", err)
	}

	data, ok := result["data"].([]interface{})
	if !ok {
		log.Fatal("Error: 'data' field not found or not an array")
	}

	var prettifiedResults []GiphyResult
	for _, item := range data {
		gif, ok := item.(map[string]interface{})
		if !ok {
			continue
		}

		prettifiedResults = append(prettifiedResults, GiphyResult{
			ID:       getStringValue(gif, "id"),
			URL:      getStringValue(gif, "url"),
			Username: getStringValue(gif, "username"),
			Title:    getStringValue(gif, "title"),
		})
	}

	writeJSONOutput(JSONoutputFileName, prettifiedResults)

	writeXMLOutput(XMLoutputFileName, prettifiedResults)

}

func writeJSONOutput(JSONoutputFilename string, results []GiphyResult) {
	prettifiedJSON, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		log.Fatalf("Error creating prettified JSON: %v\n", err)
	}

	err = os.WriteFile(JSONoutputFilename, prettifiedJSON, 0644)
	if err != nil {
		log.Fatalf("Error writing to JSON file: %v\n", err)
	}
}

func writeXMLOutput(XMLoutputFileName string, results []GiphyResult) {
	// Wrap the results in the GiphyResults struct
	giphyResults := GiphyResults{Results: results}

	xmlData, err := xml.MarshalIndent(giphyResults, "", "  ")
	if err != nil {
		log.Fatalf("Error creating XML: %v\n", err)
	}

	// Add XML header
	xmlData = []byte(xml.Header + string(xmlData))

	err = os.WriteFile(XMLoutputFileName, xmlData, 0644)
	if err != nil {
		log.Fatalf("Error writing to XML file: %v\n", err)
	}
}

func writeProtobufOutput(writeProtobufOutput string, results []GiphyResult) {

}

func getStringValue(m map[string]interface{}, key string) string {
	if val, ok := m[key].(string); ok {
		return val
	}
	return ""
}
