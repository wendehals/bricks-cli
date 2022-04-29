package model

// PartEntry represents an entry of a single part in a collection with its quantity and color.
type PartEntry struct {
	Quantity int   `json:"quantity"`
	Part     Part  `json:"part"`
	Color    Color `json:"color"`
	IsSpare  bool  `json:"is_spare"`
}
