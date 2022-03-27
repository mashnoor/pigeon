package core

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/mashnoor/pigeon/settings"
	"log"
	"sync"
)

func initNotificationSystem() {
	var wg sync.WaitGroup

	for _, service := range settings.SystemAppConfig.Services {
		currentService := service
		//fmt.Println(service.Endpoint)
		go checkHTTPHealth(&currentService, &wg)
		wg.Add(1)

	}

	wg.Wait()
}

func BootApp() {
	//settings.LoadAppConfig()
	//initNotificationSystem()
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://172.17.17.44:9200",
		},
	}
	es, err := elasticsearch.NewClient(cfg)

	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": []map[string]interface{}{
					{
						"match": map[string]interface{}{
							"kubernetes.labels.app": "ucbpg",
						},
					},
					{
						"match_phrase": map[string]string{
							"log": "Money has been successfully added",
						},
					},
				},

				"filter": map[string]interface{}{
					"range": map[string]interface{}{
						"@timestamp": map[string]string{
							"gte": "2022-03-27T00:55:47.165Z",
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
	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex("logstash-*"),
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
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
	// Print the response status, number of results, and request duration.
	log.Printf(
		"[%s] %d hits; took: %dms",
		res.Status(),
		int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
		int(r["took"].(float64)),
	)

}
