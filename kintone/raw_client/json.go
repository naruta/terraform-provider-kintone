package raw_client

import "encoding/json"

func DecodeJson(src []byte, v interface{}) error {
	return json.Unmarshal(src, v)
}

func EncodeJson(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}
