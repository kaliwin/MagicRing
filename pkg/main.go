package main

import (
	"fmt"
	"github.com/blevesearch/bleve/v2"
	"log"
)

func CreateIndex() {
	indexMapping := bleve.NewIndexMapping()
	textFieldMapping := bleve.NewTextFieldMapping()
	textFieldMapping.Analyzer = "standard" // 使用标准分词器
	indexMapping.DefaultMapping.AddFieldMappingsAt("content", textFieldMapping)

	//indexMapping.AnalyzeText()

	_, err := bleve.New("index", indexMapping)
	if err != nil {
		log.Println(err)
	}
}

func AddData() {
	index, _ := bleve.Open("index")

	index.Index("1", "hello world 另一方面，增加更多的服务器来完成系统扩展，称为水平扩展或者向外扩展（Horizontal Scale/Scaling Out）。尽管ES能够利用更强劲的硬件，但是垂直扩展毕竟还是有它的极限。真正的可扩展性来自于水平扩展，通过向集群中添加更多的节点来分担负载，增加可靠性。ES天生就是分布式的，它知道如何管理多个节点来完成扩展和实现高可用性。意味应用不需要做任何的改动。\n")
	index.Index("2", "helloElasticsearch 是一个分布式、高扩展、高实时的搜索与数据分析引擎。它能很方便的使大量数据具有搜索、分析和探索的能力。充分利用Elasticsearch的水平伸缩性，能使数据在生产环境变得更有价值。Elasticsearch 的实现原理主要分为以下几个步骤，首先用户将数据提交到Elasticsearch 数据库中，再通过分词控制器去将对应的语句分词，将其权重和分词结果一并存入数据，当用户搜索数据时候，再根据权重将结果排名，打分，再将返回结果呈现给用户。\n")
}

func Search() {
	index, _ := bleve.Open("index")

	search, err := index.Search(bleve.NewSearchRequest(bleve.NewMatchQuery("龙")))
	if err != nil {
		log.Println(err)
	}

	for _, hit := range search.Hits {
		fmt.Println(hit.ID)
		fmt.Println(hit.Score)

	}
}

func Td() {
	indexMapping := bleve.NewIndexMapping()
	//textFieldMapping := bleve.NewTextFieldMapping()
	//textFieldMapping.Analyzer = "standard" // 使用标准分词器
	//indexMapping.DefaultMapping.AddFieldMappingsAt("content", textFieldMapping)

	_, err := indexMapping.AnalyzeText("standard", []byte("cyvk"))
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(indexMapping.CustomAnalysis)
}

func main() {
	Td()

	fmt.Println("success")
}
