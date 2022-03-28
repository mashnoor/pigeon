package structures

type AppConfig struct {
	Services         []Service `yaml:"services"`
	SlackUrl         string    `yaml:"slackUrl"`
	ElasticSearchUrl string    `yaml:"elasticSearchUrl"`
}
