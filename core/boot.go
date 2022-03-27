package core

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/mashnoor/pigeon/settings"
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
	settings.LoadAppConfig()
	//initNotificationSystem()
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://172.17.17.44:9200",
		},
	}
	es, err := elasticsearch.NewClient(cfg)

}
