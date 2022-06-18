package nose

import (
	"github.com/johnhaha/nose/noseapi"
	"github.com/johnhaha/nose/nosedata"
)

type NoseDBClient struct {
	Token string `json:"token"`
	DBID  string `json:"dbID"`
}

func NewDBClient(token string, dbID string) *NoseDBClient {
	return &NoseDBClient{Token: token, DBID: dbID}
}

//save struct data to db
func (client *NoseDBClient) SaveData(data any) (string, error) {
	req := nosedata.NewCreateDatabasePageReqX(client.DBID, data)
	var res nosedata.CreateObjectRes
	httpClient := noseapi.NewHttpClient(client.Token, "POST")
	err := httpClient.Request(noseapi.CreatePageApi, req, &res)
	if err != nil {
		return "", err
	}
	return res.ID, nil
}

func (client *NoseDBClient) ToPage(pageID string) *NosePageClient {
	return &NosePageClient{Token: client.Token, PageID: pageID}
}

//create new db page with map, not recommended 🔴
// func (client *NoseDBClient) NewPage(title map[string]string, values map[string]any) (string, error) {
// 	req := nosedata.NewCreateDatabasePageReq(client.DBID, title, values)
// 	var res nosedata.CreateObjectRes
// 	httpClient := noseapi.NewHttpClient(client.Token, "POST")
// 	err := httpClient.Request(noseapi.CreatePageApi, req, &res)
// 	if err != nil {
// 		return "", err
// 	}
// 	return res.ID, nil
// }
