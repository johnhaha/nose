package nose

import (
	"log"

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

func (client *NoseDBClient) NewPage(title map[string]string, values map[string]string) string {
	req := nosedata.NewCreateDatabasePageReq(client.DBID, title, values)
	var res nosedata.CreateObjectRes
	err := noseapi.HttpPost(noseapi.CreatePageApi, req, &res, client.Token)
	if err != nil {
		log.Panic(err)
	}
	return res.ID
}

func (client *NoseDBClient) ToPage(pageID string) *NosePageClient {
	return &NosePageClient{Token: client.Token, PageID: pageID}
}
