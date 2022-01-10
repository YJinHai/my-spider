package http

import (
	"net/http"
	"strings"
)

func Get(cUrl string) *http.Response {
	requestInfo := CUrlToRequestInfo(cUrl)
	req, err := http.NewRequest("GET", requestInfo.Address, nil)
	if err != nil {
		return nil
	}

	for k,v := range requestInfo.Header {
		req.Header.Set(k,v)
	}

	clt := &http.Client{}
	res, err := clt.Do(req)

	if err != nil {
		return nil
	}

	return res
}

func Post(cUrl string)  *http.Response {
	requestInfo := CUrlToRequestInfo(cUrl)
	req, err := http.NewRequest("POST", requestInfo.Address, strings.NewReader(requestInfo.JsonPost))
	if err != nil {
		return nil
	}

	for k,v := range requestInfo.Header {
		req.Header.Set(k,v)
	}
	clt := &http.Client{}
	res, err := clt.Do(req)
	if err != nil {
		return nil
	}

	return res
}
