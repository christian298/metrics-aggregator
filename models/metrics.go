package models

type Metric struct {
	Name  string  `json:"name"`
	Value float32 `json:"value"`
	Id    string  `json:"id"`
}
