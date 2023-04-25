package es

import (
	"log"

	"github.com/elastic/go-elasticsearch/v7"
)

type IES interface {
	Get() (interface{}, error)
}

type ES struct {
	es *elasticsearch.Client
}

func NewES() IES {
	es, _ := elasticsearch.NewDefaultClient()
	log.Println(elasticsearch.Version)
	log.Println(es.Info())

	return &ES{
		es: es,
	}
}

func (e *ES) Get() (interface{}, error) {
	return nil, nil
}
