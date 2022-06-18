package noseapi

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type NotionHttpClient struct {
	Token       string
	RequestType string
}

func NewHttpClient(token string, requestType string) *NotionHttpClient {
	return &NotionHttpClient{Token: token, RequestType: requestType}
}

func (client *NotionHttpClient) Request(url string, body any, res any) error {
	postBody, _ := json.Marshal(body)
	responseBody := bytes.NewBuffer(postBody)
	req, _ := http.NewRequest(client.RequestType, url, responseBody)
	req.Header.Set("Authorization", "Bearer "+client.Token)
	req.Header.Set("Notion-Version", ApiVersion)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	bodyRes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	sb := string(bodyRes)
	json.Unmarshal([]byte(sb), res)
	return nil
}
