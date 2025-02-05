package scrape_gif

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"giphyscraper/giphy_proto"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"google.golang.org/protobuf/proto"
)

var ctx = context.Background()

var rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
})

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

func Scrape(c *gin.Context) {
	searchTerm := c.Query("q")
	outPref := c.Query("o")
	limit := 25

	cacheKey := "giphy:" + searchTerm

	cachedData, err := rdb.Get(ctx, cacheKey).Result()
	if err == nil {
		log.Println("Cache hit! Returning cached data.")
		var cachedResults []GiphyResult
		if err := json.Unmarshal([]byte(cachedData), &cachedResults); err == nil {
			sendResponse(c, cachedResults, outPref)
			return
		}
	}

	err = godotenv.Load()
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
	cacheData, _ := json.Marshal(prettifiedResults)
	rdb.Set(ctx, cacheKey, cacheData, 10*time.Minute)

	sendResponse(c, prettifiedResults, outPref)
}

func sendResponse(c *gin.Context, results []GiphyResult, outPref string) {
	JSONoutputFileName := "./out/giphy_results.json"
	XMLoutputFileName := "./out/giphy_results.xml"
	ProtoBufoutputFileName := "./out/giphy_results.pb"

	if outPref == "json" {
		writeJSONOutput(JSONoutputFileName, results)
		c.JSON(200, gin.H{"message": "Wrote JSON file"})
	} else if outPref == "xml" {
		writeXMLOutput(XMLoutputFileName, results)
		c.JSON(200, gin.H{"message": "Wrote XML file"})
	} else if outPref == "protobuf" {
		writeProtobufOutput(ProtoBufoutputFileName, results)
		c.JSON(200, gin.H{"message": "Wrote Protobuf file"})
	} else {
		c.JSON(400, gin.H{"error": "Invalid output preference"})
	}
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
	giphyResults := GiphyResults{Results: results}

	xmlData, err := xml.MarshalIndent(giphyResults, "", "  ")
	if err != nil {
		log.Fatalf("Error creating XML: %v\n", err)
	}

	xmlData = []byte(xml.Header + string(xmlData))

	err = os.WriteFile(XMLoutputFileName, xmlData, 0644)
	if err != nil {
		log.Fatalf("Error writing to XML file: %v\n", err)
	}
}

func writeProtobufOutput(ProtoBufoutputFileName string, results []GiphyResult) {
	giphyResults := &giphy_proto.GiphyResults{
		Results: make([]*giphy_proto.GiphyResult, len(results))}

	for i, result := range results {
		giphyResults.Results[i] = &giphy_proto.GiphyResult{
			Id:       result.ID,
			Url:      result.URL,
			Username: result.Username,
			Title:    result.Title,
		}
	}

	out, err := proto.Marshal(giphyResults)
	if err != nil {
		log.Fatalf("Failed to marshal protobuf: %v", err)
	}

	err = os.WriteFile(ProtoBufoutputFileName, out, 0644)
	if err != nil {
		log.Fatalf("Failed to write protobuf to file: %v", err)
	}

}

func getStringValue(m map[string]interface{}, key string) string {
	if val, ok := m[key].(string); ok {
		return val
	}
	return ""
}
