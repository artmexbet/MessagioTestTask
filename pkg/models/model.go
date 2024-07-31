package models

import "encoding/json"

// AddMessageJSON ...
type AddMessageJSON struct {
	Title string          `json:"title"`
	Data  json.RawMessage `json:"data"`
}

// Message ...
type Message struct {
	Id string `json:"id"`
	AddMessageJSON
}
