package http

import (
	"encoding/json"
	"reflect"
	"strings"
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

func GetJsonTag(v interface{}) map[string]string {
	var strM map[string]string
	val := reflect.ValueOf(v)
	typ := val.Type()
	for i := 0; i < typ.NumField(); i++ {
		sf := typ.Field(i)
		if sf.PkgPath != "" && !sf.Anonymous { // unexported
			continue
		}

		tag := sf.Tag.Get("json")
		if tag == "-" {
			continue
		}

		strM[sf.Name] = tag
	}

	return strM
}

// GetReflectTag 层级递增解析tag
func getReflectTag(reflectType reflect.Type) (string, string) {

	if reflectType.Kind() != reflect.Struct {
		return "", ""
	}

	var fields strings.Builder
	primaryKey := ""
	baseModel := ""
	for i := 0; i < reflectType.NumField(); i++ {
		tag := reflectType.Field(i).Tag.Get("json")
		if tag == "-" {
			continue
		}

		if tag == "" {
			getReflectTag(reflectType.Field(i).Type)
			continue
		}

		name, isPrimaryKey := parseTag(tag)
		if isPrimaryKey {
			primaryKey = name
			continue
		}

		if name == "" {
			continue
		}

		fields.WriteString(",")
		fields.WriteString(name)
	}

	if baseModel != "" {
		var newS strings.Builder
		newS.WriteString(",")
		newS.WriteString(baseModel)
		newS.WriteString(fields.String())
		fields.Reset()
		fields = newS
	}

	return fields.String(), primaryKey
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
