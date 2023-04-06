package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/elastic/go-elasticsearch/v7"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

// Index 在索引中创建或更新文档 索引不存在的情况下，会自动创建索引。
// 默认的type（类型）是doc，下面是指定doc类型创建添加的。
func Index() {
	addresses := []string{"http://127.0.0.1:9200", "http://127.0.0.1:9201"}
	config := elasticsearch.Config{
		Addresses: addresses,
		Username:  "",
		Password:  "",
		CloudID:   "",
		APIKey:    "",
	}
	// new client
	es, err := elasticsearch.NewClient(config)
	failOnError(err, "Error creating the client")
	// Index creates or updates a document in an index
	var buf bytes.Buffer
	doc := map[string]interface{}{
		"title":   "你看到外面的世界是什么样的？",
		"content": "外面的世界真的很精彩",
	}
	if err := json.NewEncoder(&buf).Encode(doc); err != nil {
		failOnError(err, "Error encoding doc")
	}
	res, err := es.Index("demo", &buf, es.Index.WithDocumentType("doc"))
	if err != nil {
		failOnError(err, "Error Index response")
	}
	defer res.Body.Close()
	fmt.Println(res.String())
}

// Search 搜索
func Search() {
	addresses := []string{"http://127.0.0.1:9200", "http://127.0.0.1:9201"}
	config := elasticsearch.Config{
		Addresses: addresses,
		Username:  "",
		Password:  "",
		CloudID:   "",
		APIKey:    "",
	}
	// new client
	es, err := elasticsearch.NewClient(config)
	failOnError(err, "Error creating the client")
	// info
	res, err := es.Info()
	failOnError(err, "Error getting response")
	fmt.Println(res.String())
	// search - highlight
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"title": "中国",
			},
		},
		"highlight": map[string]interface{}{
			"pre_tags":  []string{"<font color='red'>"},
			"post_tags": []string{"</font>"},
			"fields": map[string]interface{}{
				"title": map[string]interface{}{},
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		failOnError(err, "Error encoding query")
	}
	// Perform the search request.
	res, err = es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex("demo"),
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
	)
	if err != nil {
		failOnError(err, "Error getting response")
	}
	defer res.Body.Close()
	fmt.Println(res.String())
}

// DeleteByQuery 通过匹配条件删除文档
func DeleteByQuery() {
	addresses := []string{"http://127.0.0.1:9200", "http://127.0.0.1:9201"}
	config := elasticsearch.Config{
		Addresses: addresses,
		Username:  "",
		Password:  "",
		CloudID:   "",
		APIKey:    "",
	}
	// new client
	es, err := elasticsearch.NewClient(config)
	failOnError(err, "Error creating the client")
	// DeleteByQuery deletes documents matching the provided query
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"title": "外面",
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		failOnError(err, "Error encoding query")
	}
	index := []string{"demo"}
	res, err := es.DeleteByQuery(index, &buf)
	if err != nil {
		failOnError(err, "Error delete by query response")
	}
	defer res.Body.Close()
	fmt.Println(res.String())
}

// Delete通过_id删除文档
func Delete() {
	addresses := []string{"http://127.0.0.1:9200", "http://127.0.0.1:9201"}
	config := elasticsearch.Config{
		Addresses: addresses,
		Username:  "",
		Password:  "",
		CloudID:   "",
		APIKey:    "",
	}
	// new client
	es, err := elasticsearch.NewClient(config)
	failOnError(err, "Error creating the client")
	// Delete removes a document from the index
	res, err := es.Delete("demo", "POcKSHIBX-ZyL96-ywQO")
	if err != nil {
		failOnError(err, "Error delete by id response")
	}
	defer res.Body.Close()
	fmt.Println(res.String())
}

