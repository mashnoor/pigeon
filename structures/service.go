package structures

import "time"

type Service struct {
	Name                  string `yaml:"name"`
	KubernetesServiceName string `yaml:"kubernetesServiceName"`
	//ContainerName         string        `yaml:"containerName"`
	SuccessMessage       string        `yaml:"successMessage"`
	FailureMessage       string        `yaml:"failureMessage"`
	NotificationInterval time.Duration `yaml:"notificationInterval"`
}
