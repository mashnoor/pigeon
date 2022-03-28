package structures

type AppConfig struct {
	Services         []Service `yaml:"services"`
	SlackUrl         string    `yaml:"slackUrl"`
	ElasticSearchUrl string    `yaml:"elasticSearchUrl"`
	RedisHost        string    `yaml:"redisHost"`
	RedisPort        string    `yaml:"redisPort"`
}
