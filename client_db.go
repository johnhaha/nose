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

func (client *NoseDBClient) NewPage(title map[string]string, values map[string]string) (string, error) {
	req := nosedata.NewCreateDatabasePageReq(client.DBID, title, values)
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
