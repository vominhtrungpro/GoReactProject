package crawldatamodel

type CrawlDataRequest struct {
	Url string
}

type CrawlDataResponse struct {
	Url []string
}
