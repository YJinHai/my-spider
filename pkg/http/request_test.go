package http

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestGet(t *testing.T) {
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
			args: args{cUrl: "curl 'https://www.zhihu.com/api/v3/moments?limit=10&desktop=true' \\\n  -H 'authority: www.zhihu.com' \\\n  -H 'x-zse-93: 101_3_2.0' \\\n  -H 'x-ab-param: tp_zrec=1;pf_adjust=1;zr_expslotpaid=2;se_ffzx_jushen1=0;top_test_4_liguangyi=1;zr_slotpaidexp=1;tp_topic_style=0;tp_contents=1;tp_dingyue_video=0;pf_noti_entry_num=2;qap_question_author=0;qap_question_visitor= 0' \\\n  -H 'x-ab-pb: CtwBagGiA1YMzALRBjMERQSNBKwGmwtpATsCtQvgC0cAfQKgA/MD7Ar2AlcEUQUbABEFVQWABbkCtAraBOkEHAbcBioD5wXGBoQCtwPWBDMFVgXjBU8D9AM3BcIFfgaLBQcHtADrBlIFjAVmBzcMQQaUBlMHAQurBtwLpgRcBqYG9AtbBhYG6AZXB9cCoQMZBfoGUgsPCz8AdAFCBDIFPwaiBikFKgb4AwsEfAaJDNcLBwzkCjIDQAYnB0AB4AQVBQEGYAtQA4wEsgU0DNgCVAcwBjEGTwfPC0MACgSeBRJuAQABAAEAAgQDAgABAwAAAQAAAQABAQUAAAAVAAEAAAABBQAAAAAAAQAAAAMCAAAAAAEAAAABBQAAAAAAAAAAAAAAAgAAAAACAQEAAAMAAAQAAAEAAgAAAQAAFwABAgIAAAAEAAAAAQABAAsVAQE=' \\\n  -H 'x-api-version: 3.0.53' \\\n  -H 'x-zst-81: 3_2.0aR_sn77yn6O92wOB8hPZnQr0EMYxc4f18wNBUgpTQX2xE_FqM_FZQ_7qb4kYJvNMPqXTDLe1FCOm7LS1trxTA8OYLUL8iRnqB6kq24rZegO8nCSCzrxMQJxTD_tp6ukyK8FmDcpTdRtuQXtTlbXT2ccmHbS_VUPT16fq-AXmAgxL-cHT_Tf0kUQKF0ntHRX_h0PLoQL8_GOBJGHLdB38FwO1b93mout9krOO3J3m66HBPBOKMMgL-hCMWu3BF0fzqgV8ogCfwwtL-vg0s_H8jJ9_YqeXDvSqkqwBEUoG9qc_iUCyG_tCavN0SwNYo0NfJuYfrbcqVCYBEGO0xhSqLqe1jhgXrDcf3CCmb4uMVho_ChXOWBLC' \\\n  -H 'user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36' \\\n  -H 'sec-ch-ua-mobile: ?0' \\\n  -H 'x-requested-with: fetch' \\\n  -H 'x-zse-96: 2.0_aX20k798o82fb7tBMLtB2TUqb_2YcLS0ThYyb09BeL2f' \\\n  -H 'sec-ch-ua: \" Not A;Brand\";v=\"99\", \"Chromium\";v=\"96\", \"Google Chrome\";v=\"96\"' \\\n  -H 'sec-ch-ua-platform: \"Windows\"' \\\n  -H 'accept: */*' \\\n  -H 'sec-fetch-site: same-origin' \\\n  -H 'sec-fetch-mode: cors' \\\n  -H 'sec-fetch-dest: empty' \\\n  -H 'referer: https://www.zhihu.com/follow' \\\n  -H 'accept-language: zh-CN,zh;q=0.9' \\\n  -H 'cookie: _xsrf=7jmL4hlMA3AJYCi4avhFfkhE4sSlx7sc; _zap=ac43515f-7139-40c1-9aee-cfef3abb11b2; d_c0=\"ADBeqwmquBOPTtlPg1b14ezOIaOeZ4E4jww=|1631590119\"; _9755xjdesxxd_=32; YD00517437729195%3AWM_TID=hR3EvQ8CnVJEARQAEBcrm7lPSwjiGBLx; __snaker__id=kbJWKcbaamT3BKyZ; capsion_ticket=\"2|1:0|10:1637744368|14:capsion_ticket|44:OTY5Y2NiYTE2Yjg5NGM5MDkzMjQ0YWE4MmJhOWY4YWQ=|ec05d1686ada501d18ca90b628500023e0765523b3825fd8857fffbc8e01ed19\"; YD00517437729195%3AWM_NI=u2cdX5JR9kI9fuc%2FT%2B73CETaqdRkFqWO%2FyDLtczJol2%2Fxd%2FUd6r8SMZS9H%2FcLg6E3J1fcbUxi17HVm9I25Arjv%2BiE1iP4WuILYUlT%2BkOOGWSZEWtAIvGqRZRrMdGrqLibVU%3D; YD00517437729195%3AWM_NIKE=9ca17ae2e6ffcda170e2e6eed0b774f78cfe99c56db4a88eb2d55f939e8abaae80fbb1c08ff146a38e8488fc2af0fea7c3b92ab7eead85db3f8a8ef78ccc33a6abadd0d721f5ed8da6d76bb3bc9e8cc57b9591e19ac748b8bdb6abee68f2f1b885ae5fa8eabed9d2479a8df7a9ae6990b9faa6b242f3baad83d874abeff9a9d567fb97fed6c47aaaaae58ac83bab90fcb8f84390879b8dd721a99181a8d6599a9b9995c443b699bfb3d95ab293aca8f17b90b7aeb7d837e2a3; gdxidpyhxdE=S3ta3h%2Bfmw1edxRbwfHn7mMMIm5Kyr6%2Fks1cL8ctNwKC98pvTx%5CTwA60nHoqgUsgjtuCEU2P8mNsKQxPlGVLGT89e6wwxSdSQ7u8wkwtbp0xnJ%5CZ%5CWb7zilzxUBidvQbS886V8qhGROCl0%5C2qfOLPR22lrNmiAqa42yWnH6mXHTjTlzi%3A1639476113268; captcha_session_v2=\"2|1:0|10:1639475215|18:captcha_session_v2|88:UDdPM3h2cGtuWjdEcGdYRG12UWxubjJhMHBsbGJGVVJ5KzY5azhwZ1Z3bjNjTVRhalYrb3A5UkJvTlVRY2dlTg==|007b02b228192e6a7ca3a9a8357a520e8aa448088ed0b7982641f9f68eaa9d1d\"; r_cap_id=\"NzZkZDZlMzIxZjYxNDY4YWFmNzczNmQxNjRlYTc0ZWQ=|1639475220|2e6399f80b7a64e309d89c9c0b1f9d1228bc64b5\"; cap_id=\"YTFmZGE5ZDUxNmExNGI3MmEzZjI1MzE4NGIzZTQ0MmQ=|1639475220|2e9cec90b3d41dddf2e2680a4920551547b72f96\"; l_cap_id=\"MWMyNzE2ZGEyNjVlNDEwMDgyMmE1MmE5MjEzOTFmZTc=|1639475220|c0a86b1984c669c428045e1476556b6c5b04a09f\"; z_c0=Mi4xNzQzWkF3QUFBQUFBTUY2ckNhcTRFeGNBQUFCaEFsVk5OcmFsWWdEWGNUSVpJQVJucFdtN0tTODFtMWc2VVNNUmhn|1639475254|d7f4d6af37504b2dce57b67174add313b43f5847; q_c1=63430e2919854fcf8c733e592969b2bb|1639475272000|1639475272000; Hm_lvt_98beee57fd2ef70ccdd5ca52b9740c49=1639554907,1639554915,1639554930,1640081700; NOT_UNREGISTER_WAITING=1; Hm_lpvt_98beee57fd2ef70ccdd5ca52b9740c49=1640081762; KLBRSID=37f2e85292ebb2c2ef70f1d8e39c2b34|1640081774|1640081698; tst=f' \\\n  --compressed"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := Get(tt.args.cUrl)
			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(string(body))
		})
	}
}

