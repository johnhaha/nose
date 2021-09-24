package nosedata

type CreateObjectRes struct {
	Object         string   `json:"object"`
	ID             string   `json:"id"`
	CreatedTime    string   `json:"created_time"`
	LastEditedTime string   `json:"last_edited_time"`
	Parent         PageInfo `json:"parent"`
	Archived       bool     `json:"archived"`
	URL            string   `json:"url"`
}

type AppendBlockRes struct {
	Object  string      `json:"object"`
	Results []BlockInfo `json:"results"`
	HasMore bool        `json:"has_more"`
}
