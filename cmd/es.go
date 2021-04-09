package main

import (
	"context"
	"fmt"
	"math"
	"strconv"
	"time"

	elasticV7 "gopkg.in/olivere/elastic.v6"

	elasticV6 "github.com/olivere/elastic/v7"

	"github.com/alazyer/gocodes/pkg/elasticsearch"
	"github.com/alazyer/gocodes/pkg/elasticsearch/model"
)

func main() {
	Search()
}

func genLogRequest() model.LogRequest {
	req := model.LogRequest{
		RootAccount:   "",
		RegionNames:   "",
		ProjectNames:  "",
		Nodes:         "",
		K8sNamespaces: "",
		AppNames:      "",
		PodNames:      "",
		Paths:         "",
		Sources:       "",
		QueryString:   "",
		ContainerID8s: "",
	}

	sizeStr := ""
	if sizeStr != "" {
		size, _ := strconv.ParseInt(sizeStr, 10, 0)
		req.Size = int(size)
	} else {
		req.Size = 20
	}
	pageNoStr := ""
	if pageNoStr != "" {
		page, _ := strconv.ParseInt(pageNoStr, 10, 0)
		req.Page = int(page)
	} else {
		req.Page = 1
	}
	start_time := ""
	if start_time != "" {
		req.StartTime, _ = strconv.ParseFloat(start_time, 10)
	} else {
		req.StartTime = float64(time.Now().AddDate(0, 0, -5).Unix())
	}
	end_time := ""
	if end_time != "" {
		req.EndTime, _ = strconv.ParseFloat(end_time, 10)
	} else {
		req.EndTime = float64(time.Now().Unix())
	}

	return req
}

func Search() {
	var LogType = "log"
	client, err := elasticsearch.NewClient()
	if err != nil {
		fmt.Println("failed to get the es client: ", err)
		return
	}
	req := genLogRequest()

	var response *model.LogResourceList

	indexNames := client.GetIndexNames(LogType, req.StartTime, req.EndTime)
	if len(indexNames) == 0 {
		response, _ = model.EsResultToLogResources(&elasticV6.SearchResult{})
	} else {
		var from int
		if req.Page < 1 {
			from = 0
		} else {
			from = (req.Page - 1) * req.Size
		}

		boolQuery := elasticV6.NewBoolQuery()
		boolQuery.Filter(req.GetQueries()...)

		searchService := client.Search().
			Query(boolQuery).
			From(from).
			Size(req.Size).
			Index(indexNames...).
			Pretty(true)

		if req.QueryString != "" {
			highlightField := elasticV6.NewHighlighterField("log_data").NumOfFragments(0)
			highlight := elasticV6.NewHighlight().Fields(highlightField).BoundaryChars("200")
			highlight.Encoder("html")
			searchService.Highlight(highlight)
		}

		searchResult, err := searchService.Do(context.Background())
		if err != nil {
			fmt.Println("errors when query es: ", err)
		}

		response, err = model.EsResultToLogResources(searchResult)
		if err != nil {
			fmt.Println("Failed to parse es search result, ", err.Error())
		}
	}

	response.TotalPage = int64(math.Min(math.Ceil(float64(response.TotalItems)/float64(req.Size)), 200))
	fmt.Println(response)
}

