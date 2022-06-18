package nosedata

type PageInfo struct {
	PageType string `json:"-"`
	PageID   string `json:"page_id"`
}

type DBInfo struct {
	PageType   string `json:"-"`
	DatabaseID string `json:"database_id"`
}