// Create 添加文档（需要指定id，id已存在返回409）
func Create() {
	addresses := []string{"http://127.0.0.1:9200", "http://127.0.0.1:9201"}
	config := elasticsearch.Config{
		Addresses: addresses,
		Username:  "",
		Password:  "",
		CloudID:   "",
		APIKey:    "",
	}
	// new client
	es, err := elasticsearch.NewClient(config)
	failOnError(err, "Error creating the client")
	// Create creates a new document in the index.
	// Returns a 409 response when a document with a same ID already exists in the index.
	var buf bytes.Buffer
	doc := map[string]interface{}{
		"title":   "你看到外面的世界是什么样的？",
		"content": "外面的世界真的很精彩",
	}
	if err := json.NewEncoder(&buf).Encode(doc); err != nil {
		failOnError(err, "Error encoding doc")
	}
	res, err := es.Create("demo", "esd", &buf, es.Create.WithDocumentType("doc"))
	if err != nil {
		failOnError(err, "Error create response")
	}
	defer res.Body.Close()
	fmt.Println(res.String())
}

// Get 通过id获取文档
func Get() {
	addresses := []string{"http://127.0.0.1:9200", "http://127.0.0.1:9201"}
	config := elasticsearch.Config{
		Addresses: addresses,
		Username:  "",
		Password:  "",
		CloudID:   "",
		APIKey:    "",
	}
	// new client
	es, err := elasticsearch.NewClient(config)
	failOnError(err, "Error creating the client")
	res, err := es.Get("demo", "esd")
	if err != nil {
		failOnError(err, "Error get response")
	}
	defer res.Body.Close()
	fmt.Println(res.String())
}

// Update 通过_id更新文档
func Update() {
	addresses := []string{"http://127.0.0.1:9200", "http://127.0.0.1:9201"}
	config := elasticsearch.Config{
		Addresses: addresses,
		Username:  "",
		Password:  "",
		CloudID:   "",
		APIKey:    "",
	}
	// new client
	es, err := elasticsearch.NewClient(config)
	failOnError(err, "Error creating the client")
	// Update updates a document with a script or partial document.
	var buf bytes.Buffer
	doc := map[string]interface{}{
		"doc": map[string]interface{}{
			"title":   "更新你看到外面的世界是什么样的？",
			"content": "更新外面的世界真的很精彩",
		},
	}
	if err := json.NewEncoder(&buf).Encode(doc); err != nil {
		failOnError(err, "Error encoding doc")
	}
	res, err := es.Update("demo", "esd", &buf, es.Update.WithDocumentType("doc"))
	if err != nil {
		failOnError(err, "Error Update response")
	}
	defer res.Body.Close()
	fmt.Println(res.String())
}

// UpdateByQuery 通过匹配条件更新文档
func UpdateByQuery() {
	addresses := []string{"http://127.0.0.1:9200", "http://127.0.0.1:9201"}
	config := elasticsearch.Config{
		Addresses: addresses,
		Username:  "",
		Password:  "",
		CloudID:   "",
		APIKey:    "",
	}
	// new client
	es, err := elasticsearch.NewClient(config)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	// UpdateByQuery performs an update on every document in the index without changing the source,
	// for example to pick up a mapping change.
	index := []string{"demo"}
	var buf bytes.Buffer
	doc := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"title": "外面",
			},
		},
		// 根据搜索条件更新title
		/*
		   "script": map[string]interface{}{
		       "source": "ctx._source['title']='更新你看到外面的世界是什么样的？'",
		   },
		*/
		// 根据搜索条件更新title、content
		/*
		   "script": map[string]interface{}{
		       "source": "ctx._source=params",
		       "params": map[string]interface{}{
		           "title": "外面的世界真的很精彩",
		           "content": "你看到外面的世界是什么样的？",
		       },
		       "lang": "painless",
		   },
		*/
		// 根据搜索条件更新title、content
		"script": map[string]interface{}{
			"source": "ctx._source.title=params.title;ctx._source.content=params.content;",
			"params": map[string]interface{}{
				"title":   "看看外面的世界真的很精彩",
				"content": "他们和你看到外面的世界是什么样的？",
			},
			"lang": "painless",
		},
	}
	if err := json.NewEncoder(&buf).Encode(doc); err != nil {
		failOnError(err, "Error encoding doc")
	}
	res, err := es.UpdateByQuery(
		index,
		es.UpdateByQuery.WithDocumentType("doc"),
		es.UpdateByQuery.WithBody(&buf),
		es.UpdateByQuery.WithContext(context.Background()),
		es.UpdateByQuery.WithPretty(),
	)
	if err != nil {
		failOnError(err, "Error Update response")
	}
	defer res.Body.Close()
	fmt.Println(res.String())
}
