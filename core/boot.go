package core

import (
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
	settings.LoadESClient()
	settings.InitRedis()
	//initNotificationSystem()

}
