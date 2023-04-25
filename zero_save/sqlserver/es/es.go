package es

import (
	"github.com/elastic/go-elasticsearch/v7"
	"log"
)

func init() {
	es, _ := elasticsearch.NewDefaultClient()
	log.Println(elasticsearch.Version)
	log.Println(es.Info())
}
