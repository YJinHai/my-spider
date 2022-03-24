package util

import (
	"encoding/json"
	"fmt"
)

type Article struct {
	ArticleId    string `json:"article_id"`
	CategoryId   string `json:"category_id"`
	Title        string `json:"title"`
	BriefContent string `json:"brief_content"`
	Ctime        string `json:"ctime"`
	Omitempty    int64  `json:"omitempty"`
}

type WbArticle struct {
	Text       string `json:"text"`
	TextRaw    string `json:"text_raw"`
	ScreenName string `json:"screen_name"`
}

// JsonToDataList 'data' param is http body json, 'f' param is func type example:
// 	f := func(data []byte) (interface{}, error) {
//		var ret *WbArticle
//		err3 := json.Unmarshal(data, &ret)
//		if err3 != nil || (ret != nil && ret.TextRaw == "") {
//			return nil, errors.New("空字段")
//
//		}
//
//		return ret, nil
//	}

// 支持整文本和定义的结构体匹配
func JsonToDataList(data string, f func(data []byte) (interface{}, error)) []interface{} {
	var (
		keyValue map[string]interface{}
		list     []interface{}
	)

	err := json.Unmarshal([]byte(data), &keyValue)
	if err != nil {
		return nil
	}
	for _, v := range keyValue {
		temp := jsonToStructList(v, f)
		list = append(list, temp...)
	}
	fmt.Println("err:", err)

	return list
}

//func jsonToStructList(data interface{}, f func(data []byte) (interface{}, error)) []interface{} {
//	var list []interface{}
//
//	switch data.(type) {
//	case map[string]interface{}:
//		b, err := json.Marshal(data)
//		if err != nil {
//			jsonToStructList(data, f)
//			return list
//		}
//
//		ret, err2 := f(b)
//		if err2 != nil || ret == nil {
//			for _, d := range data.(map[string]interface{}) {
//				list = append(list, jsonToStructList(d, f)...)
//			}
//			return list
//		}
//		list = append(list, ret)
//
//	case []interface{}:
//		for _, d := range data.([]interface{}) {
//			list = append(list, jsonToStructList(d, f)...)
//		}
//
//	default:
//		return list
//	}
//
//	return list
//}

func jsonToStructList(data interface{}, f func(data []byte) (interface{}, error)) []interface{} {
	var list []interface{}

	switch data.(type) {
	case map[string]interface{}:
		data = jsonMergeField(data.(map[string]interface{}), nil)
		b, err := json.Marshal(data)
		if err != nil {
			jsonToStructList(data, f)
			return nil
		}

		ret, err2 := f(b)
		if err2 != nil || ret == nil {
			return nil
		}

		list = append(list, ret)

	case []interface{}:
		for _, d := range data.([]interface{}) {
			list = append(list, jsonToStructList(d, f)...)
		}

	default:
		return list
	}

	return list
}

// 拼接map[string]interface{}
func jsonMergeField(source map[string]interface{}, obj map[string]interface{}) map[string]interface{} {
	if obj == nil {
		obj = make(map[string]interface{})
	}

	for k, v := range source {
		if _, ok := obj[k]; !ok {
			switch v.(type) {
			case map[string]interface{}:
				jsonMergeField(v.(map[string]interface{}), obj)
			default:
				obj[k] = v
			}

		}
	}

	return obj

}
