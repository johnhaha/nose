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

func (client *NosePageClient) NewEmptyPage(name string) (string, error) {
	req := nosedata.NewEmptyPageReq(client.PageID, name)
	var res nosedata.CreateObjectRes
	httpClient := noseapi.NewHttpClient(client.Token, "POST")
	err := httpClient.Request(noseapi.CreatePageApi, req, &res)
	if err != nil {
		return "", err
	}
	log.Println("new page id is", res.ID)
	return res.ID, nil
}

//create db page with struct value
func (client *NosePageClient) NewDB(name string, d any) (string, error) {
	req := nosedata.NewEmptyDatabaseReqX(client.PageID, name, d)
	var res nosedata.CreateObjectRes
	httpClient := noseapi.NewHttpClient(client.Token, "POST")
	err := httpClient.Request(noseapi.CreateDatabaseApi, req, &res)
	if err != nil {
		return "", err
	}
	return res.ID, nil
}

func (client *NosePageClient) AppendTextBlock(texts ...string) error {
	req := nosedata.NewAppendTextBlockReq(texts...)
	var res nosedata.AppendBlockRes
	api := noseapi.BlockApi.AddParam(client.PageID).AddParam("children")
	httpClient := noseapi.NewHttpClient(client.Token, "PATCH")
	err := httpClient.Request(string(api), req, &res)
	if err != nil {
		return err
	}
	return nil
}

//append data text from struct
func (client *NosePageClient) AppendDataTextBlock(data any) error {
	req := nosedata.NewAppendDataTextBlockReq(data)
	var res nosedata.AppendBlockRes
	api := noseapi.BlockApi.AddParam(client.PageID).AddParam("children")
	httpClient := noseapi.NewHttpClient(client.Token, "PATCH")
	err := httpClient.Request(string(api), req, &res)
	if err != nil {
		return err
	}
	return nil
}

func (client *NosePageClient) AppendTodoBlock(text string) error {
	req := nosedata.NewAppendTodoBlockReq(text)
	var res nosedata.AppendBlockRes
	api := noseapi.BlockApi.AddParam(client.PageID).AddParam("children")
	httpClient := noseapi.NewHttpClient(client.Token, "PATCH")
	err := httpClient.Request(string(api), req, &res)
	if err != nil {
		return err
	}
	return nil
}

func (client *NosePageClient) ToDB(DBID string) *NoseDBClient {
	return &NoseDBClient{Token: client.Token, DBID: DBID}
}

//create db page with keys, not recommended ????
// func (client *NosePageClient) NewEmptyDatabase(name string, columnTitle string, columns ...string) (string, error) {
// 	req := nosedata.NewEmptyDatabaseReq(client.PageID, name, columnTitle, columns...)
// 	var res nosedata.CreateObjectRes
// 	httpClient := noseapi.NewHttpClient(client.Token, "POST")
// 	err := httpClient.Request(noseapi.CreateDatabaseApi, req, &res)
// 	if err != nil {
// 		return "", err
// 	}
// 	// log.Println("new page id is", res.ID)
// 	return res.ID, nil
// }
