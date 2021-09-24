package noseapi

const (
	CreatePageApi              = "https://api.notion.com/v1/pages"
	CreateDatabaseApi          = "https://api.notion.com/v1/databases"
	UpdateDatabaseApi ParamApi = "https://api.notion.com/v1/databases"
	BlockApi          ParamApi = "https://api.notion.com/v1/blocks"
	ApiVersion                 = "2021-08-16"
)

type ParamApi string

func (api ParamApi) AddParam(param string) ParamApi {
	return ParamApi(string(api) + "/" + param)
}

func (api ParamApi) Value() string {
	return string(api)
}
