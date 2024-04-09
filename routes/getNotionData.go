package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type changeResult struct {
	Object         string `json:"object"`
	Results        []any  `json:"results"`
	NextCursor     any    `json:"next_cursor"`
	HasMore        bool   `json:"has_more"`
	Type           string `json:"type"`
	PageOrDatabase struct {
	} `json:"page_or_database"`
	RequestID string `json:"request_id"`
}

type LastEditedTime struct {
	After time.Time `json:"after"`
}

type Filters struct {
	Timestamp      string         `json:"timestamp"`
	LastEditedTime LastEditedTime `json:"last_edited_time"`
}
type Filter struct {
	Filters Filters `json:"filter"`
}

func GetNotionData() {
	var bearer = "Bearer " + os.Getenv("NOTION_KEY")
	timeLapse := time.Now().Add(-3100 * time.Minute)
	payload, err := json.Marshal(Filter{
		Filters: Filters{
			Timestamp: "last_edited_time",
			LastEditedTime: LastEditedTime{
				After: timeLapse,
			},
		},
	})
	if err != nil {
		fmt.Printf("marshal failed, error: %v\n", err)
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("https://api.notion.com/v1/databases/%v/query", os.Getenv("NOTION_PAGE_ID")), bytes.NewBuffer(payload))
	if err != nil {
		fmt.Printf("http new failed, error: %v\n", err)
	}
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Notion-Version", "2022-06-28")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("http post failed, error: %v\n", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("ReadAll failed, error: %v\n", err)
	}
	var changeResult changeResult
	json.Unmarshal(body, &changeResult)
	fmt.Print(PrettyPrint(changeResult.Results))
	// for i := 0; i < len(changeResult.Results); i++ {
	// 	// if changeResult.Results[i] != nil {
	// 	// 	continue;
	// 	// }
	// }
}

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}