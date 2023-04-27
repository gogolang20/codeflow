package es

import (
	"log"

	"github.com/elastic/go-elasticsearch/v7"
)

type IElasticsearch interface {
	Get() (interface{}, error)
}

type Elasticsearch struct {
	es *elasticsearch.Client
}

func NewES() IElasticsearch {
	addresses := []string{"http://127.0.0.1:9200", "http://127.0.0.1:9201", "http://127.0.0.1:9202", "http://127.0.0.1:9203", "http://127.0.0.1:9204"}
	config := elasticsearch.Config{
		Addresses: addresses,
		Username:  "",
		Password:  "",
		CloudID:   "",
		APIKey:    "",
	}

	es, _ := elasticsearch.NewClient(config)

	log.Println(elasticsearch.Version)
	log.Println(es.Info())

	return &Elasticsearch{
		es: es,
	}
}

func (e *Elasticsearch) Get() (interface{}, error) {

	return nil, nil
}
