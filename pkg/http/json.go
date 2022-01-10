package http

import "encoding/json"

func JsonUnmarshal(data string, v interface{})  error {
	err := json.Unmarshal([]byte(data), v)
	return err
}


type JsonData struct {
	Data []interface{}
}

type JsonMap struct {
	Map map[string]interface{}
}
