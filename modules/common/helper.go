package common

import "encoding/json"

func ConvertThrowJson(input interface{},output interface{}) error {
	content,err := json.Marshal(input)
	if err != nil {
		return err
	}
	return json.Unmarshal(content,output)
}
