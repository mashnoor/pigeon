package core

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/mashnoor/pigeon/settings"
	"log"
)

func getTotalHits(label, logMsg string) {
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": []map[string]interface{}{
					{
						"match": map[string]interface{}{
							"kubernetes.labels.app": label,
						},
					},
					{
						"match_phrase": map[string]string{
							"log": logMsg,
						},
					},
				},

				"filter": map[string]interface{}{
					"range": map[string]interface{}{
						"@timestamp": map[string]string{
							"gte": "2022-03-27T00:55:47.165Z",
							"lte": "2022-03-27T00:55:47.165Z",
						},
					},
				},
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	fmt.Println(buf.String())

	var r map[string]interface{}
	res, err := settings.ESclient.Search(
		settings.ESclient.Search.WithContext(context.Background()),
		settings.ESclient.Search.WithIndex("logstash-*"),
		settings.ESclient.Search.WithBody(&buf),
		settings.ESclient.Search.WithTrackTotalHits(true),
		settings.ESclient.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}

	log.Printf(
		"[%s] %d hits; took: %dms",
		res.Status(),
		int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
		int(r["took"].(float64)),
	)
}
