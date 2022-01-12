package http

import (
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"unsafe"
)

func JsonUnmarshal(data string, v interface{}) error {
	err := json.Unmarshal([]byte(data), v)
	return err
}

type JsonData struct {
	Data []interface{}
}

type JsonMap struct {
	Map map[string]interface{}
}

type Article struct {
	ArticleId    string `json:"article_id"`
	CategoryId   string `json:"category_id"`
	BriefContent string `json:"brief_content"`
	Ctime        string `json:"ctime"`
	Omitempty    int64  `json:"omitempty"`
}

func JsonToStructList(data string, f func(data []byte) (interface{}, error)) []interface{} {
	var (
		jsonD JsonData
		list  []interface{}
	)
	err := json.Unmarshal([]byte(data), &jsonD)
	if err != nil {
		return nil
	}
	for _, d := range jsonD.Data {
		d2 := d.(map[string]interface{})
		temp := jsonToStruct(d2, f)
		list = append(list, temp...)
	}
	fmt.Println("err:", err)

	return list
}

func jsonToStruct(data map[string]interface{}, f func(data []byte) (interface{}, error)) []interface{} {
	var list []interface{}

	for _, d := range data {
		b, err := json.Marshal(d)
		if err != nil {
			//
			//l := make(map[string]interface{})
			//l["err"] = err
			//l["msg"] = "json.Marshal failed"
			//
			//log.Logger.Log(log.LevelWarn, l)
			d2 := d.(map[string]interface{})
			jsonToStruct(d2, f)
			continue
		}

		ret, err2 := f(b)
		if err2 != nil {
			continue
		}
		list = append(list, ret)
	}

	return list
}

func JsonRegexp(data string, v interface{}, dest uintptr) interface{} {
	//switch v.(type) {
	//case uintptr:
	//
	//}
	//vt := &v

	var strIndex []int

	tM1 := make(map[int]string)
	sM1 := make(map[int]string)
	//pM1 := make(map[string]interface{})
	val := reflect.ValueOf(v)
	typ := val.Type()
	for i := 0; i < typ.NumField(); i++ {
		sf := typ.Field(i)
		tag := sf.Tag.Get("json")
		if tag == "-" {
			continue
		}
		tM1[int(sf.Offset)] = fmt.Sprint(sf.Type)
		sM1[int(sf.Offset)] = tag
		strIndex = append(strIndex, int(sf.Offset))
	}
	data += ","
	newS := strings.Replace(strings.TrimSpace(data), " ", "", -1)
	newS = strings.Replace(strings.TrimSpace(newS), "\n", "", -1)

	for _, b := range strIndex {
		//teamp := *(*interface{})(unsafe.Pointer(dest + uintptr(b)))
		//pM1[sM1[b]] = &teamp

		regexpS := "\"" + sM1[b] + "\"" + `:(\S+?)((,"[a-zA-Z_]+":)|(},))`
		var RegexAmount = regexp.MustCompile(regexpS)

		match := RegexAmount.FindStringSubmatch(newS)
		if len(match) < 2 {
			continue
		}
		value := match[1]
		switch tM1[b] {
		case "int64":
			fmt.Println("int64")
			amount, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				return nil
			}
			*(*int64)(unsafe.Pointer(dest + uintptr(b))) = amount
		case "string":
			fmt.Println("string:", value)
			*(*string)(unsafe.Pointer(dest + uintptr(b))) = strings.Replace(strings.TrimSpace(value), "\"", "", -1)
		}
	}
	//endIndex := strIndex[len(strIndex)-1]
	////*(*int64)(unsafe.Pointer(dest + uintptr(1))) = 1
	//*(*string)(unsafe.Pointer(dest - uintptr(endIndex))) = "1"

	fmt.Println("v:", v)
	//var desC interface{}
	////copy(desC, v)
	//tmp :=
	//vt = v
	return v
}

func parseTag(tag string) (string, bool) {
	is := false
	s := strings.Split(tag, ":")
	if len(s) < 2 {
		return "", is
	}
	if strings.Contains(s[1], ";") {
		s2 := strings.Split(s[1], ";")
		if s2[1] == "primary_key" {
			is = true
		}
		return s2[0], is
	}
	return s[1], is
}
