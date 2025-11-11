package jsonpretty

import (
	"bytes"
	"encoding/json"
)

func Pretty(v any) []byte {
	var prettyJSON bytes.Buffer
	data, err := json.Marshal(v)
	if err != nil {
		return make([]byte, 0)
	}
	json.Indent(&prettyJSON, data, "", "  ")
	return prettyJSON.Bytes()
}
