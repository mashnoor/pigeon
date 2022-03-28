package core

import (
	"github.com/mashnoor/pigeon/settings"
	"sync"
)

func initSummarySystem() {
	var wg sync.WaitGroup

	for _, service := range settings.SystemAppConfig.Services {
		currentService := service
		generateSummary(&currentService, &wg)
		//fmt.Println(service.Endpoint)
		//go checkHTTPHealth(&currentService, &wg)
		//wg.Add(1)

	}

	//wg.Wait()
}

func BootApp() {
	settings.LoadAppConfig()
	settings.LoadESClient()
	settings.InitRedis()
	initSummarySystem()

}
