package nose

import (
	"log"

	"github.com/johnhaha/nose/noseapi"
	"github.com/johnhaha/nose/nosedata"
)

type NosePageClient struct {
	Token  string `json:"token"`
	PageID string `json:"pageID"`
}

func NewPageClient(token string, pageID string) *NosePageClient {
	return &NosePageClient{Token: token, PageID: pageID}
}

func (client *NosePageClient) NewEmptyPage(name string) string {
	req := nosedata.NewEmptyPageReq(client.PageID, name)
	var res nosedata.CreateObjectRes
	err := noseapi.HttpPost(noseapi.CreatePageApi, req, &res, client.Token)
	if err != nil {
		panic(err)
	}
	log.Println("new page id is", res.ID)
	return res.ID
}

func (client *NosePageClient) NewEmptyDatabase(name string, columnTitle string, columns ...string) string {
	req := nosedata.NewEmptyDatabaseReq(client.PageID, name, columnTitle, columns...)
	var res nosedata.CreateObjectRes
	err := noseapi.HttpPost(noseapi.CreateDatabaseApi, req, &res, client.Token)
	if err != nil {
		panic(err)
	}
	log.Println("new page id is", res.ID)
	return res.ID
}

func (client *NosePageClient) AppendTextBlock(texts ...string) error {
	req := nosedata.NewAppendTextBlockReq(texts...)
	var res nosedata.AppendBlockRes
	api := noseapi.BlockApi.AddParam(client.PageID).AddParam("children")
	err := noseapi.HttpPatch(string(api), req, &res, client.Token)
	if err != nil {
		return err
	}
	return nil
}

func (client *NosePageClient) AppendTodoBlock(text string) error {
	req := nosedata.NewAppendTodoBlockReq(text)
	var res nosedata.AppendBlockRes
	api := noseapi.BlockApi.AddParam(client.PageID).AddParam("children")
	err := noseapi.HttpPatch(string(api), req, &res, client.Token)
	if err != nil {
		return err
	}
	return nil
}

func (client *NosePageClient) ToDB(DBID string) *NoseDBClient {
	return &NoseDBClient{Token: client.Token, DBID: DBID}
}
