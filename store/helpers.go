package main

import (
	"bytes"
	"encoding/json"
)

func jsonRequest(data interface{}) *bytes.Buffer {
	body, _ := json.Marshal(data)
	return bytes.NewBuffer(body)
}

