package utils

import (
	"encoding/json"
	"net/http"
)

func Reverse(s string) string {
	a := []rune(s)
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return string(a)
}

func ConvertJSONToHeader(jsonStr string) (http.Header, error) {
	var headersMap map[string]string
	err := json.Unmarshal([]byte(jsonStr), &headersMap)
	if err != nil {
		return nil, err
	}
	headers := http.Header{}
	for key, value := range headersMap {
		headers.Add(key, value)
	}
	return headers, nil
}
