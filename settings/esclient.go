package settings

import "github.com/elastic/go-elasticsearch/v7"

var (
	ESclient *elasticsearch.Client
)

func LoadESClient() {
	cfg := elasticsearch.Config{
		Addresses: []string{
			SystemAppConfig.ElasticSearchUrl,
		},
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		panic("Error loading Elastic Search Client")
	}
	ESclient = es
}
