package http

import (
	"encoding/json"
	"io/ioutil"
	"my-spider/pkg/log"
	"net/http"
)

func GetResponseBodyJson(res *http.Response) string {
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Logger.Log(log.LevelWarn, err)
		return ""
	}

	return string(body)
}

func GetResponseBodyMap(res *http.Response) map[string]interface{} {
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Logger.Log(log.LevelError, err)
		return nil
	}

	m := make(map[string]interface{})
	err = json.Unmarshal(body, &m)
	if err != nil {
		log.Logger.Log(log.LevelError, err)
		return nil
	}

	return m
}
