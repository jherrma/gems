package models

type NearestItemsRequest struct {
	Phrase string `json:"phrase"`
	Limit  int64  `json:"limit"`
}
