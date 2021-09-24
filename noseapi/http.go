package noseapi

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func HttpPost(url string, body interface{}, res interface{}, token string) error {
	postBody, _ := json.Marshal(body)
	responseBody := bytes.NewBuffer(postBody)
	req, _ := http.NewRequest("POST", url, responseBody)
	req.Header.Set("Authorization", token)
	req.Header.Set("Notion-Version", ApiVersion)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()
	bodyRes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}
	sb := string(bodyRes)
	json.Unmarshal([]byte(sb), res)
	return nil
}

func HttpPatch(url string, body interface{}, res interface{}, token string) error {
	postBody, _ := json.Marshal(body)
	responseBody := bytes.NewBuffer(postBody)
	req, err := http.NewRequest("PATCH", url, responseBody)
	if err != nil {
		log.Panic(err)
	}
	req.Header.Set("Authorization", token)
	req.Header.Set("Notion-Version", ApiVersion)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()
	bodyRes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}
	sb := string(bodyRes)
	println(sb)
	err = json.Unmarshal([]byte(sb), res)
	if err != nil {
		log.Panic(err)
	}
	return nil
}