func TestPost(t *testing.T) {
	res := 	Post("curl 'https://api.juejin.cn/recommend_api/v1/article/recommend_follow_feed?aid=2608&uuid=7007317662066935304' \\\n  -H 'authority: api.juejin.cn' \\\n  -H 'sec-ch-ua: \" Not A;Brand\";v=\"99\", \"Chromium\";v=\"96\", \"Google Chrome\";v=\"96\"' \\\n  -H 'sec-ch-ua-mobile: ?1' \\\n  -H 'user-agent: Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Mobile Safari/537.36' \\\n  -H 'sec-ch-ua-platform: \"Android\"' \\\n  -H 'content-type: application/json' \\\n  -H 'accept: */*' \\\n  -H 'origin: https://juejin.cn' \\\n  -H 'sec-fetch-site: same-site' \\\n  -H 'sec-fetch-mode: cors' \\\n  -H 'sec-fetch-dest: empty' \\\n  -H 'referer: https://juejin.cn/' \\\n  -H 'accept-language: zh-CN,zh;q=0.9' \\\n  -H 'cookie: _ga=GA1.2.1103167808.1631518283; MONITOR_WEB_ID=ff79f45b-40c0-406b-ba1a-d83365e70f70; n_mh=nkk5hZarLFMvBmVvwLII7okiifRznjjtMtzWfBjPXlI; _tea_utm_cache_2608={%22utm_source%22:%22gold_browser_extension%22}; passport_csrf_token_default=f160ae3385d60879370ca98939da1d16; passport_csrf_token=f160ae3385d60879370ca98939da1d16; sid_guard=699b2f16bb29f508410fd437dd5e6fb2%7C1639473797%7C5184000%7CSat%2C+12-Feb-2022+09%3A23%3A17+GMT; uid_tt=7a027102e016fdd609fa20bfd91d2c80; uid_tt_ss=7a027102e016fdd609fa20bfd91d2c80; sid_tt=699b2f16bb29f508410fd437dd5e6fb2; sessionid=699b2f16bb29f508410fd437dd5e6fb2; sessionid_ss=699b2f16bb29f508410fd437dd5e6fb2; sid_ucp_v1=1.0.0-KGFiMGI0ODMwOGJmYjkyNzc2NjE1NjQzNDRlOGRiMGNlMDc4Mjk0OWMKFgi9i6C__fWJBxCFxeGNBhiwFDgIQAsaAmxmIiA2OTliMmYxNmJiMjlmNTA4NDEwZmQ0MzdkZDVlNmZiMg; ssid_ucp_v1=1.0.0-KGFiMGI0ODMwOGJmYjkyNzc2NjE1NjQzNDRlOGRiMGNlMDc4Mjk0OWMKFgi9i6C__fWJBxCFxeGNBhiwFDgIQAsaAmxmIiA2OTliMmYxNmJiMjlmNTA4NDEwZmQ0MzdkZDVlNmZiMg; _gid=GA1.2.2138369023.1641783448' \\\n  --data-raw '{\"id_type\":2,\"cursor\":\"0\"}' \\\n  --compressed")
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
