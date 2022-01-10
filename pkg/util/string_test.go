package util

import (
	"fmt"
	"testing"
)

func TestCUrlToSlice(t *testing.T) {
	type args struct {
		cUrl string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{
			name: "掘金",
			args: args{cUrl: "  curl 'https://api.juejin.cn/content_api/v1/article/query_list?aid=2608&uuid=7007317662066935304' \\\n  -H 'authority: api.juejin.cn' \\\n  -H 'sec-ch-ua: \" Not A;Brand\";v=\"99\", \"Chromium\";v=\"96\", \"Google Chrome\";v=\"96\"' \\\n  -H 'sec-ch-ua-mobile: ?0' \\\n  -H 'user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.45 Safari/537.36' \\\n  -H 'sec-ch-ua-platform: \"Windows\"' \\\n  -H 'content-type: application/json' \\\n  -H 'accept: */*' \\\n  -H 'origin: https://juejin.cn' \\\n  -H 'sec-fetch-site: same-site' \\\n  -H 'sec-fetch-mode: cors' \\\n  -H 'sec-fetch-dest: empty' \\\n  -H 'referer: https://juejin.cn/' \\\n  -H 'accept-language: zh-CN,zh;q=0.9' \\\n  -H 'cookie: _ga=GA1.2.1103167808.1631518283; MONITOR_WEB_ID=ff79f45b-40c0-406b-ba1a-d83365e70f70; n_mh=nkk5hZarLFMvBmVvwLII7okiifRznjjtMtzWfBjPXlI; _tea_utm_cache_2608={%22utm_source%22:%22gold_browser_extension%22}; _gid=GA1.2.1619367753.1639448540; passport_csrf_token_default=f160ae3385d60879370ca98939da1d16; passport_csrf_token=f160ae3385d60879370ca98939da1d16; sid_guard=699b2f16bb29f508410fd437dd5e6fb2%7C1639473797%7C5184000%7CSat%2C+12-Feb-2022+09%3A23%3A17+GMT; uid_tt=7a027102e016fdd609fa20bfd91d2c80; uid_tt_ss=7a027102e016fdd609fa20bfd91d2c80; sid_tt=699b2f16bb29f508410fd437dd5e6fb2; sessionid=699b2f16bb29f508410fd437dd5e6fb2; sessionid_ss=699b2f16bb29f508410fd437dd5e6fb2; sid_ucp_v1=1.0.0-KGFiMGI0ODMwOGJmYjkyNzc2NjE1NjQzNDRlOGRiMGNlMDc4Mjk0OWMKFgi9i6C__fWJBxCFxeGNBhiwFDgIQAsaAmxmIiA2OTliMmYxNmJiMjlmNTA4NDEwZmQ0MzdkZDVlNmZiMg; ssid_ucp_v1=1.0.0-KGFiMGI0ODMwOGJmYjkyNzc2NjE1NjQzNDRlOGRiMGNlMDc4Mjk0OWMKFgi9i6C__fWJBxCFxeGNBhiwFDgIQAsaAmxmIiA2OTliMmYxNmJiMjlmNTA4NDEwZmQ0MzdkZDVlNmZiMg' \\\n  --data-raw '{\"user_id\":\"3984285869016509\",\"sort_type\":2,\"cursor\":\"0\"}' \\\n  --compressed"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ret := CUrlToSlice(tt.args.cUrl)
			fmt.Println(ret)
		})
	}
}