func generateBoolQuery(req model.LogRequest) *elasticV6.BoolQuery {
	boolQuery := elasticV6.NewBoolQuery()

	if len(req.RootAccount) > 0 {
		orgTermsQuery := elasticV6.NewTermsQuery("root_account.keyword", req.RootAccount)
		boolQuery.Filter(orgTermsQuery)
	}
	if len(req.ProjectNames) > 0 {
		projectTermsQuery := elasticV6.NewTermsQuery("project_name.keyword", req.ProjectNames)
		boolQuery.Filter(projectTermsQuery)
	}
	if len(req.RegionNames) > 0 {
		regionTermsQuery := elasticV6.NewTermsQuery("region_name.keyword", req.RegionNames)
		boolQuery.Filter(regionTermsQuery)
	}
	if len(req.Nodes) > 0 {
		nodeTermsQuery := elasticV6.NewTermsQuery("node.keyword", req.Nodes)
		boolQuery.Filter(nodeTermsQuery)
	}
	if len(req.K8sNamespaces) > 0 {
		k8sNsTermsQuery := elasticV6.NewTermsQuery("kubernetes_namespace.keyword", req.K8sNamespaces)
		boolQuery.Filter(k8sNsTermsQuery)
	}
	if len(req.AppNames) > 0 {
		appNameTermsQuery := elasticV6.NewTermsQuery("application_name.keyword", req.AppNames)
		boolQuery.Filter(appNameTermsQuery)
	}
	if len(req.PodNames) > 0 {
		podNameTermsQuery := elasticV6.NewTermsQuery("pod_name.keyword", req.Nodes)
		boolQuery.Filter(podNameTermsQuery)

	}
	if len(req.ContainerID8s) > 0 {
		containerIdTermsQuery := elasticV6.NewTermsQuery("container_id8.keyword", req.ContainerID8s)
		boolQuery.Filter(containerIdTermsQuery)

	}
	if len(req.Paths) > 0 {
		pathTermsQuery := elasticV6.NewTermsQuery("file_name.keyword", req.Paths)
		boolQuery.Filter(pathTermsQuery)
	}
	if len(req.Sources) > 0 {
		sourceTermsQuery := elasticV6.NewTermsQuery("source.keyword", req.Sources)
		boolQuery.Filter(sourceTermsQuery)
	}

	if req.StartTime > 0 && req.EndTime > 0 && req.StartTime <= req.EndTime {
		rangeQuery := elasticV6.NewRangeQuery("time").Gte(req.StartTime).Lt(req.EndTime)
		boolQuery.Filter(rangeQuery)
	}

	if len(req.QueryString) > 0 {
		matchQuery := elasticV6.NewMatchQuery("log_data", req.QueryString).Operator("AND")
		boolQuery.Filter(matchQuery)

		highlightField := elasticV6.NewHighlighterField("log_data").NumOfFragments(0)
		highlight := elasticV6.NewHighlight().Fields(highlightField).BoundaryChars("200")
		highlight.Encoder("html")
	}

	return boolQuery
}

func generateBaseRequest() *elasticV6.SearchService {
	client, _ := elasticsearch.NewClient()
	searchReq := elasticV6.NewSearchService(client.Client)
	searchReq = client.Search()
	res, _ := client.IndexExists("log-20191012").Do(context.Background())
	fmt.Printf("exist check: %v, %T\n", res, res)

	boolQuery := elasticV6.NewBoolQuery()

	orgTermQuery := elasticV6.NewTermQuery("root_account.keyword", "alauda")
	pathTermQuery := elasticV6.NewTermQuery("file_name.keyword", "stdout")
	boolQuery.Filter(orgTermQuery)
	boolQuery.Filter(pathTermQuery)

	searchReq.Query(boolQuery)

	return searchReq

}

func generateSearchRequest() *elasticV6.SearchService {
	searchReq := generateBaseRequest()

	orgTermQuery := elasticV6.NewTermQuery("root_account.keyword", "alauda")
	pathTermQuery := elasticV6.NewTermQuery("file_name.keyword", "stdout")
	boolQuery := elasticV6.NewBoolQuery()
	boolQuery.Filter(orgTermQuery)
	boolQuery.Filter(pathTermQuery)

	sorter := elasticV6.NewFieldSort("time").Desc()

	searchReq.Query(boolQuery)
	searchReq.Size(10)
	searchReq.FetchSource(true)
	searchReq.SortBy(sorter)

	return searchReq
}

func generateAggRequest() *elasticV6.SearchService {
	searchReq := generateBaseRequest()
	searchReq.Size(0)

	aggs := elasticV6.NewDateHistogramAggregation()
	aggs.Field("time")
	aggs.MinDocCount(0)
	aggs.Interval("1h")
	// aggs.ExtendedBounds(1, 100)
	// aggs.ExtendedBoundsMin(1)
	// aggs.ExtendedBoundsMax(100)

	searchReq.Aggregation("aggs_by_time", aggs)
	return searchReq
}

func generateExactRequest() *elasticV6.SearchService {
	searchReq := generateBaseRequest()

	timeSorter := elasticV6.NewFieldSort("time").Desc()
	docSorter := elasticV6.NewFieldSort("_doc").Asc()

	idsQuery := elasticV6.NewIdsQuery("log").Ids("JK3_vW0B-tSnL2hBf_KQ")
	constQuery := elasticV6.NewConstantScoreQuery(idsQuery)

	searchReq.Query(constQuery)
	searchReq.SortBy(timeSorter, docSorter)
	searchReq.FetchSource(true)
	searchReq.Size(1)

	return searchReq
}

func generateContextRequest() *elasticV6.SearchService {
	searchReq := generateBaseRequest()

	timeSorter := elasticV6.NewFieldSort("time").Desc()
	docSorter := elasticV6.NewFieldSort("_doc").Asc()

	idsQuery := elasticV6.NewIdsQuery("log").Ids("JK3_vW0B-tSnL2hBf_KQ")
	constQuery := elasticV6.NewConstantScoreQuery(idsQuery)

	searchReq.Query(constQuery)
	searchReq.SortBy(timeSorter, docSorter)
	searchReq.FetchSource(true)
	searchReq.Size(1)

	return searchReq
}
