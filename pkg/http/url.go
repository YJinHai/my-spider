package http

import (
	"my-spider/pkg/util"
)

type RequestInfo struct {
	Address string
	Header map[string]string
	JsonPost string
}


func CUrlToRequestInfo(cUrl string) RequestInfo {
	info := RequestInfo{}
	info.Address = util.GetCUrlAddr(cUrl)
	info.Header = util.GetCurlHeader(cUrl)
	info.JsonPost = util.GetCUrlPost(cUrl)

	return info
}
