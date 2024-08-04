package models

import "encoding/json"

// AddMessageJSON ...
type AddMessageJSON struct {
	Title string          `json:"title"`
	Data  json.RawMessage `json:"data,omitempty"`
}

// Message ...
type Message struct {
	Id string `json:"id"`
	AddMessageJSON
}

type MsgInfo struct {
	Id int64 `json:"id"`
}
