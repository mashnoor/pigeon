package settings

import (
	"github.com/mashnoor/pigeon/structures"
	"gopkg.in/yaml.v2"
	"os"
)

var (
	SystemAppConfig structures.AppConfig
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

//
//func validateService(service structures.Service) {
//	if service.Name == "" {
//		log.Fatalln("Set service name")
//	}
//
//	if service.Endpoint == "" {
//		log.Fatal("Set service endpoint")
//	}
//
//	if service.Method == "" {
//		log.Fatal("Set service endpoint")
//	}
//
//	if service.ConsecutiveNotificationDelay == 0 {
//		log.Fatalln("Consecutive notification delay cannot be 0")
//	}
//
//	if service.MaxErrorCount == 0 {
//		log.Fatal("Service max error count cannot be 0")
//	}
//
//	if service.Timeout == 0 {
//		log.Fatalln("Service timeout cannot be 0")
//	}
//
//	if service.CheckInterval == 0 {
//		log.Fatal("Service check interval cannot be 0")
//	}
//}
//
//func validateConfig() {
//	// Validate Slack
//	if SystemAppConfig.SlackUrl == "" {
//		log.Fatalln("Slack URL missing")
//	}
//	//Validate Redis
//	if SystemAppConfig.RedisHost == "" && SystemAppConfig.RedisPort == "" {
//		log.Fatalln("Invalid redis settings")
//	}
//
//	// Validate service existance
//	if SystemAppConfig.Services == nil || len(SystemAppConfig.Services) == 0 {
//		log.Fatal("No services defined")
//	}
//
//	// Validate service
//	for _, svc := range SystemAppConfig.Services {
//		validateService(svc)
//	}
//}

func readConfigFile() string {
	dat, err := os.ReadFile("services.yaml")
	check(err)

	return string(dat)
}

func LoadAppConfig() {
	config := readConfigFile()
	//fmt.Println(config)
	err := yaml.Unmarshal([]byte(config), &SystemAppConfig)
	check(err)
	//validateConfig()
}
