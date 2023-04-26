package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/elastic/go-elasticsearch/v7"
)

// type ES struct {
// 	es *elasticsearch.Client
// }

func NewES() (*elasticsearch.Client, error) {
	addresses := []string{"http://127.0.0.1:9200", "http://127.0.0.1:9201"}
	config := elasticsearch.Config{
		Addresses: addresses,
		Username:  "",
		Password:  "",
		CloudID:   "",
		APIKey:    "",
	}
	return elasticsearch.NewClient(config)
}

// Index 在索引中创建或更新文档 索引不存在的情况下，会自动创建索引。
// 默认的type（类型）是doc，下面是指定doc类型创建添加的。
func Index() {
	es, err := NewES()
	if err != nil {
		log.Fatal("Error creating the client", err)
	}
	// Index creates or updates a document in an index
	var buf bytes.Buffer
	doc := map[string]interface{}{
		"title":   "你看到外面的世界是什么样的？",
		"content": "外面的世界真的很精彩",
	}
	if err := json.NewEncoder(&buf).Encode(doc); err != nil {
		log.Fatal("Error encoding doc: ", err)
	}
	res, err := es.Index("demo", &buf, es.Index.WithDocumentType("doc"))
	if err != nil {
		log.Fatal("[Index] error: ", err)
	}
	defer res.Body.Close()
	fmt.Println(res.String())
}

// Search 搜索
func Search() {
	es, err := NewES()
	if err != nil {
		log.Fatal("Error creating the client", err)
	}
	// info
	res, err := es.Info()
	if err != nil {
		log.Fatal("Error getting response", err)
	}
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
		log.Fatal("Error encoding query", err)
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
		log.Fatal("Error getting response", err)
	}
	defer res.Body.Close()
	fmt.Println(res.String())
}

// DeleteByQuery 通过匹配条件删除文档
func DeleteByQuery() {
	es, err := NewES()
	if err != nil {
		log.Fatal("Error creating the client", err)
	}
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
		log.Fatal("Error encoding query", err)
	}
	index := []string{"demo"}
	res, err := es.DeleteByQuery(index, &buf)
	if err != nil {
		log.Fatal(err, "Error delete by query response", err)
	}
	defer res.Body.Close()
	fmt.Println(res.String())
}

// Delete通过_id删除文档
func Delete() {
	es, err := NewES()
	if err != nil {
		log.Fatal("Error creating the client", err)
	}
	// Delete removes a document from the index
	res, err := es.Delete("demo", "POcKSHIBX-ZyL96-ywQO")
	if err != nil {
		log.Fatal("Error delete by id response", err)
	}
	defer res.Body.Close()
	fmt.Println(res.String())
}

// Create 添加文档（需要指定id，id已存在返回409）
func Create() {
	es, err := NewES()
	if err != nil {
		log.Fatal("Error creating the client", err)
	}

	// Create creates a new document in the index.
	// Returns a 409 response when a document with a same ID already exists in the index.
	var buf bytes.Buffer
	doc := map[string]interface{}{
		"title":   "你看到外面的世界是什么样的？",
		"content": "外面的世界真的很精彩",
	}
	if err := json.NewEncoder(&buf).Encode(doc); err != nil {
		log.Fatal("Error encoding doc", err)
	}
	res, err := es.Create("demo", "esd", &buf, es.Create.WithDocumentType("doc"))
	if err != nil {
		log.Fatal("Error create response", err)
	}
	defer res.Body.Close()
	fmt.Println(res.String())
}

// Get 通过id获取文档
func Get() {
	es, err := NewES()
	if err != nil {
		log.Fatal("Error creating the client", err)
	}

	res, err := es.Get("demo", "esd")
	if err != nil {
		log.Fatal("Error get response", err)
	}
	defer res.Body.Close()
	fmt.Println(res.String())
}

// Update 通过_id更新文档
func Update() {
	es, err := NewES()
	if err != nil {
		log.Fatal("Error creating the client", err)
	}
	// Update updates a document with a script or partial document.
	var buf bytes.Buffer
	doc := map[string]interface{}{
		"doc": map[string]interface{}{
			"title":   "更新你看到外面的世界是什么样的？",
			"content": "更新外面的世界真的很精彩",
		},
	}
	if err := json.NewEncoder(&buf).Encode(doc); err != nil {
		log.Fatal("Error encoding doc", err)
	}
	res, err := es.Update("demo", "esd", &buf, es.Update.WithDocumentType("doc"))
	if err != nil {
		log.Fatal("Error Update response", err)
	}
	defer res.Body.Close()
	fmt.Println(res.String())
}

// UpdateByQuery 通过匹配条件更新文档
func UpdateByQuery() {
	es, err := NewES()
	if err != nil {
		log.Fatal("Error creating the client", err)
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
		log.Fatal("Error encoding doc", err)
	}
	res, err := es.UpdateByQuery(
		index,
		es.UpdateByQuery.WithDocumentType("doc"),
		es.UpdateByQuery.WithBody(&buf),
		es.UpdateByQuery.WithContext(context.Background()),
		es.UpdateByQuery.WithPretty(),
	)
	if err != nil {
		log.Fatal("Error Update response", err)
	}
	defer res.Body.Close()
	fmt.Println(res.String())
}

func main() {
	Index()
}
