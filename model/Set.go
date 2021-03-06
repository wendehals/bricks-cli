package model

// Set represents a Lego set.
type Set struct {
	SetNum      string `json:"set_num"`
	Name        string `json:"name"`
	Year        uint   `json:"year"`
	NumParts    uint   `json:"num_parts"`
	SetURL      string `json:"set_url"`
	SetImageURL string `json:"set_img_url"`
}
