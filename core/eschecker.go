package core

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/mashnoor/pigeon/helpers"
	"github.com/mashnoor/pigeon/settings"
	"github.com/mashnoor/pigeon/structures"
	"log"
	"sync"
	"time"
)

func execute(service *structures.Service, wg *sync.WaitGroup) {
	for true {
		generateSummary(service)
		time.Sleep(time.Second * service.NotificationInterval)

	}
	wg.Done()
}

func generateSummary(service *structures.Service) {

	currentTime := time.Now()
	currentTimeUTC := currentTime.UTC()

	//currentTimeStr := currentTime.Format("2006-01-02 15:04:05")
	currentTimeStr := currentTime.Format("15:04:05")

	checkPointTimeStr := currentTime.Add(-time.Second * service.NotificationInterval).Format("15:04:05")

	checkpointTimeUTCStr := currentTimeUTC.Add(-time.Second * service.NotificationInterval).Format("2006-01-02T15:04:05.000Z")

	totalSuccessHits := getTotalHits(service.KubernetesServiceName, service.SuccessMessage, checkpointTimeUTCStr)
	//totalFailureHits := getTotalHits(service.KubernetesServiceName, service.FailureMessage, checkpointTimeUTCStr)

	//totalRecords := totalSuccessHits + totalFailureHits
	//successP := (float64(totalSuccessHits) / float64(totalRecords)) * 100
	//successPercentage := fmt.Sprintf("%.2f", successP)
	//failurePercentage := fmt.Sprintf("%.2f", 100-successP)

	//slackMsg := fmt.Sprintf("*:bird: Pigeon Got Your Summary*\n*Service Name:* %s\n*Time Range:* %s to %s\n*Total Successful Transactions:* %d\n*Total Failed Transactions:* %d\n*Percentage:* Success: %s Failure: %s\n", service.Name, checkPointTimeStr, currentTimeStr, totalSuccessHits, totalFailureHits, successPercentage, failurePercentage)
	slackMsg := fmt.Sprintf("*:bird: Pigeon Got Your Summary*\n*Service Name:* %s\n*Time Range:* %s to %s\n*Total Successful Transactions:* %d", service.Name, checkPointTimeStr, currentTimeStr, totalSuccessHits)
	helpers.SendSlackMessage(slackMsg)
	fmt.Println(slackMsg)
	//fmt.Println(totalSuccessHits, totalFailureHits, successPercentage, failurePercentage)

}

func getTotalHits(label, logMsg, checkPointTime string) int {
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
							"gte": checkPointTime,
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

	return int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64))
}
