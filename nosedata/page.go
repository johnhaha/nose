package nosedata

type PageInfo struct {
	PageType string `json:"-"`
	PageID   string `json:"page_id"`
}
